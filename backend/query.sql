-- name: GetAllUsers :many
SELECT *
FROM users
LIMIT $1
OFFSET $2;

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

-- name: GetOverlapingPackages :many
SELECT * FROM
  users_packages
WHERE
  $1 BETWEEN valid_from AND valid_until
  OR
  $2 BETWEEN valid_until AND valid_from
  OR ($1 < valid_from AND $2 > valid_until)
ORDER BY
  created_at DESC
LIMIT 5;

-- name: GetPackageByID :one
SELECT * FROM
  users_packages
WHERE
  id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CreateVerificationCode :one
INSERT INTO verification (
  user_id,
  task_type
) VALUES ($1, 'emailVerification')
RETURNING *;

-- name: CreateResetCode :one
INSERT INTO verification (
  user_id,
  task_type
) VALUES ($1, 'passwordReset')
RETURNING *;

-- name: GetVerificationCode :one
SELECT * FROM verification WHERE code = $1 LIMIT 1;

-- name: SetCodeAsUsed :exec
UPDATE verification SET used = true WHERE id = $1;

-- name: VerifyUser :exec
UPDATE users SET verified_at = now() WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1;

-- name: UpdateUserEmail :exec
UPDATE users SET email = $2 WHERE id = $1;

-- name: UpdateUserVerification :exec
UPDATE users SET verified_at = $2 WHERE id = $1;

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

-- name: RemoveUserPackage :exec
DELETE FROM users_packages WHERE id = $1;

-- name: UpdatePackage :exec
UPDATE
  users_packages
SET
  tier = $1,
  valid_from = $2,
  valid_until = $3;

-- name: GetVideosByCategory :many
SELECT
  v.id id,
  v.name name,
  v.description description,
  c.name category,
  v.upvotes upvotes,
  v.downvotes downvotes,
  v.views views,
  v.is_premium is_premium,
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
  v.is_premium is_premium,
  th.file_name as thumbnail
FROM
  video v LEFT JOIN categories c ON v.category_id = c.id
  LEFT JOIN video_thumbnails th ON th.video_id = v.id
ORDER BY RANDOM()
LIMIT $1
OFFSET $2;

-- name: GetVideosByName :many
SELECT
  v.id id,
  v.name name,
  v.description description,
  c.name category,
  v.upvotes upvotes,
  v.downvotes downvotes,
  v.views views,
  v.is_premium is_premium,
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
  v.is_premium is_premium,
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
  v.is_premium is_premium,
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

-- name: GetVideoFiles :many
SELECT
  resolution,
  duration,
  file_path
FROM
  video_files
WHERE
  video_id = $1;

-- name: AddVideo :one
INSERT INTO video (
  name,
  description,
  category_id
) VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateVideo :one
UPDATE video
SET
  name = $2,
  description = $3,
  category_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteVideo :exec
DELETE FROM video WHERE id = $1;

-- name: AddThumbnail :exec
INSERT INTO video_thumbnails (
  video_id,
  file_size,
  file_name
) VALUES ($1, $2, $3);

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

-- name: ClearVideoTags :exec
DELETE FROM video_tags WHERE video_id = $1;