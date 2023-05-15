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

-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1;

-- name: RemoveAllUserPackages :exec
DELETE FROM users_packages WHERE user_id = $1;

-- name: GiveUserPackage :exec
INSERT INTO users_packages (
  user_id,
  tier,
  valid_from,
  valid_until
) VALUES (
  $1, $2, $3, $4
);

-- name: GetVideosByCategory :many
SELECT
  v.id id,
  v.name name,
  v.description description,
  c.name category,
  v.upvotes upvotes,
  v.downvotes downvotes,
  v.views views,
  th.file_name as thumbnail
FROM
  video v LEFT JOIN categories c ON v.category_id=c.id
  LEFT JOIN video_thumbnails th ON th.video_id = v.id
WHERE
  c.id = $1
LIMIT $2
OFFSET $3;

-- name: GetRandomVideos :many
SELECT
  v.id id,
  v.name name,
  v.description description,
  c.name category,
  v.upvotes upvotes,
  v.downvotes downvotes,
  v.views views,
  th.file_name as thumbnail
FROM
  video v LEFT JOIN categories c ON v.category_id=c.id
  LEFT JOIN video_thumbnails th ON th.video_id = v.id
WHERE
  c.id = $1
ORDER BY RANDOM()
LIMIT $2
OFFSET $3;

-- name: GetVideosByName :many
SELECT
  v.id id,
  v.name name,
  v.description description,
  c.name category,
  v.upvotes upvotes,
  v.downvotes downvotes,
  v.views views,
  th.file_name as thumbnail
FROM
  video v LEFT JOIN categories c ON v.category_id=c.id
  LEFT JOIN video_thumbnails th ON th.video_id = v.id
WHERE
  v.name LIKE $1
LIMIT $2
OFFSET $3;

-- name: GetVideoByID :one
SELECT
  v.id id,
  v.name name,
  v.description description,
  c.name category,
  v.upvotes upvotes,
  v.downvotes downvotes,
  v.views views,
  th.file_name as thumbnail
FROM
  video v LEFT JOIN categories c ON v.category_id = c.id
  LEFT JOIN video_thumbnails th ON th.video_id = v.id
WHERE
  v.id = $1
LIMIT 1;

-- name: GetAllVideos :many
SELECT
  v.id id,
  v.name name,
  v.description description,
  c.name category,
  v.upvotes upvotes,
  v.downvotes downvotes,
  v.views views,
  th.file_name as thumbnail
FROM
  video v LEFT JOIN categories c ON v.category_id = c.id
  LEFT JOIN video_thumbnails th ON th.video_id = v.id
LIMIT $1
OFFSET $2;

-- name: GetVideoTags :many
SELECT
  name
FROM
  tags t INNER JOIN video_tags vt ON t.id = vt.tag_id
WHERE
  vt.video_id = $1
LIMIT 10;

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

-- name: UpdateCategory :exec
UPDATE categories SET name = $1 WHERE id = $2;

-- name: GetCategoryByName :one
SELECT * FROM categories WHERE name = $1 LIMIT 1;

-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: GetAllCategories :many
SELECT * FROM categories ORDER BY name ASC LIMIT $1 OFFSET $2;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;

-- name: GetTagByName :one
SELECT * FROM tags WHERE name = $1 LIMIT 1;

-- name: AddTag :one
INSERT INTO tags (name) VALUES ($1) RETURNING *;

-- name: AddVideoTag :exec
INSERT INTO video_tags (video_id, tag_id) VALUES ($1, $2) RETURNING *;