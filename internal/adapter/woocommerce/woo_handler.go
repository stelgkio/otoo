package woocommerce

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/port"

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

func (w WooCommerceHandler) OrderWebHook(ctx echo.Context) error {
	var order bson.M
	if err := json.NewDecoder(ctx.Request().Body).Decode(&order); err != nil {
		return err
	}
	//fmt.Println(order)
	err := w.p.InsertWoocommerceOrder(order)
	if err != nil {
		return err
	}
	return nil
}
