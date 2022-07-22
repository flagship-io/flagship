package models

type User struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserRequest struct {
	Users []User `json:"role"`
}
