package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type Format struct {
	Duration string `json:"duration"`
}

type MediaInfo struct {
	Info Format `json:"format"`
}

type VideoPostData struct {
	Name          string   `json:"name" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	FileName      string   `json:"file_name" validate:"required"`
	ThumbnailName string   `json:"thumbnail_name" validate:"required"`
	FileBucket    string   `json:"file_bucket" validate:"required"`
	CategoryId    int64    `json:"category_id" validate:"required"`
	Tags          []string `json:"tags" validate:"required,tags"`
}

func PostVideo(c *fiber.Ctx) error {
	var data VideoPostData
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := mod.Validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}

	go addVideoAsync(&data)
	return c.SendStatus(fiber.StatusAccepted)
}

func addVideoAsync(data *VideoPostData) {
	log.WithField("data", *data).Debug("Video adding task started")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	endpoint := config.GetString("APP_MINIO_S3_URL", "localhost:9000")
	accesskey := config.GetString("APP_MINIO_S3_ACCESSKEY", "")
	secretkey := config.GetString("APP_MINIO_S3_SECRETKEY", "")
	uploadBucket := config.GetString("APP_MINIO_S3_UPLOAD_BUCKET", "uploads")
	videoBucket := config.GetString("APP_MINIO_S3_VIDEO_BUCKET", "videos")
	thumbnailBucket := config.GetString("APP_MINIO_S3_THUMBNAIL_BUCKET", "thumbnails")
	useSSL := config.GetBool("APP_MINIO_S3_USESSL", true)
	storagePath := config.GetString("APP_STORAGE_PATH", "./storage")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accesskey, secretkey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.WithField("err", err).Error("Minio module error")
		return
	}

	// array for files created throughout the task
	// defered function is certain to run before any return
	var filesToCleanup []string
	defer cleanFiles(&filesToCleanup)

	// Check if video exists
	uploads := minioClient.ListObjects(ctx, uploadBucket, minio.ListObjectsOptions{})
	found := false
	for vid := range uploads {
		if vid.Err != nil {
			log.WithField("err", vid.Err).Error("File error occured (minio s3)")
			return
		}
		if vid.Key == data.FileName {
			found = true
			break
		}
	}
	if !found {
		log.WithFields(log.Fields{
			"file":   data.FileName,
			"bucket": uploadBucket,
		}).Error("Specified file couldn't be found in the bucket")
		return
	}

	// Download video to fs
	downloadedFilePath := fmt.Sprintf("%s/tmp/%s", storagePath, data.FileName)
	if err := minioClient.FGetObject(
		ctx,
		uploadBucket,
		data.FileName,
		downloadedFilePath,
		minio.GetObjectOptions{}); err != nil {

		log.WithField("err", err.Error()).Error("Could not download file from minio s3")
		return
	}
	filesToCleanup = append(filesToCleanup, downloadedFilePath)

	// Remove video from bucket
	if err := minioClient.RemoveObject(
		ctx,
		uploadBucket,
		data.FileName,
		minio.RemoveObjectOptions{}); err != nil {

		log.WithField("err", err.Error()).Error("Target file could not be removed from minio s3")
		return
	}

	// Check if thumbnail exists
	uploads = minioClient.ListObjects(ctx, uploadBucket, minio.ListObjectsOptions{})
	found = false
	for vid := range uploads {
		if vid.Err != nil {
			log.WithField("err", vid.Err).Error("File error occured (minio s3)")
			return
		}
		if vid.Key == data.ThumbnailName {
			found = true
			fmt.Println(vid.ContentType)
			break
		}
	}
	if !found {
		log.WithFields(log.Fields{
			"file":   data.ThumbnailName,
			"bucket": uploadBucket,
		}).Error("Specified file couldn't be found in the bucket")
		return
	}

	// Download thumbnail to fs
	downloadedThumbnailPath := fmt.Sprintf("%s/tmp/%s", storagePath, data.ThumbnailName)
	if err := minioClient.FGetObject(
		ctx,
		uploadBucket,
		data.ThumbnailName,
		downloadedThumbnailPath,
		minio.GetObjectOptions{}); err != nil {

		log.WithField("err", err.Error()).Error("Could not download thumbnail from minio s3")
		return
	}
	filesToCleanup = append(filesToCleanup, downloadedThumbnailPath)

	// Remove video from bucket
	if err := minioClient.RemoveObject(
		ctx,
		uploadBucket,
		data.FileName,
		minio.RemoveObjectOptions{}); err != nil {

		log.WithField("err", err.Error()).Error("Target file could not be removed from minio s3")
		return
	}

	// Remove thumbnail from bucket
	if err := minioClient.RemoveObject(
		ctx,
		uploadBucket,
		data.ThumbnailName,
		minio.RemoveObjectOptions{}); err != nil {

		log.WithField("err", err.Error()).Error("Target thumbnail could not be removed from minio s3")
		return
	}

	// Declare resolutions and unique fileId
	fileId := xid.New().String()
	resolutions := map[string]string{
		"360p": "480:360",
		"480p": "640:480",
		"720p": "1280x720",
	}

	// Mux files
	for res, ratio := range resolutions {
		localPath := fmt.Sprintf("%s/tmp/%s_%s.mp4", storagePath, fileId, res)
		if err := ffmpeg.Input(downloadedFilePath).
			Output(
				localPath,
				ffmpeg.KwArgs{
					"s":   ratio,
					"c:v": "libx265",
				}).
			OverWriteOutput().
			// ErrorToStdOut().
			Run(); err != nil {

			log.WithField("err", err.Error()).Error("FFMPEG returned an error")
			return
		}
		filesToCleanup = append(filesToCleanup, localPath)
	}

	// * START(DB BLOCK)
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		log.WithField("err", err.Error()).Error("Could not connect to database")
	}
	tx, err := db.Begin()
	if err != nil {
		log.WithField("err", err.Error()).Error("Could not start transaction")
	}
	qtx := sqlc.New(db).WithTx(tx)
	defer tx.Rollback()
	defer db.Close()
	// * END(DB BLOCK)

	_, err = qtx.GetCategoryByID(ctx, data.CategoryId)
	if err == sql.ErrNoRows {
		data.CategoryId = -1
	} else if err != nil {
		log.WithField("err", err.Error()).Error("Error occured while checking categories list")
		return
	}

	// Add video to database
	vid, err := qtx.AddVideo(ctx, sqlc.AddVideoParams{
		Name:        data.Name,
		Description: data.Description,
		CategoryID:  data.CategoryId,
	})
	if err != nil {
		log.WithField("err", err.Error()).Error("Video couldn't be added to database")
		return
	}

	// Push thumbnail image
	guid := xid.New()
	thumbnailObjectName := guid.String() + data.ThumbnailName
	_, err = minioClient.FPutObject(
		ctx,
		thumbnailBucket,
		thumbnailObjectName,
		downloadedThumbnailPath,
		minio.PutObjectOptions{})
	if err != nil {
		log.WithField("err", err.Error()).Error("Thumbnail file couldn't be uploaded!")
		return
	}

	// Iterate through files
	for res := range resolutions {
		remotePath := fmt.Sprintf("%d/%s.mp4", vid.ID, res)
		localPath := fmt.Sprintf("%s/tmp/%s_%s.mp4", storagePath, fileId, res)
		// Upload video to S3
		_, err = minioClient.FPutObject(
			ctx,
			videoBucket,
			remotePath,
			localPath,
			minio.PutObjectOptions{
				ContentType: "video/mp4",
			})
		if err != nil {
			log.WithField("err", err.Error()).Error("Video file couldn't be uploaded!")
			return
		}
		file, _ := os.Open(localPath)
		fileInfo, err := file.Stat()
		if err != nil {
			log.WithField("err", err.Error()).Error("Could not read file")
			return
		}

		// Get video information
		jsonData, err := ffmpeg.Probe(downloadedFilePath)
		if err != nil {
			log.WithField("err", err.Error()).Error("Could not probe video file")
			return
		}

		var videoInfo MediaInfo
		err = json.Unmarshal([]byte(jsonData), &videoInfo)
		if err != nil {
			log.WithField("err", err.Error()).Error("Could not get stream information")
			return
		}
		duration, err := strconv.ParseFloat(videoInfo.Info.Duration, 64)
		if err != nil {
			log.WithField("err", err.Error()).Error("Could not parse FFMPEG duration")
			return
		}

		// Add to database
		if err := qtx.AddVideoFile(ctx, sqlc.AddVideoFileParams{
			FilePath:   remotePath,
			VideoID:    vid.ID,
			FileSize:   fileInfo.Size(),
			Duration:   int64(duration),
			Resolution: sqlc.Resolution(res),
		}); err != nil {
			log.WithField("err", err.Error()).Error("Video file couldn't be added to database")
			return
		}
	}

	// Trim tags and check if exists
	// Add to database if needed and get ids to add to video
	var tagIds []int64
	for _, tag := range data.Tags {
		tag = strings.TrimSpace(tag)
		tag = strings.ToLower(tag)

		dbTag, err := qtx.GetTagByName(ctx, tag)
		if err == sql.ErrNoRows {
			dbTag, err = qtx.AddTag(ctx, tag)
			if err != nil {
				log.WithField("err", err.Error()).Error("Could not add tag to database")
				return
			}
		} else if err != nil {
			log.WithField("err", err.Error()).Error("Problem getting tags from database")
			return
		}
		tagIds = append(tagIds, dbTag.ID)
	}

	// Attach tags to video
	for _, tagId := range tagIds {
		if err := qtx.AddVideoTag(ctx, sqlc.AddVideoTagParams{
			VideoID: vid.ID,
			TagID:   tagId,
		}); err != nil {
			log.WithField("err", err.Error()).Error("Problem with attaching tag to video")
			return
		}
	}

	// Add thumbnail to database
	// Add to database
	thumbnailFileInfo, err := os.Stat(downloadedFilePath)
	if err != nil {
		log.WithField("err", err.Error()).Error("Could not stat thumbnail file")
		return
	}
	if err := qtx.AddThumbnail(ctx, sqlc.AddThumbnailParams{
		VideoID:  vid.ID,
		FileSize: int32(thumbnailFileInfo.Size()),
		FileName: thumbnailObjectName,
	}); err != nil {
		log.WithField("err", err.Error()).Error("Video file couldn't be added to database")
		return
	}

	tx.Commit()
	log.Info("Video was successfully added to database")
}

func cleanFiles(filesToClean *[]string) {
	for _, file := range *filesToClean {
		if _, err := os.Stat(file); err == nil {
			err := os.Remove(file)
			if err != nil {
				log.WithField("err", err.Error()).Error("File could not be removed")
			}
		}
	}
}
