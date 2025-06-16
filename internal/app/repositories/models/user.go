package models

import (
	"time"
)

type User struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName   string     `gorm:"type:varchar(50);not null;unique" json:"first_name"`
	LastName    string     `gorm:"type:varchar(16);not null" json:"last_name"`
	Status      string     `gorm:"type:varchar(16);not null" json:"status"`
	Email       string     `gorm:"type:varchar(16);not null" json:"email"`
	PhoneNumber string     `gorm:"type:varchar(16);not null" json:"phone_number"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at,omitempty"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (User) Table() string {
	return "users"
}
