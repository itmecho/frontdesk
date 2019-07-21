CREATE TABLE users (
    id uuid PRIMARY KEY,
    name varchar not null,
    email varchar not null unique,
    password varchar,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp
);