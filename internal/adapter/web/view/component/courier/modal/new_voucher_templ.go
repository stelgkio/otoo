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

func NewVoucher(extensions []*domain.ProjectExtension, projectId string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n\t\t.offcanvas-backdrop {\n\t\tbackdrop-filter: blur(2px); \n\t\tbackground-color: rgba(0, 0, 0, 0.3); \n\t\t}\n\t</style><!-- Start off canvas --><div class=\"offcanvas offcanvas-end w-100 w-md-75 w-lg-50\" data-bs-backdrop=\"true\" tabindex=\"-1\" id=\"newVoucherOffcanvas\" aria-labelledby=\"newVoucherOffcanvasLabel\" x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("newVoucher('%s')", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/courier/modal/new_voucher.templ`, Line: 22, Col: 53}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"newVoucherInit()\" :class=\"{ &#39;show&#39;: showOffcanvas }\" :style=\"{ visibility: showOffcanvas ? &#39;visible&#39; : &#39;hidden&#39; }\" @new-voucher.window=\"openNewVoucherOffcanva()\"><div class=\"offcanvas-header\"><h5 class=\"offcanvas-title\" id=\"newVoucherOffcanvasLabel\" data-i18n=\"offcanvas-order_processing\"><i class=\"bi bi-pencil-square\"></i>New Voucher #<span x-text=\"voucher_object.orderId\"></span></h5><button type=\"button\" class=\"btn-close\" @click=\"closeOffcanvas(&#39;newVoucherInit&#39;)\"></button></div><div class=\"offcanvas-body d-flex flex-column\"><!-- Tabs --><ul class=\"nav nav-tabs\" id=\"orderTabs\" role=\"tablist\"><li class=\"nav-item\" role=\"presentation\"><button class=\"nav-link active\" id=\"customer-tab\" data-bs-toggle=\"tab\" data-bs-target=\"#customer-info3\" type=\"button\" role=\"tab\" @click=\"setActiveTab(&#39;customer&#39;)\"><span data-i18n=\"offcanvas-nav-customer\">Customer</span></button></li><li class=\"nav-item\" role=\"presentation\"><button class=\"nav-link\" id=\"shipping-tab\" data-bs-toggle=\"tab\" data-bs-target=\"#shipping-info3\" type=\"button\" role=\"tab\" @click=\"setActiveTab(&#39;shipping&#39;)\"><span data-i18n=\"offcanvas-nav-shipping\">Shipping</span></button></li></ul><!-- Tab Content --><div class=\"tab-content flex-grow-1 overflow-auto\" id=\"orderTabContent\"><!-- Customer Information Tab --><div class=\"tab-pane fade show active\" id=\"customer-info3\" role=\"tabpanel\"><div class=\"p-3\"><div class=\"card mb-3 shadow-sm\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-customer-info-header\"><i class=\"bi bi-person\"></i> Customer Info</h6></div><div class=\"card-body\"><div class=\"row g-2\"><!-- First Name --><div class=\"col-md-6\"><label for=\"customerName\" class=\"form-label small\" data-i18n=\"off-canvas-modal-name\">First Name</label> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;shipping.first_name&#39;] }\" id=\"customerName\" x-model=\"voucher_object.shipping.first_name\" @blur=\"validateField(&#39;shipping.first_name&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;shipping.first_name&#39;]\"></div></div><!-- Last Name --><div class=\"col-md-6\"><label for=\"customerSurname\" class=\"form-label small\" data-i18n=\"off-canvas-modal-last-name\">Last Name</label> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;shipping.last_name&#39;] }\" id=\"customerSurname\" x-model=\"voucher_object.shipping.last_name\" @blur=\"validateField(&#39;shipping.last_name&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;shipping.last_name&#39;]\"></div></div><!-- Email --><div class=\"col-md-6\"><label for=\"customerEmail\" class=\"form-label small\" data-i18n=\"off-canvas-email\">Email</label> <input type=\"email\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.email&#39;] }\" id=\"customerEmail\" x-model=\"voucher_object.billing.email\" @blur=\"validateField(&#39;billing.email&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.email&#39;]\"></div></div><!-- Phone --><div class=\"col-md-6\"><label for=\"customerPhone\" class=\"form-label small\" data-i18n=\"off-canvas-modal-phone\">Phone</label> <input type=\"tel\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.phone&#39;] }\" id=\"customerPhone\" x-model=\"voucher_object.billing.phone\" @blur=\"validateField(&#39;billing.phone&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.phone&#39;]\"></div></div><!-- Address --><div class=\"col-12\"><label for=\"customerAddress\" class=\"form-label small\" data-i18n=\"off-canvas-modal-address\">Address</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-geo-alt\"></i></span> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;shipping.address_1&#39;] }\" id=\"customerAddress\" x-model=\"voucher_object.shipping.address_1\" @blur=\"validateField(&#39;shipping.address_1&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;shipping.address_1&#39;]\"></div></div></div><!-- City --><div class=\"col-md-6\"><label for=\"customerCity\" class=\"form-label small\" data-i18n=\"off-canvas-modal-city\">City</label> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;shipping.city&#39;] }\" id=\"customerCity\" x-model=\"voucher_object.shipping.city\" @blur=\"validateField(&#39;shipping.city&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;shipping.city&#39;]\"></div></div><!-- Postal Code --><div class=\"col-md-6\"><label for=\"customerPostalCode\" class=\"form-label small\" data-i18n=\"off-canvas-modal-postal\">Postal Code</label> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;shipping.postcode&#39;] }\" id=\"customerPostalCode\" x-model=\"voucher_object.shipping.postcode\" @blur=\"validateField(&#39;shipping.postcode&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;shipping.postcode&#39;]\"></div></div></div></div></div><!-- Section for Delivery Instructions --><div class=\"card mb-3 shadow-sm\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-delivery-info\"><i class=\"bi bi-chat-right-text\"></i> Delivery Info</h6></div><div class=\"card-body\"><textarea class=\"form-control form-control-sm\" id=\"deliveryInstructions\" rows=\"2\" data-i18n=\"[placeholder]off-canvas-special-instructions\" x-model=\"voucher_object.note\"></textarea></div></div></div></div><!-- Shipping Information Tab --><div class=\"tab-pane fade\" id=\"shipping-info3\" role=\"tabpanel\" aria-labelledby=\"shipping-tab3\"><div class=\"p-3\"><!-- Section for Shipping Information --><div class=\"card mb-3 shadow-sm\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-shipping-information\"><i class=\"bi bi-truck\"></i> Shipping Information</h6></div><!-- Shipping Company --><div class=\"card-body\"><div class=\"mb-3\"><label for=\"shippingCompany\" class=\"form-label small\" data-i18n=\"off-canvas-shipping-company\">Shipping Company</label> <select class=\"form-select form-select-sm\" id=\"shippingCompany\" x-model=\"selectedCourier\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;shipping.courier&#39;] }\" @change=\"validateField(&#39;shipping.courier&#39;)\"><option value=\"\" data-i18n=\"off-canvas-shipping-provider\">Select Courier Provider</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, extension := range extensions {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Code)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/courier/modal/new_voucher.templ`, Line: 214, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/courier/modal/new_voucher.templ`, Line: 214, Col: 61}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select><div class=\"invalid-feedback\" x-text=\"errors[&#39;shipping.courier&#39;]\"></div></div><!-- Order Number --><div class=\"mb-3\"><label for=\"orderId\" class=\"form-label small\" data-i18n=\"off-canvas-shipping-order-number\">Order Number</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-upc-scan\"></i></span> <input type=\"number\" class=\"form-control form-control-sm\" @blur=\"formatOrderIdWithLeadingZeros\" id=\"orderId\" x-model=\"voucher_object.orderId\"></div></div><!-- Order COD --><div class=\"mb-3\"><label for=\"parcelCOD\" class=\"form-label small\" data-i18n=\"off-canvas-parcel-cod\">COD (€)</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-cash\"></i></span> <input type=\"number\" class=\"form-control form-control-sm\" id=\"parcelCOD\" step=\"0.01\" x-model=\"voucher_object.cod\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;cod&#39;] }\" @blur=\"validateField(&#39;cod&#39;)\" required><div class=\"invalid-feedback\" x-text=\"errors[&#39;cod&#39;]\"></div><!-- Error message display --></div></div></div></div><!-- Section for ACS Courier Options (Conditional) --><div class=\"card mb-3 shadow-sm\" x-show=\"selectedCourier === &#39;asc-courier&#39;\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-delivery-options-acs\"><i class=\"bi bi-box-arrow-in-down\"></i> Delivery Options (ACS)</h6></div><div class=\"card-body\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" name=\"deliveryOption\" id=\"AcsStandardDelivery\" value=\"standard\" x-model=\"voucher_object.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"AcsSaturdayDelivery\" data-i18n=\"off-canvas-acs-standard-delivery\">Standard Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" name=\"deliveryOption\" id=\"AcsSaturdayDelivery\" value=\"saturday\" x-model=\"voucher_object.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"AcsSaturdayDelivery\" data-i18n=\"off-canvas-acs-saturday-delivery\">Saturday Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" name=\"deliveryOption\" id=\"AcsUrgentDelivery\" value=\"urgent\" x-model=\"voucher_object.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"AcsUrgentDelivery\" data-i18n=\"off-canvas-acs-urgent-delivery\">Urgent Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" name=\"deliveryOption\" id=\"AcsPickupDelivery\" value=\"pickup\" x-model=\"voucher_object.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"AcsPickupDelivery\" data-i18n=\"off-canvas-acs-pickup-delivery\">Pickup</label></div></div></div><!-- Section for Courier4U Options (Conditional) --><div class=\"card mb-3 shadow-sm\" x-show=\"selectedCourier === &#39;courier4u&#39;\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-wh-options-courier4u\"><i class=\"bi bi-box-arrow-in-down\"></i> W/D/W/H Options (Courier4U)</h6></div><div class=\"card-body\"><div class=\"row\"><!-- Parcel Weight --><div class=\"col-md-6 mb-3\"><label for=\"parcelWeight\" class=\"form-label small\">Parcel Weight (kg)</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-box-seam\"></i></span> <input type=\"number\" class=\"form-control form-control-sm\" id=\"HermesParcelWeight\" step=\"0.5\" x-model=\"voucher_object.hermes_settings.ParcelWeight\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;hermes_settings.ParcelWeight&#39;] }\" @blur=\"validateField(&#39;hermes_settings.ParcelWeight&#39;)\" required><div class=\"invalid-feedback\" x-text=\"errors[&#39;hermes_settings.ParcelWeight&#39;]\"></div></div></div><!-- Parcel Depth --><div class=\"col-md-6 mb-3\"><label for=\"parcelDepth\" class=\"form-label small\" data-i18n=\"off-canvas-shipping-parcel-depth\">Parcel Depth (cm)</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-arrows-expand\"></i></span> <input type=\"number\" class=\"form-control form-control-sm\" id=\"HermesParcelDepth\" step=\"0.5\" x-model=\"voucher_object.hermes_settings.ParcelDepth\"></div></div></div><div class=\"row\"><!-- Parcel Width --><div class=\"col-md-6 mb-3\"><label for=\"parcelWidth\" class=\"form-label small\" data-i18n=\"off-canvas-shipping-parcel-width\">Parcel Width (cm)</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-arrows-angle-expand\"></i></span> <input type=\"number\" class=\"form-control form-control-sm\" id=\"HermesParcelWidth\" step=\"0.5\" x-model=\"voucher_object.hermes_settings.ParcelWidth\"></div></div><!-- Parcel Height --><div class=\"col-md-6 mb-3\"><label for=\"parcelHeight\" class=\"form-label small\" data-i18n=\"off-canvas-shipping-parcel-height\">Parcel Height (cm)</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-arrow-up\"></i></span> <input type=\"number\" class=\"form-control form-control-sm\" id=\"HermesParcelHeight\" step=\"0.5\" x-model=\"voucher_object.hermes_settings.ParcelHeight\"></div></div></div></div></div><div class=\"card mb-3 shadow-sm\" x-show=\"selectedCourier === &#39;courier4u&#39;\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-delivery-options-courier4u\"><i class=\"bi bi-box-arrow-in-down\"></i> Delivery Options (Courier4U)</h6></div><div class=\"card-body\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" id=\"saturdayDelivery4u\" x-model=\"voucher_object.hermes_settings.ServiceSavvato\" value=\"1\"> <label class=\"form-check-label\" for=\"saturdayDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-saturday-delivery\">Saturday Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" id=\"urgentDelivery4u\" x-model=\"voucher_object.hermes_settings.ServiceEpigon\" value=\"1\"> <label class=\"form-check-label\" for=\"urgentDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-urgent-delivery\">Urgent Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" id=\"serviceEpistrofi\" x-model=\"voucher_object.hermes_settings.ServiceEpistrofi\" value=\"1\"> <label class=\"form-check-label\" for=\"serviceEpistrofi\" data-i18n=\"off-canvas-delivery-courier4u-return-delivery\">Return Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" id=\"pickupDelivery4u\" x-model=\"voucher_object.hermes_settings.ServiceReception\" value=\"1\"> <label class=\"form-check-label\" for=\"pickupDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-pickup-delivery\">Pickup from Store</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" id=\"sameDayDelivery4u\" x-model=\"voucher_object.hermes_settings.ServiceSameday\" value=\"1\"> <label class=\"form-check-label\" for=\"sameDayDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-sameday-delivery\">Same Day</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" id=\"ServiceProtocolDayDelivery4u\" x-model=\"voucher_object.hermes_settings.ServiceProtocol\" value=\"1\"> <label class=\"form-check-label\" for=\"ServiceProtocolDayDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-ServiceProtocol-delivery\">Protoco Number</label></div></div></div></div></div></div><!-- Save Button --><div class=\"d-grid gap-2\"><div x-show=\"toastMessage\" class=\"alert alert-danger mt-2\" role=\"alert\"><span x-text=\"toastMessage\"></span></div><div x-show=\"toastMessageSuuccess\" class=\"alert alert-success mt-2\" role=\"alert\"><span x-text=\"toastMessageSuuccess\"></span></div><button type=\"button\" class=\"btn btn-primary w-100\" id=\"saveBtn\" @click=\"handleSubmit()\" :disabled=\" isSubmitting\" data-i18n=\"off-canvas-save-btn\"><template x-show=\"!isSubmitting\"><span>Create Voucher</span></template><template x-show=\"isSubmitting\"><span class=\"d-flex align-items-center justify-content-center\"><span class=\"spinner-border spinner-border-sm me-2\" role=\"status\"></span> Processing...</span></template></button><!-- Close Button --><button type=\"button\" class=\"btn btn-neutral rounded-1 px-4\" id=\"closeBtn\" @click=\"closeOffcanvas(&#39;newVoucherInit&#39;)\" data-i18n=\"off-canvas-close-btn\">Close</button><!-- Save Button --></div></div><!-- Toast for notifications --></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
