package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	wp "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/progress/webhooks"
	"github.com/stelgkio/otoo/internal/core/domain"
	woo "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
)

// WooCommerceHandler represents the WooCommerce handler
type WooCommerceHandler struct {
	p  port.WoocommerceRepository
	s  port.ProjectRepository
	c  port.CustomerService
	pr port.ProductService
}

// NewWooCommerceHandler creates a new instance of WooCommerceHandler
// Injects repository, project repo, customer service, and product service
func NewWooCommerceHandler(repo port.WoocommerceRepository, projrepo port.ProjectRepository, ctm port.CustomerService, proj port.ProductService) *WooCommerceHandler {
	return &WooCommerceHandler{
		repo,
		projrepo,
		ctm,
		proj,
	}
}

// readAndResetBody reads the request body and resets it for further use
func readAndResetBody(ctx echo.Context) ([]byte, error) {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}
	// Reset the request body so it can be read again if needed
	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

// OrderCreatedWebHook handles the webhook when an order is created
// POST /webhook/order/create
func (w WooCommerceHandler) OrderCreatedWebHook(ctx echo.Context) error {
	// Read and reset request body
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("Error reading body order_created request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	// Bind the request data into WooCommerce Order struct
	req := new(woocommerce.Order)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding order_created request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	// Validate the webhook using project information
	project, err := w.validateWebhook(ctx, body, "order_created")
	if err != nil {
		slog.Error("Error validating order_created webhook", "error", err)
		return err
	}

	// Create an order record
	orderRecord := &woo.OrderRecord{
		ProjectID: project.Id.String(),
		Event:     "order.created",
		OrderID:   req.ID,
		Order:     *req,
		IsActive:  true,
		CreatedAt: time.Now().UTC(),
		Timestamp: time.Now().UTC(),
	}

	// Convert order status from string to domain status
	orderRecord.Status, err = woo.StringToOrderStatus(req.Status)
	if err != nil {
		slog.Error("Error converting order status", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	// Save the order to the database
	err = w.p.OrderCreate(orderRecord)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	// Extract customer and product data asynchronously
	go w.c.ExtractCustomerFromOrderAndUpsert(ctx, orderRecord)
	go w.pr.ExtractProductFromOrderAndUpsert(ctx, orderRecord, project)

	// Return success response
	return ctx.String(http.StatusCreated, "created")
}

// OrderUpdatesWebHook handles the webhook when an order is updated
// POST /webhook/order/update
func (w WooCommerceHandler) OrderUpdatesWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("Error reading body order_updated request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	req := new(woocommerce.Order)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding order_updated request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	project, err := w.validateWebhook(ctx, body, "order_updated")
	if err != nil {
		slog.Error("Error validating order_updated webhook", "error", err)
		return err
	}

	// Create an updated order record
	updateOrderRecord := &woo.OrderRecord{
		ProjectID: project.Id.String(),
		Event:     "order.updated",
		OrderID:   req.ID,
		Order:     *req,
		IsActive:  true,
		UpdatedAt: time.Now().UTC(),
		Timestamp: time.Now().UTC(),
	}

	// Convert the status
	updateOrderRecord.Status, err = woo.StringToOrderStatus(req.Status)
	if err != nil {
		slog.Error("Error converting order status", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	// Update the order in the database
	err = w.p.OrderUpdate(updateOrderRecord, req.ID)
	if err != nil {
		return err
	}

	// Extract and upsert customer and product data asynchronously
	go w.c.ExtractCustomerFromOrderAndUpsert(ctx, updateOrderRecord)
	go w.pr.ExtractProductFromOrderAndUpsert(ctx, updateOrderRecord, project)

	return nil
}

// OrderDeletedWebHook handles the webhook when an order is deleted
// POST /webhook/order/delete
func (w WooCommerceHandler) OrderDeletedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}

	// Delete the order
	err := w.p.OrderDelete(order)
	if err != nil {
		return err
	}
	return nil
}

// CouponCreatedWebHook handles coupon creation webhook
// POST /webhook/coupon/create
func (w WooCommerceHandler) CouponCreatedWebHook(ctx echo.Context) error {
	var coupon bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&coupon); err != nil {
		return err
	}

	// Create coupon in the system
	err := w.p.CouponCreate(coupon)
	if err != nil {
		return err
	}
	return nil
}

// CouponUpdatedWebHook handles coupon updates webhook
// POST /webhook/coupon/update
func (w WooCommerceHandler) CouponUpdatedWebHook(ctx echo.Context) error {
	var coupon bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&coupon); err != nil {
		return err
	}

	// Update the coupon
	err := w.p.CouponUpdate(coupon)
	if err != nil {
		return err
	}
	return nil
}

// CouponDeletedWebHook handles coupon deletion webhook
// POST /webhook/coupon/delete
func (w WooCommerceHandler) CouponDeletedWebHook(ctx echo.Context) error {
	var coupon bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&coupon); err != nil {
		return err
	}

	// Delete the coupon
	err := w.p.CouponDelete(coupon)
	if err != nil {
		return err
	}
	return nil
}

// CustomerCreatedWebHook handles customer creation webhook
// POST /webhook/customer/create
func (w WooCommerceHandler) CustomerCreatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("Error reading body customer_created request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	req := new(woocommerce.Customer)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding customer_created request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	project, err := w.validateWebhook(ctx, body, "customer_created")
	if err != nil {
		slog.Error("Error validating customer_created webhook", "error", err)
		return err
	}

	// Create a customer record
	customerRecord := &woo.CustomerRecord{
		ProjectID:  project.Id.String(),
		Event:      "customer.created",
		CustomerID: req.ID,
		Email:      req.Email,
		Customer:   *req,
		IsActive:   true,
		CreatedAt:  time.Now().UTC(),
		Timestamp:  time.Now().UTC(),
	}

	// Save customer data to the database
	err = w.p.CustomerCreate(customerRecord, req.Email)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	// Return success response
	return ctx.String(http.StatusCreated, "created")
}

// CustomerUpdatedWebHook handles customer updates webhook
// POST /webhook/customer/update
func (w WooCommerceHandler) CustomerUpdatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("Error reading body customer_updated request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	req := new(woocommerce.Customer)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding customer_updated request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	project, err := w.validateWebhook(ctx, body, "customer_updated")
	if err != nil {
		slog.Error("Error validating customer_updated webhook", "error", err)
		return err
	}

	// Create updated customer record
	customerRecord := &woo.CustomerRecord{
		ProjectID:  project.Id.String(),
		Event:      "customer.updated",
		CustomerID: req.ID,
		Email:      req.Email,
		Customer:   *req,
		IsActive:   true,
		UpdatedAt:  time.Now().UTC(),
		Timestamp:  time.Now().UTC(),
	}

	// Update customer data in the database
	err = w.p.CustomerUpdate(customerRecord, req.Email)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	// Return success response
	return ctx.String(http.StatusCreated, "created")
}

// CustomerDeletedWebHook handles customer deletion webhook
// POST /webhook/customer/delete
func (w WooCommerceHandler) CustomerDeletedWebHook(ctx echo.Context) error {
	var customer bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&customer); err != nil {
		return err
	}

	// Delete customer data from the database
	err := w.p.CustomerDelete(customer)
	if err != nil {
		return err
	}
	return nil
}

// ProductCreatedWebHook handles product creation webhook
// POST /webhook/product/create
func (w WooCommerceHandler) ProductCreatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("Error reading body product_created request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	req := new(woocommerce.Product)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding product_created request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if req.ID == 0 {
		return ctx.String(http.StatusOK, "bad request")
	}

	project, err := w.validateWebhook(ctx, body, "product_created")
	if err != nil {
		slog.Error("Error validating product_created webhook", "error", err)
		return err
	}

	// Create a product record
	productRecord := &woo.ProductRecord{
		ProjectID: project.Id.String(),
		Event:     "product.created",
		ProductID: req.ID,
		Product:   *req,
		IsActive:  true,
		CreatedAt: time.Now().UTC(),
		Timestamp: time.Now().UTC(),
		ParentId:  req.ParentId,
	}

	// Save product data to the database
	err = w.p.ProductUpdate(productRecord, req.ID, project.Id.String())
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	return ctx.String(http.StatusCreated, "created")
}

// ProductUpdatedWebHook handles product updates webhook
// POST /webhook/product/update
func (w WooCommerceHandler) ProductUpdatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("Error reading body product_updated request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	req := new(woocommerce.Product)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding product_updated request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if req.ID == 0 {
		return ctx.String(http.StatusOK, "bad request")
	}

	project, err := w.validateWebhook(ctx, body, "product_updated")
	if err != nil {
		slog.Error("Error validating product_updated webhook", "error", err)
		return err
	}

	// Create an updated product record
	productRecord := &woo.ProductRecord{
		ProjectID: project.Id.String(),
		Event:     "product.updated",
		ProductID: req.ID,
		Product:   *req,
		IsActive:  true,
		UpdatedAt: time.Now().UTC(),
		Timestamp: time.Now().UTC(),
		ParentId:  req.ParentId,
	}

	// Update product data in the database
	err = w.p.ProductUpdate(productRecord, req.ID, project.Id.String())
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	return ctx.String(http.StatusCreated, "created")
}

// ProductDeletedWebHook handles product deletion webhook
// POST /webhook/product/delete
func (w WooCommerceHandler) ProductDeletedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("Error reading body product_deleted request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	req := new(woocommerce.Product)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Error binding product_deleted request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if req.ID == 0 {
		return ctx.String(http.StatusOK, "bad request")
	}

	// Validate the webhook
	_, err = w.validateWebhook(ctx, body, "product_deleted")
	if err != nil {
		slog.Error("Error validating product_deleted webhook", "error", err)
		return err
	}

	// Delete product from the database
	err = w.p.ProductDelete(req.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	return ctx.String(http.StatusCreated, "deleted")
}

// WebHooksProgressPage renders the webhook progress page
// GET  /progress/:projectId
func (w WooCommerceHandler) WebHooksProgressPage(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	// Fetch webhooks associated with the project
	webhooks, err := w.p.WebhookFindByProjectID(projectID)
	if err != nil {
		return err
	}

	// If 12 webhooks are found, set the response trigger
	if len(webhooks) == 12 {
		ctx.Response().Header().Set("HX-Trigger", "done")
	}

	// Render the progress page
	return util.Render(ctx, wp.WebHooksProgress(projectID, webhooks))
}

// WebHooksProgressPageDone renders the completed progress page
// GET /progress/done/:projectId
func (w WooCommerceHandler) WebHooksProgressPageDone(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	// Fetch webhooks associated with the project
	webhooks, err := w.p.WebhookFindByProjectID(projectID)
	if err != nil {
		return err
	}

	// If 12 webhooks are found, set the response trigger
	if len(webhooks) == 12 {
		ctx.Response().Header().Set("HX-Trigger", "done")
		return util.Render(ctx, wp.WebhooksProgressDone(projectID, webhooks, util.AllErrorsEmpty(webhooks)))
	}

	return util.Render(ctx, wp.WebHooksProgress(projectID, webhooks))
}

// validateWebhook validates the webhook signature and retrieves the associated project
func (w WooCommerceHandler) validateWebhook(ctx echo.Context, body []byte, event string) (*domain.Project, error) {
	domain := ctx.Get("webhookSource").(string)

	// Fetch the project by its domain
	project, err := w.s.GetProjectByDomain(ctx, domain)
	if err != nil {
		slog.Error(fmt.Sprintf("Error fetching project by domain for event: %s", event), "error", err)
		return nil, ctx.String(http.StatusBadRequest, "bad request")
	}

	// Validate the webhook signature
	err = util.ValidateWebhookSignature(ctx, project.Id.String(), body)
	if err != nil {
		slog.Error(fmt.Sprintf("Invalid webhook signature for event: %s", event), "error", err)
		return nil, ctx.String(http.StatusUnauthorized, "unauthorized")
	}

	return project, nil
}
