package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	//configs "github.com/stelgkio/otoo/config"

	"github.com/stelgkio/otoo/handler"
	logger "github.com/stelgkio/otoo/util"
)

func NewServer(logger *logger.Logger) *echo.Echo {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())
	//e.Use(auth.TokenRefresherMiddleware)
	e.Static("/css", "css")
	e.Static("/assets", "assets")
	e.Static("/fonts", "fonts")
	// Routes

	handler.NewHandler(e, logger)
	return e
}
