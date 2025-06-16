package users

import (
	"context"

	repoModels "personalFinanceTracker/internal/app/repositories/models"
	"personalFinanceTracker/internal/app/services/models"
)

type repository interface {
	CreateUser(ctx context.Context, user *repoModels.User) error
	GetUserByID(ctx context.Context, id int) (*repoModels.User, error)
	UpdateUser(ctx context.Context, user *repoModels.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	repo repository
}

func NewUserService(r repository) *userService {
	return &userService{r}
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	return s.repo.CreateUser(ctx, user.ToRepoModel())
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	resp, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.FromRepoModel(resp), nil
}

func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.repo.UpdateUser(ctx, user.ToRepoModel())
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}
