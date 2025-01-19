package csv

import (
	"csv-producer/models"
	"csv-producer/utils"
	"encoding/csv"
	"fmt"
	"os"
)

func ProcessCSV(ch chan<- models.User) error {

	file, err := os.Open("users.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer close(ch)

	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("failed to read csv row: %s", err.Error())
		}

		if len(record) < 5 {
			return fmt.Errorf("Need at least 5 data fields for a user row")
		}

		// validate email address
		if ok := utils.ValidateEmail(record[2]); !ok {
			return fmt.Errorf("Invalid email address: %s", record[2])
		}

		// validate dob format
		if ok := utils.ValidateDOB(record[3]); !ok {
			return fmt.Errorf("Invalid date of birth: %s", record[3])
		}

		user := models.User{
			Name:  record[1],
			Email: record[2],
			DOB:   record[3],
			City:  record[4],
		}

		fmt.Println("writing to channel")

		ch <- user

	}

	return nil

}
