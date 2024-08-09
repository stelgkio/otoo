package handler

import (
	"os"

	"github.com/labstack/echo/v4"
	e "github.com/stelgkio/otoo/internal/adapter/web/view/extension"
	"github.com/stelgkio/otoo/internal/core/util"
)

// Extention get extention
func (dh *DashboardHandler) Extention(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	extensions, err := dh.extensionSvc.GetAllExtensions(ctx)
	if err != nil {
		return err
	}
	return util.Render(ctx, e.Extentions(projectID, extensions))
}

// Extention get extention
func (dh *DashboardHandler) StripeExtentionReturn(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	extensions, err := dh.extensionSvc.GetAllExtensions(ctx)
	if err != nil {
		return err
	}
	return util.Render(ctx, e.Extentions(projectID, extensions))
}

// AcsCourier get extention courier page
func (dh *DashboardHandler) AcsCourier(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "asc-courier")
	if err != nil {
		return err
	}

	return util.Render(ctx, e.ASC_Courier(os.Getenv("STRIPE_PUBLICK_KEY"), projectID, extension.ID.Hex()))
}

// WalletExpenses get extention
func (dh *DashboardHandler) WalletExpenses(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "wallet-expences")
	if err != nil {
		return err
	}
	return util.Render(ctx, e.WalletExpenses(os.Getenv("STRIPE_PUBLICK_KEY"), projectID, extension.ID.Hex()))
}

// DataSynchronizer get extention
func (dh *DashboardHandler) DataSynchronizer(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "data-synchronizer")
	if err != nil {
		return err
	}
	return util.Render(ctx, e.DataSynchronizer(os.Getenv("STRIPE_PUBLICK_KEY"), projectID, extension.ID.Hex()))
}
