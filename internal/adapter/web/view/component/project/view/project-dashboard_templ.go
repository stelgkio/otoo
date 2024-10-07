// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	l "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/list"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/component/project/template"
	"github.com/stelgkio/otoo/internal/core/domain"
)

func ProjectDashboard(projects []*domain.Project, user *domain.User) templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Err = l.ProjectListPage(projects).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <script>\n\t    var chartElement2= document.querySelector(\"#chart-users\");   \n        var options2= {\n            chart: {\n                type: \"bar\",\n                stacked: true,\n                zoom: {\n                    enabled: false\n                },\n                toolbar: {\n                    show: false\n                },\n                shadow: {\n                    enabled: false\n                },\n                offsetX: 0,\n                animations: {\n                    enabled: false,\n                    easing: \"easeinout\",\n                    speed: 800,\n                    animateGradually: {\n                        enabled: true,\n                        delay: 150\n                    },\n                    dynamicAnimation: {\n                        enabled: true,\n                        speed: 350\n                    }\n                },\n                fontFamily: \"#333\",\n\t\t\t\theight: 390,\n            },\n            colors: [ '#8957ff', '#ffc107', '#dc3545'],\n            plotOptions: {\n                bar: {\n                    columnWidth: \"23px\",\n                    borderRadius: 2\n                }\n            },\n            stroke: {\n                width: 4,\n                curve: \"smooth\"\n            },\n            series: [{\n                name: \"Revenue\",\n                data: [30, 10, 20, 10, 17, 12, 8, 20]\n            }],\n            markers: {\n                size: 0\n            },\n            xaxis: {\n\t\t\t\t  axisBorder: {\n                    show: false\n                },\n                axisTicks: {\n                    show: false\n                },\n                categories: [\"May\", \"Jun\", \"Jul\", \"Aug\", \"Sep\", \"Oct\", \"Nov\", \"Dec\"],\n                labels: {\n                    style: {\n                         colors: \"#e3dede\",\n                        fontSize: \"13px\"\n                    }\n                }\n            },\n            yaxis: {\n                labels: {\n                    style: {\n                        colors: \"#e3dede\",\n                        fontSize: \"13px\"\n                    }\n                }\n            },\n            legend: {\n                show: false\n            },\n            grid: {\n                borderColor: \"#999\",\n                strokeDashArray: 3\n            },\n            dataLabels: {\n                enabled: false\n            },\n            tooltip: {\n                shared: true,\n                intersect: false,\n                y: {\n                    formatter: function (val) {\n                        return val.toFixed(0) + \" orders\";\n                    }\n                }\n            },\n            responsive: [{\n                breakpoint: 364,\n                options: {\n                    plotOptions: {\n                        bar: {\n                            columnWidth: \"30px\",\n                            borderRadius: 2\n                        }\n                    }\n                }\n            }]\n\t\t}\n        \n\n     \n\n       var barChart = new ApexCharts(chartElement2, options2);\n        barChart.render();\n\n \n\n\n\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = t.ProjectTemplate(user).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
