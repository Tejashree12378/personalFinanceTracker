package users

import (
	"context"
	"fmt"

	repoModels "personalFinanceTracker/internal/app/repositories/models"
	"personalFinanceTracker/internal/app/services/models"
	"personalFinanceTracker/internal/app/utils"
)

type repository interface {
	CreateUser(ctx context.Context, user *repoModels.User) error
	GetUserByID(ctx context.Context, id int) (*repoModels.User, error)
	UpdateUser(ctx context.Context, user *repoModels.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUserByEmail(ctx context.Context, email string) (*repoModels.User, error)
}

type userService struct {
	repo repository
}

func NewUserService(r repository) *userService {
	return &userService{r}
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

func (s *userService) SignUp(ctx context.Context, user *models.User) error {
	// Hash password
	hashed, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = hashed

	return s.repo.CreateUser(ctx, user.ToRepoModel())
}

func (s *userService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if err := utils.CheckPasswordHash(password, user.PasswordHash); err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		return "", fmt.Errorf("could not generate token")
	}

	return token, nil
}
