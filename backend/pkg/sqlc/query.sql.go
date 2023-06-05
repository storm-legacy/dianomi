// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addCategory = `-- name: AddCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING id, name
`

func (q *Queries) AddCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRowContext(ctx, addCategory, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const addReport = `-- name: AddReport :exec
INSERT INTO Error_Reports (error_title, error_description, reported_by) VALUES ($1, $2, $3) RETURNING id, error_title, error_description, reported_by, report_date
`

type AddReportParams struct {
	ErrorTitle       string
	ErrorDescription string
	ReportedBy       string
}

func (q *Queries) AddReport(ctx context.Context, arg AddReportParams) error {
	_, err := q.db.ExecContext(ctx, addReport, arg.ErrorTitle, arg.ErrorDescription, arg.ReportedBy)
	return err
}

const addTag = `-- name: AddTag :one
INSERT INTO tags (name) VALUES ($1) RETURNING id, name
`

func (q *Queries) AddTag(ctx context.Context, name string) (Tag, error) {
	row := q.db.QueryRowContext(ctx, addTag, name)
	var i Tag
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const addThumbnail = `-- name: AddThumbnail :exec
INSERT INTO video_thumbnails (
  video_id,
  file_size,
  file_name
) VALUES ($1, $2, $3)
`

type AddThumbnailParams struct {
	VideoID  int64
	FileSize int32
	FileName string
}

func (q *Queries) AddThumbnail(ctx context.Context, arg AddThumbnailParams) error {
	_, err := q.db.ExecContext(ctx, addThumbnail, arg.VideoID, arg.FileSize, arg.FileName)
	return err
}

const addVideo = `-- name: AddVideo :one
INSERT INTO video (
  name,
  description,
  category_id
) VALUES ($1, $2, $3)
RETURNING id, name, description, category_id, upvotes, downvotes, views, is_premium, updated_at, created_at, deleted_at
`

type AddVideoParams struct {
	Name        string
	Description string
	CategoryID  int64
}

func (q *Queries) AddVideo(ctx context.Context, arg AddVideoParams) (Video, error) {
	row := q.db.QueryRowContext(ctx, addVideo, arg.Name, arg.Description, arg.CategoryID)
	var i Video
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CategoryID,
		&i.Upvotes,
		&i.Downvotes,
		&i.Views,
		&i.IsPremium,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const addVideoFile = `-- name: AddVideoFile :exec
INSERT INTO video_files (
  file_path,
  video_id,
  file_size,
  duration,
  resolution
) VALUES ($1, $2, $3, $4, $5)
`

type AddVideoFileParams struct {
	FilePath   string
	VideoID    int64
	FileSize   int64
	Duration   int64
	Resolution Resolution
}

func (q *Queries) AddVideoFile(ctx context.Context, arg AddVideoFileParams) error {
	_, err := q.db.ExecContext(ctx, addVideoFile,
		arg.FilePath,
		arg.VideoID,
		arg.FileSize,
		arg.Duration,
		arg.Resolution,
	)
	return err
}

const addVideoMertics = `-- name: AddVideoMertics :exec
INSERT INTO user_video_metrics (user_id, video_id, time_spent_watching, stopped_at, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id, user_id, video_id, time_spent_watching, stopped_at, created_at, updated_at
`

type AddVideoMerticsParams struct {
	UserID            int64
	VideoID           int64
	TimeSpentWatching int32
	StoppedAt         int32
}

func (q *Queries) AddVideoMertics(ctx context.Context, arg AddVideoMerticsParams) error {
	_, err := q.db.ExecContext(ctx, addVideoMertics,
		arg.UserID,
		arg.VideoID,
		arg.TimeSpentWatching,
		arg.StoppedAt,
	)
	return err
}

const addVideoTag = `-- name: AddVideoTag :exec
INSERT INTO video_tags (video_id, tag_id) VALUES ($1, $2) RETURNING id, video_id, tag_id
`

type AddVideoTagParams struct {
	VideoID int64
	TagID   int64
}

func (q *Queries) AddVideoTag(ctx context.Context, arg AddVideoTagParams) error {
	_, err := q.db.ExecContext(ctx, addVideoTag, arg.VideoID, arg.TagID)
	return err
}

const addVideoThumbnail = `-- name: AddVideoThumbnail :exec
INSERT INTO video_thumbnails (
  video_id,
  file_size
) VALUES ($1, $2)
`

type AddVideoThumbnailParams struct {
	VideoID  int64
	FileSize int32
}

func (q *Queries) AddVideoThumbnail(ctx context.Context, arg AddVideoThumbnailParams) error {
	_, err := q.db.ExecContext(ctx, addVideoThumbnail, arg.VideoID, arg.FileSize)
	return err
}

const clearVideoTags = `-- name: ClearVideoTags :exec
DELETE FROM video_tags WHERE video_id = $1
`

func (q *Queries) ClearVideoTags(ctx context.Context, videoID int64) error {
	_, err := q.db.ExecContext(ctx, clearVideoTags, videoID)
	return err
}

const createResetCode = `-- name: CreateResetCode :one
INSERT INTO verification (
  user_id,
  task_type
) VALUES ($1, 'passwordReset')
RETURNING id, user_id, task_type, code, used, created_at, valid_until
`

func (q *Queries) CreateResetCode(ctx context.Context, userID sql.NullInt64) (Verification, error) {
	row := q.db.QueryRowContext(ctx, createResetCode, userID)
	var i Verification
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TaskType,
		&i.Code,
		&i.Used,
		&i.CreatedAt,
		&i.ValidUntil,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password, verified_at, created_at, updated_at
`

type CreateUserParams struct {
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createVerificationCode = `-- name: CreateVerificationCode :one
INSERT INTO verification (
  user_id,
  task_type
) VALUES ($1, 'emailVerification')
RETURNING id, user_id, task_type, code, used, created_at, valid_until
`

func (q *Queries) CreateVerificationCode(ctx context.Context, userID sql.NullInt64) (Verification, error) {
	row := q.db.QueryRowContext(ctx, createVerificationCode, userID)
	var i Verification
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TaskType,
		&i.Code,
		&i.Used,
		&i.CreatedAt,
		&i.ValidUntil,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const deleteVideo = `-- name: DeleteVideo :exec
DELETE FROM video WHERE id = $1
`

func (q *Queries) DeleteVideo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteVideo, id)
	return err
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT id, name FROM categories ORDER BY name ASC LIMIT $1 OFFSET $2
`

type GetAllCategoriesParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllCategories(ctx context.Context, arg GetAllCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getAllCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, email, password, verified_at, created_at, updated_at
FROM users
LIMIT $1
OFFSET $2
`

type GetAllUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.VerifiedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllVideoMetric = `-- name: GetAllVideoMetric :many
SELECT id, user_id, video_id, time_spent_watching, stopped_at, created_at, updated_at FROM user_video_metrics LIMIT $1 OFFSET $2
`

type GetAllVideoMetricParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetAllVideoMetric(ctx context.Context, arg GetAllVideoMetricParams) ([]UserVideoMetric, error) {
	rows, err := q.db.QueryContext(ctx, getAllVideoMetric, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserVideoMetric
	for rows.Next() {
		var i UserVideoMetric
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.VideoID,
			&i.TimeSpentWatching,
			&i.StoppedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllVideos = `-- name: GetAllVideos :many
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
OFFSET $2
`

type GetAllVideosParams struct {
	Limit  int32
	Offset int32
}

type GetAllVideosRow struct {
	ID          int64
	Name        string
	Description string
	Category    sql.NullString
	Upvotes     int64
	Downvotes   int64
	Views       int64
	IsPremium   bool
	Thumbnail   sql.NullString
}

func (q *Queries) GetAllVideos(ctx context.Context, arg GetAllVideosParams) ([]GetAllVideosRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllVideos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllVideosRow
	for rows.Next() {
		var i GetAllVideosRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.Upvotes,
			&i.Downvotes,
			&i.Views,
			&i.IsPremium,
			&i.Thumbnail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryByID = `-- name: GetCategoryByID :one
SELECT id, name FROM categories WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategoryByID(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryByID, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getCategoryByName = `-- name: GetCategoryByName :one
SELECT id, name FROM categories WHERE name = $1 LIMIT 1
`

func (q *Queries) GetCategoryByName(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryByName, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getCurrentPackageForUser = `-- name: GetCurrentPackageForUser :one
SELECT id, user_id, tier, created_at, valid_from, valid_until FROM
  users_packages
WHERE
  user_id = $1
  AND
  (now()::timestamp with TIME ZONE) BETWEEN valid_from AND valid_until
LIMIT 1
`

func (q *Queries) GetCurrentPackageForUser(ctx context.Context, userID sql.NullInt64) (UsersPackage, error) {
	row := q.db.QueryRowContext(ctx, getCurrentPackageForUser, userID)
	var i UsersPackage
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Tier,
		&i.CreatedAt,
		&i.ValidFrom,
		&i.ValidUntil,
	)
	return i, err
}

const getOverlapingPackages = `-- name: GetOverlapingPackages :many
SELECT id, user_id, tier, created_at, valid_from, valid_until FROM
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
LIMIT 5
`

type GetOverlapingPackagesParams struct {
	ValidFrom  time.Time
	ValidUntil time.Time
	UserID     sql.NullInt64
}

func (q *Queries) GetOverlapingPackages(ctx context.Context, arg GetOverlapingPackagesParams) ([]UsersPackage, error) {
	rows, err := q.db.QueryContext(ctx, getOverlapingPackages, arg.ValidFrom, arg.ValidUntil, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersPackage
	for rows.Next() {
		var i UsersPackage
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Tier,
			&i.CreatedAt,
			&i.ValidFrom,
			&i.ValidUntil,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPackageByID = `-- name: GetPackageByID :one
SELECT id, user_id, tier, created_at, valid_from, valid_until FROM
  users_packages
WHERE
  id = $1
LIMIT 1
`

func (q *Queries) GetPackageByID(ctx context.Context, id int64) (UsersPackage, error) {
	row := q.db.QueryRowContext(ctx, getPackageByID, id)
	var i UsersPackage
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Tier,
		&i.CreatedAt,
		&i.ValidFrom,
		&i.ValidUntil,
	)
	return i, err
}

const getPackagesByUserID = `-- name: GetPackagesByUserID :many
SELECT id, user_id, tier, created_at, valid_from, valid_until FROM
  users_packages
WHERE
  user_id = $1
  AND
  (now()::timestamp with TIME ZONE) BETWEEN valid_from AND valid_until
ORDER BY
  created_at DESC
LIMIT $2
OFFSET $3
`

type GetPackagesByUserIDParams struct {
	UserID sql.NullInt64
	Limit  int32
	Offset int32
}

func (q *Queries) GetPackagesByUserID(ctx context.Context, arg GetPackagesByUserIDParams) ([]UsersPackage, error) {
	rows, err := q.db.QueryContext(ctx, getPackagesByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersPackage
	for rows.Next() {
		var i UsersPackage
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Tier,
			&i.CreatedAt,
			&i.ValidFrom,
			&i.ValidUntil,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRandomVideos = `-- name: GetRandomVideos :many
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
OFFSET $2
`

type GetRandomVideosParams struct {
	Limit  int32
	Offset int32
}

type GetRandomVideosRow struct {
	ID          int64
	Name        string
	Description string
	Category    sql.NullString
	Upvotes     int64
	Downvotes   int64
	Views       int64
	IsPremium   bool
	Thumbnail   sql.NullString
}

func (q *Queries) GetRandomVideos(ctx context.Context, arg GetRandomVideosParams) ([]GetRandomVideosRow, error) {
	rows, err := q.db.QueryContext(ctx, getRandomVideos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRandomVideosRow
	for rows.Next() {
		var i GetRandomVideosRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.Upvotes,
			&i.Downvotes,
			&i.Views,
			&i.IsPremium,
			&i.Thumbnail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTagByName = `-- name: GetTagByName :one
SELECT id, name FROM tags WHERE name = $1 LIMIT 1
`

func (q *Queries) GetTagByName(ctx context.Context, name string) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTagByName, name)
	var i Tag
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, verified_at, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, password, verified_at, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserVideoMerticsByUserId = `-- name: GetUserVideoMerticsByUserId :many
SELECT user_video_metrics.id, user_video_metrics.user_id, user_video_metrics.video_id, user_video_metrics.time_spent_watching, user_video_metrics.stopped_at, user_video_metrics.created_at, user_video_metrics.updated_at, video.id, video.name, video.description, video.is_premium, video_thumbnails.file_name
FROM user_video_metrics
JOIN video ON user_video_metrics.video_id = video.id JOIN video_thumbnails ON video_thumbnails.video_id = video.id
WHERE user_video_metrics.user_id = $1
`

type GetUserVideoMerticsByUserIdRow struct {
	ID                int64
	UserID            int64
	VideoID           int64
	TimeSpentWatching int32
	StoppedAt         int32
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
	ID_2              int64
	Name              string
	Description       string
	IsPremium         bool
	FileName          string
}

func (q *Queries) GetUserVideoMerticsByUserId(ctx context.Context, userID int64) ([]GetUserVideoMerticsByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserVideoMerticsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserVideoMerticsByUserIdRow
	for rows.Next() {
		var i GetUserVideoMerticsByUserIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.VideoID,
			&i.TimeSpentWatching,
			&i.StoppedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.Name,
			&i.Description,
			&i.IsPremium,
			&i.FileName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserVideoMerticsByVideoId = `-- name: GetUserVideoMerticsByVideoId :one
SELECT id, user_id, video_id, time_spent_watching, stopped_at, created_at, updated_at FROM user_video_metrics WHERE video_id = $1 LIMIT 1
`

func (q *Queries) GetUserVideoMerticsByVideoId(ctx context.Context, videoID int64) (UserVideoMetric, error) {
	row := q.db.QueryRowContext(ctx, getUserVideoMerticsByVideoId, videoID)
	var i UserVideoMetric
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.VideoID,
		&i.TimeSpentWatching,
		&i.StoppedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getVerificationCode = `-- name: GetVerificationCode :one
SELECT id, user_id, task_type, code, used, created_at, valid_until FROM verification WHERE code = $1 LIMIT 1
`

func (q *Queries) GetVerificationCode(ctx context.Context, code uuid.UUID) (Verification, error) {
	row := q.db.QueryRowContext(ctx, getVerificationCode, code)
	var i Verification
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TaskType,
		&i.Code,
		&i.Used,
		&i.CreatedAt,
		&i.ValidUntil,
	)
	return i, err
}

const getVideoByID = `-- name: GetVideoByID :one
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
LIMIT 1
`

type GetVideoByIDRow struct {
	ID          int64
	Name        string
	Description string
	Category    sql.NullString
	Upvotes     int64
	Downvotes   int64
	Views       int64
	IsPremium   bool
	Thumbnail   sql.NullString
}

func (q *Queries) GetVideoByID(ctx context.Context, id int64) (GetVideoByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getVideoByID, id)
	var i GetVideoByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Category,
		&i.Upvotes,
		&i.Downvotes,
		&i.Views,
		&i.IsPremium,
		&i.Thumbnail,
	)
	return i, err
}

const getVideoFiles = `-- name: GetVideoFiles :many
SELECT
  resolution,
  duration,
  file_path
FROM
  video_files
WHERE
  video_id = $1
`

type GetVideoFilesRow struct {
	Resolution Resolution
	Duration   int64
	FilePath   string
}

func (q *Queries) GetVideoFiles(ctx context.Context, videoID int64) ([]GetVideoFilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getVideoFiles, videoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVideoFilesRow
	for rows.Next() {
		var i GetVideoFilesRow
		if err := rows.Scan(&i.Resolution, &i.Duration, &i.FilePath); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVideoTags = `-- name: GetVideoTags :many
SELECT
  name
FROM
  tags t INNER JOIN video_tags vt ON t.id = vt.tag_id
WHERE
  vt.video_id = $1
LIMIT 10
`

func (q *Queries) GetVideoTags(ctx context.Context, videoID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getVideoTags, videoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVideosByCategory = `-- name: GetVideosByCategory :many
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
OFFSET $3
`

type GetVideosByCategoryParams struct {
	ID     int64
	Limit  int32
	Offset int32
}

type GetVideosByCategoryRow struct {
	ID          int64
	Name        string
	Description string
	Category    sql.NullString
	Upvotes     int64
	Downvotes   int64
	Views       int64
	IsPremium   bool
	Thumbnail   sql.NullString
}

func (q *Queries) GetVideosByCategory(ctx context.Context, arg GetVideosByCategoryParams) ([]GetVideosByCategoryRow, error) {
	rows, err := q.db.QueryContext(ctx, getVideosByCategory, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVideosByCategoryRow
	for rows.Next() {
		var i GetVideosByCategoryRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.Upvotes,
			&i.Downvotes,
			&i.Views,
			&i.IsPremium,
			&i.Thumbnail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVideosByName = `-- name: GetVideosByName :many
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
OFFSET $3
`

type GetVideosByNameParams struct {
	Name   string
	Limit  int32
	Offset int32
}

type GetVideosByNameRow struct {
	ID          int64
	Name        string
	Description string
	Category    sql.NullString
	Upvotes     int64
	Downvotes   int64
	Views       int64
	IsPremium   bool
	Thumbnail   sql.NullString
}

func (q *Queries) GetVideosByName(ctx context.Context, arg GetVideosByNameParams) ([]GetVideosByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getVideosByName, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetVideosByNameRow
	for rows.Next() {
		var i GetVideosByNameRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Category,
			&i.Upvotes,
			&i.Downvotes,
			&i.Views,
			&i.IsPremium,
			&i.Thumbnail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const giveUserPackage = `-- name: GiveUserPackage :exec
INSERT INTO users_packages (
  user_id,
  tier,
  valid_from,
  valid_until
) VALUES (
  $1, $2, $3, $4
)
`

type GiveUserPackageParams struct {
	UserID     sql.NullInt64
	Tier       Role
	ValidFrom  time.Time
	ValidUntil time.Time
}

func (q *Queries) GiveUserPackage(ctx context.Context, arg GiveUserPackageParams) error {
	_, err := q.db.ExecContext(ctx, giveUserPackage,
		arg.UserID,
		arg.Tier,
		arg.ValidFrom,
		arg.ValidUntil,
	)
	return err
}

const ifUserSeeThisVideo = `-- name: IfUserSeeThisVideo :one
SELECT id, user_id, video_id, time_spent_watching, stopped_at, created_at, updated_at FROM user_video_metrics WHERE user_id = $1 AND video_id=$2 LIMIT 1
`

type IfUserSeeThisVideoParams struct {
	UserID  int64
	VideoID int64
}

func (q *Queries) IfUserSeeThisVideo(ctx context.Context, arg IfUserSeeThisVideoParams) (UserVideoMetric, error) {
	row := q.db.QueryRowContext(ctx, ifUserSeeThisVideo, arg.UserID, arg.VideoID)
	var i UserVideoMetric
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.VideoID,
		&i.TimeSpentWatching,
		&i.StoppedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const removeAllUserPackages = `-- name: RemoveAllUserPackages :exec
DELETE FROM users_packages WHERE user_id = $1
`

func (q *Queries) RemoveAllUserPackages(ctx context.Context, userID sql.NullInt64) error {
	_, err := q.db.ExecContext(ctx, removeAllUserPackages, userID)
	return err
}

const removeUserPackage = `-- name: RemoveUserPackage :exec
DELETE FROM users_packages WHERE id = $1
`

func (q *Queries) RemoveUserPackage(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, removeUserPackage, id)
	return err
}

const setCodeAsUsed = `-- name: SetCodeAsUsed :exec
UPDATE verification SET used = true WHERE id = $1
`

func (q *Queries) SetCodeAsUsed(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, setCodeAsUsed, id)
	return err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories SET name = $1 WHERE id = $2
`

type UpdateCategoryParams struct {
	Name string
	ID   int64
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.Name, arg.ID)
	return err
}

const updatePackage = `-- name: UpdatePackage :exec
UPDATE
  users_packages
SET
  user_id = $2,
  tier = $3,
  valid_from = $4,
  valid_until = $5
WHERE
  id = $1
`

type UpdatePackageParams struct {
	ID         int64
	UserID     sql.NullInt64
	Tier       Role
	ValidFrom  time.Time
	ValidUntil time.Time
}

func (q *Queries) UpdatePackage(ctx context.Context, arg UpdatePackageParams) error {
	_, err := q.db.ExecContext(ctx, updatePackage,
		arg.ID,
		arg.UserID,
		arg.Tier,
		arg.ValidFrom,
		arg.ValidUntil,
	)
	return err
}

const updateUserEmail = `-- name: UpdateUserEmail :exec
UPDATE users SET email = $2 WHERE id = $1
`

type UpdateUserEmailParams struct {
	ID    int64
	Email string
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateUserEmail, arg.ID, arg.Email)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID       int64
	Password string
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPassword, arg.ID, arg.Password)
	return err
}

const updateUserVerification = `-- name: UpdateUserVerification :exec
UPDATE users SET verified_at = $2 WHERE id = $1
`

type UpdateUserVerificationParams struct {
	ID         int64
	VerifiedAt sql.NullTime
}

func (q *Queries) UpdateUserVerification(ctx context.Context, arg UpdateUserVerificationParams) error {
	_, err := q.db.ExecContext(ctx, updateUserVerification, arg.ID, arg.VerifiedAt)
	return err
}

const updateVideo = `-- name: UpdateVideo :one
UPDATE video
SET
  name = $2,
  description = $3,
  category_id = $4,
  is_premium = COALESCE($5, false)
WHERE id = $1
RETURNING id, name, description, category_id, upvotes, downvotes, views, is_premium, updated_at, created_at, deleted_at
`

type UpdateVideoParams struct {
	ID          int64
	Name        string
	Description string
	CategoryID  int64
	IsPremium   bool
}

func (q *Queries) UpdateVideo(ctx context.Context, arg UpdateVideoParams) (Video, error) {
	row := q.db.QueryRowContext(ctx, updateVideo,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CategoryID,
		arg.IsPremium,
	)
	var i Video
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CategoryID,
		&i.Upvotes,
		&i.Downvotes,
		&i.Views,
		&i.IsPremium,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateVideoMetric = `-- name: UpdateVideoMetric :exec
UPDATE user_video_metrics
SET time_spent_watching=time_spent_watching + $1, stopped_at = $2, updated_at = now()
WHERE id=$3
`

type UpdateVideoMetricParams struct {
	TimeSpentWatching int32
	StoppedAt         int32
	ID                int64
}

func (q *Queries) UpdateVideoMetric(ctx context.Context, arg UpdateVideoMetricParams) error {
	_, err := q.db.ExecContext(ctx, updateVideoMetric, arg.TimeSpentWatching, arg.StoppedAt, arg.ID)
	return err
}

const verifyUser = `-- name: VerifyUser :exec
UPDATE users SET verified_at = now() WHERE id = $1
`

func (q *Queries) VerifyUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, verifyUser, id)
	return err
}
