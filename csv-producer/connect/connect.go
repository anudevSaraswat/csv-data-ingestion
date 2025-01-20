package connect

import (
	"log"
	"os"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

var connection *amqp.Connection

// this function provides a connection handle to rabbit message queue service
func ConnectToMessageQueue() *amqp.Connection {

	if connection != nil {
		return connection
	}

	mqPort, err := strconv.Atoi(os.Getenv("RMQ_PORT"))
	if err != nil {
		log.Println("No value found for port, using default...")
		mqPort = 5672
	}

	uri := amqp.URI{
		Scheme:   "amqp",
		Host:     os.Getenv("RMQ_HOST"),
		Port:     mqPort,
		Username: os.Getenv("RMQ_USER"),
		Password: os.Getenv("RMQ_PASSWORD"),
	}

	connection, err = amqp.Dial(uri.String())
	if err != nil {
		log.Fatalln("(ConnectToMessageQueue) err in amqp.Dial:", err)
	}

	return connection

}
