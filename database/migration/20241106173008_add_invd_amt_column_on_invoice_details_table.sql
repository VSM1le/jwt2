-- +goose Up
-- +goose StatementBegin
ALTER TABLE invoice_details
ADD invd_amt DOUBLE PRECISION NOT NULL; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE invoice_details
DROP invd_amt DOUBLE PRECISION;
-- +goose StatementEnd
