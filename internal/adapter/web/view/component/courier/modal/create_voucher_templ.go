// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/stelgkio/otoo/internal/core/domain"

func CreateVoucher(extensions []*domain.ProjectExtension) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"offcanvas offcanvas-end w-100 w-md-75 w-lg-50\" data-bs-backdrop=\"true\" tabindex=\"-1\" id=\"orderProcessingOffcanvas\" x-data=\"initOrderForm()\" :class=\"{ &#39;show&#39;: showOffcanvas }\" :style=\"{ visibility: showOffcanvas ? &#39;visible&#39; : &#39;hidden&#39; }\" @show-offcanvas.window=\"openOffcanvas($event.detail.voucher)\"><div class=\"offcanvas-header\"><h5 class=\"offcanvas-title\" id=\"orderProcessingOffcanvasLabel\" data-i18n=\"offcanvas-order_processing\"><i class=\"bi bi-pencil-square\"></i>Order Processing #<span x-text=\"modalOrder.orderId\"></span></h5><button type=\"button\" class=\"btn-close\" @click=\"closeOffcanvas()\"></button></div><div class=\"offcanvas-body d-flex flex-column\"><!-- Tabs --><ul class=\"nav nav-tabs\" id=\"orderTabs\" role=\"tablist\"><li class=\"nav-item\" role=\"presentation\"><button class=\"nav-link active\" id=\"customer-tab\" data-bs-toggle=\"tab\" data-bs-target=\"#customer-info\" type=\"button\" role=\"tab\" @click=\"setActiveTab(&#39;customer&#39;)\"><i class=\"bi bi-person\"></i> Customer</button></li><li class=\"nav-item\" role=\"presentation\"><button class=\"nav-link\" id=\"shipping-tab\" data-bs-toggle=\"tab\" data-bs-target=\"#shipping-info\" type=\"button\" role=\"tab\" @click=\"setActiveTab(&#39;shipping&#39;)\"><i class=\"bi bi-truck\"></i> Shipping</button></li></ul><!-- Tab Content --><div class=\"tab-content flex-grow-1 overflow-auto\" id=\"orderTabContent\"><!-- Customer Information Tab --><div class=\"tab-pane fade show active\" id=\"customer-info\" role=\"tabpanel\"><div class=\"p-3\"><div class=\"card mb-3 shadow-sm\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-customer-info-header\"><i class=\"bi bi-person\"></i> Customer Info</h6></div><div class=\"card-body\"><div class=\"row g-2\"><!-- First Name --><div class=\"col-md-6\"><label for=\"customerName\" class=\"form-label small\">First Name</label> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.first_name&#39;] }\" id=\"customerName\" x-model=\"modalOrder.billing.first_name\" @blur=\"validateField(&#39;billing.first_name&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.first_name&#39;]\"></div></div><!-- Last Name --><div class=\"col-md-6\"><label for=\"customerSurname\" class=\"form-label small\">Last Name</label> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.last_name&#39;] }\" id=\"customerSurname\" x-model=\"modalOrder.billing.last_name\" @blur=\"validateField(&#39;billing.last_name&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.last_name&#39;]\"></div></div><!-- Email --><div class=\"col-md-6\"><label for=\"customerEmail\" class=\"form-label small\">Email</label> <input type=\"email\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.email&#39;] }\" id=\"customerEmail\" x-model=\"modalOrder.billing.email\" @blur=\"validateField(&#39;billing.email&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.email&#39;]\"></div></div><!-- Phone --><div class=\"col-md-6\"><label for=\"customerPhone\" class=\"form-label small\">Phone</label> <input type=\"tel\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.phone&#39;] }\" id=\"customerPhone\" x-model=\"modalOrder.billing.phone\" @blur=\"validateField(&#39;billing.phone&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.phone&#39;]\"></div></div><!-- Address --><div class=\"col-12\"><label for=\"customerAddress\" class=\"form-label small\">Address</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-geo-alt\"></i></span> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.address_1&#39;] }\" id=\"customerAddress\" x-model=\"modalOrder.billing.address_1\" @blur=\"validateField(&#39;billing.address_1&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.address_1&#39;]\"></div></div></div><!-- City --><div class=\"col-md-6\"><label for=\"customerCity\" class=\"form-label small\">City</label> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.city&#39;] }\" id=\"customerCity\" x-model=\"modalOrder.billing.city\" @blur=\"validateField(&#39;billing.city&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.city&#39;]\"></div></div><!-- Postal Code --><div class=\"col-md-6\"><label for=\"customerPostalCode\" class=\"form-label small\">Postal Code</label> <input type=\"text\" class=\"form-control form-control-sm\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;billing.postcode&#39;] }\" id=\"customerPostalCode\" x-model=\"modalOrder.billing.postcode\" @blur=\"validateField(&#39;billing.postcode&#39;)\"><div class=\"invalid-feedback\" x-text=\"errors[&#39;billing.postcode&#39;]\"></div></div></div></div></div><!-- Section for Products --><div class=\"card mb-3 shadow-sm\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-customer-products-header\"><i class=\"bi bi-box-seam\"></i> Products</h6></div><div class=\"card-body\"><div class=\"table-responsive\" style=\"max-height: 200px; overflow-y: auto;\"><table class=\"table table-sm\"><thead><tr><th data-i18n=\"off-canvas-customer-product-name\">Product</th><th data-i18n=\"off-canvas-customer-product-quantity\">Quantity</th><th data-i18n=\"off-canvas-customer-product-price\">Price</th></tr></thead> <tbody><template x-for=\"product in modalOrder.products\" :key=\"product.id\"><tr><td x-text=\"product.name\"></td><td x-text=\"product.quantity\"></td><td x-text=\"product.price\"></td></tr></template></tbody></table></div></div></div><!-- Section for Delivery Instructions --><div class=\"card mb-3 shadow-sm\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-delivery-info\"><i class=\"bi bi-chat-right-text\"></i> Delivery Info</h6></div><div class=\"card-body\"><textarea class=\"form-control form-control-sm\" id=\"deliveryInstructions\" rows=\"2\" data-i18n=\"[placeholder]off-canvas-special-instructions\"></textarea></div></div></div></div><!-- Shipping Information Tab --><div class=\"tab-pane fade\" id=\"shipping-info\" role=\"tabpanel\" aria-labelledby=\"shipping-tab\"><div class=\"p-3\"><!-- Section for Shipping Information --><div class=\"card mb-3 shadow-sm\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-shipping-information\"><i class=\"bi bi-truck\"></i> Shipping Information</h6></div><div class=\"card-body\"><div class=\"mb-3\"><label for=\"shippingCompany\" class=\"form-label small\" data-i18n=\"off-canvas-shipping-company\">Shipping Company</label> <select class=\"form-select form-select-sm\" id=\"shippingCompany\" x-model=\"selectedCourier\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;shipping.courier&#39;] }\" @change=\"validateField(&#39;shipping.courier&#39;)\" required><option value=\"\" data-i18n=\"off-canvas-shipping-provider\">Select Courier Provider</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, extension := range extensions {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Code)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/courier/modal/create_voucher.templ`, Line: 230, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(extension.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/courier/modal/create_voucher.templ`, Line: 230, Col: 61}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select><div class=\"invalid-feedback\" x-text=\"errors[&#39;shipping.courier&#39;]\"></div></div><div class=\"mb-3\"><label for=\"orderId\" class=\"form-label small\" data-i18n=\"off-canvas-shipping-order-number\">Order Number</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-upc-scan\"></i></span> <input type=\"text\" class=\"form-control form-control-sm\" id=\"orderId\" x-model=\"modalOrder.orderId\" readonly></div></div><div class=\"mb-3\"><label for=\"parcelWeight\" class=\"form-label small\" data-i18n=\"off-canvas-shipping-parcel-weight\">Parcel Weight (kg)</label><div class=\"input-group input-group-sm\"><span class=\"input-group-text\"><i class=\"bi bi-body-text\"></i></span> <input type=\"number\" class=\"form-control form-control-sm\" id=\"parcelWeight\" placeholder=\"0.5\" step=\"0.5\" x-model=\"modalOrder.shipping.weight\" :class=\"{ &#39;is-invalid&#39;: errors[&#39;shipping.weight&#39;] }\" @blur=\"validateField(&#39;shipping.weight&#39;)\" required><div class=\"invalid-feedback\" x-text=\"errors[&#39;shipping.weight&#39;]\"></div></div></div></div></div><!-- Section for ACS Courier Options (Conditional) --><div class=\"card mb-3 shadow-sm\" x-show=\"selectedCourier === &#39;asc-courier&#39;\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-delivery-options-acs\"><i class=\"bi bi-box-arrow-in-down\"></i> Delivery Options (ACS)</h6></div><div class=\"card-body\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"standardDeliveryACS\" value=\"standard\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"standardDelivery\" data-i18n=\"off-canvas-acs-standard-delivery\">Standard Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"saturdayDelivery\" value=\"saturday\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"saturdayDelivery\" data-i18n=\"off-canvas-acs-saturday-delivery\">Saturday Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"urgentDelivery\" value=\"urgent\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"urgentDelivery\" data-i18n=\"off-canvas-acs-urgent-delivery\">Urgent Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"pickupDelivery\" value=\"pickup\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"pickupDelivery\" data-i18n=\"off-canvas-acs-pickup-delivery\">Pickup</label></div></div></div><!-- Section for Courier4U Options (Conditional) --><div class=\"card mb-3 shadow-sm\" x-show=\"selectedCourier === &#39;courier4u&#39;\"><div class=\"card-header\"><h6 class=\"mb-0\" data-i18n=\"off-canvas-delivery-options-courier4u\"><i class=\"bi bi-box-arrow-in-down\"></i> Delivery Options (Courier4U)</h6></div><div class=\"card-body\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"standardDeliveryCourier4U\" value=\"standard\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"standardDelivery\" data-i18n=\"off-canvas-delivery-courier4u-standard-delivery\">Standard Delivery</label></div><!-- Other delivery options... --><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"saturdayDelivery4u\" value=\"saturday\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"saturdayDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-saturday-delivery\">Saturday Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"urgentDelivery4u\" value=\"urgent\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"urgentDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-urgent-delivery\">Urgent Delivery</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"pickupDelivery4u\" value=\"pickup\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"pickupDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-pickup-delivery\">Pickup</label></div><div class=\"form-check\"><input class=\"form-check-input\" type=\"radio\" name=\"deliveryOption\" id=\"sameDayDelivery4u\" value=\"sameday\" x-model=\"modalOrder.shipping.deliveryOption\"> <label class=\"form-check-label\" for=\"sameDayDelivery4u\" data-i18n=\"off-canvas-delivery-courier4u-sameday-delivery\">Same Day</label></div></div></div></div></div></div><!-- Save Button --><div class=\"mt-auto p-3 border-top\"><button type=\"button\" class=\"btn btn-primary w-100\" id=\"saveBtn\" @click=\"handleSubmit()\" :disabled=\"!isFormValid() || isSubmitting\" data-i18n=\"off-canvas-button-save\"><template x-if=\"!isSubmitting\"><span>Save Changes</span></template><template x-if=\"isSubmitting\"><span class=\"d-flex align-items-center justify-content-center\"><span class=\"spinner-border spinner-border-sm me-2\" role=\"status\"></span> Processing...</span></template></button></div></div><!-- Toast for notifications --><div class=\"toast-container position-fixed bottom-0 end-0 p-3\"><div id=\"validationToast\" class=\"toast align-items-center text-white border-0\" :class=\"toastType\" role=\"alert\" aria-live=\"assertive\" aria-atomic=\"true\"><div class=\"d-flex\"><div class=\"toast-body\" x-text=\"toastMessage\"></div><button type=\"button\" class=\"btn-close btn-close-white me-2 m-auto\" data-bs-dismiss=\"toast\"></button></div></div></div></div><script>\n// Define the initialization function\nfunction initOrderForm() {\n    return {\n        showOffcanvas: false,\n        activeTab: 'customer',\n        errors: {},\n        isSubmitting: false,\n        toastMessage: '',\n        toastType: 'bg-success',\n        selectedCourier: '',\n        modalOrder: {\n            orderId: '',\n            billing: {\n                first_name: '',\n                last_name: '',\n                email: '',\n                phone: '',\n                address_1: '',\n                city: '',\n                postcode: ''\n            },\n            shipping: {\n                first_name: '',\n                last_name: '',\n                address_1: '',\n                city: '',\n                postcode: '',\n                weight: '',\n                courier: '',\n                deliveryOption: ''\n            },\n            products: []\n        },\n\n        // Initialize the form\n        init() {\n            console.log('Initializing form');\n            this.setupValidationWatchers();\n            this.initializeBootstrapComponents();\n        },\n\n        // Handle opening the offcanvas\n        openOffcanvas(voucher) {\n            console.log('Opening offcanvas with voucher:', voucher);\n            \n            // Check if products exists and is an array\n            if (!Array.isArray(voucher.products)) {\n                voucher.products = [];\n            }\n\n            // Update modalOrder with voucher data\n            this.modalOrder = {\n                ...this.modalOrder,\n                orderId: voucher.orderId,\n                billing: { ...this.modalOrder.billing, ...voucher.billing },\n                shipping: { ...this.modalOrder.shipping, ...voucher.shipping },\n                products: [...voucher.products]\n            };\n\n            this.selectedCourier = voucher.shipping?.courier || '';\n            this.showOffcanvas = true;\n            this.activeTab = 'customer';\n            this.errors = {};\n\n            // Initialize Bootstrap offcanvas if needed\n            if (!this.offcanvas) {\n                this.initializeBootstrapComponents();\n            }\n\n            // Show the offcanvas\n            if (this.offcanvas) {\n                this.offcanvas.show();\n            }\n        },\n\n        // Initialize Bootstrap components\n        initializeBootstrapComponents() {\n            try {\n                const offcanvasElement = document.getElementById('orderProcessingOffcanvas');\n                if (offcanvasElement && typeof bootstrap !== 'undefined') {\n                    // Only initialize if not already initialized\n                    if (!this.offcanvas) {\n                        this.offcanvas = new bootstrap.Offcanvas(offcanvasElement, {\n                            backdrop: true,\n                            keyboard: true\n                        });\n\n                        // Add event listeners\n                        offcanvasElement.addEventListener('hidden.bs.offcanvas', () => {\n                            this.showOffcanvas = false;\n                            this.resetForm();\n                        });\n                    }\n                }\n            } catch (error) {\n                console.error('Error initializing Bootstrap components:', error);\n            }\n        },\n\n        // Set active tab\n        setActiveTab(tab) {\n            console.log('Setting active tab:', tab);\n            this.activeTab = tab;\n        },\n\n        // Setup validation watchers\n        setupValidationWatchers() {\n            try {\n                // Watch billing fields\n                ['first_name', 'last_name', 'email', 'phone', 'address_1', 'city', 'postcode'].forEach(field => {\n                    this.$watch(`modalOrder.billing.${field}`, (value) => {\n                        this.validateField(`billing.${field}`);\n                    });\n                });\n\n                // Watch shipping fields\n                ['first_name', 'last_name', 'address_1', 'city', 'postcode'].forEach(field => {\n                    this.$watch(`modalOrder.shipping.${field}`, (value) => {\n                        this.validateField(`shipping.${field}`);\n                    });\n                });\n\n                // Watch shipping-specific fields\n                this.$watch('selectedCourier', (value) => {\n                    this.modalOrder.shipping.courier = value;\n                    this.validateField('shipping.courier');\n                });\n\n                // this.$watch('modalOrder.shipping.weight', (value) => {\n                //     this.validateField('shipping.weight');\n                // });\n\n                this.$watch('modalOrder.shipping.deliveryOption', (value) => {\n                    this.validateField('shipping.deliveryOption');\n                });\n            } catch (error) {\n                console.error('Error setting up validation watchers:', error);\n            }\n        },\n\n        // Validation methods\n        validateField(field) {\n            console.log('Validating field:', field);\n            const value = field.includes('.') \n                ? field.split('.').reduce((obj, key) => obj?.[key], this.modalOrder)\n                : this[field] || '';\n            \n            delete this.errors[field];\n\n            const optionalFields = ['shipping.deliveryOption'];\n            if (!value && optionalFields.includes(field)) {\n                return true;\n            }\n\n            if (!value?.toString().trim()) {\n                this.errors[field] = 'This field is required';\n                return false;\n            }\n\n            switch(true) {\n                case field.endsWith('email'):\n                    const emailRegex = /^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$/;\n                    if (!emailRegex.test(value)) {\n                        this.errors[field] = 'Please enter a valid email address';\n                        return false;\n                    }\n                    break;\n\n                case field.endsWith('phone'):\n                    const phoneRegex = /^\\+?[\\d\\s-]{10,}$/;\n                    if (!phoneRegex.test(value)) {\n                        this.errors[field] = 'Please enter a valid phone number';\n                        return false;\n                    }\n                    break;\n\n                case field.endsWith('postcode'):\n                    const postcodeRegex = /^[A-Z0-9]{2,10}$/i;\n                    if (!postcodeRegex.test(value)) {\n                        this.errors[field] = 'Please enter a valid postal code';\n                        return false;\n                    }\n                    break;\n\n                // case field === 'shipping.weight':\n                //     const weight = parseFloat(value);\n                //     if (isNaN(weight) || weight <= 0) {\n                //         this.errors[field] = 'Please enter a valid weight greater than 0';\n                //         return false;\n                //     }\n                //     break;\n\n                case field === 'shipping.courier':\n                    if (!value) {\n                        this.errors[field] = 'Please select a courier';\n                        return false;\n                    }\n                    break;\n            }\n\n            return true;\n        },\n\n        validateForm() {\n            let isValid = true;\n            const billingFields = ['first_name', 'last_name', 'email', 'phone', 'address_1', 'city', 'postcode'];\n            \n            billingFields.forEach(field => {\n                if (!this.validateField(`billing.${field}`)) {\n                    isValid = false;\n                }\n            });\n\n            if (this.activeTab === 'shipping') {\n                const shippingFields = ['first_name', 'last_name', 'address_1', 'city', 'postcode'];\n                shippingFields.forEach(field => {\n                    if (!this.validateField(`shipping.${field}`)) {\n                        isValid = false;\n                    }\n                });\n\n                // ['courier', 'weight', 'deliveryOption'].forEach(field => {\n                //     if (!this.validateField(`shipping.${field}`)) {\n                //         isValid = false;\n                //     }\n                // });\n            }\n\n            return isValid;\n        },\n\n        // Form submission\n        async handleSubmit() {\n            if (this.isSubmitting) return;\n\n            if (!this.validateForm()) {\n                this.showToast('Please check the form for errors', 'bg-danger');\n                return;\n            }\n\n            this.isSubmitting = true;\n\n            try {\n                // Your API call would go here\n                await new Promise(resolve => setTimeout(resolve, 1000)); // Simulate API call\n                this.showToast('Voucher saved successfully', 'bg-success');\n                this.closeOffcanvas();\n            } catch (error) {\n                console.error('Error saving voucher:', error);\n                this.showToast('Failed to save voucher', 'bg-danger');\n            } finally {\n                this.isSubmitting = false;\n            }\n        },\n\n        // Close the offcanvas\n        closeOffcanvas() {\n            if (this.offcanvas) {\n                this.offcanvas.hide();\n            }\n            this.showOffcanvas = false;\n            this.errors = {};\n            this.resetForm();\n            \n            // Dispatch an event to notify parent component\n            this.$dispatch('offcanvas-closed');\n        },\n\n\t\t//  closeOffcanvas() {\n        //     this.showOffcanvas = false;\n        //     this.modalOrder = {\n        //         billing: { ...this.modalOrder.billing },\n        //         shipping: { ...this.modalOrder.shipping },\n        //         products: [],\n        //         payment_method: ''\n        //     };\n\t\t// \t console.log('Closing offcanvas:', this.showOffcanvas);\n        // },\n\n        // Reset the form\n        resetForm() {\n            this.modalOrder = {\n                orderId: '',\n                billing: {\n                    first_name: '',\n                    last_name: '',\n                    email: '',\n                    phone: '',\n                    address_1: '',\n                    city: '',\n                    postcode: ''\n                },\n                shipping: {\n                    first_name: '',\n                    last_name: '',\n                    address_1: '',\n                    city: '',\n                    postcode: '',\n                    weight: '',\n                    courier: '',\n                    deliveryOption: ''\n                },\n                products: []\n            };\n            this.selectedCourier = '';\n            this.activeTab = 'customer';\n        },\n\n        // Show toast notification\n        showToast(message, type = 'bg-success') {\n            this.toastMessage = message;\n            this.toastType = type;\n            const toast = new bootstrap.Toast(document.getElementById('validationToast'));\n            toast.show();\n        },\n\n        // Computed property for form validity\n        isFormValid() {\n            return Object.keys(this.errors).length === 0;\n        }\n    }\n}\n\n// Initialize Alpine.js component\n// window.addEventListener('load', () => {\n//     if (window.Alpine) {\n//         window.Alpine.data('initOrderForm', initOrderForm);\n//     }\n// });\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
