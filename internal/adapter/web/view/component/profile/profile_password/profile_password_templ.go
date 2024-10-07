// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	h "github.com/stelgkio/otoo/internal/adapter/web/view/component/profile/profile_header"
)

func ProfilePassword() templ.Component {
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
		templ_7745c5c3_Err = h.ProfileHeader("Security & Password", 2).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<hr class=\"my-6\"><div class=\"d-flex align-items-end justify-content-between\"><div><h4 class=\"fw-semibold mb-1\">Password Reset</h4><p class=\"text-sm text-muted\">By filling your data you get a much better experience using our website.</p></div><div class=\"d-none d-md-flex gap-2\"><button type=\"button\" class=\"btn btn-sm btn-neutral\">Cancel</button> <button type=\"button\" class=\"btn btn-sm btn-primary\">Save</button></div></div><hr class=\"my-6\"><form><div class=\"vstack gap-5\"><div class=\"row align-items-center g-3\"><div class=\"col-md-2\"><label class=\"form-label mb-0\">Current password</label></div><div class=\"col-md-6\"><div class=\"password-container\"><input type=\"password\" class=\"form-control\" id=\"current-password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;current-password&#39;, this)\"></i></div></div></div><div class=\"row align-items-center g-3\"><div class=\"col-md-2\"><label class=\"form-label mb-0\">New password</label></div><div class=\"col-md-6\"><div class=\"password-container\"><input type=\"password\" class=\"form-control\" id=\"new-password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;new-password&#39;, this)\"></i></div></div></div><div class=\"row align-items-center g-3\"><div class=\"col-md-2\"><label class=\"form-label mb-0\">Confirm password</label></div><div class=\"col-md-6\"><div class=\"password-container\"><input type=\"password\" class=\"form-control\" id=\"confirm-password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;confirm-password&#39;, this)\"></i></div></div></div></div><hr class=\"my-6 d-md-none\"><div class=\"d-flex d-md-none justify-content-end gap-2 mb-6\"><button type=\"button\" class=\"btn btn-sm btn-neutral\">Cancel</button> <button type=\"submit\" class=\"btn btn-sm btn-primary\">Save</button></div></form></main><script>\n    function togglePassword(inputId, icon) {\n        const input = document.getElementById(inputId);\n        if (input.type === \"password\") {\n            input.type = \"text\";\n            icon.classList.remove(\"fa-eye\");\n            icon.classList.add(\"fa-eye-slash\");\n        } else {\n            input.type = \"password\";\n            icon.classList.remove(\"fa-eye-slash\");\n            icon.classList.add(\"fa-eye\");\n        }\n    }\n</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
