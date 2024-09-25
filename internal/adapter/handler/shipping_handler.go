package handler

import (
	"github.com/stelgkio/otoo/internal/core/port"
)

// ShippingHandler handles shipping-related routes
type ShippingHandler struct {
	orderSvc port.OrderService
}

// NewShippingHandler initializes the shipping handler
func NewShippingHandler(orderSvc port.OrderService) *ShippingHandler {
	return &ShippingHandler{orderSvc: orderSvc}
}

// ShippingTable renders the shipping table for a given project and status
// func (h *ShippingHandler) ShippingTable(c echo.Context) error {
// 	projectId := c.Param("projectId")
// 	status := c.Param("status")

// 	// Fetch shipping data (e.g., shipping orders)
// 	shippingOrders, err := h.orderSvc.GetOrdersByStatus(c.Request().Context(), projectId, status)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Failed to fetch shipping orders")
// 	}

// 	// Prepare data for rendering the shipping table
// 	data := map[string]interface{}{
// 		"ShippingTable": views.ShippingTable(projectId),
// 		"Orders":        shippingOrders, // Pass the orders to the view
// 	}

// 	return util.Render(c, data)
// }

// ShippingBulkAction handles bulk actions on shipping orders
// func (h *ShippingHandler) ShippingBulkAction(c echo.Context) error {
// 	projectId := c.Param("projectId")

// 	var bulkRequest struct {
// 		Orders []string `json:"orders"`
// 		Action string   `json:"action"`
// 	}

// 	if err := c.Bind(&bulkRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, "Invalid request data")
// 	}

// 	// Process the bulk action on the shipping orders
// 	err := h.orderSvc.ProcessBulkShippingActions(c.Request().Context(), projectId, bulkRequest.Orders, bulkRequest.Action)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Failed to process bulk action")
// 	}

// 	return c.JSON(http.StatusOK, "Bulk action applied successfully")
// }
