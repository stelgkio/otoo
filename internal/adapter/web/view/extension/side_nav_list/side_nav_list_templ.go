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

func SideNavList(projectId, extensionId string, extensions []*domain.ProjectExtension) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item my-1\" id=\"extensions-dropdown\"><a class=\"nav-link d-flex align-items-center rounded-pill\" href=\"#extension-components\" data-bs-toggle=\"collapse\" role=\"button\" aria-expanded=\"false\" aria-controls=\"extension-components\" onfocus=\"handleExtensionsClick()\"><i class=\"bi bi-grid-1x2-fill\"></i> <span data-i18n=\"side-nav-extensions\">Extensions</span> <span class=\"badge badge-sm rounded-pill me-n2 bg-success-subtle text-success ms-auto\"></span></a><div class=\"collapse\" id=\"extension-components\"><ul class=\"nav nav-sm flex-column mt-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(extensions) != 0 {
			for _, extension := range extensions {
				if extension.Code == "asc-courier" || extension.Code == "courier4u" || extension.Code == "redcourier" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item extension-item\" data-extension-code=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var2 string
					templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Code)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/side_nav_list/side_nav_list.templ`, Line: 30, Col: 79}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var3 string
					templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/page/%s/%s", extension.Code, projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/side_nav_list/side_nav_list.templ`, Line: 33, Col: 81}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"nav-link\" hx-push-url=\"true\" hx-target=\"#dashboard-content\" hx-indicator=\"#spinnerCourier\"><span data-i18n=\"Courier\">Courier</span> <span id=\"spinnerCourier\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span></a></li>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else if extension.Code == "team-member" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("                 ")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item\"><a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var4 string
					templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/page/%s/%s", extension.Code, projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/side_nav_list/side_nav_list.templ`, Line: 71, Col: 81}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"nav-link\" hx-indicator=\"#spinnerD\" hx-target=\"#dashboard-content\" hx-push-url=\"true\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var5 string
					templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Title)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/side_nav_list/side_nav_list.templ`, Line: 77, Col: 26}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <span id=\"spinnerD\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span></a></li>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></li><script>\n\nfunction handleExtensionsClick() {\n\t\n    // Log the click event\n     let courierItems = document.querySelectorAll('.extension-item[data-extension-code=\"asc-courier\"], .extension-item[data-extension-code=\"courier4u\"], .extension-item[data-extension-code=\"redcourier\"]');\n    \n    console.log(courierItems); // Debugging line\n\n    if (courierItems.length > 0) {       \n        for (let i = 1; i < courierItems.length; i++) {\n            console.log(`Removing: ${courierItems[i].outerHTML}`); // Debugging line\n            courierItems[i].remove();\n        }\n    }\n}\n\n\n\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
