package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Password     string `json:"password"`
}
type UserRequest struct {
	Name         string `json:"name" validate:"alpha"`
	Email        string `json:"email" validate:"email"`
	Password     string `json:"password"` 
	MobileNumber string `json:"mobile_number"`
}
