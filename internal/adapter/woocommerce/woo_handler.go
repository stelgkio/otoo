package woocommerce

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
	woo "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
	"github.com/stelgkio/woocommerce"
	"go.mongodb.org/mongo-driver/bson"
)

type WooCommerceHandler struct {
	p port.WoocommerceRepository
	s port.ProjectRepository
}

func NewWooCommerceHandler(repo port.WoocommerceRepository ,projrepo port.ProjectRepository) *WooCommerceHandler {
	return &WooCommerceHandler{
		repo,
		projrepo,
	}
}
func readAndResetBody(ctx echo.Context) ([]byte, error) {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}
	// Print the request body
	//fmt.Println("Request Body:", string(body))
	// Reset the request body to its original state so it can be read again if needed
	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}
// Webhook Order Create
// POST /webhook/order/create
func (w WooCommerceHandler) OrderCreatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body order_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Order)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding order_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	domain := ctx.Get("webhookSource").(string)
	project ,err := w.s.GetProjectByDomain(ctx, domain)
	
	if err != nil {
		slog.Error("error GetProjectByDomain order_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	err = util.ValidateWebhookSignature(ctx, project.Id.String(),body)
	if err != nil {
		slog.Error("error invalid signature order_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	orderRecord := &woo.OrderRecord{
		ProjectID: project.Id.String(),
		Error: "",		
		Event:     "order.created",
		OrderID:   req.ID,
		Order: *req,
		IsActive:  true,
		CreatedAt: time.Now(),		
		Timestamp: time.Now(),


	}
	err = w.p.OrderCreate(orderRecord)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	
	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}

// Webhook Order Update
// POST /webhook/order/update
func (w WooCommerceHandler) OrderUpdatesWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body order_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Order)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding order_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	domain := ctx.Get("webhookSource").(string)
	project ,err := w.s.GetProjectByDomain(ctx, domain)
	if err != nil {
		slog.Error("error GetProjectByDomain order_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	err = util.ValidateWebhookSignature(ctx, project.Id.String(),body)
	if err != nil {
		slog.Error("error invalid signature order_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	updateOrderRecord := &woo.OrderRecord{
		ProjectID: project.Id.String(),
		Error: "",		
		Event:     "order.updated",
		OrderID:   req.ID,
		Order: *req,
		IsActive:  true,
		UpdatedAt: time.Now(),		
		Timestamp: time.Now(),


	}

	err = w.p.OrderUpdate(updateOrderRecord,req.ID)
	if err != nil {
		return err
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save the
	return nil
}

// Webhook Order Delete
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

// Webhook Coupon Create
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

// Webhook Coupon Update
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

// Webhook Coupon Delete
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

// Webhook Customer Create
// POST /webhook/customer/create
func (w WooCommerceHandler) CustomerCreatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body customer_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Customer)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding customer_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	domain := ctx.Get("webhookSource").(string)
	project ,err := w.s.GetProjectByDomain(ctx, domain)
	
	if err != nil {
		slog.Error("error GetProjectByDomain customer_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	err = util.ValidateWebhookSignature(ctx, project.Id.String(),body)
	if err != nil {
		slog.Error("error invalid signature customer_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	customerRecord := &woo.CustomerRecord{
		ProjectID: project.Id.String(),
		Error: "",		
		Event:     "customer.created",
		CustomerID:   req.ID,
		Customer: *req,
		IsActive:  true,
		CreatedAt: time.Now(),		
		Timestamp: time.Now(),


	}
	err = w.p.CustomerCreate(customerRecord)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}

// Webhook Customer Update
// POST /webhook/customer/update
func (w WooCommerceHandler) CustomerUpdatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body customer_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Customer)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding customer_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	domain := ctx.Get("webhookSource").(string)
	project ,err := w.s.GetProjectByDomain(ctx, domain)
	
	if err != nil {
		slog.Error("error GetProjectByDomain customer_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	err = util.ValidateWebhookSignature(ctx, project.Id.String(),body)
	if err != nil {
		slog.Error("error invalid signature customer_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	customerRecord := &woo.CustomerRecord{
		ProjectID: project.Id.String(),
		Error: "",		
		Event:     "customer.updated",
		CustomerID:   req.ID,
		Customer: *req,
		IsActive:  true,
		UpdatedAt: time.Now(),	
		Timestamp: time.Now(),


	}
	err = w.p.CustomerUpdate(customerRecord,req.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}

// Webhook Customer Delete
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

//  Webhook Product Create
// POST /webhook/product/create
func (w WooCommerceHandler) ProductCreatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body product_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Product)
	if err := ctx.Bind(req); err != nil {
		fmt.Println(req)
		slog.Error("error binding product_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	if req.ID == 0 {
		return ctx.String(http.StatusOK, "bad request")
	}
	domain := ctx.Get("webhookSource").(string)
	project ,err := w.s.GetProjectByDomain(ctx, domain)
	
	if err != nil {
		slog.Error("error GetProjectByDomain product_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	err = util.ValidateWebhookSignature(ctx, project.Id.String(),body)
	if err != nil {
		slog.Error("error invalid signature product_created request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	productRecord := &woo.ProductRecord{
		ProjectID: project.Id.String(),
		Error: "",		
		Event:     "product.created",
		ProductID:   req.ID,
		Product: *req,
		IsActive:  true,
		CreatedAt: time.Now(),		
		Timestamp: time.Now(),
		ParentId: req.ParentId,


	}
	err = w.p.ProductUpdate(productRecord,req.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}

//  Webhook Product Update
// POST /webhook/product/update
func (w WooCommerceHandler) ProductUpdatedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error("error reading body product_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	req := new(woocommerce.Product)
	if err := ctx.Bind(req); err != nil {
		slog.Error("error binding product_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	if req.ID == 0 {
		return ctx.String(http.StatusOK, "bad request")
	}
	domain := ctx.Get("webhookSource").(string)
	project ,err := w.s.GetProjectByDomain(ctx, domain)
	
	if err != nil {
		slog.Error("error GetProjectByDomain product_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	err = util.ValidateWebhookSignature(ctx, project.Id.String(),body)
	if err != nil {
		slog.Error("error invalid signature product_updated request:"+ ctx.Get("webhookTopic").(string), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	productRecord := &woo.ProductRecord{
		ProjectID: project.Id.String(),
		Error: "",		
		Event:     "product.updated",
		ProductID:   req.ID,
		Product: *req,
		IsActive:  true,
		UpdatedAt: time.Now(),		
		Timestamp: time.Now(),
		ParentId: req.ParentId,


	}
	err = w.p.ProductUpdate(productRecord,req.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	//TODO: extract cutomer from order and save them
	//TODO: extract product from order and save them
	return ctx.String(http.StatusCreated, "created")
}
//  Webhook Product Delete
// POST /webhook/product/delete
func (w WooCommerceHandler) ProductDeletedWebHook(ctx echo.Context) error {
	body, err := readAndResetBody(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("error reading body product_deleted request: %s" ,ctx.Get("webhookTopic").(string)), "error", err)
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
	domain := ctx.Get("webhookSource").(string)
	project ,err := w.s.GetProjectByDomain(ctx, domain)
	
	if err != nil {
		slog.Error( fmt.Sprintf("error GetProjectByDomain product_deleted request: %s", ctx.Get("webhookTopic").(string)), "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	err = util.ValidateWebhookSignature(ctx, project.Id.String(),body)
	if err != nil {
		slog.Error("error invalid signature product_deleted request", "error", err)
		return ctx.String(http.StatusBadRequest, "bad request")
	}
	
	
	err = w.p.ProductDelete(req.ID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	return ctx.String(http.StatusCreated, "deleted")
}

//Webhook UI Pages Endpoints
// GET  /webhook/:projectId
func (w WooCommerceHandler) FindWebHooks(ctx echo.Context) error {
	return nil
}

// GET  /progress/:projectId
func (w WooCommerceHandler) WebHooksProgressPage(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	webhooks, err := w.p.WebhookFindByProjectId(projectId)

	if err != nil {
		return err
	}
	if len(webhooks) == 12 {
		ctx.Response().Header().Set("HX-Trigger", "done")

		return util.Render(ctx, wp.WebhooksProgressDone(projectId, webhooks, util.AllErrorsEmpty(webhooks)))
	}

	return util.Render(ctx, wp.WebHooksProgress(projectId, webhooks))
}

// GET /progress/done/:projectId
func (w WooCommerceHandler) WebHooksProgressPageDone(ctx echo.Context) error {
	projectId := ctx.Param("projectId")
	webhooks, err := w.p.WebhookFindByProjectId(projectId)

	if err != nil {
		return err
	}
	if len(webhooks) == 12 {
		ctx.Response().Header().Set("HX-Trigger", "done")
		return util.Render(ctx, wp.WebhooksProgressDone(projectId, webhooks, util.AllErrorsEmpty(webhooks)))
	}

	return util.Render(ctx, wp.WebHooksProgress(projectId, webhooks))
}
