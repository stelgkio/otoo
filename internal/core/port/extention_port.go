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

	CreateProjectExtension(ctx echo.Context, projectID string, ex *domain.Extension, days int, subID string) error
	GetAllProjectExtensions(ctx echo.Context, projectID string) ([]*domain.ProjectExtension, error)
	GetProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtension, error)
	DeleteProjectExtension(ctx echo.Context, extensionID, projectID string) error

	////////////////// ACS Extension/////////////////////////

	CreateACSProjectExtension(ctx echo.Context, projectID string, ex *domain.AcsCourierExtension) error
	GetAllACSProjectExtensions(ctx echo.Context, projectID string) ([]*domain.AcsCourierExtension, error)
	GetACSProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.AcsCourierExtension, error)
	DeleteACSProjectExtension(ctx echo.Context, extensionID, projectID string) error

	////////////////// Courier4u Extension/////////////////////////
	CreateCourier4uProjectExtension(ctx echo.Context, projectID string, e *domain.Courier4uExtension) error
	GetAllCourier4uProjectExtensions(ctx echo.Context, projectID string) ([]*domain.Courier4uExtension, error)
	GetCourier4uProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.Courier4uExtension, error)
	DeleteCourier4uProjectExtension(ctx echo.Context, extensionID, projectID string) error

	////////////////// Data Synchronizer Extension/////////////////////////

	CreateSynchronizerProjectExtension(ctx echo.Context, projectID string, e *domain.DataSynchronizerExtension) error
	UpdateSynchronizerCustomerRecievedExtension(ctx echo.Context, projectID string, CustomerRecieved int) error
	UpdateSynchronizerOrderReceivedExtension(ctx echo.Context, projectID string, OrderReceived int) error
	UpdateSynchronizerProductReceivedExtension(ctx echo.Context, projectID string, ProductReceived int) error
	GetAllSynchronizerProjectExtensions(ctx echo.Context, projectID string) ([]*domain.DataSynchronizerExtension, error)
	GetSynchronizerProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.DataSynchronizerExtension, error)
	DeleteSynchronizerProjectExtension(ctx echo.Context, extensionID, projectID string) error
}

// ExtensionService defines the methods for interacting with the Extension service
type ExtensionService interface {
	CreateExtension(ctx echo.Context, e *domain.Extension) error
	GetAllExtensions(ctx echo.Context) ([]*domain.Extension, error)
	GetExtensionByID(ctx echo.Context, extensionID string) (*domain.Extension, error)
	GetExtensionByCode(ctx echo.Context, code string) (*domain.Extension, error)
	DeleteExtension(ctx echo.Context, extensionID string) error

	//////////////////PROJECT ExtensionS/////////////////////////

	CreateProjectExtension(ctx echo.Context, projectID string, e *domain.Extension, days int, subID string) error
	GetAllProjectExtensions(ctx echo.Context, projectID string) ([]*domain.ProjectExtension, error)
	GetProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtension, error)
	DeleteProjectExtension(ctx echo.Context, extensionID, projectID string) error

	////////////////// ACS Extension/////////////////////////
	CreateACSProjectExtension(ctx echo.Context, projectID string, e *domain.AcsCourierExtension) error
	GetAllACSProjectExtensions(ctx echo.Context, projectID string) ([]*domain.AcsCourierExtension, error)
	GetACSProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.AcsCourierExtension, error)
	DeleteACSProjectExtension(ctx echo.Context, extensionID, projectID string) error

	////////////////// Courier4u Extension/////////////////////////
	CreateCourier4uProjectExtension(ctx echo.Context, projectID string, e *domain.Courier4uExtension) error
	GetAllCourier4uProjectExtensions(ctx echo.Context, projectID string) ([]*domain.Courier4uExtension, error)
	GetCourier4uProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.Courier4uExtension, error)
	DeleteCourier4uProjectExtension(ctx echo.Context, extensionID, projectID string) error

	////////////////// Data Synchronizer Extension/////////////////////////

	CreateSynchronizerProjectExtension(ctx echo.Context, projectID string, e *domain.DataSynchronizerExtension) error
	UpdateSynchronizerCustomerRecievedExtension(ctx echo.Context, projectID string, CustomerRecieved int) error
	UpdateSynchronizerOrderReceivedExtension(ctx echo.Context, projectID string, OrderReceived int) error
	UpdateSynchronizerProductReceivedExtension(ctx echo.Context, projectID string, ProductReceived int) error
	GetAllSynchronizerProjectExtensions(ctx echo.Context, projectID string) ([]*domain.DataSynchronizerExtension, error)
	GetSynchronizerProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.DataSynchronizerExtension, error)
	DeleteSynchronizerProjectExtension(ctx echo.Context, extensionID, projectID string) error
}
