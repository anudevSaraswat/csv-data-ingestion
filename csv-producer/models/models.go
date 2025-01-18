package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	DOB   string `json:"dob"`
	City  string `json:"city"`
}
