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
func NewWooCommerceHandler(repo port.WoocommerceRepository, projrepo port.ProjectRepository, ctm port.CustomerService, proj port.ProductService) *WooCommerceHandler {
	return &WooCommerceHandler{
		repo,
		projrepo,
		ctm,
		proj,
	}
}
func readAndResetBody(ctx echo.Context) ([]byte, error) {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}
	// Print the request body
	// fmt.Println("Request Body:", string(body))
	// Reset the request body to its original state so it can be read again if needed
	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

// OrderCreatedWebHook POST /webhook/order/create
func (w WooCommerceHandler) OrderCreatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body order_created request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Order)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding order_created request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	project, err := w.validateWebhook(ctx, body, "order_created")
	if err != nil {
		slog.Error("error validateWebhook order_created request", "error", err)
		return err
	}
	orderRecord := &woo.OrderRecord{
		ProjectID: project.Id.String(),
		Error:     "",
		Event:     "order.created",
		OrderID:   req.ID,
		Order:     *req,
		IsActive:  true,
		CreatedAt: time.Now().UTC(),
		Timestamp: time.Now().UTC(),
	}
	orderRecord.Status, err = woo.StringToOrderStatus(req.Status)
	if err != nil {
		slog.Error("error converting order status", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	err = w.p.OrderCreate(orderRecord)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	go w.c.ExtractCustomerFromOrderAndUpsert(ctx, orderRecord)
	go w.pr.ExtractProductFromOrderAndUpsert(ctx, orderRecord, project)
	return ctx.String(http.StatusCreated, "created")
}

// OrderUpdatesWebHook POST /webhook/order/update
func (w WooCommerceHandler) OrderUpdatesWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body order_updated request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Order)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding order_updated request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	project, err := w.validateWebhook(ctx, body, "order_updated")
	if err != nil {
		slog.Error("error validateWebhook order_updated request", "error", err)
		return err
	}

	updateOrderRecord := &woo.OrderRecord{
		ProjectID: project.Id.String(),
		Error:     "",
		Event:     "order.updated",
		OrderID:   req.ID,
		Order:     *req,
		IsActive:  true,
		UpdatedAt: time.Now().UTC(),
		Timestamp: time.Now().UTC(),
	}
	updateOrderRecord.Status, err = woo.StringToOrderStatus(req.Status)
	if err != nil {
		slog.Error("error converting order status", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	err = w.p.OrderUpdate(updateOrderRecord, req.ID)
	if err != nil {
		return err
	}

	go w.c.ExtractCustomerFromOrderAndUpsert(ctx, updateOrderRecord)
	go w.pr.ExtractProductFromOrderAndUpsert(ctx, updateOrderRecord, project)
	return nil
}

// OrderDeletedWebHook Order Delete
// POST /webhook/order/delete
func (w WooCommerceHandler) OrderDeletedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.OrderDelete(order)
	if err != nil {
		return err
	}
	return nil
}

// CouponCreatedWebHook Coupon Create
// POST /webhook/coupon/create
func (w WooCommerceHandler) CouponCreatedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.CouponCreate(order)
	if err != nil {
		return err
	}
	return nil
}

// CouponUpdatedWebHook Coupon Update
// POST /webhook/coupon/update
func (w WooCommerceHandler) CouponUpdatedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.CouponUpdate(order)
	if err != nil {
		return err
	}
	return nil
}

// CouponDeletedWebHook Coupon Delete
// POST /webhook/coupon/delete
func (w WooCommerceHandler) CouponDeletedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.CouponDelete(order)
	if err != nil {
		return err
	}
	return nil
}

// CustomerCreatedWebHook Customer Create
// POST /webhook/customer/create
func (w WooCommerceHandler) CustomerCreatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body customer_created request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Customer)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding customer_created request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	project, err := w.validateWebhook(ctx, body, "customer_created")
	if err != nil {
		slog.Error("error validateWebhook customer_created request", "error", err)
		return err
	}
	customerRecord := &woo.CustomerRecord{
		ProjectID:  project.Id.String(),
		Error:      "",
		Event:      "customer.created",
		CustomerID: req.ID,
		Email:      req.Email,
		Customer:   *req,
		IsActive:   true,
		CreatedAt:  time.Now().UTC(),
		Timestamp:  time.Now().UTC(),
	}
	err = w.p.CustomerCreate(customerRecord, req.Email)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}

// CustomerUpdatedWebHook Customer Update
// POST /webhook/customer/update
func (w WooCommerceHandler) CustomerUpdatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body customer_updated request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Customer)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding customer_updated request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	project, err := w.validateWebhook(ctx, body, "customer_updated")
	if err != nil {
		slog.Error("error validateWebhook customer_updated request", "error", err)
		return err
	}
	customerRecord := &woo.CustomerRecord{
		ProjectID:  project.Id.String(),
		Error:      "",
		Event:      "customer.updated",
		CustomerID: req.ID,
		Email:      req.Email,
		Customer:   *req,
		IsActive:   true,
		UpdatedAt:  time.Now().UTC(),
		Timestamp:  time.Now().UTC(),
	}
	err = w.p.CustomerUpdate(customerRecord, req.Email)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}

// CustomerDeletedWebHook Customer Delete
// POST /webhook/customer/delete
func (w WooCommerceHandler) CustomerDeletedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.CustomerDelete(order)
	if err != nil {
		return err
	}
	return nil
}

// ProductCreatedWebHook Product Create
//
// POST /webhook/product/create
func (w WooCommerceHandler) ProductCreatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body product_created request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Product)
	if err := ctx.Bind(req); err != nil {
		fmt.Println(req)
		slog.Error("error binding product_created request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	if req.ID == 0 {
		return ctx.String(http.StatusOK, "bad request")
	}
	project, err := w.validateWebhook(ctx, body, "product_created")
	if err != nil {
		slog.Error("error validateWebhook product_created request", "error", err)
		return err
	}
	productRecord := &woo.ProductRecord{
		ProjectID: project.Id.String(),
		Error:     "",
		Event:     "product.created",
		ProductID: req.ID,
		Product:   *req,
		IsActive:  true,
		CreatedAt: time.Now().UTC(),
		Timestamp: time.Now().UTC(),
		ParentId:  req.ParentId,
	}
	err = w.p.ProductUpdate(productRecord, req.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}

// ProductUpdatedWebHook Product Update
//
// POST /webhook/product/update
func (w WooCommerceHandler) ProductUpdatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body product_updated request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Product)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding product_updated request:"+ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	if req.ID == 0 {
		return ctx.String(http.StatusOK, "bad request")
	}
	project, err := w.validateWebhook(ctx, body, "product_updated")
	if err != nil {
		slog.Error("error validateWebhook product_updated request", "error", err)
		return err
	}

	productRecord := &woo.ProductRecord{
		ProjectID: project.Id.String(),
		Error:     "",
		Event:     "product.updated",
		ProductID: req.ID,
		Product:   *req,
		IsActive:  true,
		UpdatedAt: time.Now().UTC(),
		Timestamp: time.Now().UTC(),
		ParentId:  req.ParentId,
	}
	err = w.p.ProductUpdate(productRecord, req.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}

// ProductDeletedWebHook Product Delete
//
// POST /webhook/product/delete
func (w WooCommerceHandler) ProductDeletedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("error reading body product_deleted request: %s", ctx.Get("webhookTopic").(string)), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Product)
	if err := ctx.Bind(req); err != nil {
		slog.Error(fmt.Sprintf("error binding product_deleted request: %s", ctx.Get("webhookTopic").(string)), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	if req.ID == 0 {
		return ctx.String(http.StatusOK, "bad request")
	}
	_, err = w.validateWebhook(ctx, body, "product_deleted")
	if err != nil {
		slog.Error("error validateWebhook product_deleted request", "error", err)
		return err
	}

	err = w.p.ProductDelete(req.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	return ctx.String(http.StatusCreated, "deleted")
}

// FindWebHooks Webhook UI Pages Endpoints
// GET  /webhook/:projectId
func (w WooCommerceHandler) FindWebHooks(ctx echo.Context) error {
	return nil
}

// WebHooksProgressPage GET  /progress/:projectId
func (w WooCommerceHandler) WebHooksProgressPage(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	webhooks, err := w.p.WebhookFindByProjectID(projectID)

	if err != nil {
		return err
	}
	if len(webhooks) == 12 {
		ctx.Response().Header().Set("HX-Trigger", "done")

		//	return util.Render(ctx, wp.WebhooksProgressDone(projectID, webhooks, util.AllErrorsEmpty(webhooks)))
	}

	return util.Render(ctx, wp.WebHooksProgress(projectID, webhooks))
}

// WebHooksProgressPageDone GET /progress/done/:projectId
func (w WooCommerceHandler) WebHooksProgressPageDone(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	webhooks, err := w.p.WebhookFindByProjectID(projectID)

	if err != nil {
		return err
	}
	if len(webhooks) == 12 {
		ctx.Response().Header().Set("HX-Trigger", "done")
		return util.Render(ctx, wp.WebhooksProgressDone(projectID, webhooks, util.AllErrorsEmpty(webhooks)))
	}

	return util.Render(ctx, wp.WebHooksProgress(projectID, webhooks))
}

func (w WooCommerceHandler) validateWebhook(ctx echo.Context, body []byte, event string) (*domain.Project, error) {
	domain := ctx.Get("webhookSource").(string)
	project, err := w.s.GetProjectByDomain(ctx, domain)
	if err != nil {
		slog.Error(fmt.Sprintf("error GetProjectByDomain request: %s event: %s", ctx.Get("webhookTopic").(string), event), "error", err)
		return nil, ctx.String(http.StatusBadRequest, "bad request")
	}

	err = util.ValidateWebhookSignature(ctx, project.Id.String(), body)
	if err != nil {
		slog.Error(fmt.Sprintf("error invalid signature request: %s  event: %s", ctx.Get("webhookTopic").(string), event), "error", err)
		return nil, ctx.String(http.StatusUnauthorized, "unauthorized")
	}

	return project, nil
}
