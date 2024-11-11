-- +goose Up
-- +goose StatementBegin
CREATE TYPE valid_cust AS ENUM ('person','company','gov');
CREATE TABLE customers(
    id SERIAL PRIMARY KEY,
    cust_code VARCHAR(5) NOT NULL,
    cust_name VARCHAR(255) NOT NULL,
    cust_address_1 VARCHAR(255) NOT NULL,
    cust_address_2 VARCHAR(255) NOT NULL,
    cust_zipcode VARCHAR(10) NOT NULL,
    cust_branch VARCHAR(255) NOT NULL,
    cust_type valid_cust, 
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP,
    created_by BIGINT NOT NULL REFERENCES "users" (id),
    updated_by BIGINT REFERENCES "users" (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
DROP TYPE IF EXISTS valid_cust; 
-- +goose StatementEnd
