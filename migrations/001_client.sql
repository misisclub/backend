-- +goose Up
CREATE TABLE client (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    family_name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    profession_1 VARCHAR(255),
    profession_2 VARCHAR(255),
    company VARCHAR(255),
    future VARCHAR(255),
    university VARCHAR(255),
    password_hash TEXT NOT NULL,
    UNIQUE (phone_number)
);
