-- +goose Up
-- +goose StatementBegin
-- Create the junction table for the many-to-many relationship
create table if not exists product_categories
(
    product_id UUID references products(id),
    category_id integer references categories(id),
    primary key (product_id, category_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop the junction table for the many-to-many relationship
drop table if exists product_categories;
-- +goose StatementEnd