package server

import (
	"net/http"
	"os"
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
	courierTrackingCron *cr.CourierTrackingCron,

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
	e.GET("/RunHermerTrackingCronJob/:key", func(c echo.Context) error {
		data := c.Param("key")
		exkey := os.Getenv("EXTENSION_KEY")
		if data != exkey {
			return c.JSON(http.StatusInternalServerError, "internal server error")
		}
		err := courierTrackingCron.RunCourier4uTrackingCron()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		err2 := courierTrackingCron.RunRedCourierTrackingCron()

		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, err2)
		}
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

		usergroup.Use(auth.TokenRefresherMiddleware)
		usergroup.Use(configureJWT())
		usergroup.GET("/list/:projectId", authHandler.UserList)
		usergroup.GET("/createMember/:projectId", authHandler.CreateMemberModal)
		usergroup.POST("/addmember/:projectId", authHandler.AddMember)
		usergroup.DELETE("/delete/:userId/:projectId", authHandler.RemoveMember)
		usergroup.POST("/check-email", authHandler.CheckEmail)
	}
	//Dashboard
	dashboardgroup := e.Group("/dashboard")
	{

		dashboardgroup.Use(auth.TokenRefresherMiddleware)
		dashboardgroup.Use(configureJWT())
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

		//Attach jwt token refresher.
		extensiongroup.Use(auth.TokenRefresherMiddleware)
		// Add authentication
		extensiongroup.Use(configureJWT())
		extensiongroup.GET("/:projectId", dashboardHandler.Extention)
		extensiongroup.GET("/:projectId/:extensionId/success", dashboardHandler.StripeSuccesRedirect)
		extensiongroup.GET("/:projectId/:extensionId/fail", dashboardHandler.StripeFailRedirect)

		extensiongroup.GET("/asc-courier/:projectId", dashboardHandler.AcsCourier)
		extensiongroup.POST("/asc-courier/:projectId", dashboardHandler.AcsCourierFormPost)
		extensiongroup.POST("/asc-courier/settings/:projectId", dashboardHandler.AcsCourierSettingsFormPost)
		extensiongroup.POST("/asc-courier/deactivate/:projectId", dashboardHandler.AcsCourierDeActivate)

		extensiongroup.GET("/courier4u/:projectId", dashboardHandler.Courier4u)
		extensiongroup.POST("/courier4u/:projectId", dashboardHandler.Courier4uFormPost)
		extensiongroup.POST("/courier4u/settings/:projectId", dashboardHandler.Courier4uSettingsFormPost)
		extensiongroup.POST("/courier4u/deactivate/:projectId", dashboardHandler.Courier4uDeActivate)

		extensiongroup.GET("/redcourier/:projectId", dashboardHandler.RedCourier)
		extensiongroup.POST("/redcourier/:projectId", dashboardHandler.RedCourierFormPost)
		extensiongroup.POST("/redcourier/settings/:projectId", dashboardHandler.RedCourierettingsFormPost)
		extensiongroup.POST("/redcourier/deactivate/:projectId", dashboardHandler.RedCourierDeActivate)

		extensiongroup.GET("/wallet-expences/:projectId", dashboardHandler.WalletExpenses)

		extensiongroup.GET("/data-synchronizer/:projectId", dashboardHandler.DataSynchronizer)

		extensiongroup.GET("/page/asc-courier/:projectId", dashboardHandler.CourierTable)
		extensiongroup.GET("/page/courier4u/:projectId", dashboardHandler.CourierTable)
		extensiongroup.GET("/page/redcourier/:projectId", dashboardHandler.CourierTable)
		extensiongroup.GET("/page/wallet-expences/:projectId", dashboardHandler.WalletExpensesPage)
		extensiongroup.GET("/page/data-synchronizer/:projectId", projectHandler.ProjectSynchronizePage)

		extensiongroup.GET("/project_extensions/:projectId", dashboardHandler.ProjectExtensionsList)
		extensiongroup.DELETE("/project_extension/:Id", dashboardHandler.DeleteProjectExtension)
		extensiongroup.POST("/project_extension", dashboardHandler.AddProjectExtension)

		extensiongroup.GET("/addextensionForm/:key", dashboardHandler.AddManualExtensionForm)
		extensiongroup.GET("/all", dashboardHandler.GetAllAvailableExtensios)
		extensiongroup.GET("/table", dashboardHandler.ExtensionTable)
	}
	//Payment group
	paymentgroup := e.Group("/payment")
	{
		paymentgroup.Use(auth.TokenRefresherMiddleware)
		paymentgroup.Use(configureJWT())
		paymentgroup.POST("", dashboardHandler.Payment)
		paymentgroup.GET("/table/:projectId/:page", dashboardHandler.PaymentTable)

	}
	//Project group
	projectgroup := e.Group("/project")
	{
		projectgroup.Use(auth.TokenRefresherMiddleware)
		projectgroup.Use(configureJWT())
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

		projectgroup.GET("/findbydomain", projectHandler.FindProjectByDomain)

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
			settingsroup.GET("/redcourier/:projectId", projectHandler.ProjectSettingsRedCourier)

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

		webhookgroup.Use(auth.TokenRefresherMiddleware)
		webhookgroup.Use(configureJWT())
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
		customergroupB.Use(auth.TokenRefresherMiddleware)
		customergroupB.Use(configureJWT())
		customergroupB.GET("/table/:projectId/:page", dashboardHandler.CustomerTable)

	}

	ordergroupB := e.Group("/order")
	{
		ordergroupB.Use(auth.TokenRefresherMiddleware)
		ordergroupB.Use(configureJWT())
		ordergroupB.GET("/table/:projectId/:status/:page", dashboardHandler.OrderTable)
		ordergroupB.GET("/chart/:projectId", dashboardHandler.OrderCharts)
		ordergroupB.GET("/monthy/chart/:projectId", dashboardHandler.OrderMonthlyCharts)
		ordergroupB.GET("/tablehtml/:projectId", dashboardHandler.OrderTableHTML)

		ordergroupB.POST("/bulk-action/:projectId", dashboardHandler.OrderBulkAction)
		ordergroupB.PUT("/update/:orderId/:projectId", dashboardHandler.OrderUpdate)
	}
	productgroupB := e.Group("/product")
	{

		productgroupB.Use(auth.TokenRefresherMiddleware)
		productgroupB.Use(configureJWT())
		productgroupB.GET("/table/:projectId/:page", dashboardHandler.ProductTable)
		// ordergroup.GET("/tablehtml/:projectId", dashboardHandler.ProductTableHTML)

	}

	//Profile
	profilegroup := e.Group("/profile")
	{

		profilegroup.Use(auth.TokenRefresherMiddleware)
		profilegroup.Use(configureJWT())
		profilegroup.GET("", profileHandler.Profile).Name = "profile"
		profilegroup.GET("/password", profileHandler.ProfilePassword)
		profilegroup.POST("/password/update", profileHandler.UpdatePassword)
		profilegroup.POST("/validation/currentpassword", profileHandler.ValidateCurrentPassword)
		profilegroup.POST("/validation/newpassword", profileHandler.ValidateNewPassword)
		profilegroup.POST("/user/update", profileHandler.ProfileUpdate)
		profilegroup.POST("/user/delete", profileHandler.ProfileDelete)
	}

	//Courier
	couriergroup := e.Group("/voucher")
	{

		couriergroup.Use(auth.TokenRefresherMiddleware)
		couriergroup.Use(configureJWT())
		couriergroup.GET("/table/view/:projectId", dashboardHandler.CourierTable)
		couriergroup.GET("/table/html/:projectId", dashboardHandler.VoucherTableHTML)
		couriergroup.GET("/table/:projectId/:status/:page", dashboardHandler.VoucherTable)
		couriergroup.GET("/modal/:Id", dashboardHandler.VoucherDetailModal)
		couriergroup.GET("/openOffcanvas/:Id", dashboardHandler.CreateVoucher)
		//CREATE Voucher
		couriergroup.POST("/courier4u/create/:projectId", dashboardHandler.CreateCourier4uVoucher)
		couriergroup.POST("/redcourier/create/:projectId", dashboardHandler.CreateRedCourierVoucher)
		couriergroup.POST("/acscourier/create/:projectId", dashboardHandler.CreateCourier4uVoucher)
		//DOWNLOAD Voucher
		couriergroup.GET("/courier4u/donwload/:voucherId/:projectId", dashboardHandler.DownloadCourier4uVoucher)
		couriergroup.GET("/redcourier/donwload/:voucherId/:projectId", dashboardHandler.DownloadRedCourierVoucher)
		couriergroup.GET("/acscourier/donwload/:voucherId/:projectId", dashboardHandler.DownloadCourier4uVoucher)

		//UPDATE Voucher
		couriergroup.PUT("/courier4u/update/:voucherId/:projectId", dashboardHandler.UpdateCourier4uVoucher)
		couriergroup.PUT("/redcourier/update/:voucherId/:projectId", dashboardHandler.UpdateRerCourierVoucher)

		//ownload-multiple Voucher
		couriergroup.POST("/courier4u/download-multiple/:projectId", dashboardHandler.Courier4uDownloadMmultipleVoucher)
		couriergroup.POST("/redcourier/download-multiple/:projectId", dashboardHandler.RedCourierDownloadMmultipleVoucher)

		couriergroup.GET("/validateOrderId/:orderId/:projectID", dashboardHandler.ValidateOrderID)
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
