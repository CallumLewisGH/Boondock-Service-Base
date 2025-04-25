package models

type UserRequest struct {
	Password string `json:"password"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
