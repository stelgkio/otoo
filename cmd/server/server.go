package server

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/stelgkio/otoo/internal/adapter/config"
	"github.com/stelgkio/otoo/internal/adapter/handler"
	"github.com/stelgkio/otoo/internal/adapter/storage/postgres/repository"
	"github.com/stelgkio/otoo/internal/core/service"

	logger "github.com/stelgkio/otoo/internal/core/util"
)

func NewServer(db *pg.DB, logger *logger.Logger, config *config.Container) *echo.Echo {

	s := StartServer()

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Auth
	authService := service.NewAuthService(userRepo, nil)
	authHandler := handler.NewAuthHandler(authService, userService)

	//Home
	homeHandler := handler.NewHomeHandler()

	_, err := NewRouter(s, userHandler, authHandler, homeHandler)
	if err != nil {
		return nil
	}

	return s
}

func StartServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	e.Static("/css", "css")
	e.Static("/assets", "assets")
	e.Static("/fonts", "fonts")

	return e
}
