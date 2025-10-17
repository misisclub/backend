package models

import "github.com/google/uuid"

type ClientSignUp struct {
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	FamilyName  string `json:"family_name"`
	Age         int32  `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type ClientSignIn struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type SignSuccess struct {
	Token string    `json:"token"`
	ID    uuid.UUID `json:"id"`
}

type ClientDelete struct {
	ID uuid.UUID `param:"id"`
}

type ClientGet struct {
	ID uuid.UUID `param:"id"`
}

type ClientUpdate struct {
	FirstName   *string `json:"first_name"`
	SecondName  *string `json:"second_name"`
	FamilyName  *string `json:"family_name"`
	Age         *int32  `json:"age"`
	PhoneNumber *string `json:"phone_number"`
	Profession1 *string `json:"profession_1"`
	Profession2 *string `json:"profession_2"`
	Company     *string `json:"company"`
	University  *string `json:"university"`
}

type Status struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"message"`
}

type ClientModel struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	SecondName  string    `json:"second_name"`
	FamilyName  string    `json:"family_name"`
	Age         int32     `json:"age"`
	PhoneNumber string    `json:"phone_number"`
	Profession1 *string   `json:"profession_1"`
	Profession2 *string   `json:"profession_2"`
	Company     *string   `json:"company"`
	University  *string   `json:"university"`
}
