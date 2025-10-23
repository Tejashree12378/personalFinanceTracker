package limits

import (
	"context"
	"time"

	"personalFinanceTracker/internal/app/repositories/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, limit *models.Limit) error {
	return r.db.WithContext(ctx).Create(limit).Error
}

func (r *Repository) GetByID(ctx context.Context, id int) (*models.Limit, error) {
	var limit models.Limit
	err := r.db.WithContext(ctx).First(&limit, "id = ? AND deleted_at IS NULL", id).Error

	return &limit, err
}

func (r *Repository) Update(ctx context.Context, limit *models.Limit) error {
	return r.db.WithContext(ctx).
		Table(models.Expense{}.Table()).
		Where("id = ? AND deleted_at IS NULL", limit.ID).
		Updates(limit).Error
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&models.Limit{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}
