-- +goose Up
-- +goose StatementBegin
create table categories (
    id serial primary key,
    user_id int references users(id),
    name text,
    priority int,
    CONSTRAINT unique_id_name UNIQUE (id, name)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists categories;
-- +goose StatementEnd
