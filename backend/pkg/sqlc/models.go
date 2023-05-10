// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Resolution string

const (
	Resolution360p Resolution = "360p"
	Resolution480p Resolution = "480p"
	Resolution720p Resolution = "720p"
)

func (e *Resolution) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Resolution(s)
	case string:
		*e = Resolution(s)
	default:
		return fmt.Errorf("unsupported scan type for Resolution: %T", src)
	}
	return nil
}

type NullResolution struct {
	Resolution Resolution
	Valid      bool // Valid is true if Resolution is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullResolution) Scan(value interface{}) error {
	if value == nil {
		ns.Resolution, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Resolution.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullResolution) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Resolution), nil
}

type Role string

const (
	RolePremium       Role = "premium"
	RoleAdministrator Role = "administrator"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type NullRole struct {
	Role  Role
	Valid bool // Valid is true if Role is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRole) Scan(value interface{}) error {
	if value == nil {
		ns.Role, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Role.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Role), nil
}

type VerifyEmailType string

const (
	VerifyEmailTypeEmailVerification VerifyEmailType = "emailVerification"
	VerifyEmailTypeEmailChange       VerifyEmailType = "emailChange"
	VerifyEmailTypePasswordReset     VerifyEmailType = "passwordReset"
)

func (e *VerifyEmailType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VerifyEmailType(s)
	case string:
		*e = VerifyEmailType(s)
	default:
		return fmt.Errorf("unsupported scan type for VerifyEmailType: %T", src)
	}
	return nil
}

type NullVerifyEmailType struct {
	VerifyEmailType VerifyEmailType
	Valid           bool // Valid is true if VerifyEmailType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVerifyEmailType) Scan(value interface{}) error {
	if value == nil {
		ns.VerifyEmailType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VerifyEmailType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVerifyEmailType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VerifyEmailType), nil
}

type Category struct {
	ID   int64
	Name string
}

type Tag struct {
	ID   int64
	Name string
}

type User struct {
	ID         int64
	Email      string
	Password   string
	VerifiedAt sql.NullTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserVideoMetric struct {
	ID                int64
	UserID            int64
	VideoID           int64
	TimeSpentWatching int32
	StoppedAt         int32
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
}

type UsersPackage struct {
	ID         int64
	UserID     int64
	Tier       Role
	CreatedAt  sql.NullTime
	ValidFrom  time.Time
	ValidUntil time.Time
}

type Verification struct {
	ID         int64
	UserID     int64
	TaskType   VerifyEmailType
	Code       uuid.UUID
	Used       bool
	CreatedAt  sql.NullTime
	ValidUntil sql.NullTime
}

type Video struct {
	ID          int64
	Name        string
	Description string
	AuthorID    sql.NullInt64
	CategoryID  sql.NullInt64
	Upvotes     int64
	Downvotes   int64
	Views       int64
	UpdatedAt   sql.NullTime
	CreatedAt   sql.NullTime
	DeletedAt   sql.NullTime
}

type VideoFile struct {
	ID         int64
	FilePath   string
	VideoID    int64
	FileSize   int64
	Duration   int32
	Resolution Resolution
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
	DeletedAt  sql.NullTime
}

type VideoTag struct {
	ID      int64
	VideoID int64
	TagID   int64
}

type VideoThumbnail struct {
	ID        int64
	VideoID   int64
	FileSize  int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
