package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type VideoPatchData struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	CategoryId  int64    `json:"category_id" validate:"required"`
	Tags        []string `json:"tags" validate:"required,tags"`
}

func PathVideo(c *fiber.Ctx) error {
	// Parse sent data
	var data VideoPatchData
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get video ID
	idString := c.Params("id")
	id, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = mod.Validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}

	// * CHECK AGANIST DATABASE
	// * START(DB BLOCK)
	ctx := context.Background()
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		log.WithField("err", err).Error("Could not create database connection")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	qtx := sqlc.New(db)
	defer db.Close()
	// * END(DB BLOCK)

	// Check if exists
	video, err := qtx.GetVideoByID(ctx, id)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get video from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("video_id", id).Debug("Video with specified id doesn't exist")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Update video
	if _, err := qtx.UpdateVideo(ctx, sqlc.UpdateVideoParams{
		ID:          video.ID,
		Name:        data.Name,
		Description: data.Description,
		CategoryID:  data.CategoryId,
	}); err != nil {
		log.WithField("video_id", id).Error("Problem with updating the video")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Remove tags from video
	if err := qtx.ClearVideoTags(ctx, video.ID); err != nil {
		log.WithField("video_id", id).Error("Tags could not be removed from video")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Update tags
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
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		} else if err != nil {
			log.WithField("err", err.Error()).Error("Problem getting tags from database")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		tagIds = append(tagIds, dbTag.ID)
	}

	// Attach tags to video
	for _, tagId := range tagIds {
		if err := qtx.AddVideoTag(ctx, sqlc.AddVideoTagParams{
			VideoID: video.ID,
			TagID:   tagId,
		}); err != nil {
			log.WithField("err", err.Error()).Error("Problem with attaching tag to video")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

// func addVideoAsync(data *VideoPostData) {
// 	log.WithField("data", *data).Debug("Video adding task started")
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	endpoint := config.GetString("APP_MINIO_S3_URL", "localhost:9000")
// 	accesskey := config.GetString("APP_MINIO_S3_ACCESSKEY", "")
// 	secretkey := config.GetString("APP_MINIO_S3_SECRETKEY", "")
// 	uploadBucket := config.GetString("APP_MINIO_S3_UPLOAD_BUCKET", "uploads")
// 	videoBucket := config.GetString("APP_MINIO_S3_VIDEO_BUCKET", "videos")
// 	thumbnailBucket := config.GetString("APP_MINIO_S3_THUMBNAIL_BUCKET", "thumbnails")
// 	useSSL := config.GetBool("APP_MINIO_S3_USESSL", true)
// 	storagePath := config.GetString("APP_STORAGE_PATH", "./storage")

// 	minioClient, err := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accesskey, secretkey, ""),
// 		Secure: useSSL,
// 	})
// 	if err != nil {
// 		log.WithField("err", err).Error("Minio module error")
// 		return
// 	}

// 	// array for files created throughout the task
// 	// defered function is certain to run before any return
// 	var filesToCleanup []string
// 	defer cleanFiles(&filesToCleanup)

// 	// Check if video exists
// 	uploads := minioClient.ListObjects(ctx, uploadBucket, minio.ListObjectsOptions{})
// 	found := false
// 	for vid := range uploads {
// 		if vid.Err != nil {
// 			log.WithField("err", vid.Err).Error("File error occured (minio s3)")
// 			return
// 		}
// 		if vid.Key == data.FileName {
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		log.WithFields(log.Fields{
// 			"file":   data.FileName,
// 			"bucket": uploadBucket,
// 		}).Error("Specified file couldn't be found in the bucket")
// 		return
// 	}

// 	// Download video to fs
// 	downloadedFilePath := fmt.Sprintf("%s/tmp/%s", storagePath, data.FileName)
// 	if err := minioClient.FGetObject(
// 		ctx,
// 		uploadBucket,
// 		data.FileName,
// 		downloadedFilePath,
// 		minio.GetObjectOptions{}); err != nil {

// 		log.WithField("err", err.Error()).Error("Could not download file from minio s3")
// 		return
// 	}
// 	filesToCleanup = append(filesToCleanup, downloadedFilePath)

// 	// Remove video from bucket
// 	if err := minioClient.RemoveObject(
// 		ctx,
// 		uploadBucket,
// 		data.FileName,
// 		minio.RemoveObjectOptions{}); err != nil {

// 		log.WithField("err", err.Error()).Error("Target file could not be removed from minio s3")
// 		return
// 	}

// 	// Check if thumbnail exists
// 	uploads = minioClient.ListObjects(ctx, uploadBucket, minio.ListObjectsOptions{})
// 	found = false
// 	for vid := range uploads {
// 		if vid.Err != nil {
// 			log.WithField("err", vid.Err).Error("File error occured (minio s3)")
// 			return
// 		}
// 		if vid.Key == data.ThumbnailName {
// 			found = true
// 			fmt.Println(vid.ContentType)
// 			break
// 		}
// 	}
// 	if !found {
// 		log.WithFields(log.Fields{
// 			"file":   data.ThumbnailName,
// 			"bucket": uploadBucket,
// 		}).Error("Specified file couldn't be found in the bucket")
// 		return
// 	}

// 	// Download thumbnail to fs
// 	downloadedThumbnailPath := fmt.Sprintf("%s/tmp/%s", storagePath, data.ThumbnailName)
// 	if err := minioClient.FGetObject(
// 		ctx,
// 		uploadBucket,
// 		data.ThumbnailName,
// 		downloadedThumbnailPath,
// 		minio.GetObjectOptions{}); err != nil {

// 		log.WithField("err", err.Error()).Error("Could not download thumbnail from minio s3")
// 		return
// 	}
// 	filesToCleanup = append(filesToCleanup, downloadedThumbnailPath)

// 	// Remove video from bucket
// 	if err := minioClient.RemoveObject(
// 		ctx,
// 		uploadBucket,
// 		data.FileName,
// 		minio.RemoveObjectOptions{}); err != nil {

// 		log.WithField("err", err.Error()).Error("Target file could not be removed from minio s3")
// 		return
// 	}

// 	// Remove thumbnail from bucket
// 	if err := minioClient.RemoveObject(
// 		ctx,
// 		uploadBucket,
// 		data.ThumbnailName,
// 		minio.RemoveObjectOptions{}); err != nil {

// 		log.WithField("err", err.Error()).Error("Target thumbnail could not be removed from minio s3")
// 		return
// 	}

// 	// Declare resolutions and unique fileId
// 	fileId := xid.New().String()
// 	resolutions := map[string]string{
// 		"360p": "480:360",
// 		"480p": "640:480",
// 		"720p": "1280x720",
// 	}

// 	// Mux files
// 	for res, ratio := range resolutions {
// 		localPath := fmt.Sprintf("%s/tmp/%s_%s.mp4", storagePath, fileId, res)
// 		if err := ffmpeg.Input(downloadedFilePath).
// 			Output(
// 				localPath,
// 				ffmpeg.KwArgs{
// 					"s":   ratio,
// 					"c:v": "libx265",
// 				}).
// 			OverWriteOutput().
// 			// ErrorToStdOut().
// 			Run(); err != nil {

// 			log.WithField("err", err.Error()).Error("FFMPEG returned an error")
// 			return
// 		}
// 		filesToCleanup = append(filesToCleanup, localPath)
// 	}

// 	// * START(DB BLOCK)
// 	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
// 	if err != nil {
// 		log.WithField("err", err.Error()).Error("Could not connect to database")
// 	}
// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.WithField("err", err.Error()).Error("Could not start transaction")
// 	}
// 	qtx := sqlc.New(db).WithTx(tx)
// 	defer tx.Rollback()
// 	defer db.Close()
// 	// * END(DB BLOCK)

// 	_, err = qtx.GetCategoryByID(ctx, data.CategoryId)
// 	if err == sql.ErrNoRows {
// 		data.CategoryId = -1
// 	} else if err != nil {
// 		log.WithField("err", err.Error()).Error("Error occured while checking categories list")
// 		return
// 	}

// 	// Add video to database
// 	vid, err := qtx.AddVideo(ctx, sqlc.AddVideoParams{
// 		Name:        data.Name,
// 		Description: data.Description,
// 		CategoryID:  data.CategoryId,
// 	})
// 	if err != nil {
// 		log.WithField("err", err.Error()).Error("Video couldn't be added to database")
// 		return
// 	}

// 	// Push thumbnail image
// 	guid := xid.New()
// 	thumbnailObjectName := guid.String() + data.ThumbnailName
// 	_, err = minioClient.FPutObject(
// 		ctx,
// 		thumbnailBucket,
// 		thumbnailObjectName,
// 		downloadedThumbnailPath,
// 		minio.PutObjectOptions{})
// 	if err != nil {
// 		log.WithField("err", err.Error()).Error("Thumbnail file couldn't be uploaded!")
// 		return
// 	}

// 	// Iterate through files
// 	for res := range resolutions {
// 		remotePath := fmt.Sprintf("%d/%s.mp4", vid.ID, res)
// 		localPath := fmt.Sprintf("%s/tmp/%s_%s.mp4", storagePath, fileId, res)
// 		// Upload video to S3
// 		_, err = minioClient.FPutObject(
// 			ctx,
// 			videoBucket,
// 			remotePath,
// 			localPath,
// 			minio.PutObjectOptions{
// 				ContentType: "video/mp4",
// 			})
// 		if err != nil {
// 			log.WithField("err", err.Error()).Error("Video file couldn't be uploaded!")
// 			return
// 		}
// 		file, _ := os.Open(localPath)
// 		fileInfo, err := file.Stat()
// 		if err != nil {
// 			log.WithField("err", err.Error()).Error("Could not read file")
// 			return
// 		}

// 		// Get video information
// 		jsonData, err := ffmpeg.Probe(downloadedFilePath)
// 		if err != nil {
// 			log.WithField("err", err.Error()).Error("Could not probe video file")
// 			return
// 		}

// 		var videoInfo MediaInfo
// 		err = json.Unmarshal([]byte(jsonData), &videoInfo)
// 		if err != nil {
// 			log.WithField("err", err.Error()).Error("Could not get stream information")
// 			return
// 		}
// 		duration, err := strconv.ParseFloat(videoInfo.Info.Duration, 64)
// 		if err != nil {
// 			log.WithField("err", err.Error()).Error("Could not parse FFMPEG duration")
// 			return
// 		}

// 		// Add to database
// 		if err := qtx.AddVideoFile(ctx, sqlc.AddVideoFileParams{
// 			FilePath:   remotePath,
// 			VideoID:    vid.ID,
// 			FileSize:   fileInfo.Size(),
// 			Duration:   int64(duration),
// 			Resolution: sqlc.Resolution(res),
// 		}); err != nil {
// 			log.WithField("err", err.Error()).Error("Video file couldn't be added to database")
// 			return
// 		}
// 	}

// 	// Trim tags and check if exists
// 	// Add to database if needed and get ids to add to video
// 	var tagIds []int64
// 	for _, tag := range data.Tags {
// 		tag = strings.TrimSpace(tag)
// 		tag = strings.ToLower(tag)

// 		dbTag, err := qtx.GetTagByName(ctx, tag)
// 		if err == sql.ErrNoRows {
// 			dbTag, err = qtx.AddTag(ctx, tag)
// 			if err != nil {
// 				log.WithField("err", err.Error()).Error("Could not add tag to database")
// 				return
// 			}
// 		} else if err != nil {
// 			log.WithField("err", err.Error()).Error("Problem getting tags from database")
// 			return
// 		}
// 		tagIds = append(tagIds, dbTag.ID)
// 	}

// 	// Attach tags to video
// 	for _, tagId := range tagIds {
// 		if err := qtx.AddVideoTag(ctx, sqlc.AddVideoTagParams{
// 			VideoID: vid.ID,
// 			TagID:   tagId,
// 		}); err != nil {
// 			log.WithField("err", err.Error()).Error("Problem with attaching tag to video")
// 			return
// 		}
// 	}

// 	// Add thumbnail to database
// 	// Add to database
// 	thumbnailFileInfo, err := os.Stat(downloadedFilePath)
// 	if err != nil {
// 		log.WithField("err", err.Error()).Error("Could not stat thumbnail file")
// 		return
// 	}
// 	if err := qtx.AddThumbnail(ctx, sqlc.AddThumbnailParams{
// 		VideoID:  vid.ID,
// 		FileSize: int32(thumbnailFileInfo.Size()),
// 		FileName: thumbnailObjectName,
// 	}); err != nil {
// 		log.WithField("err", err.Error()).Error("Video file couldn't be added to database")
// 		return
// 	}

// 	tx.Commit()
// 	log.Info("Video was successfully added to database")
// }

// func cleanFiles(filesToClean *[]string) {
// 	for _, file := range *filesToClean {
// 		if _, err := os.Stat(file); err == nil {
// 			err := os.Remove(file)
// 			if err != nil {
// 				log.WithField("err", err.Error()).Error("File could not be removed")
// 			}
// 		}
// 	}
// }
