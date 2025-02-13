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

func CustomerTable(projectId string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-order-table\" x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("customerTable('%s')", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/customer/table/customer_table.templ`, Line: 9, Col: 87}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"init()\"><div class=\"row align-items-center g-6 mt-0 mb-6\"><div class=\"col-sm-6\"><div class=\"d-flex gap-2\"><div class=\"input-group input-group-sm input-group-inline w-100 w-md-50\"><span class=\"input-group-text\"><i class=\"bi bi-search me-2\"></i></span> <input type=\"search\" class=\"form-control ps-0\" placeholder=\"Search all customers\" aria-label=\"Search\"></div></div></div></div><div class=\"border-top\"><div class=\"table-responsive\"><table class=\"table table-hover table-sm  table-striped  table-nowrap\"><thead><tr><th scope=\"col\"><span data-i18n=\"dashboard-table-customer_name\">Name </span></th><th @click=\"sortTable(&#39;email&#39;)\"><span data-i18n=\"dashboard-table-customer_email\">Email </span> <i :class=\"getSortIcon(&#39;email&#39;)\"></i></th><th @click=\"sortTable(&#39;order_count&#39;)\"><span data-i18n=\"dashboard-table-customer-total_orders\">Total Orders </span> <i :class=\"getSortIcon(&#39;order_count&#39;)\"></i></th><th scope=\"col\"><span data-i18n=\"dashboard-table-customer-total_spent\">Money Spend </span></th></tr></thead> <tbody><template x-if=\"!loading &amp;&amp; totalItems === 0\"><tr><td colspan=\"8\">No customer found.</td></tr></template><template x-for=\"customer in paginatedCustomers\" :key=\"customer.id\"><tr><td x-text=\"customer.name\"></td><td x-text=\"customer.email\"></td><td x-text=\"customer.totalOrders\"></td><td x-text=\"customer.totalSpent\"></td></tr></template><template x-for=\"i in 10 - paginatedCustomers.length\" :key=\"&#39;empty&#39; + i\"><tr><td colspan=\"5\" class=\"py-5\"></td></tr></template></tbody></table></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = p.PaginationControl().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><script>\nfunction customerTable(projectId) {\n    return {\n\t\tprojectID: projectId,\n        currentTab: 'all',\n        customers: [],\n        selectedcustomers: [],\n        selectAllCheckbox: false,\n        sortKey: 'order_count',\n        sortAsc: false,\n        currentPage: 1,\n        itemsPerPage: 10,\n        totalItems: 0,\n        totalPages: 0,\n        loading: false,\n\n        async init() {\n            await this.fetchcustomers(this.currentPage);\n        },\n\n        async fetchcustomers(page = 1) {\n            this.loading = true;\n            try {\n                const url = this.getUrlForTab(this.currentTab, page);\n                const response = await fetch(url);\n                const result = await response.json();\n                if (response.ok) {\n                    this.customers = result.data || [];\n                    this.totalItems = result.meta.totalItems || 0;\n                    this.currentPage = result.meta.currentPage || 1;\n                    this.itemsPerPage = result.meta.itemsPerPage || 10;\n                    this.totalPages = result.meta.totalPages || 0;\n                } else {\n                    console.error('Error fetching data:', result.message);\n                }\n            } catch (error) {\n                console.error('Error fetching data:', error);\n            } finally {\n                this.loading = false;\n            }\n        },\n\n         getUrlForTab(tab, page) {\n            const baseUrl = `${window.location.origin}/customer/table/${this.projectID}`;\n            const sortDirection = this.sortAsc ? 'asc' : 'desc'; // Determine sort direction\n            switch (tab) {\n                case 'all':\n                    return `${baseUrl}/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'completed':\n                    return `${baseUrl}/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'processing':\n                    return `${baseUrl}/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'pending':\n                    return `${baseUrl}/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                case 'cancelled':\n                    return `${baseUrl}/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n                default:\n                    return `${baseUrl}/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n            }\n        },\n\n        selectTab(tab) {\n            this.currentTab = tab;\n            this.currentPage = 1; // Reset to first page\n            this.fetchcustomers(this.currentPage);\n        },\n\n        selectAll() {\n            this.selectedcustomers = this.selectAllCheckbox ? this.customers.map(order => order.orderId) : [];\n        },\n\n      \tsortTable(key) {\n            if (this.sortKey === key) {\n                this.sortAsc = !this.sortAsc; // Toggle sort direction if the same column is clicked\n            } else {\n                this.sortKey = key; // Set new sort key\n                this.sortAsc = true; // Default to ascending if a new column is selected\n            }\n            this.fetchcustomers(this.currentPage); // Fetch sorted data\n        },\n\t\tgetSortIcon(key) {\n            if (this.sortKey !== key) return '';\n            return this.sortAsc ? 'bi bi-chevron-up' : 'bi bi-chevron-down';\n        },\n\n        changePage(page) {\n            if (page < 1 || page > this.totalPages) return;\n            this.fetchcustomers(page);\n        },\n\n        get paginatedCustomers() {\n            return this.customers;\n        },\n\n        get currentPageStart() {\n            return (this.currentPage - 1) * this.itemsPerPage + 1;\n        },\n\n        get currentPageEnd() {\n            return Math.min(this.currentPage * this.itemsPerPage, this.totalItems);\n        },\n       \tget pageNumbers() {\n            const range = 2; // Number of pages to show around the current page\n            let start = Math.max(1, this.currentPage - range);\n            let end = Math.min(this.totalPages, this.currentPage + range);\n\n            // Adjust range if there are not enough pages on one side\n            if (this.totalPages - end < range) {\n                end = this.totalPages;\n                start = Math.max(1, end - 2 * range);\n            } else if (start <= range) {\n                start = 1;\n                end = Math.min(this.totalPages, start + 2 * range);\n            }\n\t\t\t\n            return Array.from({ length: end - start + 1 }, (_, i) => start + i);\n        },\n\t\t badgeClass(status) {\n\t\t\t\tconst baseClass = 'badge bg-body-secondary badge-custom'; // Add badge-custom class\n\t\t\t\tswitch (status) {\n\t\t\t\t\tcase 'pending':\n\t\t\t\t\t\treturn `${baseClass} text-warning`;\n\t\t\t\t\tcase 'completed':\n\t\t\t\t\t\treturn `${baseClass} text-success`;\n\t\t\t\t\tcase 'cancelled':\n\t\t\t\t\t\treturn `${baseClass} text-danger`;\n\t\t\t\t\tcase 'processing':\n\t\t\t\t\t\treturn `${baseClass} text-warning`;\n\t\t\t\t\tdefault:\n\t\t\t\t\t\treturn baseClass;\n\t\t\t\t}\n\t\t}\t\t\n    };\n}\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
