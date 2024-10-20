// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	h "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/settings_header"
	"github.com/stelgkio/otoo/internal/core/domain"
)

func Team(project *domain.Project, projectExtensions []*domain.ProjectExtension, user *domain.User) templ.Component {
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
		templ_7745c5c3_Err = h.SettingsHeader("Team", 6, project.Id.String(), projectExtensions, user).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"vstack gap-5 mt-8\"><div class=\"d-flex align-items-end justify-content-between\"><div><h3 class=\"fw-semibold mb-1\">Your Team</h3><p class=\"text-sm text-muted\">The new user can not add new extesion to project.<br><strong>If you need to add admin user you have to register first.</strong></p></div><div><button type=\"button\" class=\"btn btn-neutral btn-sm\" data-bs-toggle=\"modal\" data-bs-target=\"#modalShare\"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\" hx-boost=\"validate\"></span>Add team member</button></div></div><hr class=\"my-0\"><div class=\"row justify-content-between\"><div class=\"col-md-4\"><label class=\"form-label\">Members</label></div><div class=\"col-md-8 col-xl-7\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/user/list/%s", project.Id.String()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/team/team.templ`, Line: 45, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"load\" hx-target=\"#user-list\"><div id=\"user-list\" class=\"list-group\"></div></div></div></div><div class=\"modal fade\" id=\"modalShare\" tabindex=\"-1\" aria-labelledby=\"modalShare\" aria-hidden=\"true\"><div class=\"modal-dialog modal-dialog-centered\"><div class=\"modal-content shadow-3\"><form autocomplete=\"off\" class=\"needs-validation\" novalidate><div class=\"modal-header justify-content-start\"><div class=\"icon icon-shape rounded-3 bg-primary-subtle text-primary text-lg me-4\"><i class=\"bi bi-microsoft-teams\"></i></div><div><h5 class=\"mb-1\">Add team member</h5></div></div><div class=\"modal-body\"><div><label class=\"form-label\">Name</label> <input class=\"form-control\" name=\"name\" placeholder=\"Name\" required type=\"text\"><div class=\"invalid-feedback\">Please provide a name.</div></div><div><label class=\"form-label\">Last Name</label> <input class=\"form-control\" name=\"last_name\" placeholder=\"Last Name\" required type=\"text\"><div class=\"invalid-feedback\">Please provide a last name.</div></div><div><label class=\"form-label\">Email </label> <input class=\"form-control\" name=\"email\" placeholder=\"Email address\" required type=\"email\" hx-trigger=\"keyup changed delay:250ms\" hx-post=\"/user/check-email\" hx-target=\"#user-exist\"></div><div id=\"user-exist\"><div><label class=\"form-label\">Password</label><div class=\"password-container\"><input type=\"password\" class=\"form-control\" name=\"password\" required placeholder=\"Password\" id=\"password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;password&#39;, this)\"></i></div></div><div><label class=\"form-label\">Confirmation Password</label><div class=\"password-container\"><input type=\"password\" class=\"form-control\" name=\"confirmationpassword\" required placeholder=\"Confirmation Password\" id=\"confirm-password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;confirm-password&#39;, this)\"></i></div></div></div><div><div class=\"form-check form-switch pt-5\"><label class=\"form-check-label\" for=\"flexSwitchCheckChecked\">Receive Notification</label> <input class=\"form-check-input\" type=\"checkbox\" name=\"receive_notification\" value=\"true\" checked></div></div></div><div class=\"modal-footer\"><div class=\"me-auto\"></div><button type=\"button\" class=\"btn btn-sm btn-neutral\" data-bs-dismiss=\"modal\">Close</button> <button type=\"submit\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/user/addmember/%s", project.Id.String()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/team/team.templ`, Line: 116, Col: 73}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"btn btn-sm btn-primary\" data-bs-dismiss=\"modal\" aria-label=\"Close\" hx-indicator=\"#spinner\" hx-target=\"#dashboard-content\">Add Member</button></div></form></div></div></div></main><script>   \n\n\n    function togglePassword(inputId, icon) {\n        const input = document.getElementById(inputId);\n        if (input.type === \"password\") {\n            input.type = \"text\";\n            icon.classList.remove(\"fa-eye\");\n            icon.classList.add(\"fa-eye-slash\");\n        } else {\n            input.type = \"password\";\n            icon.classList.remove(\"fa-eye-slash\");\n            icon.classList.add(\"fa-eye\");\n        }\n    }\n</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
