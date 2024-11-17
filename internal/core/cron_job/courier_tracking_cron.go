package cronjob

import (
	"log"
	"log/slog"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/stelgkio/otoo/internal/core/domain"
	cr "github.com/stelgkio/otoo/internal/core/domain/courier"
	"github.com/stelgkio/otoo/internal/core/port"
	woo_service "github.com/stelgkio/otoo/internal/core/service/woocommerce"
	"github.com/stelgkio/otoo/internal/core/util"
	woo "github.com/stelgkio/woocommerce"
)

// CourierTrackingCron represents the cron job for order analytics
type CourierTrackingCron struct {
	projectSvc   port.ProjectService
	userSvc      port.UserService
	orderSvc     port.OrderService
	extensionSvc port.ExtensionService
	voucherSvc   port.VoucherService
	hermesSvc    port.HermesService
}

// NewCourierTrackingCron creates a new OrderAnalyticsCron instance
func NewCourierTrackingCron(
	projectSvc port.ProjectService,
	userSvc port.UserService,
	orderSvc port.OrderService,
	extensionSvc port.ExtensionService,
	voucherSvc port.VoucherService,
	hermesSvc port.HermesService) *CourierTrackingCron {
	return &CourierTrackingCron{
		projectSvc:   projectSvc,
		userSvc:      userSvc,
		orderSvc:     orderSvc,
		extensionSvc: extensionSvc,
		voucherSvc:   voucherSvc,
		hermesSvc:    hermesSvc,
	}
}

// RunCourier4uTrackingCron runs the cron job for order analytics
func (c *CourierTrackingCron) RunCourier4uTrackingCron() error {
	// Get all projects
	projects, err := c.projectSvc.GetAllProjects()
	if err != nil {
		slog.Error("Error getting projects: ", "error", err)
		return err
	}
	var wg sync.WaitGroup

	for _, project := range projects {
		wg.Add(1)
		go func(project *domain.Project) {
			defer wg.Done()
			projectID := project.Id
			projectExtensions, err := c.extensionSvc.GetAllProjectExtensions(nil, projectID.String())
			if err != nil {
				slog.Error("Error getting project extensions: ", "error", err)
				return
			}

			extensions := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
				return e.Code == domain.Courier4u
			})
			if len(extensions) == 0 {
				return
			}
			courier4u := new(domain.Courier4uExtension)
			courier4u, err = c.extensionSvc.GetCourier4uProjectExtensionByID(nil, extensions[0].ExtensionID, projectID.String())
			if err != nil {
				slog.Error("Error getting courier4u extension: ", "error", err)
				return
			}
			if courier4u == nil {
				slog.Error("courier4u extension does not exist: "+project.WoocommerceProject.Domain, "error", err)
				return
			}
			wg.Add(1)
			client := woo_service.InitClient(project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain)
			vooucherCountChan := make(chan int64, 1)
			errChan := make(chan error, 1)
			go func() {
				defer wg.Done()
				c.voucherSvc.GetVoucherCountAsync(projectID.String(), cr.VoucherStatusProcessing, vooucherCountChan, errChan)
			}()
			go func() {
				wg.Wait()
				close(vooucherCountChan)
				close(errChan)
			}()
			var totalVouchers int64
			for item := range vooucherCountChan {
				if item == 0 {
					return
				}
				totalVouchers = item
			}
			for errList := range errChan {
				if errList != nil {
					return
				}
			}

			if totalVouchers == 0 {
				return
			}
			workers := int(math.Ceil(float64(totalVouchers) / 100))
			if workers == 0 {
				workers = 1
			}
			voucherListChan := make(chan []*cr.Voucher, workers)
			errListChan := make(chan error, 1)
			for i := 0; i < int(workers); i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					c.voucherSvc.FindVoucherByProjectIDAsync(projectID.String(), 100, i+1, "orderId", "desc", cr.VoucherStatusProcessing, voucherListChan, errListChan)
				}()
			}

			go func() {
				wg.Wait()
				close(voucherListChan)
				close(errListChan)
			}()
			for errList := range errListChan {
				if errList != nil {
					return
				}
			}
			if len(voucherListChan) == 0 {
				return
			}
			for item := range voucherListChan {
				c.updateHermesTracking(client, item, projectID.String(), totalVouchers, courier4u, nil)
				time.Sleep(1 * time.Second)
			}
		}(project)
	}
	go func() {
		wg.Wait()
	}()
	return nil
}

// RunRedCourierTrackingCron updates the hermes tracking for the given vouchers
func (c *CourierTrackingCron) RunRedCourierTrackingCron() error {
	// Get all projects
	projects, err := c.projectSvc.GetAllProjects()
	if err != nil {
		slog.Error("Error getting projects: ", "error", err)
		return err
	}
	var wg sync.WaitGroup

	for _, project := range projects {
		wg.Add(1)
		go func(project *domain.Project) {
			defer wg.Done()
			projectID := project.Id
			projectExtensions, err := c.extensionSvc.GetAllProjectExtensions(nil, projectID.String())
			if err != nil {
				slog.Error("Error getting project extensions: ", "error", err)
				return
			}

			extensions := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
				return e.Code == domain.RedCourier
			})
			if len(extensions) == 0 {
				return
			}
			redcourier := new(domain.RedCourierExtension)
			redcourier, err = c.extensionSvc.GetRedCourierProjectExtensionByID(nil, extensions[0].ExtensionID, projectID.String())
			if err != nil {
				slog.Error("Error getting redcourier extension: ", "error", err)
				return
			}
			if redcourier == nil {
				slog.Error("redcourier extension does not exist: "+project.WoocommerceProject.Domain, "error", err)
				return
			}
			wg.Add(1)
			client := woo_service.InitClient(project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain)
			vooucherCountChan := make(chan int64, 1)
			errChan := make(chan error, 1)
			go func() {
				defer wg.Done()
				c.voucherSvc.GetVoucherCountAsync(projectID.String(), cr.VoucherStatusProcessing, vooucherCountChan, errChan)
			}()
			go func() {
				wg.Wait()
				close(vooucherCountChan)
				close(errChan)
			}()
			var totalVouchers int64
			for item := range vooucherCountChan {
				if item == 0 {
					return
				}
				totalVouchers = item
			}
			for errList := range errChan {
				if errList != nil {
					return
				}
			}

			if totalVouchers == 0 {
				return
			}

			workers := int(math.Ceil(float64(totalVouchers) / 100))
			if workers == 0 {
				workers = 1
			}
			voucherListChan := make(chan []*cr.Voucher, workers)
			errListChan := make(chan error, 1)
			for i := 0; i < int(workers); i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					c.voucherSvc.FindVoucherByProjectIDAsync(projectID.String(), 100, i+1, "orderId", "desc", cr.VoucherStatusProcessing, voucherListChan, errListChan)
				}()
			}

			go func() {
				wg.Wait()
				close(voucherListChan)
				close(errListChan)
			}()
			for errList := range errListChan {
				if errList != nil {
					return
				}
			}
			if len(voucherListChan) == 0 {
				return
			}
			for item := range voucherListChan {
				c.updateHermesTracking(client, item, projectID.String(), totalVouchers, nil, redcourier)
				time.Sleep(1 * time.Second)
			}
		}(project)
	}
	go func() {
		wg.Wait()
	}()
	return nil
}

// updateHermesTracking saves the result of a voucher to MongoDB
func (c *CourierTrackingCron) updateHermesTracking(client *woo.Client, vouvherList []*cr.Voucher, projectID string, totalVoucher int64, courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension) error {
	var wg sync.WaitGroup
	voucherCh := make(chan *cr.Voucher, totalVoucher) // Channel to distribute products to workers
	errorCh := make(chan *cr.Voucher, 1)              // Buffered channel for error results

	workers := int(math.Ceil(float64(totalVoucher) / 100))
	if workers == 0 {
		workers = 1
	}

	// Worker pool to process products
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(projectID string, client *woo.Client) {
			defer wg.Done()
			for voucher := range voucherCh {
				err := c.saveVoucherTrackingResult(voucher, projectID, client)
				if err != nil {
					log.Printf("Failed to save product result: %v", err)
					errorCh <- voucher
				}
			}
		}(projectID, client)
	}

	// Goroutine to process error results
	go func(projectID string, client *woo.Client) {
		for result := range errorCh {
			err := c.saveVoucherTrackingResult(result, projectID, client)
			if err != nil {
				log.Printf("Failed to save error result: %v", err)
			}
		}
	}(projectID, client)

	// Create a semaphore to limit the number of concurrent fetches
	maxConcurrentFetches := 10
	sem := make(chan struct{}, maxConcurrentFetches) // Semaphore with a limit
	// Track the total number of products fetched and use mutex for safe concurrent access
	var totalFetched int = 0
	var mu sync.Mutex // Mutex to synchronize access to totalFetched
	// Fetch products concurrently by page
	go func(courier4u *domain.Courier4uExtension, redcourier *domain.RedCourierExtension) {
		var fetchWg sync.WaitGroup

		for _, voucher := range vouvherList {
			sem <- struct{}{} // Acquire a token before launching a new goroutine
			fetchWg.Add(1)

			go func(voucher *cr.Voucher) {
				defer fetchWg.Done()
				defer func() { <-sem }() // Release the token when done
				voucherIDInt, _ := strconv.ParseInt(voucher.VoucherID, 10, 64)
				resp, _ := c.hermesSvc.TrackingHermerVoucherStatus(nil, courier4u, redcourier, voucherIDInt)
				if resp.Error == true {
					voucher.UpdateVoucherError(resp.Message)
					errorCh <- voucher
					return
				}
				voucher.UpdateHermerVoucherTracking(resp)
				voucherCh <- voucher

				mu.Lock()
				totalFetched++
				mu.Unlock()

			}(voucher)

		}

		// Wait for all fetch goroutines to finish
		fetchWg.Wait()
		close(voucherCh) // Close product channel after all fetching is done
	}(courier4u, redcourier)

	// Wait for all worker goroutines to finish processing products
	wg.Wait()

	close(errorCh) // Close error channel after all errors are processed

	return nil
}

func (c *CourierTrackingCron) saveVoucherTrackingResult(vch *cr.Voucher, projectID string, client *woo.Client) error {

	_, err := c.voucherSvc.UpdateVoucherTracking(nil, vch, projectID, client)
	if err != nil {
		return errors.Wrap(err, "failed to insert product result into MongoDB")
	}

	return nil
}
