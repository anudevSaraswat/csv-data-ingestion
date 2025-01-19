package main

import (
	"csv-producer/connect"
	"csv-producer/csv"
	"csv-producer/models"
	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	conn := connect.ConnectToMessageQueue()

	defer conn.Close()

	mqChannel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	queue, err := mqChannel.QueueDeclare("users", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	dataChannel := make(chan models.User)

	go func() {
		err = csv.ProcessCSV(dataChannel)
		if err != nil {
			panic(err)
		}
	}()

	for user := range dataChannel {
		by, err := json.Marshal(user)
		if err != nil {
			panic(err)
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
			panic(err)
		}
	}

}
