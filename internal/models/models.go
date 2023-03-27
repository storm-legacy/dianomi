package models

import "time"

type User struct {
	ID         uint64    `gorm:"type:bigint;autoincrement;primary_key"`
	Email      string    `gorm:"type:varchar(100);not null;unique"`
	Password   string    `gorm:"type:varchar(100);not null"`
	VerifiedAt time.Time `gorm:"default:null"`
	CreatedAt  time.Time `gorm:"default:now()"`
	UpdatedAt  time.Time `gorm:"default:now()"`
}

type RegisterUser struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	PasswordRepeated string `json:"passwordRepeated"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
