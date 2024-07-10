package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/stelgkio/otoo/internal/adapter/config"
	"github.com/stelgkio/otoo/internal/adapter/handler"
	mongorepo "github.com/stelgkio/otoo/internal/adapter/storage/mongodb/repository"
	"github.com/stelgkio/otoo/internal/adapter/storage/postgres/repository"

	"github.com/stelgkio/otoo/internal/adapter/woocommerce"
	"github.com/stelgkio/otoo/internal/core/service"
)

func NewServer(db *pg.DB, mongodb *mongo.Client, logger *slog.Logger, config *config.Container) *echo.Echo {

	s := StartServer(logger)

	// Dependency injection

	//smtp
	smtpService := service.NewSmtpService()
	// Contact
	contactRepo := mongorepo.NewContactRepository(mongodb)
	contactService := service.NewContactService(contactRepo, smtpService)
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Auth
	authService := service.NewAuthService(userRepo, nil)
	authHandler := handler.NewAuthHandler(authService, userService)

	//WooCommerce
	woocommerceRepo := mongorepo.NewWoocommerceRepository(mongodb)
	woocommerceService := woocommerce.NewWoocommerceService()
	WooCommerceHandler := woocommerce.NewWooCommerceHandler(woocommerceRepo)

	// Project
	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo, woocommerceService)
	projectHandler := handler.NewProjectHandler(projectService, userService)

	//Home
	homeHandler := handler.NewHomeHandler(projectService, contactService)

	dashboardHandler := handler.NewDashboardHandler(projectService, userService)

	//profile handler
	profileHandler := handler.NewProfileHandler(userService, projectService, authService)
	//Router
	_, err := NewRouter(s, userHandler, authHandler, homeHandler, projectHandler, WooCommerceHandler, dashboardHandler, profileHandler)
	if err != nil {
		return nil
	}

	return s
}

func StartServer(logger *slog.Logger) *echo.Echo {
	e := echo.New()

	//e.Use(middleware.Logger())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		LogLatency:  true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			latencySeconds := float64(v.Latency) / float64(time.Millisecond)
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.Float64("latencyMS", latencySeconds),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.Float64("latencyMS", latencySeconds),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	})))
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// Take required information from error and context and send it to a service like New Relic
		//		fmt.Println(c.Path(), c.QueryParams(), err.Error())

		// Call the default handler to return the HTTP response
		e.DefaultHTTPErrorHandler(err, c)
	}

	return e
}
