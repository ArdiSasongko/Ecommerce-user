CREATE TABLE IF NOT EXISTS roles (
    level int not null unique,
    name varchar(255) not null unique,
    description varchar(255),
    created_at timestamp(0) with time zone not null default now(),
    updated_at timestamp(0) with time zone not null default now()
);

INSERT INTO roles (level, name, description) VALUES (2, 'admin', 'admin role');
INSERT INTO roles (level, name, description) VALUES (1, 'user', 'default role');