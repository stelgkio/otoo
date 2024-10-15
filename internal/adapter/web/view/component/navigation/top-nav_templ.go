// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/stelgkio/otoo/internal/core/domain"
)

func TopBar(user *domain.User, projectName, projectId string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"d-none d-lg-flex py-3\"><div class=\"flex-none\"><div class=\"input-group input-group-sm input-group-inline w-rem-64 rounded-pill\"><span class=\"input-group-text rounded-start-pill\"><i class=\"bi bi-search me-2\"></i></span> <input type=\"search\" class=\"form-control ps-0 rounded-end-pill\" placeholder=\"Search\" aria-label=\"Search\"></div></div><div class=\"hstack flex-fill justify-content-end flex-nowrap gap-6 ms-auto px-6 px-xxl-8\"><button type=\"button\" class=\"btn btn-xs btn-primary rounded-pill text-nowrap\" data-bs-target=\"#connectWalletModal\" data-bs-toggle=\"modal\"><i class=\"bi bi-cpu-fill me-3\"></i> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(projectName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/top-nav.templ`, Line: 31, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button><div class=\"dropdown\"><a href=\"#\" class=\"nav-link\" data-bs-toggle=\"dropdown\" aria-expanded=\"false\"><i class=\"bi bi-sun-fill\"></i></a><div class=\"dropdown-menu\"><button type=\"button\" class=\"dropdown-item d-flex align-items-center\" data-bs-theme-value=\"light\">Light</button> <button type=\"button\" class=\"dropdown-item d-flex align-items-center\" data-bs-theme-value=\"dark\">Dark</button></div></div><div class=\"dropdown\"><a href=\"#\" class=\"nav-link\" data-bs-toggle=\"dropdown\" aria-expanded=\"false\"><i class=\"bi bi-translate\"></i></a><div class=\"dropdown-menu\"><button type=\"button\" class=\"dropdown-item d-flex align-items-center\" onclick=\"changeLanguage(&#39;en&#39;)\">English</button> <button type=\"button\" class=\"dropdown-item d-flex align-items-center\" onclick=\"changeLanguage(&#39;el&#39;)\">Ελληνικά</button></div></div><div class=\"dropdown\" id=\"notification-list\"><a href=\"javascript:void(0)\" class=\"nav-link\" id=\"dropdown-notifications\" data-bs-toggle=\"dropdown\" aria-expanded=\"false\" hx-trigger=\"load\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dashboard/notifiaction/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/top-nav.templ`, Line: 81, Col: 66}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#notification-list\"><i class=\"bi bi-bell\"></i> <span class=\"position-absolute top-0 start-100 translate-middle bg-danger rounded-circle\" style=\"width: 8px; height: 8px; padding: 0; border-radius: 50%;\"></span></a><div class=\"dropdown-menu dropdown-menu-end px-2\" aria-labelledby=\"dropdown-notifications\"><div class=\"dropdown-item d-flex align-items-center\"><h6 class=\"dropdown-header px-0\">Notifications</h6><a href=\"javascript:void(0)\" class=\"text-sm fw-semibold ms-auto\">Clear all</a></div><div class=\"dropdown-divider\"></div><div class=\"dropdown-item py-2 text-center\"><a href=\"javascript:void(0)\" class=\"fw-semibold text-muted text-primary-hover\">View all</a></div></div></div><div class=\"dropdown\"><a class=\"avatar avatar-sm rounded-circle\" href=\"#\" role=\"button\" data-bs-toggle=\"dropdown\" aria-haspopup=\"false\" aria-expanded=\"false\"><i class=\"bi bi-person-circle h5\"></i></a><div class=\"dropdown-menu dropdown-menu-end\"><div class=\"dropdown-header\"><span class=\"d-block text-heading fw-semibold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(user.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/top-nav.templ`, Line: 124, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(user.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/navigation/top-nav.templ`, Line: 124, Col: 81}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><div class=\"dropdown-divider\"></div><a class=\"dropdown-item\" href=\"/dashboard\"><i class=\"bi bi-house me-3\"></i>Projects </a><div class=\"dropdown-divider\"></div><div class=\"dropdown-divider\"></div><a class=\"dropdown-item\" hx-get=\"/dashboard/logout\"><i class=\"bi bi-person me-3\"></i>Logout</a></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
