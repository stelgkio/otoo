// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/stelgkio/otoo/internal/core/domain"

func ContactForm(invalid, success bool, errors map[string](string), values *domain.ContactRequest) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-content\" class=\"flex-fill overflow-y-lg-auto scrollbar bg-body rounded-top-4 rounded-top-start-lg-4 rounded-top-end-lg-0 border-top border-lg shadow-2\"><main class=\"ccontainer-fluid px-3 py-5 p-lg-6 p-xxl-8\"><div class=\"mb-6 mb-xl-10\"><div class=\"row g-3 align-items-center\"><div class=\"col\"><h1 class=\"ls-tight\" data-i18n=\"getInTouch\">Get in Touch</h1></div><div class=\"col\"></div></div></div><form autocomplete=\"off\"><hr class=\"my-6\"><div class=\"row g-5\"><div class=\"row align-items-center g-3 mt-6\"><div class=\"col-md-2\"><label class=\"form-label mb-0\" data-i18n=\"off-canvas-modal-fullname\">Full Name</label></div><div class=\"col-md-8 col-xl-5\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["full_name"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input name=\"full_name\" type=\"text\" class=\"form-control is-invalid\" placeholder=\"Full Name\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(values.FullName)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 33, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(errors["full_name"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 36, Col: 30}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input name=\"full_name\" type=\"text\" class=\"form-control\" placeholder=\"Full Name\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(values.FullName)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 44, Col: 32}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><hr class=\"my-6\"><div class=\"row align-items-center g-3 mt-6\"><div class=\"col-md-2\"><label class=\"form-label mb-0\" data-i18n=\"off-canvas-email\">Email</label></div><div class=\"col-md-8 col-xl-5\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["email"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input name=\"email\" type=\"email\" class=\"form-control is-invalid\" placeholder=\"Email\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(values.Email)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 61, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(errors["email"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 65, Col: 26}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input name=\"email\" type=\"email\" class=\"form-control\" placeholder=\"Email\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(values.Email)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 73, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><hr class=\"my-6\"><div class=\"row align-items-center g-3 mt-6\"><div class=\"col-md-2\"><label class=\"form-label mb-0\" data-i18n=\"off-canvas-modal-phone\">Phone</label></div><div class=\"col-md-8 col-xl-5\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["phone"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input name=\"phone\" type=\"phone\" class=\"form-control is-invalid\" placeholder=\"Phone\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(values.Phone)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 91, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(errors["phone"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 95, Col: 26}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input name=\"phone\" type=\"phone\" class=\"form-control\" placeholder=\"Phone\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(values.Phone)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 103, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" required>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><hr class=\"my-6\"><div class=\"row align-items-center g-3 mt-6\"><div class=\"col-md-2\"><label class=\"form-label mb-0\" data-i18n=\"off-canvas-modal-message\">Message</label></div><div class=\"col-md-8 col-xl-5\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["message"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<textarea name=\"message\" class=\"form-control is-invalid\" placeholder=\"Message ...\" required rows=\"2\"></textarea><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(errors["message"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 124, Col: 28}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<textarea name=\"message\" class=\"form-control\" placeholder=\"Message ...\" required>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(values.Message)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form/contact-form.templ`, Line: 132, Col: 25}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</textarea>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><hr class=\"my-6 \"><div class=\"d-flex justify-content-end gap-2\"><button type=\"button\" class=\"btn btn-sm btn-neutral\" hx-get=\"/project/list\" hx-target=\"#dashboard-content\">Cancel</button> <button type=\"submit\" hx-post=\"/dashboard/contact\" hx-indicator=\"#spinner\" class=\"btn btn-sm btn-primary\" hx-target=\"#dashboard-content\"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> Submit</button></div><hr class=\"my-6 \"><div class=\"row justify-content-md-center\"><div class=\"col-md-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if invalid {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"invalid-feedback\" style=\"display: block;\">Something when wrong try again later!</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if success {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"alert alert-success\" role=\"alert\" style=\"display: block;\">The message is send succesfully!</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></form></main></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
