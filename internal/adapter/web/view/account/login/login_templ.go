// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import h "github.com/stelgkio/otoo/internal/adapter/web/view/component/header"

func Login(err error) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body><div class=\"row g-0 justify-content-center gradient-bottom-right start-purple middle-indigo end-pink\"><div class=\"col-md-6 col-lg-5 col-xl-5 position-fixed start-0 top-0 vh-100 overflow-y-hidden d-none d-lg-flex flex-lg-column\"><div class=\"p-12 py-xl-10 px-xl-20\"><a class=\"d-block\" href=\"/index\"><h1 class=\"display-9 mylogo text-white fw-bolder lh-tight px-sm-1\">Otoo </h1></a><div class=\"mt-16\"><h1 class=\"ls-tight fw-bolder display-6 text-white mb-5\">All in one Solution, Integration & Dashboard Analytics Faster than Ever.</h1></div></div><div class=\"mt-auto ps-16 ps-xl-20\"></div></div><div class=\"col-12 col-md-12 col-lg-7 offset-lg-5 min-vh-100 overflow-y-auto d-flex flex-column justify-content-center position-relative bg-body rounded-top-start-lg-4 border-start-lg shadow-soft-5\"><div class=\"w-md-50 mx-auto px-10 px-md-0 py-10 mt-10\"><div class=\"mb-10\"><a class=\"d-inline-block d-lg-none mb-10 d-flex justify-content-center\" href=\"/\"><h1 class=\"ls-tight fw-bolder mylogo\">Otoo</h1></a><h1 class=\"ls-tight fw-bolder h3\">Sign in to your account</h1><div class=\"mt-3 text-sm text-muted\"><span>Dont have an account?</span> <a href=\"/register\" class=\"fw-semibold\">Sign up</a></div></div><form action=\"/login\" method=\"post\" class=\"needs-validation\"><div class=\"mb-5\"><label class=\"form-label\" for=\"email\">Email address</label> <input type=\"email\" class=\"form-control\" name=\"email\" id=\"email\" required></div><div class=\"mb-5\"><div class=\"d-flex justify-content-between gap-2 mb-2 align-items-center\"><label class=\"form-label mb-0\" for=\"password\">Password</label> <a href=\"/forgotpassword\" class=\"text-sm text-muted text-primary-hover text-underline\">Forgot password?</a></div><input type=\"password\" class=\"form-control\" name=\"password\" id=\"password\" required autocomplete=\"current-password\"></div><div class=\"mb-5\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" name=\"check_example\" id=\"check_example\"> <label class=\"form-check-label\" for=\"check_example\">Keep me logged in</label></div></div><div><input type=\"submit\" class=\"btn btn-dark w-100\" value=\"Sign in \"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if err != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"d-flex justify-content-center\" style=\"color:#f36\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(err.Error())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/account/login/login.templ`, Line: 85, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</form><div class=\"py-5 text-center\"><span class=\"text-xs text-uppercase fw-semibold\">or</span></div><div class=\"row g-2\"><div class=\"col-sm-6\"><a href=\"#\" class=\"btn btn-neutral w-100\"><span class=\"icon icon-sm pe-2\"><img src=\"./assets/img/social/github.svg\" alt=\"...\"></span>Github</a></div><div class=\"col-sm-6\"><a href=\"#\" class=\"btn btn-neutral w-100\"><span class=\"icon icon-sm pe-2\"><img src=\"./assets/img/social/google.svg\" alt=\"...\"></span>Google</a></div></div></div></div></div><script src=\"./assets/js/main.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
