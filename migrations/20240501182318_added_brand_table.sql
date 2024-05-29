-- +goose Up
-- +goose StatementBegin
create table if not exists brands
(
    id  serial primary key,
    name text unique not null,
    slug text not null,
    description text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table brands;
-- +goose StatementEnd
