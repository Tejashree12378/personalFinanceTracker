package models

import (
	serviceModel "personalFinanceTracker/internal/app/services/models"
)

type LimitRequest struct {
	UserID       int `json:"user_id" binding:"required"`
	MonthlyLimit int `json:"monthly_limit"`
	YearlyLimit  int `json:"yearly_limit"`
}

func (r *LimitRequest) ToServiceModel() *serviceModel.Limit {
	return &serviceModel.Limit{
		UserID:       r.UserID,
		MonthlyLimit: r.MonthlyLimit,
		YearlyLimit:  r.YearlyLimit,
	}
}

type LimitUpdateRequest struct {
	UserID       int  `json:"user_id" binding:"required"`
	MonthlyLimit *int `json:"monthly_limit"`
	YearlyLimit  *int `json:"yearly_limit"`
}

func (r *LimitUpdateRequest) ToServiceModel() *serviceModel.Limit {
	req := &serviceModel.Limit{}

	if r.MonthlyLimit != nil {
		req.MonthlyLimit = *r.MonthlyLimit
	}

	if r.YearlyLimit != nil {
		req.YearlyLimit = *r.YearlyLimit
	}

	return req
}
