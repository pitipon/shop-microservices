package database

import (
	"context"
	"log"
	"time"

	"github.com/pitipon/shop-microservices/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatalf("Connect to mongodb errors: %s", err.Error())
	}

	// Ping to verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Ping to mongodb failed: ", err)
	}

	return client
}
