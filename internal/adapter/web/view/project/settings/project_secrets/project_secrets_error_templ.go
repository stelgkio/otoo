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

func ProjectSecretsError(project *domain.Project, projectExtensions []*domain.ProjectExtension, user *domain.User) templ.Component {
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
		templ_7745c5c3_Err = h.SettingsHeader("Project Secrets", 2, project.Id.String(), projectExtensions, user).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<hr class=\"my-6\"><form><div class=\"d-flex align-items-end justify-content-between\"><div><h4 class=\"fw-semibold mb-1\">Secrest Reset</h4><p class=\"text-sm text-muted\">By updating your project secrets you will reset all the secrets for this project.\t\t\t\t\t\t\t</p></div><div class=\"d-none d-md-flex gap-2\"><button type=\"button\" class=\"btn btn-sm btn-primary\" hx-indicator=\"#spinner\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/project/settings/secrets/update/%s", project.Id))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/project_secrets/project_secrets_error.templ`, Line: 30, Col: 79}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> Update</button></div></div><hr class=\"my-6\"><div class=\"vstack gap-5\"><div class=\"row align-items-center g-3\"><div class=\"col-md-2\"><label class=\"form-label mb-0\" for=\"consumer_key\">ConsumerKey</label></div><div class=\"col-md-6\"><div class=\"password-container\"><input type=\"password\" class=\"form-control is-invalid\" id=\"consumer_key\" name=\"consumer_key\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(project.WoocommerceProject.ConsumerKey)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/project_secrets/project_secrets_error.templ`, Line: 44, Col: 147}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;consumer_key&#39;, this)\"></i></div></div></div><div class=\"row align-items-center g-3\"><div class=\"col-md-2\"><label class=\"form-label mb-0\" for=\"consumer_secret\">ConsumerSecret</label></div><div class=\"col-md-6\"><div class=\"password-container\"><input type=\"password\" class=\"form-control is-invalid\" id=\"consumer_secret\" name=\"consumer_secret\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(project.WoocommerceProject.ConsumerSecret)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/project_secrets/project_secrets_error.templ`, Line: 53, Col: 156}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;consumer_secret&#39;, this)\"></i></div></div></div><div class=\"col-md-auto\"><div class=\"invalid-feedback\" style=\"display: block;\">Unable to connect to the server, check your secrets and try again.</div></div></div><hr class=\"my-6 d-md-none\"><div class=\"d-flex d-md-none justify-content-end gap-2 mb-6\"><button type=\"submit\" class=\"btn btn-sm btn-primary\" hx-indicator=\"#spinner\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/project/settings/secrets/update/%s", project.Id))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/project_secrets/project_secrets_error.templ`, Line: 73, Col: 78}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> Update</button></div></form></main><script>\n    function togglePassword(inputId, icon) {\n        const input = document.getElementById(inputId);\n        if (input.type === \"password\") {\n            input.type = \"text\";\n            icon.classList.remove(\"fa-eye\");\n            icon.classList.add(\"fa-eye-slash\");\n        } else {\n            input.type = \"password\";\n            icon.classList.remove(\"fa-eye-slash\");\n            icon.classList.add(\"fa-eye\");\n        }\n    }\n</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
