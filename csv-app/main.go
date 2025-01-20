package main

import (
	"csv-app/routes"
	"fmt"
	"log"
	"os"

	"csv-app/consumer"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load env file...")
	}

	engine := routes.SetupRoutes()

	go func() {
		err := consumer.ReadAndStoreData()
		if err != nil {
			fmt.Println(err)
		}
	}()

	engine.Run(os.Getenv("APP_PORT"))

}
