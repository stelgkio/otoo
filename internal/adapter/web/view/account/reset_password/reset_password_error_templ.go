// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import h "github.com/stelgkio/otoo/internal/adapter/web/view/component/header"

func ResetPasswordError() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body><div class=\"row g-0 justify-content-center gradient-bottom-right start-purple middle-indigo end-pink\"><div class=\"col-md-6 col-lg-5 col-xl-5 position-fixed start-0 top-0 vh-100 overflow-y-hidden d-none d-lg-flex flex-lg-column\"><div class=\"p-12 py-xl-10 px-xl-20\"><a class=\"d-block\" href=\"/index\"><h1 class=\"display-9 mylogo text-white fw-bolder lh-tight px-sm-1\">Otoo </h1></a><div class=\"mt-16\"><h1 class=\"ls-tight fw-bolder display-6 text-white mb-5\">All in one Solution, Integration & Dashboard Analytics Faster than Ever.</h1></div></div><div class=\"mt-auto ps-16 ps-xl-20\"></div></div><div class=\"col-12 col-md-12 col-lg-7 offset-lg-5 min-vh-100 overflow-y-auto d-flex flex-column justify-content-center position-relative bg-body rounded-top-start-lg-4 border-start-lg shadow-soft-5\"><div class=\"w-md-50 mx-auto px-10 px-md-0 py-10\"><div class=\"mb-10\"><a class=\"d-inline-block d-lg-none mb-10 d-flex justify-content-center\" href=\"/\"><h1 class=\"ls-tight fw-bolder mylogo\">Otoo</h1></a><h1 class=\"ls-tight fw-bolder h3\">Reset your password.</h1><div class=\"mt-3 text-sm text-muted\"><span>Something went wrong. Try again later.</span></div></div></div></div></div><script src=\"/assets/js/main.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
