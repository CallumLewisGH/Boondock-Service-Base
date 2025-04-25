package models

import (
	"time"

	"github.com/google/uuid"
)

type UserDbRequest struct {
	ID             *uuid.UUID `db:"id"`
	CreatedUTC     *time.Time `db:"created_utc"`
	DeletedUTC     *time.Time `db:"deleted_utc"`
	UserName       *string    `db:"user_name"`
	Email          *string    `db:"email"`
	PasswordHash   *string    `db:"password_hash"`
	ProfilePicture *string    `db:"profile_picture"`
}
