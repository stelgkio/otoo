package main

import (
	"context"
	"embed"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	server "github.com/stelgkio/otoo/cmd/server"
	"github.com/stelgkio/otoo/internal/adapter/config"
	nosql "github.com/stelgkio/otoo/internal/adapter/storage/mongodb"
	database "github.com/stelgkio/otoo/internal/adapter/storage/postgres"
)

//go:embed assets/*
var assetsFS embed.FS

func main() {
	// Load environment variables

	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelDebug)
	logger.Info("Starting the application...")

	config, err := config.New()
	if err != nil {
		logger.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Connect to database
	db := pg.Connect(&pg.Options{
		User:     config.DB.User,
		Password: config.DB.Password,
		Addr:     config.DB.Connection,
		Database: config.DB.Name,
	})
	defer db.Close()

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	logger.Info("Successfully connected to the database :", config.DB.Connection, "")

	//
	err = database.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	logger.Info("Schema created", "", "")

	mongodb, err := nosql.MongoDbConnect(config.Mongo.ConnectionUrl)
	if err != nil {
		panic(err)
	}
	defer mongodb.Disconnect(ctx)

	// Init cache service
	// cache, err := redis.New(ctx, config.Redis)
	// if err != nil {
	// 	slog.Error("Error initializing cache connection", "error", err)
	// 	os.Exit(1)
	// }
	// defer cache.Close()

	// slog.Info("Successfully connected to the cache server")

	c := server.InitCronScheduler()
	defer c.Stop()

	app := server.NewServer(db, mongodb, logger, config)
	// Serve static files from embedded FS
	app.GET("/assets/*", echo.WrapHandler(http.FileServer(http.FS(assetsFS))))

	// app.Static("/css", "./css")
	// app.Static("/assets", "./assets")
	// app.Static("/fonts", "./fonts")
	logger.Error("failed to start server", "error", app.Start(":8081"))

}
