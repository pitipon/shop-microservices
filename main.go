package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pitipon/shop-microservices/config"
	"github.com/pitipon/shop-microservices/pkg/database"
	"github.com/pitipon/shop-microservices/server"
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

	// Database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, &cfg, db)
}
