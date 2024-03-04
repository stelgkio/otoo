package server

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	h "github.com/stelgkio/otoo/internal/adapter/handler"
	auth "github.com/stelgkio/otoo/internal/core/auth"
)

type Router struct {
	*echo.Echo
}

func NewRouter(
	e *echo.Echo,
	userHandler *h.UserHandler,
	authHandler *h.AuthHandler,
	homeHandler *h.HomeHandler,
	projectHandler *h.ProjectHandler,
) (*Router, error) {

	e.GET("login", authHandler.LoginForm).Name = "SignInForm"
	e.POST("login", authHandler.Login)

	e.GET("register", authHandler.RegisterForm)
	e.POST("register", authHandler.Register)

	//Proejct group
	projectgroup := e.Group("/project")
	projectgroup.POST("/create", projectHandler.CreateProject)

	e.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("index"))
	})

	homegroup := e.Group("/")
	homegroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},

		SigningKey:   []byte(auth.GetJWTSecret()),
		TokenLookup:  "cookie:accesstoken",
		ErrorHandler: auth.JWTErrorChecker,
	}))

	//Attach jwt token refresher.
	homegroup.Use(auth.TokenRefresherMiddleware)

	homegroup.GET("index", homeHandler.Home).Name = "index"

	//ße.HTTPErrorHandler = customHTTPErrorHandler
	e.Use()
	return &Router{e}, nil
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

}
