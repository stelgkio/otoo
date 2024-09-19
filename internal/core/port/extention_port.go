package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// ExtensionRepository defines the methods for interacting with the Extension repository
type ExtensionRepository interface {
	CreateExtension(ctx echo.Context, e *domain.Extension) error
	GetAllExtensions(ctx echo.Context) ([]*domain.Extension, error)
	GetExtensionByID(ctx echo.Context, extensionID string) (*domain.Extension, error)
	GetExtensionByCode(ctx echo.Context, code string) (*domain.Extension, error)
	DeleteExtension(ctx echo.Context, extensionID string) error

	//////////////////PROJECT ExtensionS/////////////////////////

	CreateProjectExtension(ctx echo.Context, projectID string, ex *domain.Extension) error
	GetAllProjectExtensions(ctx echo.Context, projectID string) ([]*domain.ProjectExtension, error)
	GetProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtension, error)
	DeleteProjectExtension(ctx echo.Context, extensionID, projectID string) error

	////////////////// ACS Extension/////////////////////////

	CreateACSProjectExtension(ctx echo.Context, projectID string, ex *domain.AcsCourierExtension) error
	GetAllACSProjectExtensions(ctx echo.Context, projectID string) ([]*domain.AcsCourierExtension, error)
	GetACSProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.AcsCourierExtension, error)
	DeleteACSProjectExtension(ctx echo.Context, extensionID, projectID string) error
}

// ExtensionService defines the methods for interacting with the Extension service
type ExtensionService interface {
	CreateExtension(ctx echo.Context, e *domain.Extension) error
	GetAllExtensions(ctx echo.Context) ([]*domain.Extension, error)
	GetExtensionByID(ctx echo.Context, extensionID string) (*domain.Extension, error)
	GetExtensionByCode(ctx echo.Context, code string) (*domain.Extension, error)
	DeleteExtension(ctx echo.Context, extensionID string) error

	//////////////////PROJECT ExtensionS/////////////////////////

	CreateProjectExtension(ctx echo.Context, projectID string, e *domain.Extension) error
	GetAllProjectExtensions(ctx echo.Context, projectID string) ([]*domain.ProjectExtension, error)
	GetProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtension, error)
	DeleteProjectExtension(ctx echo.Context, extensionID, projectID string) error

	////////////////// ACS Extension/////////////////////////

	CreateACSProjectExtension(ctx echo.Context, projectID string, e *domain.AcsCourierExtension) error
	GetAllACSProjectExtensions(ctx echo.Context, projectID string) ([]*domain.AcsCourierExtension, error)
	GetACSProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.AcsCourierExtension, error)
	DeleteACSProjectExtension(ctx echo.Context, extensionID, projectID string) error
}
