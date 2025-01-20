package models

type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	DOB    string `json:"dob"`
	City   string `json:"city"`
}
