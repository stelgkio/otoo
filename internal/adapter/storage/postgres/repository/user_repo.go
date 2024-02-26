package repository

import (
	"github.com/go-pg/pg/v10"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	e "github.com/stelgkio/otoo/internal/core/util"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

// CreateUser creates a new user in the database
func (repo *UserRepository) CreateUser(ctx echo.Context, user *domain.User) (*domain.User, error) {
	_, err := repo.db.Model(user).Insert()
	if err != nil {
		panic(err)
	}

	return user, nil
}

// GetUserByEmail creates a user in the database
func (repo *UserRepository) GetUserByEmail(ctx echo.Context, email string) (*domain.User, error) {
	var user domain.User
	err := repo.db.Model(&user).Where("email =?", email).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, e.ErrDataNotFound
		}
		return nil, e.ErrInternal
	}

	return &user, nil
}
