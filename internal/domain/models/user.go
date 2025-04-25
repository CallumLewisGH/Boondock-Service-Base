package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `db:"id" json:"id"`
	CreatedUTC time.Time `db:"created_utc" json:"created_utc"`

	UserName       string  `db:"user_name" json:"user_name"`
	Email          string  `db:"email" json:"email"`
	ProfilePicture *string `db:"profile_picture" json:"profile_picture,omitempty"`
}
