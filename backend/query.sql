-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetPackagesByUserID :many
SELECT * FROM
  users_packages
WHERE
  user_id = $1
  AND
  (now()::timestamp with TIME ZONE) BETWEEN valid_from AND valid_until
ORDER BY
  created_at DESC
LIMIT 10;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (email, password) VALUES ($1, $2);

-- name: AddRevokedToken :exec
INSERT INTO revoked_tokens (token, user_id, valid_until) VALUES ($1, $2, $3);

-- name: CheckToken :one
SELECT * FROM revoked_tokens WHERE token = $1 LIMIT 1;
