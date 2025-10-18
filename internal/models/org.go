package models

import "github.com/google/uuid"

type OrgUpdate struct {
	Name         *string `json:"name"`
	Specs        *string `json:"specs"`
	Description  *string `json:"description"`
	WebsiteURL   *string `json:"website_url"`
	LogoURL      *string `json:"logo_url"`
	VideoURL     *string `json:"video_url"`
	AdminContact *string `json:"admin_contact"`
	Inn          *string `json:"inn"`
	Ogrn         *string `json:"ogrn"`
}

type OrgCreate struct {
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	Specs        string    `json:"specs"`
	Description  string    `json:"description"`
	WebsiteURL   string    `json:"website_url"`
	LogoURL      string    `json:"logo_url"`
	VideoURL     string    `json:"video_url"`
	AdminContact string    `json:"admin_contact"`
	Inn          string    `json:"inn"`
	Ogrn         string    `json:"ogrn"`
}

type OrgGet struct {
	ID uuid.UUID `param:"id"`
}

type OrgDelete struct {
	ID uuid.UUID `param:"id"`
}
