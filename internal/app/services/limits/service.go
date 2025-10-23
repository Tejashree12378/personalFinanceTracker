package limits

import (
	"context"

	"personalFinanceTracker/internal/app/repositories/limits"
	repoModels "personalFinanceTracker/internal/app/repositories/models"
	"personalFinanceTracker/internal/app/services/models"
)

type repository interface {
	Create(ctx context.Context, limit *repoModels.Limit) error
	GetByID(ctx context.Context, id int) (*repoModels.Limit, error)
	Update(ctx context.Context, limit *repoModels.Limit) error
	Delete(ctx context.Context, id int) error
}

type limitService struct {
	repo limits.Repository
}

func NewLimitService(r limits.Repository) *limitService {
	return &limitService{repo: r}
}

func (s *limitService) Create(ctx context.Context, limit *models.Limit) error {
	return s.repo.Create(ctx, limit.ToRepoModel())
}

func (s *limitService) GetByID(ctx context.Context, id int) (*models.Limit, error) {
	repoLimit, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return models.FromRepoModelLimit(repoLimit), nil
}

func (s *limitService) Update(ctx context.Context, limit *models.Limit) error {
	return s.repo.Update(ctx, limit.ToRepoModel())
}

func (s *limitService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
