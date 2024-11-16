package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/component/courier/overview"
	tmpl "github.com/stelgkio/otoo/internal/adapter/web/view/component/courier/template"
	"github.com/stelgkio/otoo/internal/core/domain"
	courier_domain "github.com/stelgkio/otoo/internal/core/domain/courier"
	v "github.com/stelgkio/otoo/internal/core/domain/courier"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/util"
)

type PDFResponse struct {
	Filename string `json:"filename"` // Filename of the PDF
	Data     string `json:"base64"`   // Base64 encoded PDF data
}

// CourierTable returns the order dashboard
func (dh *DashboardHandler) CourierTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extensionCodes := []string{
		domain.AcsCode,
		domain.Courier4u,
	}
	yes := domain.ContainsExtensionCodes(projectExtensions, extensionCodes)

	if !yes {
		ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/extension/%s", projectID))
		return ctx.String(http.StatusOK, "Redirecting...")
	}
	extensions := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.AcsCode || e.Code == domain.Courier4u
	})
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return util.Render(ctx, t.VoucherOverview(projectID, extensions))
	}

	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return err
	}
	return util.Render(ctx, tmpl.CourierTemplate(user, project.Name, projectID, extensions))
}

// VoucherTableHTML returns the order dashboard
func (dh *DashboardHandler) VoucherTableHTML(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	if err != nil {
		return err
	}
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extensionCodes := []string{
		domain.AcsCode,
		domain.Courier4u,
	}
	yes := domain.ContainsExtensionCodes(projectExtensions, extensionCodes)
	if !yes {
		ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/extension/%s", projectID))
		return ctx.String(http.StatusOK, "Redirecting...")
	}
	extensions := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.AcsCode || e.Code == domain.Courier4u
	})

	if ctx.Request().Header.Get("HX-Request") == "true" {

		return util.Render(ctx, t.VoucherHtml(projectID, extensions))
	}
	return util.Render(ctx, tmpl.VoucherHtmlTemplate(user, project.Name, projectID, extensions))
}

// VoucherTable returns the order dashboard
func (dh *DashboardHandler) VoucherTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	page := ctx.Param("page")
	status, err := v.StringToVoucherStatus(ctx.Param("status"))
	if page == "" {
		page = "1"
	}
	pageNum, err := strconv.Atoi(page)
	sort := ctx.QueryParam("sort")
	direction := ctx.QueryParam("direction")
	itemsPerPage := 10
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Errorf("invalid page number: %v", err))
	}

	if sort == "" {
		sort = "orderId"
	}
	if direction == "" {
		direction = "asc"
	}
	if status == "" {
		status = v.VoucherStatusNew
	}

	var wg sync.WaitGroup
	// Fetch	var wg sync.WaitGroup
	wg.Add(2)

	orderCountChan := make(chan int64, 1)
	orderListChan := make(chan []*v.Voucher, 1)
	errChan := make(chan error, 1)
	errListChan := make(chan error, 1)

	go func() {
		defer wg.Done()
		dh.voucherSvc.GetVoucherCountAsync(projectID, status, orderCountChan, errChan)
	}()

	// Fetch  10 orders
	go func() {
		defer wg.Done()
		dh.voucherSvc.FindVoucherByProjectIDAsync(projectID, itemsPerPage, pageNum, sort, direction, status, orderListChan, errListChan)
	}()
	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(orderCountChan)
		close(orderListChan)
		close(errChan)
		close(errListChan)
	}()

	var totalItems int64
	var voucherRecords []*v.Voucher

	for err := range errChan {
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"failed to fetch order count": err.Error()})
		}
	}
	for errList := range errListChan {
		if errList != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error to fetch order": errList.Error()})
		}
	}

	for item := range orderCountChan {
		totalItems = item
	}
	for item := range orderListChan {
		voucherRecords = item
	}

	var totalAmount = "0.00"

	// Convert orderRecords to OrderTableList for the response
	var vouchers []v.VoucherTableList
	if voucherRecords != nil {
		for _, record := range voucherRecords {
			if record.PaymentMethod == "cod" {
				totalAmount = record.TotalAmount
			}

			vouchers = append(vouchers, v.VoucherTableList{
				ID:                   record.ID,
				ProjectID:            record.ProjectID,
				OrderID:              *record.OrderID,
				VoucherID:            record.VoucherID,
				Status:               record.Status,
				Shipping:             *record.Shipping,
				Billing:              *record.Billing,
				UpdatedAt:            record.UpdatedAt,
				Cod:                  record.Cod,
				Products:             record.Products,
				HasError:             record.HasError,
				Error:                record.Error,
				CourierProvider:      record.CourierProvider,
				TotalAmount:          totalAmount,
				PaymentMethod:        record.PaymentMethod,
				Note:                 record.Note,
				IsPrinted:            record.IsPrinted,
				AcsVoucherRequest:    record.AcsVoucherRequest,
				HermesVoucerRequest:  record.HermesVoucerRequest,
				TrackingStatus:       record.TrackingStatus,
				HermesTrackingStages: record.HermesTrackingStages,
			})
		}
	}

	// Prepare metadata
	totalPages := int(totalItems) / itemsPerPage
	if int(totalItems)%itemsPerPage > 0 {
		totalPages++
	}

	// Create response object
	response := v.VoucherTableResponde{
		Data: vouchers,
		Meta: w.Meta{
			TotalItems:   int(totalItems),
			CurrentPage:  pageNum,
			ItemsPerPage: itemsPerPage,
			TotalPages:   totalPages,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}

// VoucherDetailModal returns the order dashboard
func (dh *DashboardHandler) VoucherDetailModal(ctx echo.Context) error {
	ID := ctx.Param("Id")
	voucher, err := dh.voucherSvc.GetVoucherByVoucherID(ctx, ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	_, err = dh.orderSvc.GetOrderByID(voucher.ProjectID, *voucher.OrderID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return nil

}

// CreateVoucher returns the order dashboard
func (dh *DashboardHandler) CreateVoucher(ctx echo.Context) error {
	ID := ctx.Param("Id")
	voucher, err := dh.voucherSvc.GetVoucherByVoucherID(ctx, ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	_, err = dh.orderSvc.GetOrderByID(voucher.ProjectID, *voucher.OrderID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return nil

}

// CreateCourier4uVoucher create courier4u voucher
func (dh *DashboardHandler) CreateCourier4uVoucher(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	req := new(courier_domain.HermesVoucerRequest)

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	validationErrors := req.Validate()
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error()})
	}
	// Convert string to int64
	orderID, err := strconv.ParseInt(req.OrderID, 10, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error converting string to int64:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid order ID"})
	}

	voucher, err := dh.voucherSvc.GetVoucherByOrderIDAndProjectID(ctx, orderID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else if voucher == nil && req.CustomOrderID == true {
		voucher, err = dh.voucherSvc.CreateHermesVoucher(ctx, req, projectID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.Courier4u
	})
	if len(extension) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Enable Courier4u extension first"})
	}

	courier4u := new(domain.Courier4uExtension)

	courier4u, err = dh.extensionSvc.GetCourier4uProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if courier4u == nil {
		courier4u = new(domain.Courier4uExtension)
	}

	respVoucher, err := dh.hermesSvc.CreateVoucher(ctx, courier4u, nil, req, projectID)
	if err != nil {
		voucher.UpdateVoucherError(err.Error())
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if respVoucher.Error == true {
		voucher.UpdateVoucherError(respVoucher.Message)
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": respVoucher.Message})
	}

	voucher.SetVoucher(respVoucher.Voucher)
	voucher.UpdateVoucherError("")
	voucher.UpdateVoucherProvider(domain.Courier4u)
	voucher.UpdateVoucherHermes(req)
	voucher.UpdateVoucherStatus(courier_domain.VoucherStatusProcessing)
	voucher.SetCustomOrderID(&orderID, req.CustomOrderID)
	dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)

	return ctx.JSON(http.StatusOK, "voucher created")

}

// CreateRedCourierVoucher create and return the pdf
func (dh *DashboardHandler) CreateRedCourierVoucher(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	req := new(courier_domain.HermesVoucerRequest)

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	validationErrors := req.Validate()
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error()})
	}
	// Convert string to int64
	orderID, err := strconv.ParseInt(req.OrderID, 10, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error converting string to int64:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid order ID"})
	}

	voucher, err := dh.voucherSvc.GetVoucherByOrderIDAndProjectID(ctx, orderID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	} else if voucher != nil && req.CustomOrderID == true {
		voucher, err = dh.voucherSvc.CreateHermesVoucher(ctx, req, projectID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.RedCourier
	})
	if len(extension) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Enable Courier4u extension first"})
	}

	redCourier := new(domain.RedCourierExtension)

	redCourier, err = dh.extensionSvc.GetRedCourierProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if redCourier == nil {
		redCourier = new(domain.RedCourierExtension)
	}

	respVoucher, err := dh.hermesSvc.CreateVoucher(ctx, nil, redCourier, req, projectID)
	if err != nil {
		voucher.UpdateVoucherError(err.Error())
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redCourier)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if respVoucher.Error == true {
		voucher.UpdateVoucherError(respVoucher.Message)
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redCourier)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": respVoucher.Message})
	}

	voucher.SetVoucher(respVoucher.Voucher)

	voucher.UpdateVoucherError("")
	voucher.UpdateVoucherProvider(domain.Courier4u)
	voucher.UpdateVoucherHermes(req)
	voucher.UpdateVoucherStatus(courier_domain.VoucherStatusProcessing)
	voucher.SetCustomOrderID(&orderID, req.CustomOrderID)
	dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redCourier)

	return ctx.JSON(http.StatusOK, "voucher created")

}

// DownloadCourier4uVoucher returns the order dashboard
func (dh *DashboardHandler) DownloadCourier4uVoucher(ctx echo.Context) error {
	voucherID := ctx.Param("voucherId")

	projectID := ctx.Param("projectId")

	voucher, err := dh.voucherSvc.GetVoucherByVoucherID(ctx, voucherID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.Courier4u
	})
	if len(extension) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Enable Courier4u extension first"})
	}
	courier4u := new(domain.Courier4uExtension)

	courier4u, err = dh.extensionSvc.GetCourier4uProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if courier4u == nil {
		courier4u = new(domain.Courier4uExtension)
	}
	vID, err := strconv.ParseInt(voucher.VoucherID, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	pdfData, err := dh.hermesSvc.PrintVoucher(ctx, courier4u, nil, vID, projectID, courier4u.PrinterType)
	if err != nil {
		voucher.UpdateVoucherIsPrinted(false)
		voucher.UpdateVoucherError(err.Error())
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// Create the response with the PDF data and filename
	pdfResponse := PDFResponse{
		Filename: fmt.Sprintf("couerier4u_voucher_%s.pdf", voucher.VoucherID), // Set your filename here
		Data:     base64.StdEncoding.EncodeToString(pdfData),                  // Encode the PDF data to Base64
	}
	voucher.UpdateVoucherIsPrinted(true)
	dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)

	return ctx.JSON(http.StatusOK, pdfResponse)

}

// DownloadRedCourierVoucher returns the order dashboard
func (dh *DashboardHandler) DownloadRedCourierVoucher(ctx echo.Context) error {
	voucherID := ctx.Param("voucherId")

	projectID := ctx.Param("projectId")

	voucher, err := dh.voucherSvc.GetVoucherByVoucherID(ctx, voucherID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.Courier4u
	})
	if len(extension) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Enable Courier4u extension first"})
	}
	redcourier := new(domain.RedCourierExtension)

	redcourier, err = dh.extensionSvc.GetRedCourierProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if redcourier == nil {
		redcourier = new(domain.RedCourierExtension)
	}
	vID, err := strconv.ParseInt(voucher.VoucherID, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	pdfData, err := dh.hermesSvc.PrintVoucher(ctx, nil, redcourier, vID, projectID, redcourier.PrinterType)
	if err != nil {
		voucher.UpdateVoucherIsPrinted(false)
		voucher.UpdateVoucherError(err.Error())
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redcourier)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// Create the response with the PDF data and filename
	pdfResponse := PDFResponse{
		Filename: fmt.Sprintf("redcourier_voucher_%s.pdf", voucher.VoucherID), // Set your filename here
		Data:     base64.StdEncoding.EncodeToString(pdfData),                  // Encode the PDF data to Base64
	}
	voucher.UpdateVoucherIsPrinted(true)
	dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redcourier)

	return ctx.JSON(http.StatusOK, pdfResponse)

}

// UpdateCourier4uVoucher create and return the pdf
func (dh *DashboardHandler) UpdateCourier4uVoucher(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	voucherID := ctx.Param("voucherId")
	req := new(courier_domain.HermesVoucerRequest)

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	validationErrors := req.Validate()
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error()})
	}
	// Convert string to int64
	orderID, err := strconv.ParseInt(req.OrderID, 10, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error converting string to int64:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid order ID"})
	}

	voucher, err := dh.voucherSvc.GetVoucherByOrderIDAndProjectID(ctx, orderID, projectID)
	if err != nil || voucher == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.Courier4u
	})
	if len(extension) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Enable Courier4u extension first"})
	}

	courier4u := new(domain.Courier4uExtension)

	courier4u, err = dh.extensionSvc.GetCourier4uProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if courier4u == nil {
		courier4u = new(domain.Courier4uExtension)
	}

	vID, err := strconv.ParseInt(voucherID, 10, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error converting string to int64:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid voucher ID"})
	}
	updateVoucherReq := courier_domain.NewHermesVoucerUpdateRequest(vID, req)

	respVoucher, err := dh.hermesSvc.UpdateVoucher(ctx, courier4u, nil, updateVoucherReq, projectID)
	if err != nil {
		voucher.UpdateVoucherError(err.Error())
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if respVoucher.Error == true {
		voucher.UpdateVoucherError(respVoucher.Message)
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": respVoucher.Message})
	}

	voucher.UpdateVoucherError("")
	voucher.UpdateVoucherProvider(domain.Courier4u)
	voucher.UpdateVoucherHermes(req)
	voucher.UpdateVoucherStatus(courier_domain.VoucherStatusProcessing)
	voucher.SetCustomOrderID(&orderID, req.CustomOrderID)
	dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)

	return ctx.JSON(http.StatusOK, "voucher updated successfully")
}

// UpdateRerCourierVoucher create and return the pdf
func (dh *DashboardHandler) UpdateRerCourierVoucher(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	voucherID := ctx.Param("voucherId")
	req := new(courier_domain.HermesVoucerRequest)

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	validationErrors := req.Validate()
	if validationErrors != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": validationErrors.Error()})
	}
	// Convert string to int64
	orderID, err := strconv.ParseInt(req.OrderID, 10, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error converting string to int64:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid order ID"})
	}

	voucher, err := dh.voucherSvc.GetVoucherByOrderIDAndProjectID(ctx, orderID, projectID)
	if err != nil || voucher == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.RedCourier
	})
	if len(extension) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Enable Courier4u extension first"})
	}

	redCourier := new(domain.RedCourierExtension)

	redCourier, err = dh.extensionSvc.GetRedCourierProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if redCourier == nil {
		redCourier = new(domain.RedCourierExtension)
	}

	vID, err := strconv.ParseInt(voucherID, 10, 64)
	if err != nil {
		// Handle the error
		fmt.Println("Error converting string to int64:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid voucher ID"})
	}
	updateVoucherReq := courier_domain.NewHermesVoucerUpdateRequest(vID, req)

	respVoucher, err := dh.hermesSvc.UpdateVoucher(ctx, nil, redCourier, updateVoucherReq, projectID)
	if err != nil {
		voucher.UpdateVoucherError(err.Error())
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redCourier)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if respVoucher.Error == true {
		voucher.UpdateVoucherError(respVoucher.Message)
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redCourier)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": respVoucher.Message})
	}

	voucher.UpdateVoucherError("")
	voucher.UpdateVoucherProvider(domain.Courier4u)
	voucher.UpdateVoucherHermes(req)
	voucher.UpdateVoucherStatus(courier_domain.VoucherStatusProcessing)
	voucher.SetCustomOrderID(&orderID, req.CustomOrderID)
	dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redCourier)

	return ctx.JSON(http.StatusOK, "voucher updated successfully")
}

type RequestBody struct {
	VoucherIds []v.VoucherTableList `json:"voucherIds"`
}

// RedCourierDownloadMmultipleVoucher returns the order dashboard
func (dh *DashboardHandler) RedCourierDownloadMmultipleVoucher(ctx echo.Context) error {
	voucherID := ctx.Param("voucherId")

	var requestBody RequestBody
	if err := ctx.Bind(&requestBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request format"})
	}

	projectID := ctx.Param("projectId")

	voucher, err := dh.voucherSvc.GetVoucherByVoucherID(ctx, voucherID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.Courier4u
	})
	if len(extension) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Enable Courier4u extension first"})
	}
	redcourier := new(domain.RedCourierExtension)

	redcourier, err = dh.extensionSvc.GetRedCourierProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if redcourier == nil {
		redcourier = new(domain.RedCourierExtension)
	}
	var voucherIDs []int64
	for _, voucher := range requestBody.VoucherIds {
		// Convert the voucherId (which is a string) to int64
		voucherIDInt, err := strconv.ParseInt(voucher.VoucherID, 10, 64)
		if err != nil {
			// Handle the error if the conversion fails
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": fmt.Sprintf("invalid voucherId %s", voucher.VoucherID)})
		}
		voucherIDs = append(voucherIDs, voucherIDInt)
	}
	pdfData, err := dh.hermesSvc.PrintMultipleVoucher(ctx, nil, redcourier, voucherIDs, projectID, redcourier.PrinterType)
	if err != nil {
		voucher.UpdateVoucherIsPrinted(false)
		voucher.UpdateVoucherError(err.Error())
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redcourier)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// Create the response with the PDF data and filename
	pdfResponse := PDFResponse{
		Filename: fmt.Sprintf("redcourier_voucher_%s.pdf", voucher.VoucherID), // Set your filename here
		Data:     base64.StdEncoding.EncodeToString(pdfData),                  // Encode the PDF data to Base64
	}
	voucher.UpdateVoucherIsPrinted(true)
	dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, nil, redcourier)

	return ctx.JSON(http.StatusOK, pdfResponse)

}

// Courier4uDownloadMmultipleVoucher returns the order dashboard
func (dh *DashboardHandler) Courier4uDownloadMmultipleVoucher(ctx echo.Context) error {

	projectID := ctx.Param("projectId")
	var requestBody RequestBody
	if err := ctx.Bind(&requestBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request format"})
	}

	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension := util.Filter(projectExtensions, func(e *domain.ProjectExtension) bool {
		return e.Code == domain.Courier4u
	})
	if len(extension) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Enable Courier4u extension first"})
	}
	courier4u := new(domain.Courier4uExtension)

	courier4u, err = dh.extensionSvc.GetCourier4uProjectExtensionByID(ctx, extension[0].ExtensionID, projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if courier4u == nil {
		courier4u = new(domain.Courier4uExtension)
	}
	var voucherIDs []int64
	for _, voucher := range requestBody.VoucherIds {
		// Convert the voucherId (which is a string) to int64
		voucherIDInt, err := strconv.ParseInt(voucher.VoucherID, 10, 64)
		if err != nil {
			// Handle the error if the conversion fails
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": fmt.Sprintf("invalid voucherId %s", voucher.VoucherID)})
		}
		voucherIDs = append(voucherIDs, voucherIDInt)
	}
	pdfData, err := dh.hermesSvc.PrintMultipleVoucher(ctx, courier4u, nil, voucherIDs, projectID, courier4u.PrinterType)
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	for _, v := range requestBody.VoucherIds {
		voucher, err := dh.voucherSvc.GetVoucherByVoucherID(ctx, v.ID.Hex())

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		voucher.UpdateVoucherIsPrinted(true)
		dh.voucherSvc.UpdateVoucherNewDetails(ctx, voucher, projectID, courier4u, nil)
	}
	// Create the response with the PDF data and filename
	pdfResponse := PDFResponse{
		Filename: "courier4u_multi_voucher.pdf",              // Set your filename here
		Data:     base64.StdEncoding.EncodeToString(pdfData), // Encode the PDF data to Base64
	}

	return ctx.JSON(http.StatusOK, pdfResponse)

}
