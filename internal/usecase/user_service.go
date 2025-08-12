package usecase

import (
	"Architecture/internal/domain"
	"context"
	"errors"
)

type UserService interface {
	Get(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (int, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int) error
}

type userService struct {
	userRepo domain.UserRepo
}

func NewUserService(userRepo domain.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Get(ctx context.Context) ([]domain.User, error) {
	users, err := s.userRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, errors.New("user not found")
	}

	return users, nil
}

func (s *userService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *userService) Create(ctx context.Context, user *domain.User) (int, error) {
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (s *userService) Update(ctx context.Context, user *domain.User) error {
	err := s.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Delete(ctx context.Context, id int) error {
	err := s.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
