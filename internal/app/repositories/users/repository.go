package users

import (
	"context"
	"time"

	"personalFinanceTracker/internal/app/repositories/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

//nolint:always return struct with proper initialisation.
func New(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *repository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).
		Table(models.User{}.Table()).
		Where("id = ? AND deleted_at IS NULL", user.ID).
		Updates(user).Error
}

func (r *repository) DeleteUser(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).
		Table(models.User{}.Table()).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}
