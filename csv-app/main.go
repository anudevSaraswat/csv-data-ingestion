package main

import (
	"csv-app/routes"
	"os"

	"csv-app/consumer"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	engine := routes.SetupRoutes()

	go func() {
		err := consumer.ReadAndStoreData()
		if err != nil {
			panic(err)
		}
	}()

	engine.Run(os.Getenv("APP_PORT"))

}
