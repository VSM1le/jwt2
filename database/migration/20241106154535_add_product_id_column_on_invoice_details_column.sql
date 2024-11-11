-- +goose Up
-- +goose StatementBegin
ALTER TABLE invoice_details
ADD product_id BIGINT NOT NULL REFERENCES "product_services" (id) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE invoice_details
DROP product_id BIGINT;
-- +goose StatementEnd
