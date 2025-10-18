-- +goose Up
CREATE TABLE subscription (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    club_id uuid NOT NULL
);
