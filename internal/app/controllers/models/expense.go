package models

import (
	"time"

	serviceModel "personalFinanceTracker/internal/app/services/models"
)

type ExpenseCreateRequest struct {
	UserID      int       `json:"user_id" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
}

type ExpenseUpdateRequest struct {
	ID          int
	Amount      *float64   `json:"amount" binding:"required"`
	Description *string    `json:"description"`
	Name        *string    `json:"name"`
	Category    *string    `json:"category"`
	Date        *time.Time `json:"date"`
}

func (e *ExpenseCreateRequest) ToServiceModel() *serviceModel.Expense {
	return &serviceModel.Expense{
		UserID:      e.UserID,
		Amount:      e.Amount,
		Description: e.Description,
		Name:        e.Name,
		Category:    e.Category,
		Date:        e.Date,
	}
}

func (e *ExpenseUpdateRequest) ToServiceModel() *serviceModel.Expense {
	exp := &serviceModel.Expense{
		ID: e.ID,
	}

	if e.Amount != nil {
		exp.Amount = *e.Amount
	}

	if e.Description != nil {
		exp.Description = *e.Description
	}

	if e.Name != nil {
		exp.Name = *e.Name
	}

	if e.Category != nil {
		exp.Category = *e.Category
	}

	if e.Date != nil {
		exp.Date = *e.Date
	}

	return exp
}
