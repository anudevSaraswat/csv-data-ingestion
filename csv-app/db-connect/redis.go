package dbconnect

import (
	"context"
	"log"
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/redis/go-redis/v9"
)

var cDB = cacheDB{}

type cacheDB struct {
	DB       *redis.Client
	SearchDB *redisearch.Client
}

func (cdb cacheDB) IsInitialised() bool {
	return cdb.DB != nil && cdb.SearchDB != nil
}

// this function provides a connection handle to redis cache service
func ConnectToCache() cacheDB {

	if cDB.IsInitialised() {
		return cDB
	}

	cDB.DB = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     os.Getenv("CACHE_ADDR"),
		Password: "",
		DB:       0,
	})

	status := cDB.DB.Ping(context.TODO())
	if status.Val() != "PONG" {
		log.Default().Println("datastore is not running...")
	}

	cDB.SearchDB = redisearch.NewClient(os.Getenv("CACHE_ADDR"), "idx:users")

	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("user_id")).
		AddField(redisearch.NewTextField("name")).
		AddField(redisearch.NewTextField("email")).
		AddField(redisearch.NewTextField("dob")).
		AddField(redisearch.NewTextField("city"))

	if err := cDB.SearchDB.CreateIndex(schema); err != nil {
		panic(err)
	}

	return cDB

}
