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

func OrderTable(projectId string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-order\"><div class=\"mb-1 mb-xl-10\"><div class=\"col\"><div class=\"hstack gap-2 justify-content-start\"><button class=\"btn btn-sm btn-neutral \" type=\"button\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/voucher/table/html/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/order/table/order_table.templ`, Line: 24, Col: 63}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-order\"><i class=\"bi bi-box-seam-fill\"></i> <span class=\"ms-2\">Vouchers</span></button></div></div></div><div id=\"dashboard-order-table\" x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("orderTable('%s')", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/order/table/order_table.templ`, Line: 32, Col: 85}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"init()\"><div class=\"px-6 px-lg-7 pt-1 border-bottom\"><ul class=\"nav nav-tabs nav-tabs-flush gap-8 overflow-x border-0 mt-4\"><li class=\"nav-item\"><a data-i18n=\"order-table-all-orders\" href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;all&#39;}\" @click.prevent=\"selectTab(&#39;all&#39;)\">All</a></li><li class=\"nav-item\"><a data-i18n=\"order-table-completed-orders\" href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;completed&#39;}\" @click.prevent=\"selectTab(&#39;completed&#39;)\">Completed</a></li><li class=\"nav-item\"><a data-i18n=\"order-table-pending-orders\" href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;pending&#39;}\" @click.prevent=\"selectTab(&#39;pending&#39;)\">Pending</a></li><li class=\"nav-item\"><a data-i18n=\"order-table-processing-orders\" href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;processing&#39;}\" @click.prevent=\"selectTab(&#39;processing&#39;)\">Processing</a></li><li class=\"nav-item\"><a data-i18n=\"order-table-cancelled-orders\" href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;cancelled&#39;}\" @click.prevent=\"selectTab(&#39;cancelled&#39;)\">Canceled</a></li></ul></div><div class=\"d-flex gap-2 py-3 px-7 border-bottom\"><div class=\"dropdown\" x-data=\"{ showDropdown: false }\" @click.outside=\"showDropdown = false\"><button class=\"btn btn-sm btn-neutral dropdown-toggle\" type=\"button\" id=\"dropdownMenuButton2\" @click=\"showDropdown = !showDropdown\" :aria-expanded=\"showDropdown.toString()\"><i class=\"bi bi-plus-circle\"></i> <span class=\"ms-2\">Bulk Action</span></button><div class=\"dropdown-menu\" :class=\"{ &#39;show&#39;: showDropdown }\" aria-labelledby=\"dropdownMenuButton2\"><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statusCompleted\" value=\"completed\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusTodo\">Change status to completed</label></div></div><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statuSpending\" value=\"pending\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusInProgress\">Change status to pending</label></div></div><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statusProdessing\" value=\"processing\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusDone\">Change status to processing</label></div></div><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statusCancelled\" value=\"cancelled\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusCancelled\">Change status to cancelled</label></div></div><!-- Add other status options here --><div class=\"mt-3\"><button type=\"button\" class=\"btn btn-sm btn-primary d-sm-inline-flex position-relative\" @click=\"applyAction\"><span x-show=\"!loading\">Apply </span> <span x-show=\"loading\" class=\"spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> <i class=\"px-3 bi bi-arrow-right\"></i></button></div><div class=\"mt-3\"><div x-show=\"errorMessage\" class=\"alert alert-danger\" role=\"alert\"><span x-text=\"errorMessage\"></span></div></div></div></div></div><div class=\"table-responsive\"><table class=\"table table-hover table-striped table-sm table-nowrap\"><thead><tr><th><div class=\"text-base\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" @change=\"selectAll()\" x-model=\"selectAllCheckbox\"></div></div></th><th @click=\"sortTable(&#39;orderId&#39;)\"><span data-i18n=\"dashboard-table-order_id\">Order ID</span> <i :class=\"getSortIcon(&#39;orderId&#39;)\"></i></th><th @click=\"sortTable(&#39;order_created&#39;)\"><span data-i18n=\"dashboard-table-order-created\">Order Created</span> <i :class=\"getSortIcon(&#39;order_created&#39;)\"></i></th><th @click=\"sortTable(&#39;total_amount&#39;)\"><span data-i18n=\"dashboard-table-order_total\">Total Amount</span> <i :class=\"getSortIcon(&#39;total_amount&#39;)\"></i></th><th><span data-i18n=\"dashboard-table-order_status\">Status</span></th><th><span data-i18n=\"dashboard-table-order_action\">Action</span></th></tr></thead> <tbody><!-- Display this row if no orders are found --><template x-if=\"!loading &amp;&amp; totalItems === 0\"><tr><td colspan=\"6\">No orders found.</td></tr></template><!-- Loop through paginated orders --><template x-for=\"order in paginatedOrders\" :key=\"order.id\"><tr><td><div class=\"text-base\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" :value=\"order.orderId\" x-model=\"selectedOrders\"></div></div></td><td x-text=\"&#39;#&#39; + order.orderId\"></td><td x-text=\"new Date(order.order_created).toLocaleString()\"></td><td x-text=\"order.total_amount + &#39; &#39; + order.currency_symbol\"></td><td><span :class=\"badgeClass(order.status)\" x-text=\"order.status\"></span></td><td><button type=\"button\" class=\"btn btn-sm btn-neutral\" @click=\"loadModalData(order)\"><i class=\"bi bi-eye\"></i></button></td></tr></template><template x-for=\"i in 10 - paginatedOrders.length\" :key=\"&#39;empty&#39; + i\"><tr><td colspan=\"7\" class=\"py-5\"></td></tr></template></tbody></table></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = OrderModal().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = p.PaginationControl().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><script>\n\n\n\tfunction orderTable(projectId) {\n    return {\n        projectID: projectId,\n        currentTab: 'all',\n        orders: [],\n        selectedOrders: [],\n        selectAllCheckbox: false,\n        sortKey: 'orderId',\n        sortAsc: false,\n        currentPage: 1,\n        itemsPerPage: 10,\n        totalItems: 0,\n        totalPages: 0,\n        loading: false,\n        selectedStatus: '',\n        errorMessage: '',\n        showDropdown: false, // Initialize showDropdown\n        modalOrder: {\n            billing: {\n                first_name: '',\n                last_name: '',\n                address_1: '',\n                city: '',\n                postcode: '',\n                email: '',\n                phone: '',\n            },\n            shipping: {\n                first_name: '',\n                last_name: '',\n                address_1: '',\n                city: '',\n                postcode: '',\n            },\n            products: [],\n            payment_method: '',\n\t\t\tcustomer_note:''\n        },\n        showModal: false,\n\n        async init() {\n            await this.fetchOrders(this.currentPage);\n        },\n\n        async fetchOrders(page = 1) {\n            this.loading = true;\n            try {\n                const url = this.getUrlForTab(this.currentTab, page);\n                const response = await fetch(url);\n                const result = await response.json();\n                if (response.ok) {\n                    this.orders = result.data || [];\n                    this.totalItems = result.meta.totalItems || 0;\n                    this.currentPage = result.meta.currentPage || 1;\n                    this.itemsPerPage = result.meta.itemsPerPage || 10;\n                    this.totalPages = result.meta.totalPages || 0;\n                } else {\n                    console.error('Error fetching data:', result.message);\n\t\t\t\t\tthis.errorMessage = result.message || 'An error occurred while fetching data.';\n                }\n            } catch (error) {\n                console.error('Error fetching data:', error);\n            } finally {\n                this.loading = false;\n            }\n        },\n\n        getUrlForTab(tab, page) {\n            const baseUrl = `${window.location.origin}/order/table/${this.projectID}`;\n            const sortDirection = this.sortAsc ? 'asc' : 'desc'; \n            switch (tab) {\n                case 'all':\n                    return `${baseUrl}/all/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'completed':\n                    return `${baseUrl}/completed/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'processing':\n                    return `${baseUrl}/processing/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'pending':\n                    return `${baseUrl}/pending/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'cancelled':\n                    return `${baseUrl}/cancelled/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                default:\n                    return `${baseUrl}/completed/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n            }\n        },\n\n        selectTab(tab) {\n            this.currentTab = tab;\n            this.currentPage = 1;\n            this.fetchOrders(this.currentPage);\n        },\n\n        selectAll() {\n            this.selectedOrders = this.selectAllCheckbox ? this.orders.map(order => order.orderId) : [];\n        },\n\n        sortTable(key) {\n            if (this.sortKey === key) {\n                this.sortAsc = !this.sortAsc;\n            } else {\n                this.sortKey = key;\n                this.sortAsc = true;\n            }\n            this.fetchOrders(this.currentPage);\n        },\n\n        getSortIcon(key) {\n            if (this.sortKey !== key) return '';\n            return this.sortAsc ? 'bi bi-chevron-up' : 'bi bi-chevron-down';\n        },\n\n        changePage(page) {\n            if (page < 1 || page > this.totalPages) return;\n            this.fetchOrders(page);\n        },\n\n        get paginatedOrders() {\n            return this.orders;\n        },\n\n        get currentPageStart() {\n            return (this.currentPage - 1) * this.itemsPerPage + 1;\n        },\n\n        get currentPageEnd() {\n            return Math.min(this.currentPage * this.itemsPerPage, this.totalItems);\n        },\n\n        get pageNumbers() {\n            const range = 2;\n            let start = Math.max(1, this.currentPage - range);\n            let end = Math.min(this.totalPages, this.currentPage + range);\n\n            if (this.totalPages - end < range) {\n                end = this.totalPages;\n                start = Math.max(1, end - 2 * range);\n            } else if (start <= range) {\n                start = 1;\n                end = Math.min(this.totalPages, start + 2 * range);\n            }\n\n            return Array.from({ length: end - start + 1 }, (_, i) => start + i);\n        },\n\n        badgeClass(status) {\n            const baseClass = 'badge bg-body-secondary badge-custom';\n            switch (status) {\n                case 'pending':\n                    return `${baseClass} text-warning`;\n                case 'completed':\n                    return `${baseClass} text-success`;\n                case 'cancelled':\n                    return `${baseClass} text-danger`;\n                case 'processing':\n                    return `${baseClass} text-warning`;\n                default:\n                    return baseClass;\n            }\n        },\n\n        async applyAction() {\n            this.loading = true;\n            this.errorMessage = \"\"; \n            this.showDropdown = false; \n\n            const selectedOrderIds = this.selectedOrders.join(',');\n            if (!selectedOrderIds || !this.selectedStatus) {\n                this.errorMessage = 'Please select at least one order and a status.';\n                this.loading = false;\n                this.showDropdown = true;\n                return;\n            }\n\n            try {\n                const response = await fetch(`/order/bulk-action/${this.projectID}`, {\n                    method: 'POST',\n                    headers: { 'Content-Type': 'application/json' },\n                    body: JSON.stringify({\n                        status: this.selectedStatus,\n                        orders: selectedOrderIds.split(','),\n                    }),\n                });\n                const result = await response.json();\n\n                if (response.ok) {\n                    this.fetchOrders(this.currentPage);\n                    console.log('Bulk action successful:', result);\n                } else {\n                    this.errorMessage = result.message || 'An error occurred while processing the request.';\n                }\n            } catch (error) {\n                console.error('Error during bulk action:', error);\n                this.errorMessage = 'An error occurred while processing the request.';\n            } finally {\n                this.loading = false;\n            }\n        },\n\n        loadModalData(order) {\n            this.modalOrder = {\n                ...this.modalOrder,\n                ...order,\n                billing: { ...this.modalOrder.billing, ...order.billing },\n                shipping: { ...this.modalOrder.shipping, ...order.shipping }\n            };\n            this.showModal = true;\n        },\n\n        closeModal() {\n            this.showModal = false;\n\t\t\t  this.errorMessage ='';\n            this.modalOrder = {\n                billing: { ...this.modalOrder.billing },\n                shipping: { ...this.modalOrder.shipping },\n                products: [],\n                payment_method: '',\n\t\t\t\tcustomer_note:''\n            };\n        },\n\t\t async saveChanges() {\n            this.loading = true;\n            try {\n                const response = await fetch(`/order/update/${this.modalOrder.orderId}/${this.projectID}`, {\n                    method: 'PUT', // Or 'PUT' based on your API\n                    headers: {\n                        'Content-Type': 'application/json',\n                    },\n                    body: JSON.stringify({\n                        billing: this.modalOrder.billing,\n                        shipping: this.modalOrder.shipping,\n                        products: this.modalOrder.products,\n                        payment_method: this.modalOrder.payment_method,\n                        customer_note: this.modalOrder.customer_note,\n                    }),\n                });\n\n                const result = await response.json();\n\n                if (response.ok) {\n                    // Successfully saved, close the modal and refresh the order list\n                    this.closeModal();\n                    this.fetchOrders(this.currentPage);  // Refresh orders list\n                    console.log('Order updated successfully:', result);\n                } else {\n\t\t\t\t\t this.errorMessage =  'An error occurred while updating data.';\n                    console.error('Error saving changes:', result.message);\n                }\n            } catch (error) {\n                console.error('Error saving changes:', error);\n            } finally {\n                this.loading = false;\n            }\n        },\n    };\n}\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
