package models

import "time"

type Expense struct {
	ID          int        `gorm:"primaryKey;autoIncrement"`
	UserID      int        `gorm:"not null;index"`
	Amount      float64    `gorm:"type:decimal(10,2);not null"`
	Description string     `gorm:"type:varchar(255)"`
	Name        string     `gorm:"type:varchar(255)"`
	Category    string     `gorm:"type:varchar(100)"`
	Date        time.Time  `gorm:"default:current_timestamp"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `gorm:"index"`
}

func (e Expense) Table() string {
	return "expenses"
}
