package service

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	e "github.com/stelgkio/otoo/internal/core/util"
)

/**
 * UserService implements port.UserService interface
 * and provides an access to the user repository
 * and cache service
 */
type UserService struct {
	repo port.UserRepository
}

// NewUserService creates a new user service instance
func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

// CreateUser creates a new user
func (us *UserService) CreateUser(ctx echo.Context, user *domain.User) (*domain.User, error) {
	u, err := us.repo.GetUserByEmail(ctx, user.Email)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	if u != nil {
		return nil, e.ErrEmailAlreadyExist
	}
	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		if err == e.ErrConflictingData {
			return nil, err
		}
		return nil, e.ErrInternal
	}

	return user, nil
}

// GetAdminUserByProjectId retrieves a user by project id
func (us *UserService) GetAdminUserByProjectId(ctx echo.Context, projectid uuid.UUID) ([]*domain.User, error) {
	users, err := us.repo.GetAdminUserByProjectId(ctx, projectid)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return users, nil
}

// GetUserById retrieves a user by id
func (us *UserService) GetUserById(ctx echo.Context, id uuid.UUID) (*domain.User, error) {
	user, err := us.repo.GetUserById(ctx, id)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return user, nil

}

// GetUserByEmail retrieves a user by email
func (us *UserService) GetUserByEmail(ctx echo.Context, email string) (*domain.User, error) {
	user, err := us.repo.GetUserByEmail(ctx, email)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return user, nil

}

// UpdateUser updates a user
func (us *UserService) UpdateUser(ctx echo.Context, user *domain.User) (*domain.User, error) {
	user, err := us.repo.UpdateUser(ctx, user)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return user, nil

}

// DeleteUser deletes a user
func (us *UserService) DeleteUser(ctx echo.Context, userID uuid.UUID) error {

	err := us.repo.DeleteUser(ctx, userID)
	if err != nil && err != e.ErrDataNotFound {
		return e.ErrInternal
	}
	return nil

}

// FindUsersByProjectId retrieves a user by project id
func (us *UserService) FindUsersByProjectId(ctx echo.Context, id uuid.UUID) ([]*domain.User, error) {
	users, err := us.repo.FindUsersByProjectId(ctx, id)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return users, nil

}

// FindProjectsByUserId retrieves a user by project id
func (us *UserService) FindProjectsByUserId(ctx echo.Context, userId uuid.UUID) ([]*domain.Project, error) {
	users, err := us.repo.FindProjectsByUserId(ctx, userId)
	if err != nil && err != e.ErrDataNotFound {
		return nil, e.ErrInternal
	}
	return users, nil
}
