package consumer

import (
	database "csv-app/db-connect"
	"csv-app/models"
	"encoding/json"
)

// this function will read data from message broker's queue
// and populate it into the PostgreSQL database
// TODO: add the data in cache too
func ReadAndStoreData() error {

	conn := database.ConnectToMessageQueue()
	db := database.ConnectToDatabase()

	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	queue, err := channel.QueueDeclare("users", false, false, false, false, nil)
	if err != nil {
		return err
	}

	dataChannel, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	insertSQL := `INSERT INTO users (name, email, dob, city) VALUES 
	(?, ?, ?, ?);`

	for data := range dataChannel {
		user := models.User{}
		err = json.Unmarshal(data.Body, &user)
		if err != nil {
			return err
		}

		_, err := db.Exec(insertSQL, user.Name, user.Email, user.DOB, user.City)
		if err != nil {
			return err
		}
	}

	return nil

}
