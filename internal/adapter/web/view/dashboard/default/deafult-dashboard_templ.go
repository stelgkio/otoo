// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	o "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/order/history"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/dashboard/product/best_seller"
	woo "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
)

func DeafultDashboard(projectId string, counts map[string]string, orders []*woo.OrderRecord, bestSeller []*woo.ProductBestSellerRecord, weeklyBalance *woo.WeeklyAnalytics) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dashboard-content\" class=\"flex-fill overflow-y-lg-auto scrollbar bg-body rounded-top-4 rounded-top-start-lg-4 rounded-top-end-lg-0 border-top border-lg shadow-2\"><main class=\"container-fluid px-3 py-5 p-lg-6 p-xxl-8\"><div class=\"mb-6 mb-xl-10\"><div class=\"row g-3 align-items-center\"><div class=\"col\"><h1 class=\"ls-tight\" data-i18n=\"Dashboard\"></h1></div><div class=\"col\"><div class=\"hstack gap-2 justify-content-end\"><button type=\"button\" class=\"btn btn-sm btn-square btn-neutral rounded-circle d-xxl-none\" data-bs-toggle=\"offcanvas\" data-bs-target=\"#responsiveOffcanvas\" aria-controls=\"responsiveOffcanvas\"><i class=\"bi bi-three-dots\"></i></button></div></div></div></div><div class=\"row g-3 g-xxl-6\"><div class=\"col-xxl-8\"><div class=\"vstack gap-3 gap-md-6\"><div class=\"row g-3\"><div class=\"col-md col-sm-6\"><div class=\"card border-primary-hover\"><div class=\"card-body p-4\"><div class=\"d-flex align-items-center gap-2\"><i class=\"bi bi-bag-plus\"></i> <a href=\"javascript:void(0)\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dashboard/order/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 47, Col: 66}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\" class=\"h6 stretched-link\" hx-push-url=\"true\">Orders</a></div><div class=\"text-sm fw-semibold mt-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(counts["order_count"])
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 55, Col: 71}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- <div class=\"d-flex align-items-center gap-2 mt-1 text-xs\"><span\n                                                        class=\"badge badge-xs bg-success\"><i\n                                                            class=\"bi bi-arrow-up-right\"></i> </span><span>+13.7%</span>\n                                                </div> --></div></div></div><div class=\"col-md col-sm-6\"><div class=\"card border-primary-hover\"><div class=\"card-body p-4\"><div class=\"d-flex align-items-center gap-2\"><i class=\"bi bi-file-earmark-person\"></i> <a href=\"javascript:void(0)\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dashboard/customer/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 70, Col: 69}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\" class=\"h6 stretched-link\" hx-push-url=\"true\">Customers</a></div><div class=\"text-sm fw-semibold mt-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(counts["customer_count"])
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 76, Col: 74}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- <div class=\"d-flex align-items-center gap-2 mt-1 text-xs\"><span\n                                                        class=\"badge badge-xs bg-danger\"><i\n                                                            class=\"bi bi-arrow-up-right\"></i> </span><span>-3.2%</span>\n                                                </div> --></div></div></div><div class=\"col-md col-sm-6\"><div class=\"card border-primary-hover\"><div class=\"card-body p-4\"><div class=\"d-flex align-items-center gap-2\"><i class=\"bi bi-shop\"></i> <a href=\"javascript:void(0)\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/dashboard/product/%s", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 91, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dashboard-content\" class=\"h6 stretched-link\" hx-push-url=\"true\">Products</a></div><div class=\"text-sm fw-semibold mt-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(counts["product_count"])
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 97, Col: 73}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- <div class=\"d-flex align-items-center gap-2 mt-1 text-xs\"><span\n                                                        class=\"badge badge-xs bg-danger\"><i\n                                                            class=\"bi bi-arrow-up-right\"></i> </span><span>-2.2%</span>\n                                                </div> --></div></div></div><!-- <div class=\"col-md-1 d-none d-md-block\">\n                                        <div\n                                            class=\"card h-md-100 d-flex flex-column align-items-center justify-content-center py-4 bg-body-secondary bg-opacity-75 bg-opacity-100-hover\">\n                                            <a href=\"#cryptoModal\" class=\"stretched-link text-body-secondary\"\n                                                data-bs-toggle=\"modal\"><i class=\"bi bi-plus-lg\"></i></a>\n                                        </div>\n                                    </div> --></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = MonthlyChart(projectId).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = o.LatestOrderHistory(orders).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"col-xxl-4\"><div class=\"offcanvas-xxl m-xxl-0 rounded-sm-4 rounded-xxl-0 offcanvas-end overflow-hidden m-sm-4\" tabindex=\"-1\" id=\"responsiveOffcanvas\" aria-labelledby=\"responsiveOffcanvasLabel\"><div class=\"offcanvas-header rounded-top-4\"><h5 class=\"offcanvas-title\" id=\"responsiveOffcanvasLabel\">Quick Stats</h5><button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"offcanvas\" data-bs-target=\"#responsiveOffcanvas\" aria-label=\"Close\"></button></div><div class=\"offcanvas-body d-flex flex-column p-3 p-sm-6 p-xxl-0 gap-3 gap-xxl-6\"><div class=\"vstack gap-6 gap-xxl-6\"><div class=\"card border-0 border-xxl\"><div class=\"card-body d-flex flex-column p-0 p-xxl-6\"><div class=\"d-flex justify-content-between align-items-center mb-3\"><div><h5>Weekly Balance</h5></div><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if weeklyBalance != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"text-heading fw-bold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%2.f", weeklyBalance.AnalyticsBase.ActiveOrderRate))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 145, Col: 80}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" %</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"text-heading fw-bold\">0%</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if weeklyBalance != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-2xl fw-bolder text-heading ls-tight\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("€%.2f", weeklyBalance.AnalyticsBase.TotalRevenue))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 156, Col: 78}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-2xl fw-bolder text-heading ls-tight\">€ 0</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"d-flex align-items-center justify-content-between mt-8\"><!-- <div class=\"\">\n                                                        <div class=\"d-flex gap-3 align-items-center\">\n                                                            <div\n                                                                class=\"icon icon-sm icon-shape text-sm rounded-circle bg-dark text-success\">\n                                                                <i class=\"bi bi-arrow-down\"></i>\n                                                            </div><span class=\"h6 fw-semibold text-muted\">Income</span>\n                                                        </div>\n                                                        <div class=\"fw-bold text-heading mt-3\">$23.863,21 USD</div>\n                                                    </div><span class=\"vr bg-dark bg-opacity-10\"></span>\n                                                    <div class=\"\">\n                                                        <div class=\"d-flex gap-3 align-items-center\">\n                                                            <div\n                                                                class=\"icon icon-sm icon-shape text-sm rounded-circle bg-dark text-danger\">\n                                                                <i class=\"bi bi-arrow-up\"></i>\n                                                            </div><span\n                                                                class=\"h6 fw-semibold text-muted\">Expenses</span>\n                                                        </div>\n                                                        <div class=\"fw-bold text-heading mt-3\">$5.678,45 USD</div>\n                                                    </div> --></div></div></div><hr class=\"my-0 d-xxl-none\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = p.ProductBestSeller(bestSeller).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<hr class=\"my-0 d-xxl-none\"></div></div></div></div></div></main></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func MonthlyChart(projectId string) templ.Component {
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
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("chartComponent('%s')", projectId))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/adapter/web/view/dashboard/default/deafult-dashboard.templ`, Line: 199, Col: 61}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-init=\"init()\"><div class=\"card\"><div class=\"card-body pb-0\"><div class=\"d-flex justify-content-between align-items-center\"><div><h5>Orders</h5></div><div class=\"hstack align-items-center\"><a href=\"#\" class=\"text-muted\" @click=\"refreshChart\"><i class=\"bi bi-arrow-repeat\"></i></a></div></div><div class=\"mx-n4\"><div id=\"chart-users\" data-height=\"270\"></div></div></div></div></div><script>\n  \t\t function chartComponent(projectId) {\n    return {\n            projectID: projectId,  // Store project ID\n            chartData: {\n                months: [],  // Categories (Months)\n                orders: []  // Data for the chart\n            },\n        \tchartInstance: null,  // Store the ApexCharts instance\n \t\t\t\n            // Initialize and fetch data\n            async init() {\n                await this.fetchChartData();  // Fetch data first\n            },\n\n            // Fetch data from the server\n            async fetchChartData() {\n                try {\n                    const response = await fetch(`${window.location.origin}/order/monthy/chart/${this.projectID}`);\n                    const data = await response.json();                   \n\n                    // Assuming the API returns \"months\" and \"orders\" arrays\n                    this.chartData.months = data.months || [];\n                    this.chartData.orders = data.orders || [];\n\n                    // Render or update the chart after data is fetched\n                    this.renderChart();\n                } catch (error) {\n                    console.error('Error fetching chart data:', error);\n                }\n            },\n\n            // Render the chart using ApexCharts\n            renderChart() {\n                const chartElement = document.querySelector(\"#chart-users\");\n\n                const options = {\n                    chart: {\n                        type: \"bar\",\n                        stacked: true,\n                        zoom: { enabled: true },\n                        toolbar: { show: true },\n                        height: 390,\n                        animations: {\n                            enabled: true,\n                            speed: 800,\n                        },\n                    },\n                    colors: ['#8957ff', '#ffc107', '#dc3545'],\n                    plotOptions: {\n                        bar: {\n                            columnWidth: \"23px\",\n                            borderRadius: 2\n                        }\n                    },\n                    series: [{\n                        name: \"Orders\",\n                        data: this.chartData.orders  // Fetched orders data\n                    }],\n                    xaxis: {\n                        categories: this.chartData.months,  // Fetched months (categories)\n                        labels: {\n                            style: {\n                                colors: \"#827f7f\",\n                                fontSize: \"13px\"\n                            }\n                        }\n                    },\n                    yaxis: {\n                        labels: {\n                            style: {\n                                colors: \"#827f7f\",\n                                fontSize: \"13px\"\n                            }\n                        }\n                    },\n                    grid: {\n                        borderColor: \"#999\",\n                        strokeDashArray: 3\n                    },\n                    dataLabels: {\n                        enabled: false\n                    },\n                    tooltip: {\n                        shared: true,\n                        intersect: false,\n                        y: {\n                            formatter: function (val) {\n                                return val.toFixed(0) + \" orders\";\n                            }\n                        }\n                    }\n                };\n\n                // Initialize chart only once\n                if (!this.chartInstance) {\n                    this.chartInstance = new ApexCharts(chartElement, options);\n                    this.chartInstance.render();\n                } else {\n                    this.chartInstance.updateOptions(options);  // Update chart if it already exists\n                }\n            },\n\n            // Method to refresh the chart (if needed)\n            refreshChart() {\n                this.fetchChartData();  // Fetch new data and re-render chart\n            }\n\t\t\t// Display progress message if current month is in progress\n       \n       \n    }};\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
