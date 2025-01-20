package main

import (
	"csv-producer/connect"
	"csv-producer/csv"
	"csv-producer/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("failed to load env file...")
	}

	conn := connect.ConnectToMessageQueue()

	defer conn.Close()

	mqChannel, err := conn.Channel()
	if err != nil {
		log.Fatalln("(main) err in conn.Channel:", err)
	}

	queue, err := mqChannel.QueueDeclare("users", false, false, false, false, nil)
	if err != nil {
		log.Fatalln("(main) err in mqChannel.QueueDeclare:", err)
	}

	dataChannel := make(chan models.User)

	go func() {
		err = csv.ProcessCSV(dataChannel)
		if err != nil {
			log.Fatalln("(main) err in csv.ProcessCSV:", err)
		}
	}()

	counter := 1

	for user := range dataChannel {
		by, err := json.Marshal(user)
		if err != nil {
			log.Printf("(main) err in json.Marshal for user %s:\n", user.FirstName)
			continue
		}

		err = mqChannel.Publish(
			"",
			queue.Name,
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        by,
			},
		)
		if err != nil {
			log.Printf("(main) err in mqChannel.Publish for user %s:\n", user.FirstName)
			continue
		}

		counter++
	}

	fmt.Printf("%d user records sent to message broker.\n", counter)

}
