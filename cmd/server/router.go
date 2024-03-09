package server

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	h "github.com/stelgkio/otoo/internal/adapter/handler"
	w "github.com/stelgkio/otoo/internal/adapter/woocommerce"
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
	WooCommerceHandler *w.WooCommerceHandler,
) (*Router, error) {

	e.GET("login", authHandler.LoginForm).Name = "SignInForm"
	e.POST("login", authHandler.Login)

	e.GET("register", authHandler.RegisterForm)
	e.POST("register", authHandler.Register)
	e.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("index"))
	})

	homegroup := e.Group("/")
	homegroup.Use(configureJWT())
	//Attach jwt token refresher.
	homegroup.Use(auth.TokenRefresherMiddleware)

	homegroup.GET("index", homeHandler.Home).Name = "index"

	//Proejct group
	projectgroup := e.Group("/project")
	// Add authentication
	projectgroup.Use(configureJWT())
	//Attach jwt token refresher.
	projectgroup.Use(auth.TokenRefresherMiddleware)
	projectgroup.POST("/create", projectHandler.CreateProject)

	//Proejct group
	woocommercegroup := e.Group("/woocommerce")
	woocommercegroup.POST("/create", WooCommerceHandler.OrderWebHook)

	return &Router{e}, nil
}

func configureJWT() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},

		SigningKey:   []byte(auth.GetJWTSecret()),
		TokenLookup:  "cookie:accesstoken",
		ErrorHandler: auth.JWTErrorChecker,
	})
}

// func customHTTPErrorHandler(err error, c echo.Context) {
// 	code := http.StatusInternalServerError
// 	if he, ok := err.(*echo.HTTPError); ok {
// 		c.Logger().Info(he.Code)
// 		code = he.Code
// 	}
// 	errorPage := fmt.Sprintf("%d.html", code)
// 	if err := c.File(errorPage); err != nil {
// 		c.Logger().Error(err)
// 	}

// }
