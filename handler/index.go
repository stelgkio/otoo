package handler

import (
	"github.com/labstack/echo/v4"
	t "github.com/stelgkio/otoo/template"
	r "github.com/stelgkio/otoo/util"
)

func (i Handler) HandleIndex(c echo.Context) error {
	return r.Render(c, t.Index())

}
