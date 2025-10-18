-- +goose Up
CREATE TABLE club (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL,
    name TEXT NOT NULL,                           
    specs TEXT NOT NULL,                    
    description TEXT NOT NULL,                    
    tg_url TEXT NOT NULL,                              
    logo_url TEXT NOT NULL,                                
    admin_contact TEXT NOT NULL,                  
    from_org BOOL NOT NULL,                  
    UNIQUE (owner_id)                              
);
