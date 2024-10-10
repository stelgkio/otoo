package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	m "github.com/stelgkio/otoo/internal/adapter/web/view/component/courier/modal"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/component/courier/overview"
	v "github.com/stelgkio/otoo/internal/core/domain/courier"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/util"
)

// CourierTable returns the order dashboard
func (dh *DashboardHandler) CourierTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	return util.Render(ctx, t.VoucherOverview(projectID))
}

// VoucherTableHTML returns the order dashboard
func (dh *DashboardHandler) VoucherTableHTML(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	return util.Render(ctx, t.VoucherHtml(projectID))
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

	// Convert orderRecords to OrderTableList for the response
	var vouchers []v.VoucherTableList
	if voucherRecords != nil {
		for _, record := range voucherRecords {
			vouchers = append(vouchers, v.VoucherTableList{
				ID:        record.ID,
				ProjectID: record.ProjectID,
				OrderID:   record.OrderID,
				VoucherID: record.VoucherID,
				Status:    record.Status,
				Shipping:  *record.Shipping,
				Billing:   *record.Billing,
				UpdatedAt: record.UpdatedAt,
				Cod:       record.Cod,
				Products:  record.Products,
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

	order, err := dh.orderSvc.GetOrderByID(voucher.ProjectID, voucher.OrderID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return util.Render(ctx, m.VoucherDetails(voucher.ProjectID, voucher, order))

}

// CreateVoucher returns the order dashboard
func (dh *DashboardHandler) CreateVoucher(ctx echo.Context) error {
	ID := ctx.Param("Id")
	voucher, err := dh.voucherSvc.GetVoucherByVoucherID(ctx, ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	_, err = dh.orderSvc.GetOrderByID(voucher.ProjectID, voucher.OrderID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return util.Render(ctx, m.CreateVoucher())

}
