package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pitipon/shop-microservices/config"
	"github.com/pitipon/shop-microservices/pkg/database/migration"
)

func main() {
	ctx := context.Background()

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		fmt.Println("Using config file:", os.Args[1])
		return os.Args[1]
	}())

	switch cfg.App.Name {
	case "player":
		migration.PlayerMigrate(ctx, &cfg)
	case "auth":
		migration.AuthMigrate(ctx, &cfg)
	case "item":
		migration.ItemMigrate(ctx, &cfg)
	case "inventory":
		migration.InventoryMigrate(ctx, &cfg)
	case "payment":
		migration.InventoryMigrate(ctx, &cfg)
	}
}
