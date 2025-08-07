package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pitipon/shop-microservices/config"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		fmt.Println("Using config file:", os.Args[1])
		return os.Args[1]
	}())

	log.Println(cfg)
}
