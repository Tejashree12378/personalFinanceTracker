package models

import "time"

type Limit struct {
	ID           int        `gorm:"primaryKey"`
	UserID       int        `gorm:"not null;index"`
	MonthlyLimit int        `gorm:"column:monthly_limit"`
	YearlyLimit  int        `gorm:"column:yearly_limit"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `gorm:"index"`
}
