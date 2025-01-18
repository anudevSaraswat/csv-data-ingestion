package dbconnect

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

// this function provides a connection handle to redis cache service
func ConnectToCache() *redis.Client {

	if client != nil {
		return client
	}

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: "",
		DB:       0,
	})

	status := client.Ping(context.TODO())
	if status.Val() != "PONG" {
		log.Default().Println("datastore is not running...")
	}

	return client

}
