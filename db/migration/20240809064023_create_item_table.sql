-- +goose Up
-- +goose StatementBegin
create table items (
    id serial primary key,
    category_id int references categories(id),
    name text,
    description text,
    done boolean,
    created_at timestamp 
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists items;
-- +goose StatementEnd
