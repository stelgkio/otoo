package handler

import (
	"github.com/labstack/echo/v4"
	t "github.com/stelgkio/otoo/internal/adapter/web/template/components/home"
	r "github.com/stelgkio/otoo/internal/core/util"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h HomeHandler) Home(c echo.Context) error {
	return r.Render(c, t.HomeComponent())
}
