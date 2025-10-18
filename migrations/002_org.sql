-- +goose Up
CREATE TABLE org (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    name TEXT NOT NULL,                           
    specs TEXT NOT NULL,                    
    description TEXT NOT NULL,                    
    website_url TEXT NOT NULL,                              
    logo_url TEXT NOT NULL,                                
    video_url TEXT NOT NULL,                                
    admin_contact TEXT NOT NULL,                  
    inn VARCHAR(255) NOT NULL,                     
    ogrn VARCHAR(255) NOT NULL,                    
    UNIQUE (user_id)                              
);

