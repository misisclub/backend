package models

import (
	"time"

	"github.com/google/uuid"
)

// FeedRequest represents the request for feed operations with pagination
type FeedRequest struct {
	Limit  int32  `json:"limit" query:"limit" validate:"min=1,max=100"`
	Offset int32  `json:"offset" query:"offset" validate:"min=0"`
	Tag    string `json:"tag" query:"tag"`
}

// SubscriptionFeedRequest represents the request for subscription feed with pagination
type SubscriptionFeedRequest struct {
	UserID uuid.UUID `json:"user_id" param:"user_id" validate:"required"`
	Limit  int32     `json:"limit" query:"limit" validate:"min=1,max=100"`
	Offset int32     `json:"offset" query:"offset" validate:"min=0"`
	Tag    string    `json:"tag" query:"tag"`
}

// ClubFeedRequest represents the request for club feed with pagination
type ClubFeedRequest struct {
	ClubID uuid.UUID `json:"club_id" param:"club_id" validate:"required"`
	Limit  int32     `json:"limit" query:"limit" validate:"min=1,max=100"`
	Offset int32     `json:"offset" query:"offset" validate:"min=0"`
	Tag    string    `json:"tag" query:"tag"`
}

// FeedResponse represents the response for feed operations
type FeedResponse struct {
	Posts []PostResponse `json:"posts"`
	Total int64          `json:"total"`
}

// PostResponse represents a post in feed responses (reusing from post.go)
type PostResponse struct {
	ID          uuid.UUID `json:"id"`
	Tag         string    `json:"tag"`
	OwnerID     uuid.UUID `json:"owner_id"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	FromOrg     bool      `json:"from_org"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
