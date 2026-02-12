package entity

import (
	"time"
)

// User represents the user entity in the domain
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"` // Password không được expose trong JSON
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}

// UserFilter represents filter options for querying users
type UserFilter struct {
	Email    *string
	Username *string
	IsActive *bool
	Limit    int
	Offset   int
}

