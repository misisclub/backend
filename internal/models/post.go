package models

import (
	"time"

	"github.com/google/uuid"
)

// Post represents a post in the system
type Post struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Tag         string    `json:"tag" db:"tag"`
	OwnerID     uuid.UUID `json:"owner_id" db:"owner_id"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	FromOrg     bool      `json:"from_org" db:"from_org"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CreatePostRequest represents the request to create a new post
type CreatePostRequest struct {
	Tag         string    `json:"tag"`
	OwnerID     uuid.UUID `json:"owner_id"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	FromOrg     bool      `json:"from_org"`
}

// UpdatePostRequest represents the request to update a post
type UpdatePostRequest struct {
	Tag         string `json:"tag" `
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	FromOrg     bool   `json:"from_org"`
}
