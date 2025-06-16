package expenses

import (
	"context"

	"personalFinanceTracker/internal/app/repositories/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, expense *models.Expense) error {
	return r.db.WithContext(ctx).Create(expense).Error
}

func (r *Repository) Update(ctx context.Context, expense *models.Expense) error {
	return r.db.WithContext(ctx).
		Table(models.Expense{}.Table()).
		Where("id = ? AND deleted_at IS NULL", expense.ID).
		Updates(expense).Error
}

func (r *Repository) GetByID(ctx context.Context, id int) (*models.Expense, error) {
	var expense models.Expense
	if err := r.db.WithContext(ctx).First(&expense, id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Model(&models.Expense{}).
		Where("id = ?", id).
		Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}
