ALTER TABLE users
ADD COLUMN role integer references roles(level) not null default 1;