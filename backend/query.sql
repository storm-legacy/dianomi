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
LIMIT $2
OFFSET $3;

-- name: GetOverlapingPackages :many
SELECT * FROM
  users_packages
WHERE
  ($1 BETWEEN valid_from AND valid_until
  OR
  $2 BETWEEN valid_until AND valid_from
  OR ($1 < valid_from AND $2 > valid_until))
  AND
  user_id = $3
ORDER BY
  created_at DESC
LIMIT 5;

-- name: GetPackageByID :one
SELECT * FROM
  users_packages
WHERE
  id = $1
LIMIT 1;

-- name: GetCurrentPackageForUser :one
SELECT * FROM
  users_packages
WHERE
  user_id = $1
  AND
  (now()::timestamp with TIME ZONE) BETWEEN valid_from AND valid_until
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

-- name: BanUser :exec
UPDATE users SET banned_at = now() WHERE id = $1;

-- name: UnbanUser :exec
UPDATE users SET banned_at = NULL WHERE id = $1;

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
  user_id = $2,
  tier = $3,
  valid_from = $4,
  valid_until = $5
WHERE
  id = $1;

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
  category_id = $4,
  is_premium = COALESCE($5, false)
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

-- name: AddReport :exec
INSERT INTO Error_Reports (error_title, error_description, reported_by) VALUES ($1, $2, $3) RETURNING *;

-- name: GetVideoIDByName :many
SELECT id FROM video WHERE LOWER(name) LIKE LOWER($1);
-- name: AddVideoMertics :exec
INSERT INTO user_video_metrics (user_id, video_id, time_spent_watching, stopped_at, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING *;

-- name: GetUserVideoMerticsByUserId :many
SELECT user_video_metrics.*, video.id, video.name, video.description, video.is_premium, video_thumbnails.file_name
FROM user_video_metrics
JOIN video ON user_video_metrics.video_id = video.id JOIN video_thumbnails ON video_thumbnails.video_id = video.id
WHERE user_video_metrics.user_id = $1;
-- name: GetUserVideoMerticsByVideoId :one
SELECT * FROM user_video_metrics WHERE video_id = $1 LIMIT 1;

-- name: IfUserSeeThisVideo :one
SELECT * FROM user_video_metrics WHERE user_id = $1 AND video_id=$2 LIMIT 1;

-- name: UpdateVideoMetric :exec
UPDATE user_video_metrics
SET time_spent_watching=time_spent_watching + $1, stopped_at = $2, updated_at = now()
WHERE id=$3;

-- name: UpdateVoteUpVideo :exec
UPDATE video
SET upvotes = upvotes + 1,
    downvotes = GREATEST(downvotes - 1, 0)
WHERE id = $1;

-- name: UpdateVoteDownVideo :exec
UPDATE video
SET downvotes = downvotes + 1,
    upvotes = GREATEST(upvotes - 1, 0)
WHERE id = $1;

-- name: GetAllVideoMetric :many
SELECT * FROM user_video_metrics LIMIT $1 OFFSET $2;


-- name: AddComments :exec
INSERT INTO comments (user_id, video_id, comment) VALUES ($1, $2, $3);

-- name: UpdateComments :exec
UPDATE comments
SET comment = $1,
    updated_at = NOW()
WHERE id = $2;

-- name: GetCommentsForVideo :many
SELECT
  u.email, c.id, c.comment, c.upvotes, c.downvotes, c.updated_at
FROM
  comments c
  INNER JOIN users u ON u.id = c.user_id
WHERE
  c.video_id = $1;

-- name: GetCommentsAll :many
SELECT
  u.email, v.name, c.id, c.comment, c.upvotes, c.downvotes, c.updated_at
FROM
  comments c
  INNER JOIN users u ON u.id = c.user_id
  INNER JOIN video v ON v.id = c.video_id
LIMIT $1 OFFSET $2;

-- name: GetCommentById :one
SELECT * FROM comments WHERE id=$1 LIMIT 1;

-- name: DeleteComment :exec
DELETE FROM comments  WHERE id = $1;

-- name: IfUserReactionThisVideo :one
SELECT * FROM video_reaction WHERE user_id = $1 AND video_id=$2 LIMIT 1;

-- name: AddVideoReaction :exec
INSERT INTO video_reaction (user_id, video_id, value)
VALUES ($1, $2, $3);

-- name: UpdateVideoReaction :exec
UPDATE video_reaction
SET value = $1, updated_at = NOW()
WHERE user_id =$2 AND video_id=$3;