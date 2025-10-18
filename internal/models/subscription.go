package models

import "github.com/google/uuid"

type SubscriptionCreate struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	ClubID uuid.UUID `json:"club_id" validate:"required"`
}

type SubscriptionGet struct {
	ID uuid.UUID `param:"id"`
}

type SubscriptionGetByUser struct {
	UserID uuid.UUID `param:"user_id"`
}

type SubscriptionGetByClub struct {
	ClubID uuid.UUID `param:"club_id"`
}

type SubscriptionGetByUserAndClub struct {
	UserID uuid.UUID `param:"user_id"`
	ClubID uuid.UUID `param:"club_id"`
}

type SubscriptionDelete struct {
	ID uuid.UUID `param:"id"`
}

type SubscriptionDeleteByUserAndClub struct {
	UserID uuid.UUID `param:"user_id"`
	ClubID uuid.UUID `param:"club_id"`
}

type SubscriptionModel struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	ClubID uuid.UUID `json:"club_id"`
}
