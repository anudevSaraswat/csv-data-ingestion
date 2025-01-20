package main

import (
	"csv-app/routes"
	"fmt"
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
			fmt.Println(err)
		}
	}()

	engine.Run(os.Getenv("APP_PORT"))

}
