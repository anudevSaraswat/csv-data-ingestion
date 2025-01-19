package connect

import (
	"log"
	"os"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

// this function provides a connection handle to rabbit message queue service
// TODO:  convert this to singleton
func ConnectToMessageQueue() *amqp.Connection {

	var mqPort int
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
		Password: os.Getenv("RMQ_PASSWORD"),
	}

	connection, err := amqp.Dial(uri.String())
	if err != nil {
		panic(err)
	}

	return connection

}
