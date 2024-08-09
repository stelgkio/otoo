package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// ExtentionRepository defines the methods for interacting with the Extention repository
type ExtentionRepository interface {
	CreateExtention(ctx echo.Context, e *domain.Extention) error
	GetAllExtentions(ctx echo.Context) ([]*domain.Extention, error)
	GetExtentionsByID(ctx echo.Context, extensionID string) (*domain.Extention, error)
	DeleteExtention(ctx echo.Context, extensionID string) error

	//////////////////PROJECT EXTENTIONS/////////////////////////

	CreateProjectExtention(ctx echo.Context, projectID string, ex *domain.Extention) error
	GetAllProjectExtentions(ctx echo.Context, projectID string) ([]*domain.ProjectExtention, error)
	GetProjectExtentionsByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtention, error)
	DeleteProjectExtention(ctx echo.Context, extensionID, projectID string) error
}

// ExtentionService defines the methods for interacting with the Extention service
type ExtentionService interface {
	CreateExtention(ctx echo.Context, e *domain.Extention) error
	GetAllExtentions(ctx echo.Context) ([]*domain.Extention, error)
	GetExtentionsByID(ctx echo.Context, extensionID string) (*domain.Extention, error)
	DeleteExtention(ctx echo.Context, extensionID string) error

	//////////////////PROJECT EXTENTIONS/////////////////////////

	CreateProjectExtention(ctx echo.Context, projectID string, e *domain.Extention) error
	GetAllProjectExtentions(ctx echo.Context, projectID string) ([]*domain.ProjectExtention, error)
	GetProjectExtentionsByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtention, error)
	DeleteProjectExtention(ctx echo.Context, extensionID, projectID string) error
}
