package repository

import (
	"time"

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

// GetAdminUserByProjectId creates a user in the database
func (repo *UserRepository) GetAdminUserByProjectId(ctx echo.Context, id uuid.UUID) ([]*domain.User, error) {
	var users []*domain.User
	err := repo.db.Model(&users).Where("projectId =?", id).Where("is_active =true").Where("reseve_notification = true").Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, e.ErrDataNotFound
		}
		return nil, e.ErrInternal
	}

	return users, nil
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
	currentTime := time.Now()
	user := &domain.User{
		Base: domain.Base{
			DeletedAt: &currentTime,
			IsActive:  false,
		},
	}
	_, err := repo.db.Model(user).
		Where("id = ?", id).
		Set("deleted_at = ?", user.Base.DeletedAt).
		Set("is_active = ?", user.Base.IsActive).
		Update()
	if err != nil {
		panic(err)
	}

	return nil
}

// FindUsersByProjectId retrieves all users associated with a given project
func (repo *UserRepository) FindUsersByProjectId(ctx echo.Context, projectID uuid.UUID) ([]*domain.User, error) {
	var users []*domain.User
	// Write your raw SQL query
	query := `
		SELECT u.*
			FROM "user" u 
			Left JOIN "user_projects" up  ON u.id = up.user_id
			where up.project_id = ?::uuid 
			AND u.is_active = ?;
		`

	// Execute the raw SQL query
	_, err := repo.db.Query(&users, query, projectID, true)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindProjectsByUserId retrieves all projects associated with a given user
func (repo *UserRepository) FindProjectsByUserId(ctx echo.Context, userId uuid.UUID) ([]*domain.Project, error) {
	var projects []*domain.Project

	// Query to join UserProject and Project tables
	err := repo.db.Model(&projects).
		Join("JOIN user_projects ON user_projects.project_id = project.id").
		Where("user_projects.user_id = ?::uuid", userId).
		Where("project.is_active = true"). // Make sure to get only active projects
		Order("project.name ASC").
		Select()

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, e.ErrDataNotFound
		}
		return nil, e.ErrInternal
	}

	return projects, nil
}
