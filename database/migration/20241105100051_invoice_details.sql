-- +goose Up
-- +goose StatementBegin
CREATE TYPE receipt_flag AS ENUM('yes','no');
CREATE TABLE invoice_details(
    id SERIAL PRIMARY KEY,
    invoice_header_id BIGINT NOT NULL REFERENCES "invoice_headers" (id),
    invd_ps_code varchar(4) NOT NULL,
    invd_ps_name_th varchar(255) NOT NULL,
    invd_ps_name_en varchar(255) NOT NULL,
    invd_vat DECIMAL(3,2) NOT NULL,
    invd_whtax DECIMAL(3,2) NOT NULL,
    invd_vat_amt DOUBLE PRECISION NOT NULL, 
    invd_whtax_amt DOUBLE PRECISION NOT NULL,
    invd_net_amt DOUBLE PRECISION NOT NULL,
    invd_receipt_flag receipt_flag DEFAULT ('no'),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    created_by BIGINT NOT NUll REFERENCES "users" (id),
    updated_by BIGINT REFERENCES "users" (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS invoice_details;
DROP TYPE IF EXISTS receipt_flag;
-- +goose StatementEnd
