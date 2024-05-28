package server

import (
	"context"
	"log/slog"

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
	projectHandler := handler.NewProjectHandler(projectService)

	//Home
	homeHandler := handler.NewHomeHandler(projectService)

	dashboardHandler := handler.NewDashboardHandler(projectService, userService)

	//Router
	_, err := NewRouter(s, userHandler, authHandler, homeHandler, projectHandler, WooCommerceHandler, dashboardHandler)
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
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.Duration("latency", v.Latency),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.Duration("latency", v.Latency),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// Take required information from error and context and send it to a service like New Relic
		//		fmt.Println(c.Path(), c.QueryParams(), err.Error())

		// Call the default handler to return the HTTP response
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.Static("/css", "css")
	e.Static("/assets", "assets")
	e.Static("/fonts", "fonts")

	return e
}
