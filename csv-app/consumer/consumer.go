package consumer

import (
	database "csv-app/db-connect"
	"csv-app/models"
	"encoding/json"
	"fmt"

	"github.com/RediSearch/redisearch-go/redisearch"
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

	insertSQL := `INSERT INTO user_data (user_id, first_name, last_name, sex, email, phone, dob, job_title) VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8);`

	counter := 1
	for data := range dataChannel {
		user := models.User{}
		err = json.Unmarshal(data.Body, &user)
		if err != nil {
			return err
		}

		fmt.Printf("user-%d:%v\n", counter, user)
		// insert user record into the database
		_, err := db.Exec(insertSQL, user.UserID, user.FirstName, user.LastName, user.Sex, user.Email, user.Phone, user.DOB, user.JobTitle)
		if err != nil {
			return err
		}

		// store user json in cache
		document := redisearch.NewDocument(fmt.Sprintf("user:%d", counter), 1.0)

		document.Set("user_id", user.UserID)
		document.Set("first_name", user.FirstName)
		document.Set("last_name", user.LastName)
		document.Set("sex", user.Sex)
		document.Set("email", user.Email)
		document.Set("phone", user.Phone)
		document.Set("dob", user.DOB)
		document.Set("job_title", user.JobTitle)

		if err := cache.SearchDB.Index(document); err != nil {
			return err
		}

		counter++

	}

	return nil

}
