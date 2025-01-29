CREATE TABLE IF NOT EXISTS user_session (
    id serial primary key,
    user_id int not null,
    active_token text not null,
    refresh_token text not null,
    active_token_expires_at timestamp not null,
    refresh_token_expires_at timestamp not null,
    created_at timestamp(0) with time zone not null default now(),
    updated_at timestamp(0) with time zone not null default now()
);