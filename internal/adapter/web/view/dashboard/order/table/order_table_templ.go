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
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/order/chart/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/order/table/order_table.templ`, Line: 16, Col: 56}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-order\"><i class=\"bi bi-bar-chart-fill\"></i> <span class=\"ms-2\">Charts</span></button> <button class=\"btn btn-sm btn-neutral \" type=\"button\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/voucher/table/view/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/order/table/order_table.templ`, Line: 24, Col: 63}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\"><i class=\"bi bi-box-seam-fill\"></i> <span class=\"ms-2\">Voucher</span></button></div></div></div><div id=\"dashboard-order-table\" x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("orderTable('%s')", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/order/table/order_table.templ`, Line: 32, Col: 85}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"init()\"><div class=\"px-6 px-lg-7 pt-1 border-bottom\"><ul class=\"nav nav-tabs nav-tabs-flush gap-8 overflow-x border-0 mt-4\"><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;all&#39;}\" @click.prevent=\"selectTab(&#39;all&#39;)\">All</a></li><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;completed&#39;}\" @click.prevent=\"selectTab(&#39;completed&#39;)\">Completed</a></li><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;pending&#39;}\" @click.prevent=\"selectTab(&#39;pending&#39;)\">Pending</a></li><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;processing&#39;}\" @click.prevent=\"selectTab(&#39;processing&#39;)\">Processing</a></li><li class=\"nav-item\"><a href=\"#\" :class=\"{&#39;nav-link&#39;: true, &#39;active&#39;: currentTab === &#39;cancelled&#39;}\" @click.prevent=\"selectTab(&#39;cancelled&#39;)\">Canceled</a></li></ul></div><div class=\"d-flex gap-2 py-3 px-7 border-bottom\"><div class=\"dropdown\" x-data=\"{ showDropdown: false }\" @click.outside=\"showDropdown = false\"><button class=\"btn btn-sm btn-neutral dropdown-toggle\" type=\"button\" id=\"dropdownMenuButton2\" @click=\"showDropdown = !showDropdown\" :aria-expanded=\"showDropdown.toString()\"><i class=\"bi bi-plus-circle\"></i> <span class=\"ms-2\">Bulk Action</span></button><div class=\"dropdown-menu\" :class=\"{ &#39;show&#39;: showDropdown }\" aria-labelledby=\"dropdownMenuButton2\"><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statusBacklog\" value=\"asc_courier\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusBacklog\">Send to Couries and update to completed</label></div></div><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statusCompleted\" value=\"completed\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusTodo\">Change status to completed</label></div></div><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statuSpending\" value=\"pending\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusInProgress\">Change status to pending</label></div></div><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statusProdessing\" value=\"processing\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusDone\">Change status to processing</label></div></div><div class=\"dropdown-item py-1 px-2 d-flex align-items-center\"><div class=\"text-lg\"><input class=\"form-check-input\" type=\"radio\" name=\"statusOption\" id=\"statusCancelled\" value=\"cancelled\" x-model=\"selectedStatus\"></div><div class=\"ms-3 me-5\"><label for=\"statusCancelled\">Change status to cancelled</label></div></div><!-- Add other status options here --><div class=\"mt-3\"><button type=\"button\" class=\"btn btn-sm btn-primary d-sm-inline-flex position-relative\" @click=\"applyAction\"><span x-show=\"!loading\">Apply </span> <span x-show=\"loading\" class=\"spinner-border spinner-border-sm\" role=\"status\" aria-hidden=\"true\"></span> <i class=\"px-3 bi bi-arrow-right\"></i></button></div><div class=\"mt-3\"><div x-show=\"errorMessage\" class=\"alert alert-danger\" role=\"alert\"><span x-text=\"errorMessage\"></span></div></div></div></div></div><div class=\"table-responsive\"><table class=\"table table-hover table-nowrap\"><thead><tr><th><div class=\"text-base\"><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" @change=\"selectAll()\" x-model=\"selectAllCheckbox\"></div></div></th><th @click=\"sortTable(&#39;orderId&#39;)\">Order ID <i :class=\"getSortIcon(&#39;orderId&#39;)\"></i></th><th @click=\"sortTable(&#39;timestamp&#39;)\">Date <i :class=\"getSortIcon(&#39;timestamp&#39;)\"></i></th><th @click=\"sortTable(&#39;total_amount&#39;)\">Total Amount <i :class=\"getSortIcon(&#39;total_amount&#39;)\"></i></th><th>Status</th></tr></thead> <tbody><template x-if=\"!loading &amp;&amp; totalItems === 0\"><tr><td colspan=\"5\">No orders found.</td></tr></template><template x-for=\"order in paginatedOrders\" :key=\"order.id\"><tr><td><div class=\"form-check\"><input class=\"form-check-input\" type=\"checkbox\" :value=\"order.orderId\" x-model=\"selectedOrders\"></div></td><td x-text=\"order.orderId\"></td><td x-text=\"new Date(order.timestamp).toLocaleString()\"></td><td x-text=\"order.total_amount\"></td><td><span :class=\"badgeClass(order.status)\" x-text=\"order.status\"></span></td></tr></template><!-- Fill remaining rows if less than 10 --><template x-for=\"i in 10 - paginatedOrders.length\" :key=\"&#39;empty&#39; + i\"><tr><td colspan=\"5\" class=\"py-7\"></td></tr></template></tbody></table></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = p.PaginationControl().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><script>\n\tfunction orderTable(projectId) {\n\t\treturn {\n\t\t\tprojectID: projectId,\n\t\t\tcurrentTab: 'all',\n\t\t\torders: [],\n\t\t\tselectedOrders: [],\n\t\t\tselectAllCheckbox: false,\n\t\t\tsortKey: 'orderId',\n\t\t\tsortAsc: false,\n\t\t\tcurrentPage: 1,\n\t\t\titemsPerPage: 10,\n\t\t\ttotalItems: 0,\n\t\t\ttotalPages: 0,\n\t\t\tloading: false,\n\t\t\tselectedStatus: '',\n\t\t\terrorMessage: '',\n\n\t\t\tasync init() {\n\t\t\t\tawait this.fetchOrders(this.currentPage);\n\t\t\t},\n\n\t\t\tasync fetchOrders(page = 1) {\n\t\t\t\tthis.loading = true;\n\t\t\t\ttry {\n\t\t\t\t\tconst url = this.getUrlForTab(this.currentTab, page);\n\t\t\t\t\tconst response = await fetch(url);\n\t\t\t\t\tconst result = await response.json();\n\t\t\t\t\tif (response.ok) {\n\t\t\t\t\t\tthis.orders = result.data || [];\n\t\t\t\t\t\tthis.totalItems = result.meta.totalItems || 0;\n\t\t\t\t\t\tthis.currentPage = result.meta.currentPage || 1;\n\t\t\t\t\t\tthis.itemsPerPage = result.meta.itemsPerPage || 10;\n\t\t\t\t\t\tthis.totalPages = result.meta.totalPages || 0;\n\t\t\t\t\t} else {\n\t\t\t\t\t\tconsole.error('Error fetching data:', result.message);\n\t\t\t\t\t}\n\t\t\t\t} catch (error) {\n\t\t\t\t\tconsole.error('Error fetching data:', error);\n\t\t\t\t} finally {\n\t\t\t\t\tthis.loading = false;\n\t\t\t\t}\n\t\t\t},\n\n\t\t\tgetUrlForTab(tab, page) {\n\t\t\t\tconst baseUrl = `${window.location.origin}/order/table/${this.projectID}`;\n\t\t\t\tconst sortDirection = this.sortAsc ? 'asc' : 'desc'; // Determine sort direction\n\t\t\t\tswitch (tab) {\n\t\t\t\t\tcase 'all':\n\t\t\t\t\t\treturn `${baseUrl}/all/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n\t\t\t\t\tcase 'completed':\n\t\t\t\t\t\treturn `${baseUrl}/completed/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n\t\t\t\t\tcase 'processing':\n\t\t\t\t\t\treturn `${baseUrl}/processing/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n\t\t\t\t\tcase 'pending':\n\t\t\t\t\t\treturn `${baseUrl}/pending/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n\t\t\t\t\tcase 'cancelled':\n\t\t\t\t\t\treturn `${baseUrl}/cancelled/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n\t\t\t\t\tdefault:\n\t\t\t\t\t\treturn `${baseUrl}/completed/${page}?sort=${this.sortKey}&direction=${sortDirection}`;\n\t\t\t\t}\n\t\t\t},\n\n\t\t\tselectTab(tab) {\n\t\t\t\tthis.currentTab = tab;\n\t\t\t\tthis.currentPage = 1; // Reset to first page\n\t\t\t\tthis.fetchOrders(this.currentPage);\n\t\t\t},\n\n\t\t\tselectAll() {\n\t\t\t\tthis.selectedOrders = this.selectAllCheckbox ? this.orders.map(order => order.orderId) : [];\n\t\t\t},\n\n\t\t\tsortTable(key) {\n\t\t\t\tif (this.sortKey === key) {\n\t\t\t\t\tthis.sortAsc = !this.sortAsc; // Toggle sort direction if the same column is clicked\n\t\t\t\t} else {\n\t\t\t\t\tthis.sortKey = key; // Set new sort key\n\t\t\t\t\tthis.sortAsc = true; // Default to ascending if a new column is selected\n\t\t\t\t}\n\t\t\t\tthis.fetchOrders(this.currentPage); // Fetch sorted data\n\t\t\t},\n\t\t\tgetSortIcon(key) {\n\t\t\t\tif (this.sortKey !== key) return '';\n\t\t\t\treturn this.sortAsc ? 'bi bi-chevron-up' : 'bi bi-chevron-down';\n\t\t\t},\n\n\t\t\tchangePage(page) {\n\t\t\t\tif (page < 1 || page > this.totalPages) return;\n\t\t\t\tthis.fetchOrders(page);\n\t\t\t},\n\n\t\t\tget paginatedOrders() {\n\t\t\t\treturn this.orders;\n\t\t\t},\n\n\t\t\tget currentPageStart() {\n\t\t\t\treturn (this.currentPage - 1) * this.itemsPerPage + 1;\n\t\t\t},\n\n\t\t\tget currentPageEnd() {\n\t\t\t\treturn Math.min(this.currentPage * this.itemsPerPage, this.totalItems);\n\t\t\t},\n\t\t\tget pageNumbers() {\n\t\t\t\tconst range = 2; // Number of pages to show around the current page\n\t\t\t\tlet start = Math.max(1, this.currentPage - range);\n\t\t\t\tlet end = Math.min(this.totalPages, this.currentPage + range);\n\n\t\t\t\t// Adjust range if there are not enough pages on one side\n\t\t\t\tif (this.totalPages - end < range) {\n\t\t\t\t\tend = this.totalPages;\n\t\t\t\t\tstart = Math.max(1, end - 2 * range);\n\t\t\t\t} else if (start <= range) {\n\t\t\t\t\tstart = 1;\n\t\t\t\t\tend = Math.min(this.totalPages, start + 2 * range);\n\t\t\t\t}\n\n\t\t\t\treturn Array.from({ length: end - start + 1 }, (_, i) => start + i);\n\t\t\t},\n\t\t\tbadgeClass(status) {\n\t\t\t\tconst baseClass = 'badge bg-body-secondary badge-custom'; // Add badge-custom class\n\t\t\t\tswitch (status) {\n\t\t\t\t\tcase 'pending':\n\t\t\t\t\t\treturn `${baseClass} text-warning`;\n\t\t\t\t\tcase 'completed':\n\t\t\t\t\t\treturn `${baseClass} text-success`;\n\t\t\t\t\tcase 'cancelled':\n\t\t\t\t\t\treturn `${baseClass} text-danger`;\n\t\t\t\t\tcase 'processing':\n\t\t\t\t\t\treturn `${baseClass} text-warning`;\n\t\t\t\t\tdefault:\n\t\t\t\t\t\treturn baseClass;\n\t\t\t\t}\n\t\t\t},\n\t\t\tasync applyAction() {\n\t\t\t\tthis.loading = true;\n\t\t\t\tthis.errorMessage = \"\"; // Clear previous error message\n\t\t\t\tthis.showDropdown = true; // Close dropdown on action\n\n\t\t\t\tconst selectedOrderIds = this.selectedOrders.join(',');\n\t\t\t\tif (!selectedOrderIds || !this.selectedStatus) {\n\t\t\t\t\tthis.errorMessage = 'Please select at least one order and a status.';\n\t\t\t\t\tthis.loading = false;\n\t\t\t\t\tthis.showDropdown = true;\n\t\t\t\t\treturn;\n\t\t\t\t}\n\n\t\t\t\ttry {\n\t\t\t\t\tconst response = await fetch(`/order/bulk-action/${this.projectID}`, {\n\t\t\t\t\t\tmethod: 'POST',\n\t\t\t\t\t\theaders: { 'Content-Type': 'application/json' },\n\t\t\t\t\t\tbody: JSON.stringify({\n\t\t\t\t\t\t\tstatus: this.selectedStatus,\n\t\t\t\t\t\t\torders: selectedOrderIds.split(','),\n\t\t\t\t\t\t}),\n\t\t\t\t\t});\n\t\t\t\t\tconst result = await response.json();\n\n\t\t\t\t\tif (response.ok) {\n\t\t\t\t\t\t// Handle success (e.g., show a success message or refresh the table)\n\t\t\t\t\t\tthis.fetchOrders(this.currentPage);\n\t\t\t\t\t\tconsole.log('Bulk action successful:', result);\n\t\t\t\t\t\tthis.showDropdown = false;\n\t\t\t\t\t} else {\n\t\t\t\t\t\tthis.errorMessage = result.message || 'An error occurred while processing the request.';\n\t\t\t\t\t}\n\t\t\t\t} catch (error) {\n\t\t\t\t\tconsole.error('Error during bulk action:', error);\n\t\t\t\t\tthis.errorMessage = 'An error occurred while processing the request.';\n\t\t\t\t\tthis.showDropdown = true;\n\t\t\t\t} finally {\n\t\t\t\t\tthis.loading = false;\n\n\t\t\t\t}\n\t\t\t},\n\n\t\t};\n\t}\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
