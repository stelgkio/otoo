package handler

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/auth"
	error_messages "github.com/stelgkio/otoo/template/pages/error"
	logger "github.com/stelgkio/otoo/util"
	r "github.com/stelgkio/otoo/util"
)

type Handler struct{}

func NewHandler(e *echo.Echo, logger *logger.Logger) *Handler {
	h := &Handler{}

	apiRoutes(e, h)
	authRoutes(e, h)
	httpRoutes(e, h)
	e.HTTPErrorHandler = customHTTPErrorHandler
	return h
}

func httpRoutes(e *echo.Echo, h *Handler) {
	homegroup := e.Group("/")
	homegroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},

		SigningKey:   []byte(auth.GetJWTSecret()),
		TokenLookup:  "cookie:access-token",
		ErrorHandler: auth.JWTErrorChecker,
	}))

	//Attach jwt token refresher.

	homegroup.GET("", h.HandleIndex)
}

func apiRoutes(e *echo.Echo, h *Handler) {
	api := e.Group("api")
	api.GET("/hello", hello)
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func authRoutes(e *echo.Echo, h *Handler) {

	e.GET("/login", h.HandleLogin).Name = "userSignInForm"
	e.POST("/submit-login", h.HandleSubmitLogin)
	e.GET("/account/login", h.HandleLogin)
	e.POST("/login/email", h.HandleLoginEmainValidation)
	e.GET("/register", h.HandleRegister)

}
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		c.Logger().Info(he.Code)
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	r.Render(c, error_messages.ErrorPage())
}
