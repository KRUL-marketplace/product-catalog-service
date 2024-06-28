-- +goose Up
-- +goose StatementBegin
alter table products
add column quantity integer default 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table products
drop column quantity;
-- +goose StatementEnd
