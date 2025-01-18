package main

import (
	"csv-app/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	engine := routes.SetupRoutes()

	engine.Run(os.Getenv("APP_PORT"))

}
