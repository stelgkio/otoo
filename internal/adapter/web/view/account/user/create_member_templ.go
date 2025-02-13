// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

func CreateMeember(projectId string, errors map[string](string)) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"modals-here\" class=\"modal modal-blur show\" tabindex=\"-1\" style=\"display: block;\" aria-modal=\"true\" role=\"dialog\"><div class=\"modal-dialog modal-dialog-centered\"><div class=\"modal-content shadow-3\"><form autocomplete=\"off\" class=\"needs-validation\" novalidate><div class=\"modal-header justify-content-start\"><div class=\"icon icon-shape rounded-3 bg-primary-subtle text-primary text-lg me-4\"><i class=\"bi bi-microsoft-teams\"></i></div><div><h5 class=\"mb-1\">Add team member</h5></div><!-- Close Button in Header --><button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"modal\" aria-label=\"Close\" onclick=\"closeModal()\"></button></div><div class=\"modal-body\"><div><label class=\"form-label\">Name</label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["name"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"form-control is-invalid\" name=\"name\" placeholder=\"Name\" required type=\"text\"><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(errors["name"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/account/user/create_member.templ`, Line: 24, Col: 78}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"form-control\" name=\"name\" placeholder=\"Name\" required type=\"text\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div><label class=\"form-label\">Last Name</label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["lastname"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"form-control is-invalid\" name=\"last_name\" placeholder=\"Last Name\" required type=\"text\"><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(errors["lastname"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/account/user/create_member.templ`, Line: 33, Col: 82}
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"form-control\" name=\"last_name\" placeholder=\"Last Name\" required type=\"text\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div><label class=\"form-label\">Email </label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["email"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"form-control is-invalid\" name=\"email\" placeholder=\"Email address\" required type=\"email\"><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(errors["email"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/account/user/create_member.templ`, Line: 42, Col: 79}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"form-control\" name=\"email\" placeholder=\"Email address\" required type=\"email\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div><label class=\"form-label\">Password</label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["password"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"password-container\"><input type=\"password\" class=\"form-control is-invalid\" name=\"password\" required placeholder=\"Password\" id=\"password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;password&#39;, this)\"></i></div><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(errors["password"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/account/user/create_member.templ`, Line: 54, Col: 82}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"password-container\"><input type=\"password\" class=\"form-control\" name=\"password\" required placeholder=\"Password\" id=\"password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;password&#39;, this)\"></i></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div><label class=\"form-label\">Confirmation Password</label> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if errors["confirmation_password"] != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"password-container\"><input type=\"password\" class=\"form-control is-invalid\" name=\"confirmationpassword\" required placeholder=\"Confirmation Password\" id=\"confirm-password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;confirm-password&#39;, this)\"></i></div><div class=\"invalid-feedback\" style=\"display: block;\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(errors["confirmation_password"])
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/account/user/create_member.templ`, Line: 69, Col: 95}
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"password-container\"><input type=\"password\" class=\"form-control\" name=\"confirmationpassword\" required placeholder=\"Confirmation Password\" id=\"confirm-password\"> <i class=\"fas fa-eye toggle-password\" onclick=\"togglePassword(&#39;confirm-password&#39;, this)\"></i></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div><div class=\"form-check form-switch pt-5\"><label class=\"form-check-label\" for=\"flexSwitchCheckChecked\">Receive Notification</label> <input type=\"hidden\" name=\"receive_notification\" value=\"false\"> <input class=\"form-check-input\" type=\"checkbox\" name=\"receive_notification\" value=\"true\" checked></div></div></div><div class=\"modal-footer\"><div class=\"me-auto\"></div><button type=\"button\" class=\"btn btn-sm btn-neutral\" data-bs-dismiss=\"modal\" aria-label=\"Close\" onclick=\"closeModal()\">Close</button> <button type=\"submit\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/user/addmember/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/account/user/create_member.templ`, Line: 90, Col: 61}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"btn btn-sm btn-primary\" data-bs-dismiss=\"modal\" hx-indicator=\"#spinner\" hx-target=\"#modals-here\" hx-swap=\"innerHTML\"><span id=\"spinner\" class=\"htmx-indicator spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\" hx-boost=\"validate\"></span>Add Member</button></div></form></div></div><script>\n        \t// Function to close the modal\nfunction closeModal() {\n\tconsole.log(\"Closing modal...\");\n   var modalElement = document.getElementById('modals-here');\n    var modal = bootstrap.Modal.getInstance(modalElement);\n\n    if (modal) {\n        modal.hide();\n    } else {\n        var newModal = new bootstrap.Modal(modalElement);\n        newModal.hide();\n    }\n}\n</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
