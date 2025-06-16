package models

import serviceModel "personalFinanceTracker/internal/app/services/models"

type UserCreateRequest struct {
	ID          uint
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UserUpdateRequest struct {
	ID          uint
	FirstName   *string `json:"first_name" binding:"required"`
	LastName    *string `json:"last_name" binding:"required"`
	Email       *string `json:"email" binding:"required,email"`
	PhoneNumber *string `json:"phone_number" binding:"required"`
}

func (req *UserCreateRequest) ToServiceModel() *serviceModel.User {
	return &serviceModel.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Status:      "active",
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
