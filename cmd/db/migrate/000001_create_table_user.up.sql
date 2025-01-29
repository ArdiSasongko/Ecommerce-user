CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    username varchar(255) not null unique,
    email varchar(255) not null unique,
    phone_number varchar(255) not null unique,
    address varchar(255) not null,
    dob DATE not null,
    password varchar(255) not null,
    fullname varchar(255) not null,
    created_at timestamp(0) with time zone not null default now(),
    updated_at timestamp(0) with time zone not null default now()
);