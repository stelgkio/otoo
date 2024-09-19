package handler

import (
	"log/slog"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	e "github.com/stelgkio/otoo/internal/adapter/web/view/extension"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/util"
)

// Extention get extention
func (dh *DashboardHandler) Extention(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	extensions, err := dh.extensionSvc.GetAllExtensions(ctx)
	if err != nil {
		return err
	}
	return util.Render(ctx, e.Extensions(projectID, extensions))
}

// StripeSuccesRedirect  the redirect handler for Stripe success
func (dh *DashboardHandler) StripeSuccesRedirect(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extensionID := ctx.Param("extensionId")

	extension, err := dh.extensionSvc.GetExtensionByID(ctx, extensionID)
	if err != nil {
		return err
	}
	if extension.PriceID == "" {
		return util.Render(ctx, e.Extensions(projectID, nil))
	}
	return util.Render(ctx, e.Extensions(projectID, nil))
}

// StripeFailRedirect the redirect handler for Stripe fail
func (dh *DashboardHandler) StripeFailRedirect(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	extensions, err := dh.extensionSvc.GetAllExtensions(ctx)
	if err != nil {
		return err
	}
	return util.Render(ctx, e.Extensions(projectID, extensions))
}

// AcsCourier get extention courier page
func (dh *DashboardHandler) AcsCourier(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "asc-courier")
	if err != nil {
		return err
	}
	//TODO:  by projectID and extensionID find AcsCourierExtension
	acs, err := dh.extensionSvc.GetACSProjectExtensionByID(ctx, extension.ID.Hex(), projectID)
	if err != nil {
		return err
	}
	if acs == nil {
		acs = new(domain.AcsCourierExtension)
	}

	return util.Render(ctx, e.ASC_Courier(projectID, extension.ID.Hex(), nil, acs))
}

// AcsCourierFormPost post form with acs courier data
func (dh *DashboardHandler) AcsCourierFormPost(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "asc-courier")
	if err != nil {
		return err
	}
	req := new(domain.AcsCourierExtension)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Create project binding error", "error", err)
		return util.Render(ctx, e.ASC_Courier(projectID, extension.ID.Hex(), nil, req))
	}
	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return util.Render(ctx, e.ASC_Courier(projectID, extension.ID.Hex(), validationErrors, req))
	}

	req.ProjectID = projectID
	req.ExtensionID = extension.ID.Hex()
	req.CreatedAt = time.Now().UTC()
	req.IsActive = true

	dh.extensionSvc.CreateACSProjectExtension(ctx, projectID, req)

	return util.Render(ctx, e.ASC_Courier_Subscription(os.Getenv("STRIPE_PUBLICK_KEY"), projectID, extension.ID.Hex()))
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

// AcsCourierPage get extention courier page
func (dh *DashboardHandler) AcsCourierPage(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "asc-courier")
	if err != nil {
		return err
	}

	return util.Render(ctx, e.ASC_Courier(projectID, extension.ID.Hex(), nil, new(domain.AcsCourierExtension)))
}

// WalletExpensesPage get extention
func (dh *DashboardHandler) WalletExpensesPage(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "wallet-expences")
	if err != nil {
		return err
	}
	return util.Render(ctx, e.WalletExpenses(os.Getenv("STRIPE_PUBLICK_KEY"), projectID, extension.ID.Hex()))
}

// DataSynchronizerPage get extention
func (dh *DashboardHandler) DataSynchronizerPage(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "data-synchronizer")
	if err != nil {
		return err
	}
	return util.Render(ctx, e.DataSynchronizer(os.Getenv("STRIPE_PUBLICK_KEY"), projectID, extension.ID.Hex()))
}
