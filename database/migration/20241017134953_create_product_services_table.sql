-- +goose Up
-- +goose StatementBegin
CREATE table product_services(
    id SERIAL PRIMARY KEY, 
    ps_code VARCHAR(4) NOT NULL,
    ps_name_th VARCHAR(255) NOT NULL,
    ps_name_en VARCHAR(255) NOT NULL,
    ps_vat DECIMAL(3,2) NOT NULL,
    ps_whtax DECIMAL(3,2) NOT NULL,
    ps_gov_whtax DECIMAL(3,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    created_by BIGINT NOT NULL REFERENCES "users" (id),
    updated_by BIGINT REFERENCES "users" (id) 
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_services;
-- +goose StatementEnd
