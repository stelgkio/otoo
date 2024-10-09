// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	h "github.com/stelgkio/otoo/internal/adapter/web/view/component/header"
)

func ResetPasswordForm(BadRequest int, email string, verr validator.ValidationErrors, err error) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-theme=\"light\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = h.HeaderComponent().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body><div class=\"row g-0 justify-content-center gradient-bottom-right start-purple middle-indigo end-pink\"><div class=\"col-md-6 col-lg-5 col-xl-5 position-fixed start-0 top-0 vh-100 overflow-y-hidden d-none d-lg-flex flex-lg-column\"><div class=\"p-12 py-xl-10 px-xl-20\"><a class=\"d-block\" href=\"/index\"><h1 class=\"display-9 mylogo text-white fw-bolder lh-tight px-sm-1\">KonektorX </h1></a><div class=\"mt-16\"><h1 class=\"ls-tight fw-bolder display-6 text-white mb-5\">All in one Solution, Integration & Dashboard Analytics Faster than Ever.</h1></div></div><div class=\"mt-auto ps-16 ps-xl-20\"></div></div><div class=\"col-12 col-md-12 col-lg-7 offset-lg-5 min-vh-100 overflow-y-auto d-flex flex-column justify-content-center position-relative bg-body rounded-top-start-lg-4 border-start-lg shadow-soft-5\"><div class=\"w-md-50 mx-auto px-10 px-md-0 py-10\"><div class=\"mb-10\"><a class=\"d-inline-block d-lg-none mb-10 d-flex justify-content-center\" href=\"/\"><h1 class=\"ls-tight fw-bolder mylogo\">KonektorX</h1></a><h1 class=\"ls-tight fw-bolder h3\">Reset your password.</h1><div class=\"mt-3 text-sm text-muted\"><span>Do you remeber your account?</span> <a href=\"/login\" class=\"fw-semibold\">Sign in</a> to your account.</div></div><form action=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.URL(fmt.Sprintf("/resetpassword/%s", email))
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" method=\"post\" class=\"needs-validation\" novalidate><div class=\"row g-5\"><div class=\"col-sm-12\"><label class=\"form-label\">Password</label> <input id=\"password\" name=\"password\" type=\"password\" class=\"form-control\" required min=\"8\" onkeyup=\"validatePassword()\" autocomplete=\"off\"><div id=\"PasswordMessage\" class=\"invalid-feedback\">Please enter password.</div></div><div class=\"col-sm-12\"><label class=\"form-label\">Confirmation Password</label> <input id=\"ConfirmationPassword\" name=\"ConfirmationPassword\" type=\"password\" class=\"form-control\" required min=\"8\" oninput=\"validateConfirmationPassword()\"><div id=\"confirmMessage\" class=\"invalid-feedback\">Please enter confirmation password.</div></div><div id=\"lengthMessage\" class=\"invalid-feedback\">Please Enter Password.</div><div class=\"col-sm-12\"><input type=\"submit\" class=\"btn btn-dark w-100\" value=\"Reset Password\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if BadRequest != 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"d-flex justify-content-center\" style=\"color:#f36\">Something went wrong, please try again later</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if verr != nil {
			for _, e := range verr {
				if e.Field() == "Password" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"d-flex justify-content-center\" style=\"color:#f36\">Invalid password. Password length must be 8 characters or more.</div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if e.Field() == "Confirmation" {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"d-flex justify-content-center\" style=\"color:#f36\">Invalid confirmation password</div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
		}
		if err != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"d-flex justify-content-center\" style=\"color:#f36\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/account/reset_password/reset_password.templ`, Line: 109, Col: 23}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></form></div></div></div><script src=\"/assets/js/main.js\"></script><script>\r\n            // Example starter JavaScript for disabling form submissions if there are invalid fields\r\n(function () {\r\n  'use strict'\r\n\r\n  // Fetch all the forms we want to apply custom Bootstrap validation styles to\r\n  var forms = document.querySelectorAll('.needs-validation')\r\n\r\n  // Loop over them and prevent submission\r\n  Array.prototype.slice.call(forms)\r\n    .forEach(function (form) {\r\n      form.addEventListener('submit', function (event) {\r\n        if (!form.checkValidity()) {\r\n          event.preventDefault()\r\n          event.stopPropagation()\r\n        }\r\n       if( !submitConfirmationPassword()){\r\n          event.preventDefault()\r\n          event.stopPropagation()\r\n       }\r\n\r\n        form.classList.add('was-validated')\r\n      }, false)\r\n    })\r\n    //   // Add event listener to password field for validation\r\n    //     var passwordField = document.getElementById(\"password\");\r\n    //     passwordField.addEventListener(\"onchange\", validatePassword);\r\n})()\r\n\r\n            // Function to validate confirmation password when leaving password field\r\n    function validateConfirmationPassword() {\r\n        console.log(\"validateConfirmationPassword\");\r\n        var password = document.getElementById(\"password\").value;\r\n        var confirm_password = document.getElementById(\"ConfirmationPassword\");\r\n        var message = document.getElementById(\"confirmMessage\");\r\n\r\n\r\n         if (confirm_password.value.length < 8) {\r\n            confirm_password.classList.add(\"is-invalid\");\r\n            message.style.display = \"block\";\r\n        } else {\r\n             confirm_password.classList.remove(\"is-invalid\");\r\n            message.style.display = \"none\";\r\n        }\r\n        if (password !== confirm_password.value) {\r\n            confirm_password.classList.add(\"is-invalid\");\r\n            message.style.display = \"block\";\r\n        } else {\r\n            confirm_password.classList.remove(\"is-invalid\");\r\n            message.style.display = \"none\";\r\n        }\r\n       \r\n  \r\n        var result=password.value !== \"\" && password.value === confirm_password.value;\r\n        return result;\r\n    }\r\n    \r\n\r\n    function validatePassword() {\r\n    console.log(\"validatePassword\");\r\n    var password = document.getElementById(\"password\");\r\n    var confirm_password = document.getElementById(\"ConfirmationPassword\");\r\n    var message = document.getElementById(\"confirmMessage\");\r\n    var passwordMessage = document.getElementById(\"PasswordMessage\");\r\n    \r\n    // Check if password and confirmation password match\r\n    if (confirm_password.value !== \"\" && password.value !== confirm_password.value) {\r\n        confirm_password.classList.add(\"is-invalid\");\r\n        message.innerText = \"Passwords do not match.\";\r\n        message.style.display = \"block\";\r\n    } else {\r\n        confirm_password.classList.remove(\"is-invalid\");\r\n        message.style.display = \"none\";\r\n    }\r\n    \r\n    // Check if password is at least 8 characters long\r\n    if (password.value.length < 8) {\r\n        password.classList.add(\"is-invalid\");\r\n        passwordMessage.innerText = \"Password must be at least 8 characters long.\";\r\n        passwordMessage.style.display = \"block\";\r\n    } else {\r\n        password.classList.remove(\"is-invalid\");\r\n        passwordMessage.style.display = \"none\";\r\n    }\r\n}\r\n     function submitConfirmationPassword() {\r\n        console.log(\"validateConfirmationPassword\");\r\n        var password = document.getElementById(\"password\");\r\n        var confirm_password = document.getElementById(\"ConfirmationPassword\");\r\n         if(password.value.length >= 8) {\r\n            if(password.value === confirm_password.value){\r\n                return true;\r\n            }else {\r\n                return false;\r\n            }\r\n        }\r\n\r\n     }\r\n</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
