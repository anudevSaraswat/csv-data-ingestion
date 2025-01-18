package utils

import (
	"net/mail"
	"time"
)

func ValidateEmail(email string) bool {

	_, err := mail.ParseAddress(email)
	return err == nil

}

// expecting DOB in format YYYY-MM-DD
func ValidateDOB(dob string) bool {
	_, err := time.Parse("2006-01-02", dob)
	return err == nil
}
