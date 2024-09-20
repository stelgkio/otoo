// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

func ASC_Courier(stripePublishableKey string, projectId, extensionId string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-content\" class=\"flex-fill overflow-y-lg-auto scrollbar bg-body rounded-top-4 rounded-top-start-lg-4 rounded-top-end-lg-0 border-top border-lg shadow-2\"><main class=\"container-fluid px-6 pb-10\"><form autocomplete=\"off\" x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("subscriptionForm('%s','%s','%s')", stripePublishableKey, projectId, extensionId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/asc_courier.templ`, Line: 11, Col: 130}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><header class=\"py-4 border-bottom\"><div class=\"row align-items-center\"><div class=\"col\"><div class=\"d-flex align-items-center gap-4\"><div><button type=\"button\" class=\"btn-close text-xs\" aria-label=\"Close\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/asc_courier.templ`, Line: 21, Col: 58}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\"></button></div><div class=\"vr opacity-20 my-1\"></div><h1 class=\"h4 ls-tight\">Acs Courier</h1></div></div><div class=\"col\"><div class=\"hstack gap-2 justify-content-end\"><button type=\"button\" @click=\"createSubscription\" class=\"btn btn-sm btn-neutral d-none d-sm-inline-flex\"><span class=\"pe-2\"><i class=\"bi bi-plus-circle\"></i></span><span>Subscription</span></button></div></div></div></header></form><article class=\"article mw-read\"><h2>What is Acs Courier Extention</h2><p>Accelerate your shipping process with Otoo ASC Courier. Manage and send orders with speed and ease.</p><ul class=\"list-unstyled mt-7\"><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Automatic order status updates</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Download shipping vouchers</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Send customer notifications via email</p></li></ul></article></main><script>\r\n\t\tfunction subscriptionForm(stripePublishableKey,projectId,extensionId) {\r\n\t\t\treturn {\r\n\r\n\t\t\t\tcreateSubscription() {\r\n\t\t\t\t\tconst stripe = Stripe(stripePublishableKey);\r\n\r\n\t\t\t\t\tconst value = {\r\n\t\t\t\t\t\t\r\n\t\t\t\t\t\tprojectId: projectId, // Add additional data if needed\r\n\t\t\t\t\t\textensionId: extensionId // Add additional data if needed\r\n\t\t\t\t\t};\r\n\r\n\t\t\t\t\tfetch(\"/payment\", {\r\n\t\t\t\t\t\tmethod: 'POST',\r\n\t\t\t\t\t\theaders: {\r\n\t\t\t\t\t\t\t'Content-Type': 'application/json'\r\n\t\t\t\t\t\t},\r\n\t\t\t\t\t\tbody: JSON.stringify(value)\r\n\t\t\t\t\t})\r\n\t\t\t\t\t.then(async response => {\r\n\t\t\t\t\t\tconst res = await response.json();\r\n\t\t\t\t\t\tconst id = res.id;\r\n\t\t\t\t\t\tstripe.redirectToCheckout({ sessionId: id });\r\n\t\t\t\t\t})\r\n\t\t\t\t\t.catch(error => {\r\n\t\t\t\t\t\tconsole.error('Error creating subscription:', error);\r\n\t\t\t\t\t});\r\n\t\t\t\t}\r\n\t\t\t};\r\n\t\t}\r\n\t</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
