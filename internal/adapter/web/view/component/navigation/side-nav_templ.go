// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	c "github.com/stelgkio/otoo/internal/adapter/web/view/component/contact/contact-btn"
	"github.com/stelgkio/otoo/internal/core/domain"
	"unicode/utf8"
)

func SideBar(user *domain.User, projectName, projectId string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"flex-none navbar navbar-vertical navbar-expand-lg navbar-light bg-transparent show vh-lg-100 px-0 py-2\" id=\"sidebar\"><div class=\"container-fluid px-3 px-md-4 px-lg-6\"><button class=\"navbar-toggler ms-n2\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#sidebarCollapse\" aria-controls=\"sidebarCollapse\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><span class=\"navbar-toggler-icon\"></span></button> <a class=\"navbar-brand d-inline-block py-lg-1 mb-lg-5\" href=\"/dashboard\"><h1 class=\"ls-tight mylogo\">Otoo</h1></a><div class=\"navbar-user d-lg-none\"><div class=\"dropdown\"><a class=\"d-flex align-items-center\" href=\"#\" role=\"button\" data-bs-toggle=\"dropdown\" aria-haspopup=\"false\" aria-expanded=\"false\"><div><div class=\"avatar avatar-sm text-bg-secondary rounded-circle\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(getInitials(user.Name, user.LastName))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 46, Col: 108}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"d-none d-sm-block ms-3\"><span class=\"h6\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(user.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 48, Col: 70}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(user.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 48, Col: 87}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><div class=\"d-none d-md-block ms-md-2\"><i class=\"bi bi-chevron-down text-muted text-xs\"></i></div></a><div class=\"dropdown-menu dropdown-menu-end\"><div class=\"dropdown-header\"><a href=\"javascript:void(0)\" class=\"dropdown-item\" hx-get=\"/profile\" hx-target=\"#dashboard-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(user.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 61, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(user.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 61, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div><div class=\"dropdown-divider\"></div><a class=\"dropdown-item\" href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 templ.SafeURL = templ.URL("/dashboard/project/" + projectId)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var7)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><i class=\"bi bi-cpu-fill me-3\"></i> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(projectName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 68, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a> <a class=\"dropdown-item\" class=\"dropdown-item\" href=\"/dashboard\"><i class=\"bi bi-house me-3\"></i><span>Projects </span></a><div class=\"dropdown-divider\"></div><a class=\"dropdown-item\" hx-get=\"/dashboard/logout\"><i class=\"bi bi-person me-3\"></i><span data-i18n=\"side-nav-logout\">Logout </span></a></div></div></div><div class=\"collapse navbar-collapse overflow-x-hidden\" id=\"sidebarCollapse\"><ul class=\"navbar-nav\"><li class=\"nav-item my-1\"><a class=\"nav-link d-flex align-items-center rounded-pill \" href=\"#sidebar-dashboards\" data-bs-toggle=\"collapse\" role=\"button\" aria-expanded=\"true\" aria-controls=\"sidebar-dashboards\"><i class=\"bi bi-house-fill\"></i> <span data-i18n=\"side-nav-dashboard\">Dashboard</span> <span class=\"badge badge-sm rounded-pill me-n2 bg-success-subtle text-success ms-auto\"></span></a><div class=\"collapse show\" id=\"sidebar-dashboards\"><ul class=\"nav nav-sm flex-column mt-1\"><li class=\"nav-item\"><a href=\"javascript:void(0)\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dashboard/default/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 114, Col: 66}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" hx-target=\"#dashboard-content\" class=\"nav-link fw-bold\" hx-indicator=\"#spinnerDefault\" data-i18n=\"side-nav-default\">Default <span id=\"spinnerDefault\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span></a></li><li class=\"nav-item\"><a href=\"javascript:void(0)\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dashboard/customer/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 133, Col: 67}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" hx-target=\"#dashboard-content\" class=\"nav-link\" hx-indicator=\"#spinnerCustomers\" data-i18n=\"side-nav-customers\">Customers  <span id=\"spinnerCustomers\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span></a></li><li class=\"nav-item\"><a href=\"javascript:void(0)\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dashboard/product/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 152, Col: 66}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" hx-target=\"#dashboard-content\" class=\"nav-link\" hx-indicator=\"#spinnerProducts\" data-i18n=\"side-nav-products\">Products <span id=\"spinnerProducts\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span></a></li><li class=\"nav-item\"><a href=\"javascript:void(0)\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dashboard/order/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 171, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" hx-target=\"#dashboard-content\" class=\"nav-link\" hx-indicator=\"#spinnerOrders\" data-i18n=\"side-nav-orders\">Orders <span id=\"spinnerOrders\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span></a></li></ul></div></li></ul><hr class=\"navbar-divider my-5 opacity-70\"><ul class=\"navbar-nav\"><li><span class=\"nav-link text-xs fw-semibold text-uppercase text-muted ls-wide\" data-i18n=\"side-nav-resources\">Resources</span></li><li class=\"nav-item my-1\" id=\"extensions-dropdown\"><a class=\"nav-link d-flex align-items-center rounded-pill\" href=\"#extension-components\" data-bs-toggle=\"collapse\" role=\"button\" aria-expanded=\"false\" aria-controls=\"extension-components\" hx-trigger=\"load\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/project_extensions/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 208, Col: 74}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#extensions-dropdown\"><i class=\"bi bi-grid-1x2-fill\"></i> <span data-i18n=\"side-nav-extensions\">Extensions</span> <span class=\"badge badge-sm rounded-pill me-n2 bg-success-subtle text-success ms-auto\"></span></a><div class=\"collapse\" id=\"extension-components\"></div></li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if ConvertUserRole(user.Role) == "admin" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item my-1\"><a class=\"nav-link d-flex align-items-center rounded-pill\" href=\"javascript:void(0)\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/%s", projectId))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 224, Col: 56}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" hx-target=\"#dashboard-content\"><i class=\"bi bi-calendar2-plus-fill\"></i> <span data-i18n=\"side-nav-add-extensions\">Add Extensions</span> <span class=\"badge badge-sm rounded-pill me-n2 bg-warning-subtle text-warning ms-auto\"></span></a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item my-1\"><a class=\"nav-link d-flex align-items-center rounded-pill\" href=\"javascript:void(0)\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/project/settings/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/side-nav.templ`, Line: 243, Col: 62}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\" hx-push-url=\"true\"><i class=\"bi bi-gear-fill\"></i> <span data-i18n=\"side-nav-settings\">Settings</span></a></li></ul><div class=\"mt-auto\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = c.ContactBtn().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ConvertUserRole(role domain.UserRole) string {
	if role == "client" {
		return "admin"
	} else {
		return "user"
	}
}

func getInitials(name, lastName string) (string, error) {
	if name == "" || lastName == "" {
		return "", fmt.Errorf("name and last name must not be empty")
	}

	// Extract the first rune (Unicode code point) from each string
	nameInitial, _ := utf8.DecodeRuneInString(name)
	lastNameInitial, _ := utf8.DecodeRuneInString(lastName)

	// Combine the initials into a single string
	initials := string(nameInitial) + string(lastNameInitial)
	return initials, nil
}

var _ = templruntime.GeneratedTemplate
