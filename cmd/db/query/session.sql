-- name: InsertSession :one
INSERT INTO user_session (user_id, active_token, refresh_token, active_token_expires_at, refresh_token_expires_at) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: GetSessionByActiveToken :one
SELECT id, user_id, active_token, refresh_token, active_token_expires_at, refresh_token_expires_at, created_at, updated_at FROM user_session WHERE active_token = $1;

-- name: GetSessionByUserID :one
SELECT id, user_id, active_token, refresh_token, active_token_expires_at, refresh_token_expires_at, created_at, updated_at FROM user_session WHERE user_id = $1;