package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/models"
)

var db *gorm.DB

func init() {
	// Generate connection string
	databaseUrl := fmt.Sprintf(
		`postgresql://%s:%s@%s:%d/%s`,
		config.GetString("APP_PG_USER"),
		config.GetString("APP_PG_PASSWORD"),
		config.GetString("APP_PG_HOST"),
		config.GetInt("APP_PG_PORT"),
		config.GetString("APP_PG_DB"),
	)
	log.WithField("databaseUrl", databaseUrl).Debug("Loaded database configuration")

	// Connect to database
	var err error
	db, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.WithField("err", err.Error()).Fatal("Could not connect to database")
	} else {
		log.Info("Connected to database!")
	}

	// Set log level for database
	db.Logger = logger.Default.LogMode(logger.Info)

	// Migrate db
	migrate()
}

func migrate() {
	db.AutoMigrate(&models.User{})
}
