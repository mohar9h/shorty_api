package migrations

import (
	"fmt"
	"shorty_api/internal/database"
	"shorty_api/internal/model"

	"log"
)

func RunMigrations() {
	db := database.GetDB()
	if db == nil {
		log.Fatal("❌ DB not initialized. Did you call database.InitPostgres()?")
	}

	models := []any{
		&model.Link{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("✅ Migrations completed successfully!")
}
