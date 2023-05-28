package controllers

import "time"

type Package struct {
	ID         int64     `json:"id"`
	UserID     uint64    `json:"user_id" validate:"required"`
	Tier       string    `json:"tier" validate:"required"`
	ValidFrom  time.Time `json:"valid_from" validate:"required"`
	ValidUntil time.Time `json:"valid_until" validate:"required"`
}
