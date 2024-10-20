// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/stelgkio/otoo/internal/core/domain"

func ProjectListPage(projects []*domain.Project, user *domain.User) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-content\" class=\"flex-fill overflow-y-lg-auto scrollbar bg-body rounded-top-4 rounded-top-start-lg-4 rounded-top-end-lg-0 border-top border-lg shadow-2\"><main class=\"container-fluid px-3 py-5 p-lg-6 p-xxl-8\"><div class=\"mb-6 mb-xl-10\"><div class=\"row g-3 align-items-center\"><div class=\"col\"><h1 class=\"ls-tight\">Project</h1></div><div class=\"col\"></div></div></div><div class=\"row g-3 g-xxl-6 d-flex justify-content-center\"><div class=\"col-xxl-6\"><div id=\"cryptoModal\" tabindex=\"-1\" aria-labelledby=\"cryptoModalLabel\" aria-hidden=\"true\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(projects) == 0 {
			if ConvertUserRole(user.Role) == "admin" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card border-primary-hover shadow-soft-3-hover d-flex justify-content-center\"><div class=\"modal-content overflow-hidden\"><div class=\"card-body p-0\"><div class=\"px-6 py-5 bg-body-secondary d-flex justify-content-center\"><button hx-get=\"/project/createform\" hx-push-url=\"true\" hx-target=\"#dashboard-content\" class=\"btn btn-sm btn-dark\"><i class=\"bi bi-plus-circle me-2\"></i>Create New Project</button></div></div></div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card border-primary-hover shadow-soft-3-hover d-flex justify-content-center\"><div class=\"modal-content overflow-hidden\"><div class=\"card-body p-0\"><div class=\"px-6 py-5 border-bottom\"><input id=\"project-search\" type=\"text\" class=\"form-control\" placeholder=\"Search project\" aria-label=\"Search\"></div><div id=\"project-list\" class=\"p-2\" style=\"max-height: 400px; overflow-y: auto;\"><div class=\"vstack\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, p := range projects {
				if p.ProjectType == "Shopify" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div data-project-type=\"pro\" class=\"position-relative d-flex gap-3 p-4 rounded bg-body-secondary-hover\"><div class=\"d-flex flex-fill justify-content-center\"><div class=\"justify-content-center\"><img src=\"/assets/img/marketing/shopify-logo.svg\" class=\"w-rem-5 flex-none \" alt=\"...\"> <a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var2 templ.SafeURL = templ.URL("/dashboard/project/" + p.Id.String())
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"stretched-link text-heading fw-bold\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var3 string
					templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(p.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/list/project-list.templ`, Line: 73, Col: 26}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div data-project-type=\"pro\" class=\"position-relative d-flex gap-3 p-4 rounded bg-body-secondary-hover\"><div class=\"d-flex flex-fill justify-content-center\"><div class=\"justify-content-center\"><img src=\"/assets/img/marketing/WooCommerce_logo.svg.png\" class=\"w-rem-5 flex-none \" alt=\"...\"> <a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var4 templ.SafeURL = templ.URL("/dashboard/project/" + p.Id.String())
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"stretched-link text-heading fw-bold\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var5 string
					templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(p.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/project/list/project-list.templ`, Line: 92, Col: 26}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if ConvertUserRole(user.Role) == "admin" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"px-6 py-5 bg-body-secondary d-flex justify-content-center\"><button class=\"btn btn-sm btn-dark\" hx-get=\"/project/createform\" hx-push-url=\"true\" hx-target=\"#dashboard-content\"><i class=\"bi bi-plus-circle me-2\"></i>Create New Project</button></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></main></div>")
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
