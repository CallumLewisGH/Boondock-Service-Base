package models

import (
	"time"

	"github.com/google/uuid"
)

type UserDbRequest struct {
	ID             *uuid.UUID `db:"id" json:"id,omitempty"`
	CreatedUTC     *time.Time `db:"created_utc" json:"created_utc,omitempty"`
	DeletedUTC     *time.Time `db:"deleted_utc" json:"deleted_utc,omitempty"`
	UserName       *string    `db:"user_name" json:"user_name,omitempty"`
	Email          *string    `db:"email" json:"email,omitempty"`
	PasswordHash   *string    `db:"password_hash" json:"password_hash,omitempty"`
	ProfilePicture *string    `db:"profile_picture" json:"profile_picture,omitempty"`
}
