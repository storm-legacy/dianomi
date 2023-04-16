// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleFree          Role = "free"
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

type RevokedToken struct {
	ID         int64
	Token      string
	UserID     int64
	ValidUntil time.Time
}

type User struct {
	ID         int64
	Email      string
	Password   string
	VerifiedAt sql.NullTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
