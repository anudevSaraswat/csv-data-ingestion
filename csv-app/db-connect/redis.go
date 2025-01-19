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
		Addr:     os.Getenv("CACHE_ADDR"),
		Password: "",
		DB:       0,
	})

	status := client.Ping(context.TODO())
	if status.Val() != "PONG" {
		log.Default().Println("datastore is not running...")
	}

	_, err := client.FTCreate(
		context.TODO(),
		"idx:users",
		&redis.FTCreateOptions{
			OnJSON: true,
			Prefix: []interface{}{"user:"},
		},
		&redis.FieldSchema{
			FieldName: "$.user_id",
			As:        "user_id",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.name",
			As:        "name",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.email",
			As:        "email",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.dob",
			As:        "dob",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.city",
			As:        "city",
			FieldType: redis.SearchFieldTypeText,
		},
	).Result()
	if err != nil {
		panic(err)
	}

	return client

}
