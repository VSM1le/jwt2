-- +goose Up
-- +goose StatementBegin
CREATE TYPE invoice_status AS ENUM ('active','inactive');
CREATE TABLE invoice_headers(
    id SERIAL PRIMARY KEY, 
    inv_status invoice_status DEFAULT ('active'),
    inv_no VARCHAR(20) NOT NULL,
    inv_date date NOT NULL,
    customer_id BIGINT NOT NULL REFERENCES "customers" (id),
    inv_cust_name varchar(255) NOT NULL,
    inv_cust_address_1 varchar(255) NOT NULL,
    inv_cust_address_2 varchar(255) NOT NULL,
    inv_cust_zipcode varchar(10) NOT NULL,
    inv_cust_branch varchar(255) NOT NULL,
    inv_remark varchar(255),
    created_at TIMESTAMP DEFAULT NOW(), 
    updated_at TIMESTAMP,
    created_by BIGINT NOT NULL REFERENCES "users" (id),
    updated_by BIGINT REFERENCES "users" (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS invoice_headers;
DROP TYPE IF EXISTS invoice_status;
-- +goose StatementEnd
