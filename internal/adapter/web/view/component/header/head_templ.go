// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func HeaderComponent() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width,initial-scale=1,viewport-fit=cover\"><meta name=\"color-scheme\" content=\"dark light\"><title>Otoo: Streamline Your E-commerce Data Integration Effortlessly</title><link rel=\"preload\" href=\"/assets/css/DAGGERSQUARE.otf\" as=\"font\" crossorigin><link rel=\"stylesheet\" type=\"text/css\" href=\"/assets/css/main.css\"><link rel=\"stylesheet\" type=\"text/css\" href=\"/assets/css/utility.css\"><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.2/font/bootstrap-icons.css\"><link rel=\"stylesheet\" href=\"https://api.fontshare.com/v2/css?f=satoshi@900,700,500,300,401,400&amp;display=swap\"><link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css\"><link rel=\"icon\" type=\"image/png\" href=\"/assets/img/favicon1.png\"><script src=\"https://unpkg.com/htmx.org@2.0.0\" integrity=\"sha384-wS5l5IKJBvK6sPTKa2WZ1js3d947pvWXbPJ1OmWfEuxLgeHcEbjUUA5i9V5ZkpCw\" crossorigin=\"anonymous\"></script><link rel=\"stylesheet\" type=\"text/css\" href=\"/assets/css/form.css\"><script src=\"https://cdn.jsdelivr.net/npm/apexcharts\"></script><script src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\" defer></script><script src=\"https://js.stripe.com/v3/\"></script></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
