package woocommerce

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	wp "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/progress/webhooks"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
	"go.mongodb.org/mongo-driver/bson"
)

type WooCommerceHandler struct {
	p port.WoocommerceRepository
}

func NewWooCommerceHandler(repo port.WoocommerceRepository) *WooCommerceHandler {
	return &WooCommerceHandler{
		repo,
	}
}

// Order
func (w WooCommerceHandler) OrderCreatedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.OrderCreate(order)
	if err != nil {
		return err
	}
	return nil
}
func (w WooCommerceHandler) OrderUpdatesWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.OrderUpdate(order)
	if err != nil {
		return err
	}
	return nil
}
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

// Coupon
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

// Customer
func (w WooCommerceHandler) CustomerCreatedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.CustomerCreate(order)
	if err != nil {
		return err
	}
	return nil
}

func (w WooCommerceHandler) CustomerUpdatedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.CustomerUpdate(order)
	if err != nil {
		return err
	}
	return nil
}

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

// Product
func (w WooCommerceHandler) ProductCreatedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.ProductCreate(order)
	if err != nil {
		return err
	}
	return nil
}

func (w WooCommerceHandler) ProductUpdatedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.ProductUpdate(order)
	if err != nil {
		return err
	}
	return nil
}

func (w WooCommerceHandler) ProductDeletedWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.ProductDelete(order)
	if err != nil {
		return err
	}
	return nil
}

//Webhook UI Pages Endpoints

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
