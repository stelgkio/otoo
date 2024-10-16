package server

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	h "github.com/stelgkio/otoo/internal/adapter/handler"
	v "github.com/stelgkio/otoo/internal/adapter/web/view"
	con "github.com/stelgkio/otoo/internal/adapter/web/view/component/contact"
	conf "github.com/stelgkio/otoo/internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form"
	auth "github.com/stelgkio/otoo/internal/core/auth"
	cr "github.com/stelgkio/otoo/internal/core/cron_job"
	"github.com/stelgkio/otoo/internal/core/domain"
	r "github.com/stelgkio/otoo/internal/core/util"
)

// Router is the router for the application
type Router struct {
	*echo.Echo
}

// NewRouter creates a new router
func NewRouter(
	e *echo.Echo,
	userHandler *h.UserHandler,
	authHandler *h.AuthHandler,
	homeHandler *h.HomeHandler,
	projectHandler *h.ProjectHandler,
	WooCommerceHandler *h.WooCommerceHandler,
	dashboardHandler *h.DashboardHandler,
	profileHandler *h.ProfileHandler,
	orderAnalyticsCron *cr.OrderAnalyticsCron,
	productBestSellerCron *cr.ProductBestSellerCron,

) (*Router, error) {

	e.GET("/index", func(c echo.Context) error {
		return r.Render(c, v.IndexTemplate())
	}).Name = "index"

	e.GET("/contact", func(c echo.Context) error {
		return r.Render(c, con.ContantComponent())
	})

	e.POST("/webhook", dashboardHandler.PaymentEvent)

	e.GET("/RunOrderWeeklyBalanceJob", func(c echo.Context) error {
		err := orderAnalyticsCron.RunOrderWeeklyBalanceJob()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusAccepted, "OK")
	})
	e.GET("/RunOrderMonthlyCountJob", func(c echo.Context) error {
		err := orderAnalyticsCron.RunOrderMonthlyCountJob()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusAccepted, "OK")
	})

	e.GET("/RunAProductBestSellerDailyJob", func(c echo.Context) error {
		err := productBestSellerCron.RunAProductBestSellerDailyJob()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusAccepted, "OK")
	})

	e.GET("/RunCustomerBestBuyerJob", func(c echo.Context) error {
		return c.JSON(http.StatusAccepted, "OK")
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

	//user

	usergroup := e.Group("/user")
	{
		usergroup.Use(configureJWT())
		usergroup.Use(auth.TokenRefresherMiddleware)
		usergroup.GET("/list/:projectId", authHandler.UserList)
		usergroup.POST("/addmember/:projectId", authHandler.AddMember)
	}
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

		dashboardgroup.GET("/project/:projectId", dashboardHandler.DefaultDashboard)
		dashboardgroup.GET("/default/:projectId", dashboardHandler.DefaultDashboardOverView)

		customergroup := dashboardgroup.Group("/customer")
		{
			customergroup.GET("/:projectId", dashboardHandler.CustomerDashboard)
		}

		ordergroup := dashboardgroup.Group("/order")
		{
			ordergroup.GET("/:projectId", dashboardHandler.OrderDashboard)
			ordergroup.GET("/table/:projectId/:status", dashboardHandler.OrderTable)
		}

		productgroup := dashboardgroup.Group("/product")
		{
			productgroup.GET("/:projectId", dashboardHandler.ProductDashboard)
		}

		notificationgroup := dashboardgroup.Group("/notifiaction")
		{
			notificationgroup.GET("/:projectId", dashboardHandler.FindNotification)
			notificationgroup.DELETE("/delete/:projectId/:notifiactionId", dashboardHandler.DeleteNotification)
			notificationgroup.DELETE("/settings/delete/:projectId/:notifiactionId", dashboardHandler.DeleteNotificationSettings)
			notificationgroup.DELETE("/delete/all/:projectId", dashboardHandler.DeleteAllNotification)

		}

	}
	//Extension group
	extensiongroup := e.Group("/extension")
	{
		// Add authentication
		extensiongroup.Use(configureJWT())
		//Attach jwt token refresher.
		extensiongroup.Use(auth.TokenRefresherMiddleware)
		extensiongroup.GET("/:projectId", dashboardHandler.Extention)
		extensiongroup.GET("/:projectId/:extensionId/success", dashboardHandler.StripeSuccesRedirect)
		extensiongroup.GET("/:projectId/:extensionId/fail", dashboardHandler.StripeFailRedirect)

		extensiongroup.GET("/asc-courier/:projectId", dashboardHandler.AcsCourier)
		extensiongroup.POST("/asc-courier/:projectId", dashboardHandler.AcsCourierFormPost)
		extensiongroup.POST("/asc-courier/deactivate/:projectId", dashboardHandler.AcsCourierDeActivate)

		extensiongroup.GET("/courier4u/:projectId", dashboardHandler.Courier4u)
		extensiongroup.POST("/courier4u/:projectId", dashboardHandler.Courier4uFormPost)
		extensiongroup.POST("/courier4u/deactivate/:projectId", dashboardHandler.Courier4uDeActivate)

		extensiongroup.GET("/wallet-expences/:projectId", dashboardHandler.WalletExpenses)

		extensiongroup.GET("/data-synchronizer/:projectId", dashboardHandler.DataSynchronizer)

		extensiongroup.GET("/page/asc-courier/:projectId", dashboardHandler.CourierTable)
		extensiongroup.GET("/page/courier4u/:projectId", dashboardHandler.CourierTable)
		extensiongroup.GET("/page/redcourier/:projectId", dashboardHandler.CourierTable)
		extensiongroup.GET("/page/wallet-expences/:projectId", dashboardHandler.WalletExpensesPage)
		extensiongroup.GET("/page/data-synchronizer/:projectId", projectHandler.ProjectSynchronizePage)

		extensiongroup.GET("/project_extensions/:projectId", dashboardHandler.ProjectExtensionsList)
	}

	//Payment group
	paymentgroup := e.Group("/payment")
	{
		// Add authentication
		paymentgroup.Use(configureJWT())
		//Attach jwt token refresher.
		paymentgroup.Use(auth.TokenRefresherMiddleware)
		paymentgroup.POST("", dashboardHandler.Payment)
		paymentgroup.GET("/table/:projectId/:page", dashboardHandler.PaymentTable)

	}
	//Project group
	projectgroup := e.Group("/project")
	{
		// Add authentication
		projectgroup.Use(configureJWT())
		//Attach jwt token refresher.
		projectgroup.Use(auth.TokenRefresherMiddleware)

		projectgroup.GET("/list", projectHandler.ProjectListPage)
		projectgroup.GET("/test/synchronize/:projectId", projectHandler.ProjectSynchronizeTest)
		projectgroup.GET("/synchronize/:projectId", projectHandler.ProjectSynchronize)
		projectgroup.POST("/synchronize/start/:projectId/:customerTotal/:productTotal/:orderTotal", projectHandler.ProjectSynchronizeStart)
		projectgroup.POST("/page/synchronize/start/:projectId/:customerTotal/:productTotal/:orderTotal", projectHandler.ProjectSynchronizeStartPage)
		projectgroup.GET("/synchronize/done/:projectId/:customerTotal/:productTotal/:orderTotal", projectHandler.ProjectSynchronizeDone)
		projectgroup.GET("/page/synchronize/done/:projectId/:customerTotal/:productTotal/:orderTotal", projectHandler.ProjectSynchronizeDonePage)
		projectgroup.GET("/createform", projectHandler.ProjectCreateForm)
		projectgroup.POST("/create", projectHandler.CreateProject)
		projectgroup.POST("/validation/name", projectHandler.ProjectNameValidation)
		projectgroup.POST("/validation/domain", projectHandler.ProjectDomainValidation)
		projectgroup.POST("/validation/key", projectHandler.ProjectKeyValidation)

		projectgroup.GET("/webhooks/:projectId", projectHandler.CheckWebHooks)
		projectgroup.GET("/:projectId", projectHandler.CheckWebHooks)

		settingsroup := projectgroup.Group("/settings")
		{
			settingsroup.GET("/:projectId", projectHandler.ProjectSettings)
			settingsroup.GET("/secret/:projectId", projectHandler.ProjectSettingsSercrets)
			settingsroup.GET("/notification/:projectId", projectHandler.ProjectSettingsNotification)
			settingsroup.GET("/webhook/:projectId", projectHandler.ProjectSettingsWebHook)
			settingsroup.POST("/update/:projectId", projectHandler.ProjectUpdate)
			settingsroup.POST("/delete/:projectId", projectHandler.ProjectDelete)
			settingsroup.POST("/secrets/update/:projectId", projectHandler.ProjectSecretsUpdate)
			settingsroup.DELETE("/webhook/delete/:projectId", WooCommerceHandler.DeleteAllWebhooks)
			settingsroup.GET("/webhook/createall/:projectId", WooCommerceHandler.WebhookCreateAll)
			settingsroup.GET("/team/:projectId", projectHandler.ProjectSettingsTeam)

			settingsroup.GET("/asc-courier/:projectId", projectHandler.ProjectSettingsAcsCourier)
			settingsroup.GET("/courier4u/:projectId", projectHandler.ProjectSettingsCourier4u)

			settingsroup.GET("/payment/:projectId", projectHandler.ProjectSettingsPayment)
		}
	}
	//woocommerce group
	woocommercegroup := e.Group("/woocommerce")
	{
		woocommercegroup.POST("/order/created", r.ExtractWebhookHeaders(WooCommerceHandler.OrderCreatedWebHook))
		woocommercegroup.POST("/order/updated", r.ExtractWebhookHeaders(WooCommerceHandler.OrderUpdatesWebHook))
		woocommercegroup.POST("/order/deleted", r.ExtractWebhookHeaders(WooCommerceHandler.OrderDeletedWebHook))

		woocommercegroup.POST("/coupon/created", r.ExtractWebhookHeaders(WooCommerceHandler.CouponCreatedWebHook))
		woocommercegroup.POST("/coupon/updated", r.ExtractWebhookHeaders(WooCommerceHandler.CouponUpdatedWebHook))
		woocommercegroup.POST("/coupon/deleted", r.ExtractWebhookHeaders(WooCommerceHandler.CouponDeletedWebHook))

		woocommercegroup.POST("/customer/created", r.ExtractWebhookHeaders(WooCommerceHandler.CustomerCreatedWebHook))
		woocommercegroup.POST("/customer/updated", r.ExtractWebhookHeaders(WooCommerceHandler.CustomerUpdatedWebHook))
		woocommercegroup.POST("/customer/deleted", r.ExtractWebhookHeaders(WooCommerceHandler.CustomerDeletedWebHook))

		woocommercegroup.POST("/product/created", r.ExtractWebhookHeaders(WooCommerceHandler.ProductCreatedWebHook))
		woocommercegroup.POST("/product/updated", r.ExtractWebhookHeaders(WooCommerceHandler.ProductUpdatedWebHook))
		woocommercegroup.POST("/product/deleted", r.ExtractWebhookHeaders(WooCommerceHandler.ProductDeletedWebHook))

	}

	webhookgroup := e.Group("/webhook")
	{
		webhookgroup.Use(configureJWT())
		webhookgroup.Use(auth.TokenRefresherMiddleware)
		webhookgroup.GET("/:projectId/:webhookId", nil)
		webhookgroup.GET("/:projectId", WooCommerceHandler.FindWebHooks)
		webhookgroup.GET("/progress/:projectId", WooCommerceHandler.WebHooksProgressPage)
		webhookgroup.GET("/progress/done/:projectId", WooCommerceHandler.WebHooksProgressPageDone)
		webhookgroup.POST("/update/:projectId/:webhookId", nil)
		webhookgroup.POST("/delete/:projectId/:webhookId", nil)
		webhookgroup.POST("/create/:projectId/:webhookId", nil)
		webhookgroup.GET("/table/:projectId/:page", WooCommerceHandler.WebHookTable)
		webhookgroup.GET("/:projectId/:eventname", func(c echo.Context) error {
			time.Sleep(2 * time.Second)
			return c.NoContent(286)
		})
		webhookgroup.DELETE("/:projectId", WooCommerceHandler.DeleteAllWebhooks)
		webhookgroup.POST("/bulk-action/:projectId", WooCommerceHandler.WebhookBulkAction)
	}

	customergroupB := e.Group("/customer")
	{
		customergroupB.Use(configureJWT())
		customergroupB.Use(auth.TokenRefresherMiddleware)
		customergroupB.GET("/table/:projectId/:page", dashboardHandler.CustomerTable)

	}

	ordergroupB := e.Group("/order")
	{
		ordergroupB.Use(configureJWT())
		ordergroupB.Use(auth.TokenRefresherMiddleware)
		ordergroupB.GET("/table/:projectId/:status/:page", dashboardHandler.OrderTable)
		ordergroupB.GET("/chart/:projectId", dashboardHandler.OrderCharts)
		ordergroupB.GET("/monthy/chart/:projectId", dashboardHandler.OrderMonthlyCharts)
		ordergroupB.GET("/tablehtml/:projectId", dashboardHandler.OrderTableHTML)

		ordergroupB.POST("/bulk-action/:projectId", dashboardHandler.OrderBulkAction)
		ordergroupB.PUT("/update/:orderId/:projectId", dashboardHandler.OrderUpdate)
	}
	productgroupB := e.Group("/product")
	{
		productgroupB.Use(configureJWT())
		productgroupB.Use(auth.TokenRefresherMiddleware)
		productgroupB.GET("/table/:projectId/:page", dashboardHandler.ProductTable)
		// ordergroup.GET("/tablehtml/:projectId", dashboardHandler.ProductTableHTML)

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

	//Courier
	couriergroup := e.Group("/voucher")
	{
		couriergroup.Use(configureJWT())
		couriergroup.Use(auth.TokenRefresherMiddleware)
		couriergroup.GET("/table/view/:projectId", dashboardHandler.CourierTable)
		couriergroup.GET("/table/html/:projectId", dashboardHandler.VoucherTableHTML)
		couriergroup.GET("/table/:projectId/:status/:page", dashboardHandler.VoucherTable)
		couriergroup.GET("/modal/:Id", dashboardHandler.VoucherDetailModal)
		couriergroup.GET("/openOffcanvas/:Id", dashboardHandler.CreateVoucher)

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
