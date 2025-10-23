package models

import (
	"time"

	repoModels "personalFinanceTracker/internal/app/repositories/models"
)

type Expense struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description"`
	Name        string     `json:"name"`
	Category    string     `json:"category"`
	Date        time.Time  `json:"date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// ToRepoModel converts service model to repository model
func (e *Expense) ToRepoModel() *repoModels.Expense {
	return &repoModels.Expense{
		ID:          e.ID,
		UserID:      e.UserID,
		Amount:      e.Amount,
		Description: e.Description,
		Name:        e.Name,
		Category:    e.Category,
		Date:        e.Date,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// FromRepoModelExpense converts repo model to service model
func FromRepoModelExpense(repo *repoModels.Expense) *Expense {
	if repo == nil {
		return nil
	}
	return &Expense{
		ID:          repo.ID,
		UserID:      repo.UserID,
		Amount:      repo.Amount,
		Description: repo.Description,
		Name:        repo.Name,
		Category:    repo.Category,
		Date:        repo.Date,
		CreatedAt:   repo.CreatedAt,
		UpdatedAt:   repo.UpdatedAt,
		DeletedAt:   repo.DeletedAt,
	}
}
