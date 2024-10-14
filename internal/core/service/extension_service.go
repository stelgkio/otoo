package service

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
)

// ExtensionService defines the methods for interacting with the Extension service
type ExtensionService struct {
	repo port.ExtensionRepository
}

// NewExtensionService creates a new user service instance
func NewExtensionService(repo port.ExtensionRepository) *ExtensionService {
	return &ExtensionService{
		repo,
	}
}

// CreateExtension creates a new Extension
func (ex *ExtensionService) CreateExtension(ctx echo.Context, c *domain.Extension) error {
	panic("unimplemented")
}

// GetAllExtensions gets all Extensions
func (ex *ExtensionService) GetAllExtensions(ctx echo.Context) ([]*domain.Extension, error) {
	return ex.repo.GetAllExtensions(ctx)
}

// GetExtensionByID gets a Extension by ID
func (ex *ExtensionService) GetExtensionByID(ctx echo.Context, extensionID string) (*domain.Extension, error) {
	return ex.repo.GetExtensionByID(ctx, extensionID)
}

// GetExtensionByCode gets a Extension by code
func (ex *ExtensionService) GetExtensionByCode(ctx echo.Context, code string) (*domain.Extension, error) {
	return ex.repo.GetExtensionByCode(ctx, code)
}

// DeleteExtension deletes a Extension by ID
func (ex *ExtensionService) DeleteExtension(ctx echo.Context, extensionID string) error {
	return ex.repo.DeleteExtension(ctx, extensionID)
}

//////////////////PROJECT Extension/////////////////////////

// CreateProjectExtension creates a new ProjectExtension
func (ex *ExtensionService) CreateProjectExtension(ctx echo.Context, projectID string, e *domain.Extension) error {
	return ex.repo.CreateProjectExtension(ctx, projectID, e)
}

// GetAllProjectExtensions gets all ProjectExtensions
func (ex *ExtensionService) GetAllProjectExtensions(ctx echo.Context, projectID string) ([]*domain.ProjectExtension, error) {
	return ex.repo.GetAllProjectExtensions(ctx, projectID)
}

// GetProjectExtensionByID gets a ProjectExtension by ID
func (ex *ExtensionService) GetProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.ProjectExtension, error) {
	return ex.repo.GetProjectExtensionByID(ctx, extensionID, projectID)
}

// DeleteProjectExtension deletes a ProjectExtension by ID
func (ex *ExtensionService) DeleteProjectExtension(ctx echo.Context, extensionID, projectID string) error {
	return ex.repo.DeleteProjectExtension(ctx, extensionID, projectID)
}

////////////////// ACS Extension/////////////////////////

// CreateACSProjectExtension creates a new ProjectExtension
func (ex *ExtensionService) CreateACSProjectExtension(ctx echo.Context, projectID string, e *domain.AcsCourierExtension) error {

	return ex.repo.CreateACSProjectExtension(ctx, projectID, e)
}

// GetAllACSProjectExtensions gets all ProjectExtensions
func (ex *ExtensionService) GetAllACSProjectExtensions(ctx echo.Context, projectID string) ([]*domain.AcsCourierExtension, error) {
	return ex.repo.GetAllACSProjectExtensions(ctx, projectID)
}

// GetACSProjectExtensionByID gets a ProjectExtension by ID
func (ex *ExtensionService) GetACSProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.AcsCourierExtension, error) {
	return ex.repo.GetACSProjectExtensionByID(ctx, extensionID, projectID)
}

// DeleteACSProjectExtension deletes a ProjectExtension by ID
func (ex *ExtensionService) DeleteACSProjectExtension(ctx echo.Context, extensionID, projectID string) error {
	return ex.repo.DeleteACSProjectExtension(ctx, extensionID, projectID)
}

////////////////// Courier4u Extension/////////////////////////

// CreateCourier4uProjectExtension creates a new ProjectExtension
func (ex *ExtensionService) CreateCourier4uProjectExtension(ctx echo.Context, projectID string, e *domain.Courier4uExtension) error {

	return ex.repo.CreateCourier4uProjectExtension(ctx, projectID, e)
}

// GetAllCourier4uProjectExtensions gets all ProjectExtensions
func (ex *ExtensionService) GetAllCourier4uProjectExtensions(ctx echo.Context, projectID string) ([]*domain.Courier4uExtension, error) {
	return ex.repo.GetAllCourier4uProjectExtensions(ctx, projectID)
}

// GetCourier4uProjectExtensionByID gets a ProjectExtension by ID
func (ex *ExtensionService) GetCourier4uProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.Courier4uExtension, error) {
	return ex.repo.GetCourier4uProjectExtensionByID(ctx, extensionID, projectID)
}

// DeleteCourier4uProjectExtension deletes a ProjectExtension by ID
func (ex *ExtensionService) DeleteCourier4uProjectExtension(ctx echo.Context, extensionID, projectID string) error {
	return ex.repo.DeleteCourier4uProjectExtension(ctx, extensionID, projectID)
}

////////////////// Data Synchronizer Extension/////////////////////////

// CreateSynchronizerProjectExtension creates a new ProjectExtension
func (ex *ExtensionService) CreateSynchronizerProjectExtension(ctx echo.Context, projectID string, e *domain.DataSynchronizerExtension) error {

	return ex.repo.CreateSynchronizerProjectExtension(ctx, projectID, e)
}

// UpdateSynchronizerCustomerRecievedExtension creates a new ProjectExtension
func (ex *ExtensionService) UpdateSynchronizerCustomerRecievedExtension(ctx echo.Context, projectID string, CustomerRecieved int) error {

	return ex.repo.UpdateSynchronizerCustomerRecievedExtension(ctx, projectID, CustomerRecieved)
}

// UpdateSynchronizerOrderReceivedExtension creates a new ProjectExtension
func (ex *ExtensionService) UpdateSynchronizerOrderReceivedExtension(ctx echo.Context, projectID string, OrderReceived int) error {

	return ex.repo.UpdateSynchronizerOrderReceivedExtension(ctx, projectID, OrderReceived)
}

// UpdateSynchronizerProductReceivedExtension creates a new ProjectExtension
func (ex *ExtensionService) UpdateSynchronizerProductReceivedExtension(ctx echo.Context, projectID string, ProductReceived int) error {

	return ex.repo.UpdateSynchronizerProductReceivedExtension(ctx, projectID, ProductReceived)
}

// GetAllSynchronizerProjectExtensions gets all ProjectExtensions
func (ex *ExtensionService) GetAllSynchronizerProjectExtensions(ctx echo.Context, projectID string) ([]*domain.DataSynchronizerExtension, error) {
	return ex.repo.GetAllSynchronizerProjectExtensions(ctx, projectID)
}

// GetSynchronizerProjectExtensionByID gets a ProjectExtension by ID
func (ex *ExtensionService) GetSynchronizerProjectExtensionByID(ctx echo.Context, extensionID, projectID string) (*domain.DataSynchronizerExtension, error) {
	return ex.repo.GetSynchronizerProjectExtensionByID(ctx, extensionID, projectID)
}

// DeleteSynchronizerProjectExtension deletes a ProjectExtension by ID
func (ex *ExtensionService) DeleteSynchronizerProjectExtension(ctx echo.Context, extensionID, projectID string) error {
	return ex.repo.DeleteSynchronizerProjectExtension(ctx, extensionID, projectID)
}
