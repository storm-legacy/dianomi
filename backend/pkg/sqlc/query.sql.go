// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
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
RETURNING id, name, description, category_id, upvotes, downvotes, views, updated_at, created_at, deleted_at
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

const createUser = `-- name: CreateUser :exec
INSERT INTO users (email, password) VALUES ($1, $2)
`

type CreateUserParams struct {
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Email, arg.Password)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
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

const getAllVideos = `-- name: GetAllVideos :many
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

const getPackagesByUserID = `-- name: GetPackagesByUserID :many
SELECT id, user_id, tier, created_at, valid_from, valid_until FROM
  users_packages
WHERE
  user_id = $1
  AND
  (now()::timestamp with TIME ZONE) BETWEEN valid_from AND valid_until
ORDER BY
  created_at DESC
LIMIT 10
`

func (q *Queries) GetPackagesByUserID(ctx context.Context, userID int64) ([]UsersPackage, error) {
	rows, err := q.db.QueryContext(ctx, getPackagesByUserID, userID)
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

const getVideoByID = `-- name: GetVideoByID :one
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
	UserID     int64
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

const removeAllUserPackages = `-- name: RemoveAllUserPackages :exec
DELETE FROM users_packages WHERE user_id = $1
`

func (q *Queries) RemoveAllUserPackages(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, removeAllUserPackages, userID)
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
