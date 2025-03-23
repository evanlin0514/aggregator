-- +goose Up
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    create_at TIMESTAMP NOT NULL DEFAULT NOW(),
    update_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR(255) UNIQUE NOT NULL
);

-- +goose Down 
DROP TABLE users;