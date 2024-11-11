-- +goose Up
-- +goose StatementBegin
CREATE TABLE receipt_details(
    id SERIAL PRIMARY KEY,
    receipt_header_id BIGINT NOT NULL REFERENCES "receipt_headers" (id),
    invoice_detail_id BIGINT REFERENCES "invoice_details" (id),
    recd_inv_no VARCHAR(20) NOT NULL, 
    recd_ps_code VARCHAR(4) NOT NULL,
    recd_ps_name_th VARCHAR(255) NOT NULL,
    recd_ps_name_en VARCHAR(255) NOT NULL,
    recd_vat DECIMAL(3,2) NOT NULL,
    recd_whtax DECIMAL(3,2) NOT NULL,
    recd_amt DOUBLE PRECISION NOT NULL, 
    recd_vat_amt DOUBLE PRECISION NOT NULL,
    recd_whtax_amt DOUBLE PRECISION NOT NULL,
    recd_net_amt DOUBLE PRECISION NOT NULL,
    recd_wh_pay DOUBLE PRECISION NOT NULl DEFAULT (0),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    created_by BIGINT REFERENCES "users" (id),
    updated_by BIGINT NOT NULL REFERENCES "users" (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS receipt_details;
-- +goose StatementEnd
