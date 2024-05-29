-- +goose Up
-- +goose StatementBegin
alter table products
    add slug TEXT not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table products
    drop column slug;
-- +goose StatementEnd
