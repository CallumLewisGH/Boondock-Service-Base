package models

import (
	"time"

	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	ID             *uuid.UUID `json:"id,omitempty"`
	CreatedUTC     *time.Time `json:"created_utc,omitempty"`
	DeletedUTC     *time.Time `json:"deleted_utc,omitempty"`
	UserName       *string    `json:"user_name,omitempty"`
	Email          *string    `json:"email,omitempty"`
	PasswordHash   *string    `json:"password_hash,omitempty"`
	ProfilePicture *string    `json:"profile_picture,omitempty"`
}
