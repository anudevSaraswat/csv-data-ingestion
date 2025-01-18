package dbconnect

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

var connection *amqp.Connection

// this function provides a connection handle to rabbit message queue service
func ConnectToMessageQueue() *amqp.Connection {

	if connection != nil {
		return connection
	}

	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	mqPort, err := strconv.Atoi(os.Getenv("RMQ_PORT"))
	if err != nil {
		log.Default().Println("No value found for port, using default...")
		mqPort = 5672
	}

	uri := amqp.URI{
		Scheme:   "amqp",
		Host:     os.Getenv("RMQ_HOST"),
		Port:     mqPort,
		Username: os.Getenv("RMQ_USER"),
	}

	connection, err = amqp.Dial(uri.String())
	if err != nil {
		panic(err)
	}

	return connection

}
