-- name: InsertUser :one
INSERT INTO users (username, email, phone_number, address, dob, password, fullname, role) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;

-- name: GetUserByUsername :one
SELECT id, username, email, phone_number, address, dob, password, fullname, role, created_at, updated_at FROM users WHERE username = $1;

-- name: GetUserByEmail :one
SELECT id, username, email, phone_number, address, dob, password, fullname, role, created_at, updated_at FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT id, username, email, phone_number, address, dob, password, fullname, role, created_at, updated_at FROM users WHERE id = $1;