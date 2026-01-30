-- +goose Up
create table users (
    id serial primary key,
    username text unique,
    password_hash text
);

-- +goose Down
drop table if exists users;
