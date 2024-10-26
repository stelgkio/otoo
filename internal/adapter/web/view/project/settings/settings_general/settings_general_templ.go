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

var settingGeneralPoperHandle = templ.NewOnceHandle()

func SettingsGeneral(user *domain.User, project *domain.Project, projectExtensions []*domain.ProjectExtension) templ.Component {
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
		templ_7745c5c3_Err = h.SettingsHeader("Project Settings", 1, project.Id.String(), projectExtensions, user).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form><div class=\"d-flex align-items-end justify-content-between\"><div><h4 class=\"fw-semibold mb-1\">General</h4><p class=\"text-sm text-muted\">Update your project data.</p></div><div class=\"d-none d-md-flex gap-2\"><button type=\"submit\" hx-indicator=\"#spinner\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/project/settings/update/%s", project.Id))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/settings_general/settings_general.templ`, Line: 30, Col: 71}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\" class=\"btn btn-sm btn-primary\"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> Update</button></div></div><hr class=\"my-6\"><div class=\"row align-items-center\"><div class=\"col-md-2\"><label class=\"form-label\" for=\"name\">Name</label></div><div class=\"col-md-8 col-xl-5\"><div class=\"\"><input type=\"text\" class=\"form-control\" name=\"name\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(project.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/settings_general/settings_general.templ`, Line: 43, Col: 92}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div></div><hr class=\"my-6\"><div class=\"row align-items-center\"><div class=\"col-md-2\"><label class=\"form-label\" for=\"description\">Description</label></div><div class=\"col-md-8 col-xl-5\"><div class=\"\"><input type=\"text\" class=\"form-control\" name=\"description\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(project.Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/settings_general/settings_general.templ`, Line: 50, Col: 106}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div></div><hr class=\"my-6\"><div class=\"row align-items-center\"><div class=\"col-md-2\"><label class=\"form-label\" for=\"domain\">Domain</label></div><div class=\"col-md-8 col-xl-5\"><div class=\"\"><input type=\"text\" class=\"form-control\" name=\"domain\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(project.WoocommerceProject.Domain)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/settings_general/settings_general.templ`, Line: 58, Col: 102}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" disabled></div></div><div class=\"col-md-2\"><span class=\"text-muted text-opacity-60 text-opacity-100-hover\" tabindex=\"0\" data-bs-toggle=\"popover\" data-bs-trigger=\"hover focus\" data-bs-placement=\"right\" data-bs-html=\"true\" data-bs-content=\"Contact our support team to change your domain.\" role=\"button\"><i class=\"bi bi-info-circle\"></i></span></div></div><hr class=\"my-6 d-md-none\"><div class=\"d-flex d-md-none justify-content-end gap-2 mb-6\"><button type=\"submit\" hx-indicator=\"#spinner\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/project/settings/update/%s", project.Id))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/settings_general/settings_general.templ`, Line: 79, Col: 70}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\" class=\"btn btn-sm btn-primary \"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> Update</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if ConvertUserRole(user.Role) == "admin" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<hr class=\"my-6\"><div class=\"d-flex align-items-end justify-content-between\"><div><h4 class=\"fw-semibold mb-1\">Delete Project</h4><p class=\"text-sm text-muted\">By deleting this project you are lossing all of your data!<br><strong>Make sure you have deactivate all of your extensions!</strong></p></div><div class=\"d-md-flex gap-2\"><button data-bs-target=\"#deleteProjectModal\" data-bs-toggle=\"modal\" class=\"btn btn-sm btn-danger\">Delete </button></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main><div class=\"modal fade\" id=\"deleteProjectModal\" tabindex=\"-1\" aria-labelledby=\"deleteProjectModalLabel\" aria-hidden=\"true\"><div class=\"modal-dialog modal-dialog-centered\"><div class=\"modal-content overflow-hidden\"><div class=\"modal-header pb-0 border-0\"><h1 class=\"modal-title h4\" id=\"deleteProjectModalLabel\">Delete Project</h1><button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"modal\" aria-label=\"Close\"></button></div><div class=\"modal-body p-0\"><div class=\"px-6 py-5 border-bottom\"><h3 class=\"modal-title h4\" id=\"deleteProjectModalLabel\">Are you sure you want to delete this project?</h3></div><div class=\"px-6 py-5 bg-body-secondary d-flex justify-content-center\"><button type=\"submit\" id=\"spinner\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/project/settings/delete/%s", project.Id))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/settings/settings_general/settings_general.templ`, Line: 135, Col: 72}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"btn btn-sm btn-danger\"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> Delete </button></div></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n\tvar popoverTriggerList2 = document.querySelectorAll('[data-bs-toggle=\"popover\"]')\n\tvar popoverList2 = [...popoverTriggerList2].map(popoverTriggerEl2 => new bootstrap.Popover(popoverTriggerEl2))\n\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = settingGeneralPoperHandle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
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

var _ = templruntime.GeneratedTemplate
