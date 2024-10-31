package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/shelojara/collection-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	slog.Info("Migrating models...")
	for _, model := range model.Models {
		if err := db.AutoMigrate(model); err != nil {
			slog.Error("Failed to migrate model",
				slog.String("model", fmt.Sprintf("%T", model)),
				slog.Any("error", err),
			)

			os.Exit(1)
		}

		slog.Info("Migrated model", slog.String("model", fmt.Sprintf("%T", model)))
	}

	slog.Info("All models migrated")
}
