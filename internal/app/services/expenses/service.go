package expenses

import (
	"context"

	repoModels "personalFinanceTracker/internal/app/repositories/models"
	"personalFinanceTracker/internal/app/services/models"
)

type repository interface {
	Create(ctx context.Context, expense *repoModels.Expense) error
	Update(ctx context.Context, expense *repoModels.Expense) error
	GetByID(ctx context.Context, id int) (*repoModels.Expense, error)
	Delete(ctx context.Context, id int) error
}

type expenseService struct {
	repo repository
}

func NewExpenseService(repo repository) *expenseService {
	return &expenseService{repo: repo}
}

func (s *expenseService) CreateExpense(ctx context.Context, expense *models.Expense) error {
	return s.repo.Create(ctx, expense.ToRepoModel())
}

func (s *expenseService) UpdateExpense(ctx context.Context, expense *models.Expense) error {
	return s.repo.Update(ctx, expense.ToRepoModel())
}

func (s *expenseService) GetExpenseByID(ctx context.Context, id int) (*models.Expense, error) {
	repoModel, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return models.FromRepoModelExpense(repoModel), nil
}

func (s *expenseService) DeleteExpense(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
