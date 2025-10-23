package models

import (
	"time"

	repoModels "personalFinanceTracker/internal/app/repositories/models"
)

type Limit struct {
	ID           int        `json:"id"`
	UserID       int        `json:"user_id"`
	MonthlyLimit int        `json:"monthly_limit"`
	YearlyLimit  int        `json:"yearly_limit"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

func (l *Limit) ToRepoModel() *repoModels.Limit {
	return &repoModels.Limit{
		ID:           l.ID,
		UserID:       l.UserID,
		MonthlyLimit: l.MonthlyLimit,
		YearlyLimit:  l.YearlyLimit,
		CreatedAt:    l.CreatedAt,
		UpdatedAt:    l.UpdatedAt,
		DeletedAt:    l.DeletedAt,
	}
}

func FromRepoModelLimit(repo *repoModels.Limit) *Limit {
	if repo == nil {
		return nil
	}
	return &Limit{
		ID:           repo.ID,
		UserID:       repo.UserID,
		MonthlyLimit: repo.MonthlyLimit,
		YearlyLimit:  repo.YearlyLimit,
		CreatedAt:    repo.CreatedAt,
		UpdatedAt:    repo.UpdatedAt,
		DeletedAt:    repo.DeletedAt,
	}
}
