// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	h "github.com/stelgkio/otoo/internal/adapter/web/view/component/profile/profile_header"
	"github.com/stelgkio/otoo/internal/core/domain"
)

func Profile(user *domain.User) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-content\" class=\"flex-fill overflow-y-lg-auto scrollbar bg-body rounded-top-4 rounded-top-start-lg-4 rounded-top-end-lg-0 border-top border-lg shadow-2\"><main class=\"container-fluid px-3 py-5 p-lg-6 p-xxl-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = h.ProfileHeader("Account Settings", 1).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form><div class=\"d-flex align-items-end justify-content-between\"><div><h4 class=\"fw-semibold mb-1\">General</h4><p class=\"text-sm text-muted\">Update your personal data.</p></div><div class=\"d-none d-md-flex gap-2\"><button type=\"submit\" hx-indicator=\"#spinner\" hx-post=\"/profile/user/update\" hx-target=\"#dashboard-content\" class=\"btn btn-sm btn-primary\"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> Update</button></div></div><hr class=\"my-6\"><div class=\"row align-items-center\"><div class=\"col-md-2\"><label class=\"form-label\" for=\"name\">Name</label></div><div class=\"col-md-8 col-xl-5\"><div class=\"\"><input type=\"text\" class=\"form-control\" name=\"name\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(user.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/profile/profile.templ`, Line: 40, Col: 89}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div></div><hr class=\"my-6\"><div class=\"row align-items-center\"><div class=\"col-md-2\"><label class=\"form-label\" for=\"last_name\">Last Name</label></div><div class=\"col-md-8 col-xl-5\"><div class=\"\"><input type=\"text\" class=\"form-control\" name=\"last_name\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(user.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/profile/profile.templ`, Line: 47, Col: 98}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div></div><hr class=\"my-6\"><div class=\"row align-items-center\"><div class=\"col-md-2\"><label class=\"form-label\" for=\"email\">Email</label></div><div class=\"col-md-8 col-xl-5\"><div class=\"\"><input type=\"email\" class=\"form-control\" name=\"email\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(user.Email)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/profile/profile.templ`, Line: 56, Col: 79}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div></div><hr class=\"my-6 d-md-none\"><div class=\"d-flex d-md-none justify-content-end gap-2 mb-6\"><button type=\"submit\" hx-indicator=\"#spinner\" hx-post=\"/profile/user/update\" hx-target=\"#dashboard-content\" class=\"btn btn-sm btn-primary \"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> Update</button></div></form><hr class=\"my-6\"><div class=\"d-flex align-items-end justify-content-between\"><div><h4 class=\"fw-semibold mb-1\">Delete Account</h4><p class=\"text-sm text-muted\">By deleting this account you are lossing all of your data!</p></div><div class=\"d-md-flex gap-2\"><button data-bs-target=\"#deleteAccountModal\" data-bs-toggle=\"modal\" class=\"btn btn-sm btn-danger\">Delete </button></div></div></main><div class=\"modal fade\" id=\"deleteAccountModal\" tabindex=\"-1\" aria-labelledby=\"deleteAccountModalLabel\" aria-hidden=\"true\"><div class=\"modal-dialog modal-dialog-centered\"><div class=\"modal-content overflow-hidden\"><div class=\"modal-header pb-0 border-0\"><h1 class=\"modal-title h4\" id=\"deleteAccountModalLabel\">Delete Account</h1><button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"modal\" aria-label=\"Close\"></button></div><div class=\"modal-body p-0\"><div class=\"px-6 py-5 border-bottom\"><h3 class=\"modal-title h4\" id=\"deleteAccountModalLabel\">Are you sure you want to delete this account?</h3></div><div class=\"px-6 py-5 bg-body-secondary d-flex justify-content-center\"><button type=\"submit\" hx-post=\"/profile/user/delete\" class=\"btn btn-sm btn-danger\">Delete </button></div></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
