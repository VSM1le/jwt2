-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL, 
    password varchar(255) UNIQUE NOT NULL,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    token varchar(255),
    refresh_token varchar(255),
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
