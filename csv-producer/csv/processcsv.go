package csv

import (
	"csv-producer/models"
	"csv-producer/utils"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ProcessCSV(ch chan<- models.User) error {

	file, err := os.Open("people-10000.csv")
	if err != nil {
		log.Fatalln("(ProcessCSV) err in os.Open:", err)
	}

	defer file.Close()
	defer close(ch)

	csvReader := csv.NewReader(file)

	_, _ = csvReader.Read()

	counter := 1
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Println("(ProcessCSV) err in csv.Reader:", err)
			return err
		}

		if len(record) < 5 {
			return fmt.Errorf("Need at least 5 data fields for a user row")
		}

		// validate email address
		if ok := utils.ValidateEmail(record[5]); !ok {
			return fmt.Errorf("Invalid email address: %s", record[2])
		}

		// validate dob format
		if ok := utils.ValidateDOB(record[7]); !ok {
			return fmt.Errorf("Invalid date of birth: %s", record[3])
		}

		user := models.User{
			UserID:    record[1],
			FirstName: record[2],
			LastName:  record[3],
			Sex:       record[4],
			Email:     record[5],
			Phone:     record[6],
			DOB:       record[7],
			JobTitle:  record[8],
		}

		ch <- user

		counter++

	}

	return nil

}
