package models

import "github.com/google/uuid"

// Item represents a shopping item
// @Description Shopping item information
type Item struct {
	Id   uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name string    `json:"name" example:"Milk"`
}
