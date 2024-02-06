package main

import (
	server "github.com/stelgkio/otoo/cmd/prepare"
	//configs "github.com/stelgkio/otoo/config"
	logger "github.com/stelgkio/otoo/util"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()

	// cfg, err := configs.ReadConfigs("./config/dev.yml")
	// if err != nil {
	// 	logger.Fatal("failed to load config", err)
	// }

	app := server.NewServer(logger)

	logger.Fatal("failed to start server", app.Start(":8081"))
}
