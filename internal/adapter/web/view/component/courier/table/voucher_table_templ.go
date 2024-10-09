// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/component/pagination"
)

func VoucherTable(projectId string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-voucher\"><div id=\"dashboard-voucher-table\" x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("voucherTable('%s')", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/component/courier/table/voucher_table.templ`, Line: 10, Col: 89}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"init()\"><div class=\"px-6 px-lg-7 pt-1 border-bottom\"><ul class=\"nav nav-tabs nav-tabs-flush gap-8 overflow-x border-0 mt-4\"><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;new&#39;}\" @click.prevent=\"selectTab(&#39;new&#39;)\">New</a></li><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;completed&#39;}\" @click.prevent=\"selectTab(&#39;completed&#39;)\">Completed</a></li><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;processing&#39;}\" @click.prevent=\"selectTab(&#39;processing&#39;)\">Processing</a></li><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;cancelled&#39;}\" @click.prevent=\"selectTab(&#39;cancelled&#39;)\">Canceled</a></li><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;all&#39;}\" @click.prevent=\"selectTab(&#39;all&#39;)\">All</a></li></ul></div><div class=\"table-responsive\"><table class=\"table table-hover table-nowrap\"><thead><tr><th><div class=\"text-base\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" @change=\"selectAll()\" x-model=\"selectAllCheckbox\"></div></div></th><template x-if=\"!isNewTab\"><th @click=\"sortTable(&#39;Id&#39;)\">Voucher ID <i :class=\"getSortIcon(&#39;Id&#39;)\"></i></th></template><th @click=\"sortTable(&#39;orderId&#39;)\">Order ID <i :class=\"getSortIcon(&#39;orderId&#39;)\"></i></th><th @click=\"sortTable(&#39;created_at&#39;)\">Date <i :class=\"getSortIcon(&#39;created_at&#39;)\"></i></th><th @click=\"sortTable(&#39;cod&#39;)\">COD <i :class=\"getSortIcon(&#39;cod&#39;)\"></i></th><th>Status</th><template x-if=\"isPrinted\"><th @click=\"sortTable(&#39;is_printed&#39;)\">Printed <i :class=\"getSortIcon(&#39;is_printed&#39;)\"></i></th></template><th>Action</th></tr></thead> <tbody><template x-if=\"!loading &amp;&amp; totalItems === 0\"><tr><td colspan=\"8\">No vouchers found.</td></tr></template><template x-for=\"voucher in paginatedVouchers\" :key=\"voucher.Id\"><tr @click=\"openOffcanvas(voucher)\"><td><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" :value=\"voucher.Id\" x-model=\"selectedVouchers\" @click.stop></div></td><template x-if=\"!isNewTab\"><td x-text=\"voucher.Id\"></td></template><td x-text=\"voucher.orderId\"></td><td x-text=\"new Date(voucher.created_at).toLocaleString()\"></td><td x-text=\"voucher.cod\"></td><td><span :class=\"badgeClass(voucher.status)\" x-text=\"voucher.status\"></span></td><template x-if=\"isPrinted\"><td><template x-if=\"voucher.is_printed\"><span x-text=\"&#39;Yes&#39;\"></span></template><template x-if=\"!voucher.is_printed\"><span x-text=\"&#39;No&#39;\"></span></template></td></template><td><button type=\"button\" class=\"btn btn-sm btn-neutral\" @click=\"openOffcanvas(voucher)\"><i class=\"bi bi-plus-circle me-2\"></i>Create Voucher</button> <button type=\"button\" class=\"btn btn-sm btn-neutral\" @click.stop=\"downloadVoucher(voucher.Id)\"><i class=\"fas fa-download ml-1\" style=\"cursor: pointer;\" title=\"Download Voucher\"></i></button></td></tr></template><template x-for=\"i in 10 - paginatedVouchers.length\" :key=\"&#39;empty&#39; + i\"><tr><td colspan=\"8\" class=\"py-8\"></td></tr></template></tbody></table></div><!-- Offcanvas Section --><div class=\"offcanvas offcanvas-end\" tabindex=\"-1\" id=\"offcanvasTop\" aria-labelledby=\"offcanvasTopLabel\" :class=\"{ &#39;show&#39;: showOffcanvas }\" :style=\"{ visibility: showOffcanvas ? &#39;visible&#39; : &#39;hidden&#39; }\"><div class=\"offcanvas-header\"><h5 class=\"offcanvas-title\" id=\"offcanvasOrderDetails\">Voucher Details</h5><button type=\"button\" class=\"btn-close\" @click=\"closeOffcanvas()\"></button></div><div class=\"offcanvas-body\"><!-- Display order details inside the offcanvas --><p><strong>Shipping Information:</strong></p><p>Name: <span x-text=\"modalOrder.shipping.first_name + &#39; &#39; + modalOrder.shipping.last_name\"></span></p><p>Address: <span x-text=\"modalOrder.shipping.address_1\"></span></p><p>City: <span x-text=\"modalOrder.shipping.city\"></span></p><p>Postcode: <span x-text=\"modalOrder.shipping.postcode\"></span></p><!-- Display billing information --><br><p><strong>Billing Information:</strong></p><p>Name: <span x-text=\"modalOrder.billing.first_name + &#39; &#39; + modalOrder.billing.last_name\"></span></p><p>Address: <span x-text=\"modalOrder.billing.address_1\"></span></p><p>City: <span x-text=\"modalOrder.billing.city\"></span></p><p>State: <span x-text=\"modalOrder.billing.state\"></span></p><p>Postcode: <span x-text=\"modalOrder.billing.postcode\"></span></p><p>Country: <span x-text=\"modalOrder.billing.country\"></span></p><p>Email: <span x-text=\"modalOrder.billing.email\"></span></p><p>Phone: <span x-text=\"modalOrder.billing.phone\"></span></p><br><p><strong>Products:</strong></p><ul><template x-for=\"product in modalOrder.products\" :key=\"product.id\"><li x-text=\"product.name + &#39; - Quantity: &#39; + product.quantity\"></li></template></ul><p><strong>Payment Method:</strong> <span x-text=\"modalOrder.payment_method\"></span></p><br><p><strong>Payment Amount:</strong> <span x-text=\"modalOrder.cod\"></span></p></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = p.PaginationControl().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><script>\nfunction voucherTable(projectId) {\n    return {\n        projectID: projectId,\n        currentTab: 'new',\n        vouchers: [],\n        selectedVouchers: [],\n        selectAllCheckbox: false,\n        sortKey: 'created_at',\n        sortAsc: false,\n        currentPage: 1,\n        itemsPerPage: 10,\n        totalItems: 0,\n        totalPages: 0,\n        loading: false,\n        selectedStatus: '',\n        errorMessage: '',\n\t\tisNewTab: true, \n\t\tisPrinted: false,\n        showOffcanvas: false,\n\t\t modalOrder: {\n            billing: {\n                first_name: '',\n                last_name: '',\n                address_1: '',\n                city: '',\n                postcode: '',\n                email: '',\n                phone: '',\n            },\n            shipping: {\n                first_name: '',\n                last_name: '',\n                address_1: '',\n                city: '',\n                postcode: '',\n            },\n            products: [],\n            payment_method: ''\n        },\n\n        async init() {\n            await this.fetchVouchers(this.currentPage);\n        },\n\n        async fetchVouchers(page = 1) {\n            this.loading = true;\n            try {\n                const url = this.getUrlForTab(this.currentTab, page);\n                const response = await fetch(url);\n                const result = await response.json();\n                if (response.ok) {\n                    this.vouchers = result.data || [];\n                    this.totalItems = result.meta.totalItems || 0;\n                    this.currentPage = result.meta.currentPage || 1;\n                    this.itemsPerPage = result.meta.itemsPerPage || 10;\n                    this.totalPages = result.meta.totalPages || 0;\n                } else {\n                    console.error('Error fetching data:', result.message);\n                }\n            } catch (error) {\n                console.error('Error fetching data:', error);\n            } finally {\n                this.loading = false;\n            }\n        },\n\n        getUrlForTab(tab, page) {\n            const baseUrl = `${window.location.origin}/voucher/table/${this.projectID}`;\n            const sortDirection = this.sortAsc ? 'asc' : 'desc';\n            switch (tab) {\n                case 'all':\n                    return `${baseUrl}/all/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'pending':\n                    return `${baseUrl}/pending/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'completed':\n                    return `${baseUrl}/completed/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'processing':\n                    return `${baseUrl}/processing/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'cancelled':\n                    return `${baseUrl}/cancelled/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n\t\t\t\tcase 'new':\n                    return `${baseUrl}/new/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                default:\n                    return `${baseUrl}/all/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n            }\n        },\n\n        selectTab(tab) {\n\t\t\tif (tab != 'new') {\n\t\t\t\tthis.isNewTab = false;\n\t\t\t\tthis.isPrinted = true;\n\t\t\t}else {\n\t\t\t\tthis.isNewTab = true;\n\t\t\t\tthis.isPrinted = false;\n\t\t\t}\n            this.currentTab = tab;\n            this.currentPage = 1; \n            this.fetchVouchers(this.currentPage);\n        },\n\n        selectAll() {\n            this.selectedVouchers = this.selectAllCheckbox ? this.vouchers.map(voucher => voucher.Id) : [];\n        },\n\n        sortTable(key) {\n            if (this.sortKey === key) {\n                this.sortAsc = !this.sortAsc;\n            } else {\n                this.sortKey = key;\n                this.sortAsc = true;\n            }\n            this.fetchVouchers(this.currentPage);\n        },\n        getSortIcon(key) {\n            if (this.sortKey !== key) return '';\n            return this.sortAsc ? 'bi bi-chevron-up' : 'bi bi-chevron-down';\n        },\n\n        changePage(page) {\n            if (page < 1 || page > this.totalPages) return;\n            this.fetchVouchers(page);\n        },\n\t\t downloadVoucher(voucherId) {\n        \n       \t console.log(`Downloading voucher with ID: ${voucherId}`);\n        \n    \t},\n\n        get paginatedVouchers() {\n            return this.vouchers;\n        },\n\t\tget currentPageStart() {\n\t\t\treturn (this.currentPage - 1) * this.itemsPerPage + 1;\n\t\t},\n\n\t\tget currentPageEnd() {\n\t\t\treturn Math.min(this.currentPage * this.itemsPerPage, this.totalItems);\n\t\t},\n\t\tget pageNumbers() {\n\t\t\tconst range = 2; // Number of pages to show around the current page\n\t\t\tlet start = Math.max(1, this.currentPage - range);\n\t\t\tlet end = Math.min(this.totalPages, this.currentPage + range);\n\n\t\t\t// Adjust range if there are not enough pages on one side\n\t\t\tif (this.totalPages - end < range) {\n\t\t\t\tend = this.totalPages;\n\t\t\t\tstart = Math.max(1, end - 2 * range);\n\t\t\t} else if (start <= range) {\n\t\t\t\tstart = 1;\n\t\t\t\tend = Math.min(this.totalPages, start + 2 * range);\n\t\t\t}\n\n\t\t\treturn Array.from({ length: end - start + 1 }, (_, i) => start + i);\n\t\t},\n\n        badgeClass(status) {\n            const baseClass = 'badge bg-body-secondary badge-custom';\n            switch (status) {\n                case 'pending':\n                    return `${baseClass} text-info`;\n\t\t\t\t case 'new':\n                    return `${baseClass} text-info`;\t\n                case 'completed':\n                    return `${baseClass} text-success`;\n                case 'cancelled':\n                    return `${baseClass} text-danger`;\n                case 'processing':\n                    return `${baseClass} text-warning`;\n                default:\n                    return baseClass;\n            }\n        },\n\n        async applyAction() {\n            if (!this.selectedStatus || !this.selectedVouchers.length) {\n                this.errorMessage = 'Please select a status and at least one voucher.';\n                return;\n            }\n            this.loading = true;\n            try {\n                const response = await fetch('/voucher/bulk-action', {\n                    method: 'POST',\n                    headers: {\n                        'Content-Type': 'application/json',\n                    },\n                    body: JSON.stringify({\n                        voucherIds: this.selectedVouchers,\n                        status: this.selectedStatus,\n                    }),\n                });\n                const result = await response.json();\n                if (response.ok) {\n                    this.selectedVouchers = [];\n                    this.selectedStatus = '';\n                    this.selectAllCheckbox = false;\n                    this.fetchVouchers(this.currentPage);\n                } else {\n                    this.errorMessage = result.message;\n                }\n            } catch (error) {\n                this.errorMessage = 'An error occurred while processing the request.';\n            } finally {\n                this.loading = false;\n            }\n        },\n\t\t openOffcanvas(voucher) {\n\t\t\t// Check if products exists and is an array\n\t\t\tif (!Array.isArray(voucher.products)) {\t\t\t\n\t\t\t\t// Set products to an empty array if it's not valid\n\t\t\t\tvoucher.products = []; \n\t\t\t\t\n    \t\t}\n\t\t this.modalOrder = {\n                   ...this.modalOrder,\n                ...voucher,\n                billing: { ...this.modalOrder.billing, ...voucher.billing },\n                shipping: { ...this.modalOrder.shipping, ...voucher.shipping },\n                products: [...voucher.products],\n                payment_method: voucher.payment_method\n            };\n\t\t\t\n            this.showOffcanvas = true;\n\t\t\t  console.log('Opening offcanvas:', this.showOffcanvas);\n\t\t\t\n        },\n        closeOffcanvas() {\n            this.showOffcanvas = false;\n            this.modalOrder = {\n                billing: { ...this.modalOrder.billing },\n                shipping: { ...this.modalOrder.shipping },\n                products: [],\n                payment_method: ''\n            };\n\t\t\t console.log('Closing offcanvas:', this.showOffcanvas);\n        },\n    };\n}\n\n</script><!-- End Tab Content -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
