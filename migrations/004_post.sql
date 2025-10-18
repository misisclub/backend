-- +goose Up
CREATE TABLE post (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    tag TEXT NOT NULL,
    owner_id uuid NOT NULL,
    description TEXT NOT NULL,                           
    image_url TEXT NOT NULL,                                
    from_org BOOL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()                  
);
