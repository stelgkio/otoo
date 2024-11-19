// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

func Courier4uSubscriptionSuccess(projectId, extensionId string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-content\" class=\"flex-fill overflow-y-lg-auto scrollbar bg-body rounded-top-4 rounded-top-start-lg-4 rounded-top-end-lg-0 border-top border-lg shadow-2\"><main class=\"container-fluid px-6 pb-10\"><form autocomplete=\"off\"><header class=\"py-4 border-bottom\"><div class=\"row align-items-center\"><div class=\"col\"><div class=\"d-flex align-items-center gap-4\"><div><button type=\"button\" class=\"btn-close text-xs\" aria-label=\"Close\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/courier4u/courier4u_subscription_success.templ`, Line: 23, Col: 58}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\"></button></div><div class=\"vr opacity-20 my-1\"></div><h1 class=\"h4 ls-tight\">Courier4u</h1></div></div></div></header></form><div class=\"col-12\"><div class=\"col d-flex align-items-center justify-content-center\"><article class=\"article mw-read\"><!-- Success Message --><div class=\"alert alert-success text-center mb-4\" role=\"alert\">The installation of Courier4u was successful!</div><div class=\"card card-pricing text-bg-secondary border-0 shadow-4 shadow-6-hover\"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">Courier4u</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">29€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\"><p>Accelerate your shipping process with KonektorX Courier4u. Manage and send orders with speed and ease.</p></p></div></div><ul class=\"list-unstyled mt-7\"><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Automatic order status updates</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Download shipping vouchers</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Send customer notifications via email</p></li></ul></article></div></div></main></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
