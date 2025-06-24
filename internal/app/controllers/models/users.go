package models

import (
	"time"

	serviceModel "personalFinanceTracker/internal/app/services/models"
)

type UserUpdateRequest struct {
	ID          uint
	FirstName   *string `json:"first_name" binding:"required"`
	LastName    *string `json:"last_name" binding:"required"`
	Email       *string `json:"email" binding:"required,email"`
	PhoneNumber *string `json:"phone_number" binding:"required"`
}

type SignUpRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Status      string `json:"status"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *SignUpRequest) ToServiceModel() *serviceModel.User {
	updatedAt := time.Now()

	return &serviceModel.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Status:       req.Status,
		Email:        req.Email,
		PasswordHash: req.Password,
		PhoneNumber:  req.PhoneNumber,
		CreatedAt:    time.Now(),
		UpdatedAt:    &updatedAt,
	}
}

func (req *UserUpdateRequest) ToServiceModel() *serviceModel.User {
	u := &serviceModel.User{}

	if req.FirstName != nil {
		u.FirstName = *req.FirstName
	}

	if req.LastName != nil {
		u.LastName = *req.LastName
	}

	if req.Email != nil {
		u.Email = *req.Email
	}

	if req.PhoneNumber != nil {
		u.PhoneNumber = *req.PhoneNumber
	}

	return u
}
