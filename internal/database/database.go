package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/internal/models"
)

var db *gorm.DB

// Function connecting to database
func ConnectToDatabase(databaseUrl string) (err error) {
	log.WithField("databaseUrl", databaseUrl).Debug("Loaded database configuration")

	db, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	db.Logger = logger.Default.LogMode(logger.Error)

	if err != nil {
		log.WithField("val", err).Fatal("Could not connect to database")
	} else {
		log.Info("Connected to database!")
	}

	return nil
}

// Migrate all changes to database
func AutoMigrate() {
	db.AutoMigrate(&models.User{})
}

// Create user
func CreateUser(user *models.User) error {
	// Add to database
	tx := db.Raw("INSERT INTO users(email, password) VALUES (?,?);", user.Email, user.Password)
	return tx.Error
}

// Check if user exists in database
func IsNewUser(user *models.RegisterUser) (result bool, err error) {
	// Check if user is present in database
	var count int = -1
	tx := db.Raw("SELECT count(*) FROM users WHERE email=?;", user.Email).Scan(&count)
	return (count == 0), tx.Error
}
