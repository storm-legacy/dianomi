package database

import "github.com/storm-legacy/dianomi/pkg/models"

// Create user
func CreateUser(user *models.User) error {
	// Add to database
	tx := db.Exec("INSERT INTO users(email, password) VALUES (?,?);", user.Email, user.Password)
	return tx.Error
}

// Check if user exists in database
func IsNewUser(user *models.FormRegisterUser) (result bool, err error) {
	// Check if user is present in database
	var count int = -1
	tx := db.Exec("SELECT count(*) FROM users WHERE email=?;", user.Email).Scan(&count)
	return (count == 0), tx.Error
}

// Get user from database via email
func GetUser(email string) (user *models.User, err error) {
	// Check if user is present in database
	var count int = -1
	tx := db.Raw("SELECT count(*) FROM users WHERE email=?;", email).Scan(&count)

	// Return error if value isn't equal 1
	if tx.Error != nil || count != 1 {
		return nil, tx.Error
	}

	// Get user by email
	tx = db.Raw("SELECT * FROM users WHERE email=?;", email).Scan(&user)
	return user, tx.Error
}
