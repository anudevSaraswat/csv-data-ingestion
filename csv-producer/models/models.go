package models

type User struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       string `json:"sex"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	DOB       string `json:"dob"`
	JobTitle  string `json:"job_title"`
}
