package models

// this struct can be used to scan user data in and for sending filters for querying data
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

type Response struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	Data         any    `json:"data"`
}
