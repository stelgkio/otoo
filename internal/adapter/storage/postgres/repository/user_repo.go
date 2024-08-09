package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	e "github.com/stelgkio/otoo/internal/core/util"
)

// UserRepository is the repository for the user model
type UserRepository struct {
	db *pg.DB
}

// NewUserRepository returns a new UserRepository
func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

// CreateUser creates a new user in the database
func (repo *UserRepository) CreateUser(ctx echo.Context, user *domain.User) (*domain.User, error) {
	_, err := repo.db.Model(user).Insert()
	if err != nil {
		// Handle unique constraint violation
		if pgErr, ok := err.(pg.Error); ok {
			if pgErr.IntegrityViolation() {
				return nil, e.ErrEmailAlreadyExist
			}
		}
		return nil, err
	}

	return user, nil
}

// GetUserByEmail creates a user in the database
func (repo *UserRepository) GetUserByEmail(ctx echo.Context, email string) (*domain.User, error) {
	var user domain.User
	err := repo.db.Model(&user).Where("email =?", email).Where("is_active =true").Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, e.ErrDataNotFound
		}
		return nil, e.ErrInternal
	}

	return &user, nil
}

// GetUserById creates a user in the database
func (repo *UserRepository) GetUserById(ctx echo.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := repo.db.Model(&user).Where("id =?", id).Where("is_active =true").Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, e.ErrDataNotFound
		}
		return nil, e.ErrInternal
	}

	return &user, nil
}

// UpdateUser creates a new user in the database
func (repo *UserRepository) UpdateUser(ctx echo.Context, user *domain.User) (*domain.User, error) {
	_, err := repo.db.Model(user).Where("email =?", user.Email).Where("is_active =true").Update()
	if err != nil {
		panic(err)
	}

	return user, nil
}

// DeleteUser delete a user in the database
func (repo *UserRepository) DeleteUser(ctx echo.Context, id uuid.UUID) error {
	user := &domain.User{}
	_, err := repo.db.Model(user).Where("id = ?", id).Delete()

	if err != nil {
		panic(err)
	}

	return nil
}
