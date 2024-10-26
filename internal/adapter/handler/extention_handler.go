package handler

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	page "github.com/stelgkio/otoo/internal/adapter/web/view/admin/page"
	e "github.com/stelgkio/otoo/internal/adapter/web/view/extension"
	ac "github.com/stelgkio/otoo/internal/adapter/web/view/extension/acs_courier"
	cu "github.com/stelgkio/otoo/internal/adapter/web/view/extension/courier4u"
	nv "github.com/stelgkio/otoo/internal/adapter/web/view/extension/side_nav_list"
	et "github.com/stelgkio/otoo/internal/adapter/web/view/extension/template"
	acscourier "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/acs-courier"
	scu "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/courier4u"
	stm "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/template"

	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/util"
	r "github.com/stelgkio/otoo/internal/core/util"
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
		dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, "")

		return util.Render(ctx, et.ExtentionAcsSubscriptionSuccessTemplate(user, project.Name, projectID, extensionID))
	}
	if extension.Code == "courier4u" {
		dh.extensionSvc.CreateProjectExtension(ctx, projectID, extension, 30, "")
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

// AcsCourierSettingsFormPost post form with acs courier data
func (dh *DashboardHandler) AcsCourierSettingsFormPost(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := dh.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	project, err := dh.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "asc-courier")
	if err != nil {
		return err
	}
	req := new(domain.AcsCourierExtension)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Create project binding error", "error", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": errors.New("Invalid request").Error(),
		})
	}
	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": util.ConcatenateErrors(validationErrors),
		})
	}

	req.ProjectID = projectID
	req.ExtensionID = extension.ID.Hex()
	req.CreatedAt = time.Now().UTC()
	req.IsActive = true

	err = dh.extensionSvc.CreateACSProjectExtension(ctx, projectID, req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, acscourier.SettingsACSCourier(projectID, projectExtensions, user, make(map[string]string), req))
	}
	return r.Render(ctx, stm.AcsCourierTemplate(user, project.Name, projectID, project, projectExtensions, make(map[string]string), req))
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

	}
	log.Printf("Subscription %s canceled successfully", subscriptionID)

	err = dh.extensionSvc.DeleteProjectExtension(ctx, extension.ID.Hex(), projectID)
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/extension/%s", projectID))
	return ctx.String(http.StatusOK, "Redirecting...")

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

// Courier4uFormPost post form with acs courier data
func (dh *DashboardHandler) Courier4uSettingsFormPost(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := dh.userSvc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	project, err := dh.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	projectExtensions, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	extension, err := dh.extensionSvc.GetExtensionByCode(ctx, "courier4u")
	if err != nil {
		return err
	}
	req := new(domain.Courier4uExtension)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Create project binding error", "error", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": errors.New("Invalid request").Error(),
		})
	}
	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": util.ConcatenateErrors(validationErrors),
		})
	}

	req.ProjectID = projectID
	req.ExtensionID = extension.ID.Hex()
	req.CreatedAt = time.Now().UTC()
	req.IsActive = true

	err = dh.extensionSvc.CreateCourier4uProjectExtension(ctx, projectID, req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, scu.SettingsCourier4u(projectID, projectExtensions, user, make(map[string]string), req))
	}
	return r.Render(ctx, stm.Courier4uTemplate(user, project.Name, projectID, project, projectExtensions, make(map[string]string), req))
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
	}
	log.Printf("Subscription %s canceled successfully", subscriptionID)

	err = dh.extensionSvc.DeleteProjectExtension(ctx, extension.ID.Hex(), projectID)
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("HX-Redirect", fmt.Sprintf("/extension/%s", projectID))
	return ctx.String(http.StatusOK, "Redirecting...")
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

// AddManualExtensionForm get extention
func (dh *DashboardHandler) AddManualExtensionForm(ctx echo.Context) error {
	key := ctx.Param("key")
	if key == os.Getenv("EXTENSION_KEY") {
		return util.Render(ctx, page.AddExtensionForm())
	}
	return ctx.JSON(http.StatusBadRequest, "Invalid key")
}

// AddManualExtensionForm get extention
func (dh *DashboardHandler) GetAllAvailableExtensios(ctx echo.Context) error {
	extension, err := dh.extensionSvc.GetAllExtensions(ctx)
	if err != nil {
		return err
	}
	return util.Render(ctx, page.AvailableExtension(extension))

}

// AddManualExtensionForm get extention
func (dh *DashboardHandler) ExtensionTable(ctx echo.Context) error {
	projectId := ctx.QueryParam("project")
	extension, err := dh.extensionSvc.GetAllProjectExtensions(ctx, projectId)
	if err != nil {
		return err
	}
	return util.Render(ctx, page.ExtensionTable(extension, projectId))

}

func (dh *DashboardHandler) DeleteProjectExtension(ctx echo.Context) error {
	projectextensionID := ctx.Param("Id")

	err := dh.extensionSvc.DeleteProjectExtensionByID(ctx, projectextensionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (dh *DashboardHandler) AddProjectExtension(ctx echo.Context) error {

	type extension struct {
		ProjectID   string `form:"project"`
		ExtensionID string `form:"extension"`
		Period      int    `form:"period"`
	}
	req := new(extension)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Create project binding error", "error", err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ext, err := dh.extensionSvc.GetExtensionByID(ctx, req.ExtensionID)
	err = dh.extensionSvc.CreateProjectExtension(ctx, req.ProjectID, ext, req.Period, domain.CustomSubsctiption)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	projExtension, err := dh.extensionSvc.GetAllProjectExtensions(ctx, req.ProjectID)
	if err != nil {
		return err
	}
	return util.Render(ctx, page.ExtensionTable(projExtension, req.ProjectID))
}
