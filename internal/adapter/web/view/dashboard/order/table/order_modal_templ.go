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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"modal fade\" tabindex=\"-1\" :class=\"{ &#39;show d-block&#39;: showModal }\" aria-labelledby=\"orderDetailsLabel\" aria-hidden=\"true\" style=\"background-color: rgba(0, 0, 0, 0.5); transition: all 0.3s ease-in-out;\"><div class=\"modal-dialog modal-lg\"><div class=\"modal-content shadow-lg border-0 rounded-3\"><div class=\"modal-header border-0 pb-0\"><h5 class=\"modal-title text-primary\" id=\"orderDetailsLabel\"><i class=\"bi bi-receipt\"></i> Order Details</h5><button type=\"button\" class=\"btn-close\" aria-label=\"Close\" @click=\"closeModal\" style=\"transition: all 0.3s ease;\"></button></div><div class=\"modal-body pt-0 pb-3\"><div class=\"row\"><!-- Billing Column --><div class=\"col-md-6 border-end\"><h5 class=\"text-secondary\">Billing Information</h5><hr class=\"my-2\"><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\" x-for=\"$index in Object.keys(modalOrder.billing).length\"><strong>First Name:</strong> <span x-text=\"modalOrder.billing.first_name\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>Last Name:</strong> <span x-text=\"modalOrder.billing.last_name\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>Address:</strong> <span x-text=\"modalOrder.billing.address_1\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>City:</strong> <span x-text=\"modalOrder.billing.city\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>Postcode:</strong> <span x-text=\"modalOrder.billing.postcode\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>Email:</strong> <span x-text=\"modalOrder.billing.email\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>Phone:</strong> <span x-text=\"modalOrder.billing.phone\"></span></div></div><!-- Shipping Column --><div class=\"col-md-6\"><h5 class=\"text-secondary\">Shipping Information</h5><hr class=\"my-2\"><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>First Name:</strong> <span x-text=\"modalOrder.shipping.first_name\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>Last Name:</strong> <span x-text=\"modalOrder.shipping.last_name\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>Address:</strong> <span x-text=\"modalOrder.shipping.address_1\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>City:</strong> <span x-text=\"modalOrder.shipping.city\"></span></div><div class=\"py-2\" :class=\"$index % 2 === 0 ? &#39;bg-light&#39; : &#39;&#39;\"><strong>Postcode:</strong> <span x-text=\"modalOrder.shipping.postcode\"></span></div></div></div><!-- Order Details Section --><div class=\"mt-4\"><h5 class=\"text-secondary\">Order Details</h5><hr class=\"my-2\"><div class=\"d-flex mb-2\"><strong>Order ID:</strong> <span class=\"ms-auto\" x-text=\"modalOrder.orderId\"></span></div><div class=\"d-flex mb-2\"><strong>Total Amount:</strong> <span class=\"ms-auto\" x-text=\"modalOrder.total_amount + &#39; &#39; + modalOrder.currency_symbol\"></span></div><div class=\"d-flex mb-2\"><strong>Status:</strong> <span class=\"ms-auto\" x-text=\"modalOrder.status\"></span></div><div class=\"d-flex mb-2\"><strong>Payment Method:</strong> <span class=\"ms-auto\" x-text=\"modalOrder.payment_method\"></span></div><!-- Products List --><div class=\"mt-3 bg-light p-3 rounded\"><h6>Products</h6><template x-for=\"product in modalOrder.products\" :key=\"product.id\"><div class=\"d-flex justify-content-between\"><div><strong>Product:</strong> <span x-text=\"product.name\"></span></div><div><strong>Qty:</strong> <span x-text=\"product.quantity\"></span></div><div><strong>Price:</strong> <span x-text=\"product.price + &#39; &#39; + modalOrder.currency_symbol\"></span></div></div><hr class=\"my-2\"></template></div></div></div><div class=\"modal-footer border-0\"><button type=\"button\" class=\"btn btn-secondary\" @click=\"closeModal\">Close</button></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
