package main

import (
	"fmt"
	"shorty_api/internal/config"
	"shorty_api/internal/database"
	"shorty_api/internal/server"
	"shorty_api/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewZapLogger(cfg)

	err := database.InitPostgres(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer database.CloseDB()

	r := server.SetupServer(cfg)
	address := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := r.Run(address); err != nil {
		logger.Fatal(logging.Internal, logging.Startup, err.Error(), nil)
	}
}
