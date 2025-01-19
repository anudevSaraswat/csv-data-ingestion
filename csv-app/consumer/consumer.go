package consumer

import (
	"context"
	database "csv-app/db-connect"
	"csv-app/models"
	"encoding/json"
	"fmt"
)

// this function will read data from message broker's queue
// and populate it into the PostgreSQL database and cache
func ReadAndStoreData() error {

	conn := database.ConnectToMessageQueue()
	db := database.ConnectToDatabase()
	cache := database.ConnectToCache()

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

	insertSQL := `INSERT INTO users (user_id, name, email, dob, city) VALUES 
	($1, $2, $3, $4, $5);`

	counter := 1
	for data := range dataChannel {
		user := models.User{}
		err = json.Unmarshal(data.Body, &user)
		if err != nil {
			return err
		}

		// insert user record into the database
		_, err := db.Exec(insertSQL, user.UserID, user.Name, user.Email, user.DOB, user.City)
		if err != nil {
			return err
		}

		// store user json in cache
		documentKey := fmt.Sprintf("user:%d", counter)
		_, err = cache.JSONSet(context.Background(), documentKey, "$", user).Result()
		if err != nil {
			return err
		}

		counter++

	}

	return nil

}
