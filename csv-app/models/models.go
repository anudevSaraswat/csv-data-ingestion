package models

// this struct can be used to scan user data in and for sending filters for querying data
type User struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	DOB    string `json:"dob"`
	City   string `json:"city"`
}

type Response struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	Data         any    `json:"data"`
}
