-- +goose Up
-- +goose StatementBegin
CREATE TYPE receipt_status AS ENUM ('active','inactive');
CREATE TABLE receipt_headers(
    id SERIAL PRIMARY KEY,
    rec_status receipt_status DEFAULT ('active'),
    rec_no VARCHAR(20) NOT NULL,
    rec_date DATE NOT NULL,
    rec_payment_amt DOUBLE PRECISION NOT NULL,
    customer_id BIGINT NOT NULL REFERENCES "customers" (id),
    rec_cust_name VARCHAR(255) NOT NULL,
    rec_cust_address_1 VARCHAR(255) NOT NULL,
    rec_cust_address_2 VARCHAR(255) NOT NULL,
    rec_cust_zipcode VARCHAR(10) NOT NULL,
    rec_cust_branch VARCHAR(255) NOT NULL,
    rec_remark VARCHAR(255) NOT NULL,
    created_by BIGINT NOT NULL REFERENCES "users" (id),
    updated_by BIGINT REFERENCES "users" (id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP 
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS receipt_headers;
DROP TYPE IF EXISTS receipt_flag;
-- +goose StatementEnd
