package models

type CreateUserRequest struct {
	Password *string
	UserName *string
	Email    *string
}
