package server

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	h "github.com/stelgkio/otoo/internal/adapter/handler"
	v "github.com/stelgkio/otoo/internal/adapter/web/view"
	con "github.com/stelgkio/otoo/internal/adapter/web/view/component/contact"
	conf "github.com/stelgkio/otoo/internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form"
	w "github.com/stelgkio/otoo/internal/adapter/woocommerce"
	auth "github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	r "github.com/stelgkio/otoo/internal/core/util"
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
	dashboardHandler *h.DashboardHandler,
	profileHandler *h.ProfileHandler,
) (*Router, error) {

	e.GET("/index", func(c echo.Context) error {
		return r.Render(c, v.IndexTemplate())
	}).Name = "index"

	e.GET("/contact", func(c echo.Context) error {
		return r.Render(c, con.ContantComponent())
	})
	e.POST("/contact", homeHandler.ContactForm)

	e.GET("login", authHandler.LoginForm).Name = "SignInForm"
	e.POST("login", authHandler.Login)

	e.GET("forgotpassword", authHandler.ForgotPasswordForm)
	e.POST("forgotpassword", authHandler.ForgotPassword)

	e.GET("resetpassword/:token", authHandler.ResetPasswordForm)
	e.POST("resetpassword/:email", authHandler.ResetPassword)
	//e.GET("projectlist", homeHandler.ProjectList)
	e.GET("register", authHandler.RegisterForm)
	e.POST("register", authHandler.Register)
	e.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("index"))
	})

	//Dashboard
	dashboardgroup := e.Group("/dashboard")
	{
		dashboardgroup.Use(configureJWT())
		dashboardgroup.Use(auth.TokenRefresherMiddleware)

		dashboardgroup.GET("", projectHandler.GetProjectDashboardPage).Name = "dashboard"

		dashboardgroup.GET("/logout", authHandler.Logout)
		dashboardgroup.GET("/contact", func(c echo.Context) error {
			return r.Render(c, conf.ContactForm(false, false, nil, new(domain.ContactRequest)))
		})
		dashboardgroup.POST("/contact", homeHandler.DashboardContactForm)

	}
	//Project group
	projectgroup := e.Group("/project")
	{
		// Add authentication
		projectgroup.Use(configureJWT())
		//Attach jwt token refresher.
		projectgroup.Use(auth.TokenRefresherMiddleware)

		projectgroup.GET("/list", projectHandler.ProjectListPage)
		projectgroup.GET("/createform", projectHandler.ProjectCreateForm)
		projectgroup.POST("/create", projectHandler.CreateProject)
		projectgroup.POST("/validation/name", projectHandler.ProjectNameValidation)
		projectgroup.POST("/validation/domain", projectHandler.ProjectDomainValidation)
	}
	//woocommerce group
	woocommercegroup := e.Group("/woocommerce")
	{
		woocommercegroup.POST("/order/created", WooCommerceHandler.OrderCreatedWebHook)
		woocommercegroup.POST("/order/updated", WooCommerceHandler.OrderUpdatesWebHook)
		woocommercegroup.POST("/order/deleted", WooCommerceHandler.OrderDeletedWebHook)

		woocommercegroup.POST("/coupon/created", WooCommerceHandler.CouponCreatedWebHook)
		woocommercegroup.POST("/coupon/updated", WooCommerceHandler.CouponUpdatedWebHook)
		woocommercegroup.POST("/coupon/deleted", WooCommerceHandler.CouponDeletedWebHook)

		woocommercegroup.POST("/customer/created", WooCommerceHandler.CustomerCreatedWebHook)
		woocommercegroup.POST("/customer/updated", WooCommerceHandler.CustomerUpdatedWebHook)
		woocommercegroup.POST("/customer/deleted", WooCommerceHandler.CustomerDeletedWebHook)

		woocommercegroup.POST("/product/created", WooCommerceHandler.ProductCreatedWebHook)
		woocommercegroup.POST("/product/updated", WooCommerceHandler.ProductUpdatedWebHook)
		woocommercegroup.POST("/product/deleted", WooCommerceHandler.ProductDeletedWebHook)

	}

	//Profile
	profilegroup := e.Group("/profile")
	{
		profilegroup.Use(configureJWT())
		profilegroup.Use(auth.TokenRefresherMiddleware)
		profilegroup.GET("", profileHandler.Profile).Name = "profile"
		profilegroup.GET("/password", profileHandler.ProfilePassword)
		profilegroup.POST("/user/update", profileHandler.ProfileUpdate)
		profilegroup.POST("/user/delete", profileHandler.ProfileDelete)
	}
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
