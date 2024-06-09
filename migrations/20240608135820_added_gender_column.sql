-- +goose Up
-- +goose StatementBegin
ALTER TABLE products ADD COLUMN gender TEXT NOT NULL CHECK (gender IN ('men', 'women', 'unisex'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE products DROP COLUMN gender;
-- +goose StatementEnd
