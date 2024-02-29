package service

import (
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

// Register creates a new user
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

	////
	///    REDIS
	//

	// cacheKey := util.GenerateCacheKey("user", user.ID)
	// userSerialized, err := util.Serialize(user)
	// if err != nil {
	// 	return nil, util.ErrInternal
	// }

	// err = us.cache.Set(ctx, cacheKey, userSerialized, 0)
	// if err != nil {
	// 	return nil, util.ErrInternal
	// }

	// err = us.cache.DeleteByPrefix(ctx, "users:*")
	// if err != nil {
	// 	return nil, util.ErrInternal
	// }

	return user, nil
}
