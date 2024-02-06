package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	u "github.com/stelgkio/otoo/models"
	t "github.com/stelgkio/otoo/template/pages/account"
	r "github.com/stelgkio/otoo/util"
)

func (i Handler) HandleLogin(c echo.Context) error {
	return r.Render(c, t.Login())

}

func (i Handler) HandleRegister(c echo.Context) error {
	return r.Render(c, t.Register())

}
func (i Handler) HandleSubmitLogin(c echo.Context) error {

	u := new(u.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.String(http.StatusCreated, "success")

}

func (i Handler) HandleLoginEmainValidation(c echo.Context) error {
	return c.String(http.StatusCreated, "success")
}
