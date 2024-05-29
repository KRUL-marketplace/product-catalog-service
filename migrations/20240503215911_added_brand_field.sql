-- +goose Up
-- +goose StatementBegin
alter table products
    add brand_id INTEGER REFERENCES brands(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table products
    drop column brand_id;
-- +goose StatementEnd
