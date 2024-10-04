package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/component/courier/table"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/util"
)

// FindNotification find all notification by projectID
func (dh *DashboardHandler) CourierTable(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	return util.Render(ctx, t.VoucherTable(projectID))
}

// OrderTable returns the order dashboard
func (dh *DashboardHandler) OrderTable2(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	page := ctx.Param("page")
	status, err := w.StringToOrderStatus(ctx.Param("status"))
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

	var wg sync.WaitGroup
	// Fetch	var wg sync.WaitGroup
	wg.Add(2)

	orderCountChan := make(chan int64, 1)
	orderListChan := make(chan []*w.OrderRecord, 1)
	errChan := make(chan error, 1)
	errListChan := make(chan error, 1)

	go func() {
		defer wg.Done()
		dh.orderSvc.GetOrderCountAsync(ctx, projectID, status, "", orderCountChan, errChan)
	}()

	// Fetch  10 orders
	go func() {
		defer wg.Done()
		dh.orderSvc.FindOrderByProjectIDAsync(projectID, itemsPerPage, pageNum, status, sort, direction, orderListChan, errListChan)
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
	var orderRecords []*w.OrderRecord

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
		orderRecords = item
	}

	// Convert orderRecords to OrderTableList for the response
	var orders []w.OrderTableList
	if orderRecords != nil {
		for _, record := range orderRecords {
			orders = append(orders, w.OrderTableList{
				ID:          record.ID,
				ProjectID:   record.ProjectID,
				Timestamp:   record.Timestamp,
				OrderID:     record.OrderID,
				TotalAmount: record.Order.Total,
				Status:      record.Status,
				Billing:     *record.Order.Billing,
				Shipping:    *record.Order.Shipping,
			})
		}
	}

	// Prepare metadata
	totalPages := int(totalItems) / itemsPerPage
	if int(totalItems)%itemsPerPage > 0 {
		totalPages++
	}

	// Create response object
	response := w.OrderTableResponde{
		Data: orders,
		Meta: w.Meta{
			TotalItems:   int(totalItems),
			CurrentPage:  pageNum,
			ItemsPerPage: itemsPerPage,
			TotalPages:   totalPages,
		},
	}

	return ctx.JSON(http.StatusOK, response)
}
