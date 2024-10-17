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

func Extensions(projectId string, extensions []*domain.Extension, projectExtensions []*domain.ProjectExtension) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-content\" class=\"flex-fill overflow-y-lg-auto scrollbar bg-body rounded-top-4 rounded-top-start-lg-4 rounded-top-end-lg-0 border-top border-lg shadow-2\"><main class=\"container-fluid px-3 py-5 p-lg-6 p-xxl-8\"><div class=\"mb-6 mb-xl-10\"><div class=\"row g-3 align-items-center\"><div class=\"col\"><h1 class=\"ls-tight\">Extensions</h1></div></div></div><div class=\"row row-cols-1 row-cols-sm-2 row-cols-lg-3 g-6 justify-content-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, extension := range extensions {
			if extension.Code == "asc-courier" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- ASC Courier Card --> <div class=\"col\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if domain.ContainsExtensionID(projectExtensions, extension.ID.Hex()) {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-primary border-0 shadow-4 shadow-6-hover card-disabled\"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">ASC Courier</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">29€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var2 string
					templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 34, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Start Shipping with KonnektorX!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var3 string
					templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/asc-courier/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 40, Col: 72}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" class=\"btn btn-sm btn-square btn-dark stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-download\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-primary border-0 shadow-4 shadow-6-hover \"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">ASC Courier</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">29€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var4 string
					templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 58, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Start Shipping with KonektorX!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var5 string
					templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/asc-courier/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 64, Col: 72}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" class=\"btn btn-sm btn-square btn-dark stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-arrow-right\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"list-unstyled mt-7\"><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Automatic order status updates</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Download shipping vouchers</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Send customer notifications via email</p></li></ul></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if extension.Code == "courier4u" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Wallet & Expenses Overview Card --> <div class=\"col\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if domain.ContainsExtensionID(projectExtensions, extension.ID.Hex()) {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-secondary border-0 shadow-4 shadow-6-hover card-disabled\"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var6 string
					templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Title)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 102, Col: 64}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">29€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 string
					templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 107, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Start Shipping with KonektorX!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/courier4u/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 113, Col: 70}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" class=\"btn btn-sm btn-square btn-dark stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-download\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-secondary border-0 shadow-4 shadow-6-hover \"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Title)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 126, Col: 64}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">29€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var10 string
					templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 131, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Start Shipping with KonektorX!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var11 string
					templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/courier4u/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 137, Col: 70}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" class=\"btn btn-sm btn-square btn-dark stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-arrow-right\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"list-unstyled mt-7\"><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Automatic order status updates</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Download shipping vouchers</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Send customer notifications via email</p></li></ul></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if extension.Code == "wallet-expences" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Wallet & Expenses Overview Card --> <div class=\"col\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if domain.ContainsExtensionID(projectExtensions, extension.ID.Hex()) {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-secondary border-0 shadow-4 shadow-6-hover card-disabled\"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">Wallet & Expenses</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">23€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var12 string
					templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 180, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Optimize Your Finances with Otoo!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var13 string
					templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/wallet-expences/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 186, Col: 76}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" class=\"btn btn-sm btn-square btn-dark stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-download\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-secondary border-0 shadow-4 shadow-6-hover \"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">Wallet & Expenses</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">23€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var14 string
					templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 204, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Optimize Your Finances with Otoo!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var15 string
					templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/wallet-expences/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 210, Col: 76}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"true\" class=\"btn btn-sm btn-square btn-dark stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-arrow-right\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"list-unstyled mt-7\"><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-secondary-subtle text-secondary me-3\"><i class=\"bi bi-check\"></i></div><p>Integrate with Facebook campaigns</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-secondary-subtle text-secondary me-3\"><i class=\"bi bi-check\"></i></div><p>Monitor courier charges</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-secondary-subtle text-secondary me-3\"><i class=\"bi bi-check\"></i></div><p>Automatic tax calculations</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-secondary-subtle text-secondary me-3\"><i class=\"bi bi-check\"></i></div><p>Priority support for Otoo users</p></li></ul></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if extension.Code == "data-synchronizer" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Data Synchronizer Card --> <div class=\"col\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if domain.ContainsExtensionID(projectExtensions, extension.ID.Hex()) {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-dark shadow-4 shadow-6-hover card-disabled\"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">Data Synchronizer</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">290€/yr</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var16 string
					templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 259, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Synchronize with Otoo!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var17 string
					templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/data-synchronizer/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 265, Col: 78}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var18 string
					templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/data-synchronizer/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 266, Col: 83}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"btn btn-sm btn-square btn-white stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-download\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-dark shadow-4 shadow-6-hover \"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">Data Synchronizer</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">290€/yr</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var19 string
					templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 283, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Synchronize with Otoo!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var20 string
					templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/data-synchronizer/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 289, Col: 78}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var21 string
					templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/data-synchronizer/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 290, Col: 83}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"btn btn-sm btn-square btn-white stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-arrow-right\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"list-unstyled mt-7\"><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Retrieve customer data seamlessly</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Sync product information and variations</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Access full order history</p></li><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Build custom analytics and reports</p></li></ul></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if extension.Code == "team-member" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Data Synchronizer Card --> <div class=\"col\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if domain.ContainsExtensionID(projectExtensions, extension.ID.Hex()) {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-dark shadow-4 shadow-6-hover card-disabled\"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var22 string
					templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Title)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 333, Col: 64}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">20€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var23 string
					templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 338, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Add Team Member to your project!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var24 string
					templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/data-synchronizer/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 344, Col: 78}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var25 string
					templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/data-synchronizer/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 345, Col: 83}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"btn btn-sm btn-square btn-white stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-download\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-pricing text-bg-dark shadow-4 shadow-6-hover \"><div class=\"p-6\"><h3 class=\"text-reset ls-tight mb-1\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var26 string
					templ_7745c5c3_Var26, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Title)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 357, Col: 64}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var26))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><div class=\"d-flex align-items-center my-5\"><span class=\"d-block display-5 text-reset\">20€/mo</span></div><p class=\"text-reset text-opacity-75 mb-4\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var27 string
					templ_7745c5c3_Var27, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 362, Col: 34}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var27))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"mt-7 mb-2 d-flex justify-content-between align-items-center\"><span class=\"text-sm fw-semibold\">Add Team Member to your project!</span> <a href=\"javascript:void(0)\" hx-get=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var28 string
					templ_7745c5c3_Var28, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/data-synchronizer/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 368, Col: 78}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var28))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-push-url=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var29 string
					templ_7745c5c3_Var29, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/extension/data-synchronizer/%s", projectId))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/extension/extension.templ`, Line: 369, Col: 83}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var29))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"btn btn-sm btn-square btn-white stretched-link\" hx-target=\"#dashboard-content\"><i class=\"bi bi-arrow-right\"></i></a></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"list-unstyled mt-7\"><li class=\"py-2 d-flex align-items-center\"><div class=\"icon icon-xs text-base icon-shape rounded-circle bg-primary-subtle text-primary me-3\"><i class=\"bi bi-check\"></i></div><p>Add as many users you want to your project</p></li></ul></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></main></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
