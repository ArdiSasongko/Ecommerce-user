-- name: GetRoleByLevel :one
SELECT level, name, description FROM roles WHERE level = $1;

-- name: GetRoleByName :one
SELECT level, name, description FROM roles WHERE name = $1;