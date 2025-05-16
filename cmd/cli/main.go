package main

import (
	"fmt"
	"os"
	"shorty_api/internal/config"
	"shorty_api/internal/database"
	"shorty_api/internal/database/migrations"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "migrate":
		cfg := config.GetConfig()
		if err := database.InitPostgres(cfg); err != nil {
			fmt.Println("Failed to connect to Postgres:", err)
			os.Exit(1)
		}
		defer database.CloseDB()
		migrations.RunMigrations()
	default:
		fmt.Println("Unknown command:", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/cli/main.go migrate run               Run model-based migrations")
	fmt.Println("  go run cmd/cli/main.go migrate create <name>     Create new SQL migration files")
}
