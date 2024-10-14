package handler

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	e "github.com/stelgkio/otoo/internal/adapter/web/view/extension"
	ac "github.com/stelgkio/otoo/internal/adapter/web/view/extension/acs_courier"
	cu "github.com/stelgkio/otoo/internal/adapter/web/view/extension/courier4u"
	nv "github.com/stelgkio/otoo/internal/adapter/web/view/extension/side_nav_list"
	et "github.com/stelgkio/otoo/internal/adapter/web/view/extension/template"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/util"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/sub"
)

// Extention get extention
func (dh *DashboardHandler) Extention(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	extensions, err := dh.extensionSvc.GetAllExtensions(ctx)
	if err != nil {
		return err
	}

	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if ctx.Request().Header.Get("HX-Request") == "true" {
		return util.Render(ctx, e.Extensions(projectID, extensions, projectExtensions))
	}
	project, user, projectID, err := GetProjectAndUser(ctx, dh)
	return util.Render(ctx, et.ExtensionTemplate(user, project.Name, projectID, extensions, projectExtensions))

}

// StripeSuccesRedirect  the redirect handler for Stripe success
func (dh *DashboardHandler) StripeSuccesRedirect(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extensionID := ctx.Param("extensionId")

	extension, err := dh.extensionSvc.GetExtensionByID(ctx, extensionID)
	if err != nil {
		return err
	}

	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := dh.userSvc.GetUserById(ctx, userID)
	project, err := dh.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if extension.Code == "asc-courier" {
		//dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, "")

		return util.Render(ctx, et.ExtentionAcsSubscriptionSuccessTemplate(user, project.Name, projectID, extensionID))
	}
	if extension.Code == "courier4u" {
		//dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, "")
		return util.Render(ctx, et.ExtentionCourier4uSubscriptionSuccessTemplate(user, project.Name, projectID, extensionID))
	}
	if extension.Code == "wallet-expences" {
		//dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, "")
	}
	if extension.Code == "team-member" {
		//dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, "")
	}
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	return util.Render(ctx, e.Extensions(projectID, nil, projectExtensions))
}

// StripeFailRedirect the redirect handler for Stripe fail
func (dh *DashboardHandler) StripeFailRedirect(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extensionID := ctx.Param("extensionId")

	extension, err := dh.extensionSvc.GetExtensionByID(ctx, extensionID)
	if err != nil {
		return err
	}

	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := dh.userSvc.GetUserById(ctx, userID)
	project, err := dh.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	if extension.Code == "asc-courier" {
		dh.extensionSvc.DeleteProjectExtension(ctx, extensionID, projectID)

		return util.Render(ctx, et.ExtentionAcsSubscriptionFailTemplate(user, project.Name, projectID, extensionID))
	}
	if extension.Code == "courier4u" {
		dh.extensionSvc.DeleteProjectExtension(ctx, extensionID, projectID)
		return util.Render(ctx, et.ExtentionCourier4uSubscriptionFailTemplate(user, project.Name, projectID, extensionID))
	}
	if extension.Code == "wallet-expences" {
		dh.extensionSvc.DeleteProjectExtension(ctx, extensionID, projectID)
	}
	if extension.Code == "team-member" {
		dh.extensionSvc.DeleteProjectExtension(ctx, extensionID, projectID)
	}
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	return util.Render(ctx, e.Extensions(projectID, nil, projectExtensions))
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

	return util.Render(ctx, ac.ASC_Courier(projectID, extension.ID.Hex(), nil, acs))
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
		return util.Render(ctx, ac.ASC_Courier(projectID, extension.ID.Hex(), nil, req))
	}
	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return util.Render(ctx, ac.ASC_Courier(projectID, extension.ID.Hex(), validationErrors, req))
	}

	req.ProjectID = projectID
	req.ExtensionID = extension.ID.Hex()
	req.CreatedAt = time.Now().UTC()
	req.IsActive = true

	dh.extensionSvc.CreateACSProjectExtension(ctx, projectID, req)

	return util.Render(ctx, ac.ASC_Courier_Subscription(os.Getenv("STRIPE_PUBLICK_KEY"), projectID, extension.ID.Hex()))
}

// AcsCourierDeActivate post form with acs courier data
func (dh *DashboardHandler) AcsCourierDeActivate(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "asc-courier")
	if err != nil {
		return err
	}
	projectExtensions, err := dh.extensionSvc.GetProjectExtensionByID(ctx, extension.ID.Hex(), projectID)
	if err != nil {
		return err
	}
	subscriptionID := projectExtensions.SubscriptionID
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	params := &stripe.SubscriptionCancelParams{}
	_, err = sub.Cancel(subscriptionID, params)
	if err != nil {
		log.Printf("Error canceling subscription: %v", err)
		return err
	}
	log.Printf("Subscription %s canceled successfully", subscriptionID)

	err = dh.extensionSvc.DeleteProjectExtension(ctx, extension.ID.Hex(), projectID)
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusMovedPermanently, ctx.Request().RequestURI)

}

// Courier4u get extention courier page
func (dh *DashboardHandler) Courier4u(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "courier4u")
	if err != nil {
		return err
	}
	//TODO:  by projectID and extensionID find AcsCourierExtension
	acs, err := dh.extensionSvc.GetCourier4uProjectExtensionByID(ctx, extension.ID.Hex(), projectID)
	if err != nil {
		return err
	}
	if acs == nil {
		acs = new(domain.Courier4uExtension)
	}

	return util.Render(ctx, cu.Courier4u(projectID, extension.ID.Hex(), nil, acs))
}

// Courier4uFormPost post form with acs courier data
func (dh *DashboardHandler) Courier4uFormPost(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "courier4u")
	if err != nil {
		return err
	}
	req := new(domain.Courier4uExtension)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Create project binding error", "error", err)
		return util.Render(ctx, cu.Courier4u(projectID, extension.ID.Hex(), nil, req))
	}
	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return util.Render(ctx, cu.Courier4u(projectID, extension.ID.Hex(), validationErrors, req))
	}

	req.ProjectID = projectID
	req.ExtensionID = extension.ID.Hex()
	req.CreatedAt = time.Now().UTC()
	req.IsActive = true

	dh.extensionSvc.CreateCourier4uProjectExtension(ctx, projectID, req)

	return util.Render(ctx, cu.Courier4uSubscriptio(os.Getenv("STRIPE_PUBLICK_KEY"), projectID, extension.ID.Hex()))
}

// Courier4uDeActivate post form with acs courier data
func (dh *DashboardHandler) Courier4uDeActivate(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "courier4u")
	if err != nil {
		return err
	}
	projectExtensions, err := dh.extensionSvc.GetProjectExtensionByID(ctx, extension.ID.Hex(), projectID)
	if err != nil {
		return err
	}
	subscriptionID := projectExtensions.SubscriptionID
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	params := &stripe.SubscriptionCancelParams{}
	_, err = sub.Cancel(subscriptionID, params)
	if err != nil {
		log.Printf("Error canceling subscription: %v", err)
		return err
	}
	log.Printf("Subscription %s canceled successfully", subscriptionID)

	err = dh.extensionSvc.DeleteProjectExtension(ctx, extension.ID.Hex(), projectID)
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusMovedPermanently, ctx.Request().RequestURI)
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

	return util.Render(ctx, ac.ASC_Courier(projectID, extension.ID.Hex(), nil, new(domain.AcsCourierExtension)))
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

// ProjectExtensionsList get extention
func (dh *DashboardHandler) ProjectExtensionsList(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	if len(projectExtensions) == 0 {
		return util.Render(ctx, nv.SideNavList(projectID, "", nil))
	}
	return util.Render(ctx, nv.SideNavList(projectID, projectExtensions[0].ExtensionID, projectExtensions))
}
