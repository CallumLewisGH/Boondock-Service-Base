package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	CreatedUTC time.Time  `db:"created_utc" json:"created_utc"`
	DeletedUTC *time.Time `db:"deleted_utc" json:"deleted_utc,omitempty"`

	UserName       string  `db:"user_name" json:"user_name"`
	Email          string  `db:"email" json:"email"`
	PasswordHash   string  `db:"password_hash" json:"-"`
	ProfilePicture *string `db:"profile_picture" json:"profile_picture,omitempty"`
}
