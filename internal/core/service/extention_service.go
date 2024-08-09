package service

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
)

// ExtentionService defines the methods for interacting with the Extention service
type ExtentionService struct {
	repo port.ExtentionRepository
}

// NewExtentionService creates a new user service instance
func NewExtentionService(repo port.ExtentionRepository) *ExtentionService {
	return &ExtentionService{
		repo,
	}
}

// CreateExtention creates a new Extention
func (ex *ExtentionService) CreateExtention(ctx echo.Context, c *domain.Extention) error {
	panic("unimplemented")
}

// GetAllExtentions gets all Extentions
func (ex *ExtentionService) GetAllExtentions(ctx echo.Context) ([]*domain.Extention, error) {
	panic("unimplemented")
}

// GetExtentionsByID gets a Extention by ID
func (ex *ExtentionService) GetExtentionsByID(ctx echo.Context, extensionID string) (*domain.Extention, error) {
	panic("unimplemented")
}

// DeleteExtention deletes a Extention by ID
func (ex *ExtentionService) DeleteExtention(ctx echo.Context, extensionID string) error {
	panic("unimplemented")
}

//////////////////PROJECT EXTENTIONS/////////////////////////

// CreateProjectExtention creates a new ProjectExtention
func (ex *ExtentionService) CreateProjectExtention(ctx echo.Context, projectID string, e *domain.Extention) error {
	panic("unimplemented")
}

// GetAllProjectExtentions gets all ProjectExtentions
func (ex *ExtentionService) GetAllProjectExtentions(ctx echo.Context, projectID string) ([]*domain.ProjectExtention, error) {
	panic("unimplemented")
}

// GetProjectExtentionsByID gets a ProjectExtention by ID
func (ex *ExtentionService) GetProjectExtentionsByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtention, error) {
	panic("unimplemented")
}

// DeleteProjectExtention deletes a ProjectExtention by ID
func (ex *ExtentionService) DeleteProjectExtention(ctx echo.Context, extensionID, projectID string) error {
	panic("unimplemented")
}
