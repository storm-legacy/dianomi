package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
	"golang.org/x/exp/slices"
)

type VideoFiles struct {
	Resolution string `json:"resolution"`
	Duration   uint64 `json:"duration"`
	FilePath   string `json:"file_path"`
}

type VideoResponse struct {
	ID           uint64       `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Category     string       `json:"category"`
	CategoryID   uint64       `json:"category_id"`
	Upvotes      uint64       `json:"upvotes"`
	Downvotes    uint64       `json:"downvotes"`
	Views        uint64       `json:"views"`
	IsPremium    bool         `json:"IsPremium"`
	ThumbnailUrl string       `json:"thumbnail_url"`
	Files        []VideoFiles `json:"videos"`
	Tags         []string     `json:"tags"`
}

func VideoSearch(c *fiber.Ctx) error {
	searchString := c.Query("phrase")
	search := strings.Fields(searchString)

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

	// Look for video with phrase
	var results []int64
	for _, key := range search {
		items, err := qtx.GetVideoIDByName(ctx, fmt.Sprintf("%%%s%%", key))
		if err != nil {
			log.WithField("err", err).Error("Could not get data from database")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		// Add missing items
		for _, item := range items {
			if !slices.Contains(results, item) {
				results = append(results, item)
			}
		}
	}

	// Get specified videos
	videos := make([]VideoResponse, 0)
	for _, id := range results {
		// Check if exists
		vid, err := qtx.GetVideoByID(ctx, id)
		if err != sql.ErrNoRows && err != nil {
			log.WithField("err", err.Error()).Error("Could not get video from database")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if err == sql.ErrNoRows {
			log.WithField("category_id", id).Debug("Video with specified id doesn't exist")
			return c.SendStatus(fiber.StatusNotFound)
		}
		tags, err := qtx.GetVideoTags(ctx, vid.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get tags from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		dbFiles, err := qtx.GetVideoFiles(ctx, vid.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get tags from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		files := make([]VideoFiles, 0)
		fileUrlPrefix := config.GetString("APP_VIDEOS_URL", "https://localhost/videos")

		for _, file := range dbFiles {
			var newFile = VideoFiles{
				Resolution: string(file.Resolution),
				Duration:   uint64(file.Duration),
				FilePath:   fileUrlPrefix + "/" + file.FilePath,
			}
			files = append(files, newFile)
		}

		videos = append(videos, VideoResponse{
			ID:           uint64(vid.ID),
			Name:         vid.Name,
			Description:  vid.Description,
			Category:     vid.Category.String,
			CategoryID:   uint64(vid.ID),
			Upvotes:      uint64(vid.Upvotes),
			Downvotes:    uint64(vid.Downvotes),
			Views:        uint64(vid.Views),
			IsPremium:    bool(vid.IsPremium),
			ThumbnailUrl: vid.Thumbnail.String,
			Tags:         tags,
			Files:        files,
		})
	}

	return c.Status(fiber.StatusOK).JSON(videos)
}

func GetVideo(c *fiber.Ctx) error {

	idString := c.Params("id")
	id, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
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
	vid, err := qtx.GetVideoByID(ctx, id)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get video from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("category_id", id).Debug("Video with specified id doesn't exist")
		return c.SendStatus(fiber.StatusNotFound)
	}

	tags, err := qtx.GetVideoTags(ctx, vid.ID)
	if err != nil {
		log.WithField("err", err).Error("Could not get tags from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	dbFiles, err := qtx.GetVideoFiles(ctx, vid.ID)
	if err != nil {
		log.WithField("err", err).Error("Could not get tags from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	files := make([]VideoFiles, 0)
	fileUrlPrefix := config.GetString("APP_VIDEOS_URL", "https://localhost/videos")
	thumbnailUrlPrefix := config.GetString("APP_THUMBNAILS_URL", "https://localhost/thumbnails")

	for _, file := range dbFiles {
		var newFile = VideoFiles{
			Resolution: string(file.Resolution),
			Duration:   uint64(file.Duration),
			FilePath:   fileUrlPrefix + "/" + file.FilePath,
		}
		files = append(files, newFile)
	}

	return c.Status(fiber.StatusOK).JSON(VideoResponse{
		ID:           uint64(vid.ID),
		Name:         vid.Name,
		Description:  vid.Description,
		Category:     vid.Category.String,
		CategoryID:   uint64(vid.ID),
		Upvotes:      uint64(vid.Upvotes),
		Downvotes:    uint64(vid.Downvotes),
		Views:        uint64(vid.Views),
		IsPremium:    bool(vid.IsPremium),
		ThumbnailUrl: thumbnailUrlPrefix + "/" + vid.Thumbnail.String,
		Tags:         tags,
		Files:        files,
	})
}

func GetAllVideos(c *fiber.Ctx) error {
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

	res, err := qtx.GetAllVideos(ctx, sqlc.GetAllVideosParams{
		Limit:  25,
		Offset: 0,
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get videos from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	videos := make([]VideoResponse, 0)
	for _, vid := range res {
		tags, err := qtx.GetVideoTags(ctx, vid.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get tags from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		dbFiles, err := qtx.GetVideoFiles(ctx, vid.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get tags from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		fileUrlPrefix := config.GetString("APP_VIDEOS_URL", "https://localhost/videos")
		thumbnailUrlPrefix := config.GetString("APP_THUMBNAILS_URL", "https://localhost/thumbnails")

		files := make([]VideoFiles, 0)
		for _, file := range dbFiles {
			var newFile = VideoFiles{
				Resolution: string(file.Resolution),
				Duration:   uint64(file.Duration),
				FilePath:   fileUrlPrefix + "/" + file.FilePath,
			}
			files = append(files, newFile)
		}

		video := VideoResponse{
			ID:           uint64(vid.ID),
			Name:         vid.Name,
			Description:  vid.Description,
			Category:     vid.Category.String,
			CategoryID:   uint64(vid.ID),
			Upvotes:      uint64(vid.Upvotes),
			Downvotes:    uint64(vid.Downvotes),
			Views:        uint64(vid.Views),
			IsPremium:    bool(vid.IsPremium),
			ThumbnailUrl: thumbnailUrlPrefix + "/" + vid.Thumbnail.String,
			Tags:         tags,
			Files:        files,
		}
		videos = append(videos, video)
	}

	return c.Status(fiber.StatusOK).JSON(videos)
}

func GetRecommendedVideos(c *fiber.Ctx) error {
	// Check if limit and offset is resonable
	var offset int32 = 0
	offsetArray := c.Query("offset")
	fmt.Print(string(offsetArray))
	if len(offsetArray) > 1 {
		result, err := strconv.ParseInt(string(offsetArray), 10, 64)
		if err != nil {
			log.WithField("err", err).Debug("Offset value could not be parsed")
			return c.SendStatus(fiber.StatusBadRequest)
		}
		offset = int32(result)
	}

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

	res, err := qtx.GetRandomVideos(ctx, sqlc.GetRandomVideosParams{
		Limit:  25,
		Offset: offset,
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get videos from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	videos := make([]VideoResponse, 0)
	for _, vid := range res {
		tags, err := qtx.GetVideoTags(ctx, vid.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get tags from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		dbFiles, err := qtx.GetVideoFiles(ctx, vid.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get tags from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		files := make([]VideoFiles, 0)
		for _, file := range dbFiles {
			var newFile = VideoFiles{
				Resolution: string(file.Resolution),
				Duration:   uint64(file.Duration),
				FilePath:   file.FilePath,
			}
			files = append(files, newFile)
		}

		video := VideoResponse{
			ID:           uint64(vid.ID),
			Name:         vid.Name,
			Description:  vid.Description,
			Category:     vid.Category.String,
			CategoryID:   uint64(vid.ID),
			Upvotes:      uint64(vid.Upvotes),
			Downvotes:    uint64(vid.Downvotes),
			Views:        uint64(vid.Views),
			IsPremium:    bool(vid.IsPremium),
			ThumbnailUrl: vid.Thumbnail.String,
			Tags:         tags,
			Files:        files,
		}
		videos = append(videos, video)
	}

	return c.Status(fiber.StatusOK).JSON(videos)
}
