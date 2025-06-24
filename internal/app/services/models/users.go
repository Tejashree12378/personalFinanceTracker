package models

import (
	"time"

	repoModels "personalFinanceTracker/internal/app/repositories/models"
)

type User struct {
	ID           int
	FirstName    string
	LastName     string
	Status       string
	Email        string
	PasswordHash string
	PhoneNumber  string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

func (u *User) ToRepoModel() *repoModels.User {
	return &repoModels.User{
		ID:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Status:       u.Status,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		PhoneNumber:  u.PhoneNumber,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		DeletedAt:    u.DeletedAt,
	}
}

func FromRepoModel(repoUser *repoModels.User) *User {
	if repoUser == nil {
		return nil
	}
	return &User{
		ID:          repoUser.ID,
		FirstName:   repoUser.FirstName,
		LastName:    repoUser.LastName,
		Status:      repoUser.Status,
		Email:       repoUser.Email,
		PhoneNumber: repoUser.PhoneNumber,
		CreatedAt:   repoUser.CreatedAt,
		UpdatedAt:   repoUser.UpdatedAt,
		DeletedAt:   repoUser.DeletedAt,
	}
}
