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

-- name: AddVideo :one
INSERT INTO video (
  name,
  description,
  category_id
) VALUES ($1, $2, $3)
RETURNING *;

-- name: AddVideoFile :exec
INSERT INTO video_files (
  file_path,
  video_id,
  file_size,
  duration,
  resolution
) VALUES ($1, $2, $3, $4, $5);

-- name: AddVideoThumbnail :exec
INSERT INTO video_thumbnails (
  video_id,
  file_size
) VALUES ($1, $2);

-- name: AddCategory :one 
INSERT INTO categories (name) VALUES ($1) RETURNING *;

-- name: GetAllCategories :many
SELECT * FROM categories LIMIT $1 OFFSET $2;

-- name: GetTagByName :one
SELECT * FROM tags WHERE name = $1 LIMIT 1;

-- name: AddTag :one
INSERT INTO tags (name) VALUES ($1) RETURNING *;

-- name: AddVideoTag :exec
INSERT INTO video_tags (video_id, tag_id) VALUES ($1, $2) RETURNING *;