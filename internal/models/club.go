package models

import "github.com/google/uuid"

type ClubUpdate struct {
	Name         *string `json:"name"`
	Specs        *string `json:"specs"`
	Description  *string `json:"description"`
	TgURL        *string `json:"tg_url"`
	LogoURL      *string `json:"logo_url"`
	AdminContact *string `json:"admin_contact"`
	FromOrg      *bool   `json:"from_org"`
}

type ClubCreate struct {
	OwnerID      uuid.UUID `json:"owner_id"`
	Name         string    `json:"name"`
	Specs        string    `json:"specs"`
	Description  string    `json:"description"`
	TgURL        string    `json:"tg_url"`
	LogoURL      string    `json:"logo_url"`
	AdminContact string    `json:"admin_contact"`
	FromOrg      bool      `json:"from_org"`
}

type ClubGet struct {
	ID uuid.UUID `param:"id" `
}

type ClubDelete struct {
	ID uuid.UUID `param:"id" `
}
