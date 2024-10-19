// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func OrderModal() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"modal fade\" tabindex=\"-1\" :class=\"{ &#39;show d-block&#39;: showModal }\" aria-labelledby=\"orderDetailsLabel\" aria-hidden=\"false\" style=\"background-color: rgba(0, 0, 0, 0.7); transition: all 0.3s ease-in-out;\"><div class=\"modal-dialog modal-xl modal-dialog-centered \"><div class=\"modal-content shadow-lg border-0 rounded-2 \"><div class=\"modal-header border-0 pb-0 pt-4 px-4\"><h5 class=\"modal-title text-primary fw-bold\" id=\"orderDetailsLabel\"><i class=\"bi bi-receipt-cutoff me-2\"></i>Order Details</h5><button type=\"button\" class=\"btn-close\" aria-label=\"Close\" @click=\"closeModal\"></button></div><div class=\"modal-body pt-0 pb-4 px-4\"><div class=\"row g-4 mt-2\"><!-- Billing Column --><div class=\"col-md-6 pe-md-4\"><div class=\" rounded-2 h-100 px-6 py-5 border\"><div class=\"d-flex justify-content-between align-items-center mb-3\"><h6 class=\"text-muted fw-bold\"><i class=\"bi bi-credit-card me-2\"></i>Billing Information</h6></div><form><div class=\"row\"><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">First Name</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.billing.first_name\"></div><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">Last Name</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.billing.last_name\"></div></div><div class=\"mb-3\"><label class=\"form-label text-muted\">Address</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.billing.address_1\"></div><div class=\"row\"><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">City</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.billing.city\"></div><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">Postcode</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.billing.postcode\"></div></div><div class=\"mb-3\"><label class=\"form-label text-muted\">Email</label> <input type=\"email\" class=\"form-control form-control-sm\" x-model=\"modalOrder.billing.email\"></div><div class=\"mb-3\"><label class=\"form-label text-muted\">Phone</label> <input type=\"tel\" class=\"form-control form-control-sm\" x-model=\"modalOrder.billing.phone\"></div></form></div></div><!-- Shipping Column --><div class=\"col-md-6 ps-md-4\"><div class=\" p-3 rounded-2 h-100 border\"><div class=\"d-flex justify-content-between align-items-center mb-3\"><h6 class=\"text-muted fw-bold mb-0\"><i class=\"bi bi-truck me-2 text-muted\"></i>Shipping Information</h6></div><form><div class=\"row\"><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">First Name</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.shipping.first_name\"></div><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">Last Name</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.shipping.last_name\"></div></div><div class=\"row\"><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">Address</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.shipping.address_1\"></div><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">Address 2</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.shipping.address_2\"></div></div><div class=\"row\"><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">City</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.shipping.city\"></div><div class=\"col-md-6 mb-3\"><label class=\"form-label text-muted\">Postcode</label> <input type=\"text\" class=\"form-control form-control-sm\" x-model=\"modalOrder.shipping.postcode\"></div></div><div class=\"mb-3\"><label class=\"form-label text-muted\">Customer Notes</label> <textarea class=\"form-control form-control-sm\" rows=\"4\" x-model=\"modalOrder.customer_note\"></textarea></div></form></div></div></div><!-- Order Details Section --><div class=\"mt-4\"><h6 class=\"text-muted fw-bold mb-3\"><i class=\"bi bi-info-circle me-2 text-muted\"></i>Order Details</h6><div class=\"p-3 rounded-2 shadow-sm border\"><div class=\"row g-3\"><div class=\"col-md-4\"><span class=\"d-block text-muted  text-muted\">Order ID</span> <strong x-text=\"modalOrder.orderId\"></strong></div><div class=\"col-md-4\"><span class=\"d-block text-muted  text-muted\">Total Amount</span> <strong x-text=\"modalOrder.total_amount + &#39; &#39; + modalOrder.currency_symbol\"></strong></div><div class=\"col-md-4\"><span class=\"d-block text-muted text-muted\">Status</span> <span :class=\"badgeClass(modalOrder.status)\" x-text=\"modalOrder.status\"></span></div></div><div class=\"mt-3\"><span class=\"d-block text-muted text-muted\">Payment Method</span> <strong x-text=\"modalOrder.payment_method\"></strong></div></div><!-- Products List --><div class=\"mt-4\"><h6 class=\"text-muted fw-bold mb-3\"><i class=\"bi bi-box-seam me-2 text-muted\"></i>Products</h6><div class=\"table-responsive\" style=\"max-height: 200px;\"><table class=\"table table-sm table-hover align-middle\"><thead class=\"body-secondary\"><tr><th class=\"text-muted fw-bold text-muted\">Product</th><th class=\"text-muted fw-bold text-center text-muted\">Qty</th><th class=\"text-muted fw-bold text-end text-muted\">Price</th></tr></thead> <tbody><template x-for=\"product in modalOrder.products\" :key=\"product.id\"><tr><td x-text=\"product.name\"></td><td x-text=\"product.quantity\" class=\"text-center\"></td><td x-text=\"product.price + &#39; &#39; + modalOrder.currency_symbol\" class=\"text-end\"></td></tr></template></tbody></table></div></div></div></div><div class=\"modal-footer border-0 pt-0\"><button type=\"button\" class=\"btn btn-sm btn-neutral rounded-1 px-4\" @click=\"closeModal\">Close</button><!-- Save Changes Button with Spinner --><button type=\"button\" class=\"btn btn-sm btn-primary rounded-1 px-4\" @click=\"saveChanges\"><!-- Show this when not loading --><span x-show=\"!loading\">Save Changes</span><!-- Spinner (Bootstrap spinner or custom) --><span x-show=\"loading\" class=\"spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span><!-- Optionally, show 'Saving...' text when loading --><span x-show=\"loading\" class=\"ms-2\">Saving...</span></button></div><div x-show=\"errorMessage\" class=\"alert alert-danger mt-2\" role=\"alert\"><span x-text=\"errorMessage\"></span></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
