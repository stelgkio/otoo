// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func IndexTemplate() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-theme=\"light\"><head><meta charset=\"UTF-8\"><meta name=\"description\" content=\"Streamline Your E-commerce Data Integration Effortlessly.\"><meta name=\"viewport\" content=\"width=device-width,initial-scale=1,viewport-fit=cover\"><meta name=\"color-scheme\" content=\"dark light\"><title>Otoo: Streamline Your E-commerce Data Integration Effortlessly</title><link rel=\"preload\" href=\"/assets/css/DAGGERSQUARE.otf\" as=\"font\" crossorigin><link rel=\"stylesheet\" type=\"text/css\" href=\"assets/css/main.css\"><link rel=\"stylesheet\" type=\"text/css\" href=\"assets/css/utility.css\"><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.2/font/bootstrap-icons.css\"><link rel=\"stylesheet\" href=\"https://api.fontshare.com/v2/css?f=satoshi@900,700,500,300,401,400&amp;display=swap\"><link rel=\"icon\" type=\"image/png\" href=\"assets/img/favicon1.png\"></head><body class=\"p-1 p-lg-2\"><div class=\"overflow-x-hidden rounded-top-4 pt-2 pt-lg-4\"><header><div class=\"w-lg-75 mx-2 mx-lg-auto position-relative z-2 px-lg-3 py-0 shadow-5 rounded-3 rounded-lg-pill bg-dark\"><nav class=\"navbar navbar-expand-lg navbar-dark p-0\" id=\"navbar\"><div class=\"container px-sm-0\"><a class=\"navbar-brand d-inline-block w-lg-64\" href=\"/index\"><h1 class=\"display-9 mylogo text-white fw-bolder lh-tight px-sm-1\">Otoo </h1></a> <button class=\"navbar-toggler\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarCollapse\" aria-controls=\"navbarCollapse\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><span class=\"navbar-toggler-icon\"></span></button><div class=\"collapse navbar-collapse\" id=\"navbarCollapse\"><ul class=\"navbar-nav gap-2 mx-lg-auto\"><li class=\"nav-item\"><a class=\"nav-link rounded-pill\" href=\"#howtiworks\" aria-current=\"page\">How it works</a></li><li class=\"nav-item\"><a class=\"nav-link rounded-pill\" href=\"#speed\">Speed</a></li><li class=\"nav-item\"><a class=\"nav-link rounded-pill\" href=\"#custom\">Customization</a></li><li class=\"nav-item\"><a class=\"nav-link rounded-pill\" href=\"#problem\">Problem Solving</a></li><li class=\"nav-item\"><a class=\"nav-link rounded-pill\" href=\"#contact\">Contact</a></li></ul><div class=\"navbar-nav align-items-lg-center justify-content-end gap-2 ms-lg-4 w-lg-64\"><a class=\"nav-item nav-link rounded-pill d-none d-lg-block\" href=\"/login\">Sign in</a> <a href=\"/register\" class=\"btn btn-sm btn-white bg-dark-hover border-0 rounded-pill w-100 w-lg-auto mb-4 mb-lg-0\">Get started</a></div></div></div></nav></div></header><main><div class=\"pt-56 pb-10 pt-lg-56 pb-lg-10 mt-n40 position-relative gradient-bottom-right start-indigo middle-purple end-yellow\"><div class=\"container\"><div class=\"row align-items-center g-10\" style=\"padding-top: 2%; padding-bottom: 2%;\"><div class=\"col-lg-8\"><h1 class=\"ls-tight fw-bolder display-3 text-white mb-5\">All in one Solution, Integration & Dashboard Analytics Faster than Ever.</h1><p class=\"w-xl-75 lead text-white\">With Otoo you can connect any e-commerce platform quicker than ever.</p></div><div class=\"col-lg-4 align-self-end\"><div class=\"hstack gap-3 justify-content-lg-end\"><a href=\"/dashboard\" class=\"btn btn-lg btn-dark rounded-pill border-0 shadow-none px-lg-8\">Explore <span class=\"mylogo fw-bolder lh-tight px-sm-1\" style=\"color: #8957ff\">Otoo </span></a></div></div></div><div class=\"mt-10 d-none d-lg-block\"><div class=\"col-lg-8\"><p class=\"w-xl-75 lead text-white\"></p></div></div><div class=\"mt-10  d-lg-block\"><img alt=\"otoo dashboard landing page with sample data\" style=\"border-radius: 10px 10px;  width: 100%;\" src=\"./assets/img/landing2.jpeg\"></div></div></div><div id=\"howtiworks\" class=\"mt-2 py-20 pt-lg-32 bg-dark rounded-bottom-4 overflow-hidden position-relative z-1\"><div class=\"container mw-screen-xl\"><div class=\"row\"><div class=\"col-lg-6 col-md-10\"><h5 class=\"h5 mb-5 text-uppercase text-primary\">How it works</h5><h1 class=\"display-4 font-display text-white fw-bolder lh-tight mb-4\">It's time to take action</h1><p class=\"text-lg text-white text-opacity-75\">With Otoo, you can effortlessly connect your Shopify or WooCommerce store with any ERP system. Our intuitive dashboard provides real-time analytics and insights, giving you full control over your business operations.</p></div></div><div class=\"row g-6 g-lg-20 my-10\"><div class=\"col-md-4\"><div class=\"card shadow-none border-0\"><div class=\"card-body p-7\"><div class=\"mt-4 mb-7\"><div class=\"icon icon-shape text-white bg-primary rounded-circle text-lg\"><i class=\"bi bi-regex\"></i></div></div><div class=\"pt-2 pb-3\"><h5 class=\"h3 font-display fw-bold mb-3\">E-commerce Integration</h5><p class=\"text-muted\">Unlock the full potential of your online store with seamless integration solutions. Whether you're using Shopify, WooCommerce, or another platform, we streamline your e-commerce operations for optimal performance and customer experience.</p></div></div></div></div><div class=\"col-md-4\"><div class=\"card shadow-none border-0\"><div class=\"card-body p-7\"><div class=\"mt-4 mb-7\"><div class=\"icon icon-shape text-white bg-primary rounded-circle text-lg\"><i class=\"bi bi-person-workspace\"></i></div></div><div class=\"pt-2 pb-3\"><h5 class=\"h3 font-display fw-bold mb-3\">AI Analysis</h5><p class=\"text-muted\">Turn data into insights and opportunities with our advanced AI solutions. From data collection to analysis and visualization, we help you harness the power of data to make informed decisions and drive business growth. </p></div></div></div></div><div class=\"col-md-4\"><div class=\"card shadow-none border-0\"><div class=\"card-body p-7\"><div class=\"mt-4 mb-7\"><div class=\"icon icon-shape text-white bg-primary rounded-circle text-lg\"><i class=\"bi bi-stars\"></i></div></div><div class=\"pt-2 pb-3\"><h5 class=\"h3 font-display fw-bold mb-3\">Cloud Infrastructure</h5><p class=\"text-muted\">Scale your business securely and efficiently with our cloud infrastructure services. Whether you're migrating to the cloud or optimizing your existing setup, we provide reliable solutions tailored to your requirements, ensuring flexibility, scalability, and cost-effectiveness. .</p></div></div></div></div></div></div></div><div class=\"py-20 pt-lg-32 pb-lg-20\"><div class=\"container mw-screen-xl\"><div class=\"row justify-content-lg-end mb-10 mb-lg-24\"><div class=\"col-md-6\"><h1 class=\"font-display lh-tight fw-bolder display-5 mb-5\">Why Devs and startups are already using <a class=\"mylogo\" href=\"/index\">otoo</a></h1></div></div><div id=\"speed\" class=\"section-step-lg\"><div class=\"row justify-content-between align-items-center\"><div class=\"col-lg-5 mb-7 mb-lg-0\"><h5 class=\"h5 mb-5 text-uppercase fw-bolder text-primary\">Speed</h5><h1 class=\"ls-tight font-display fw-bolder mb-5\">Start Collaboration in hours, not weeks</h1><p class=\"lead\"><b>Faster Collaboration:</b> Say goodbye to lengthy email threads and hello to real-time collaboration. With everyone on the same page, projects move forward faster than ever before.. Clients can actively participate in the project, providing valuable insights and feedback that drive innovation and ensure the final product meets their needs and expectations.</p><ul class=\"list-unstyled mt-6 mb-0\"><li class=\"py-2\"><div class=\"d-flex align-items-center\"><div><div class=\"icon icon-xs icon-shape bg-success text-white text-base rounded-circle me-3\"><i class=\"bi bi-check\"></i></div></div><div><span class=\"fw-semibold\">Quickly start collaboration</span></div></div></li><li class=\"py-2\"><div class=\"d-flex align-items-center\"><div><div class=\"icon icon-xs icon-shape bg-success text-white text-base rounded-circle me-3\"><i class=\"bi bi-check\"></i></div></div><div><span class=\"fw-semibold\">Real-Time Feedback </span></div></div></li><li class=\"py-2\"><div class=\"d-flex align-items-center\"><div><div class=\"icon icon-xs icon-shape bg-success text-white text-base rounded-circle me-3\"><i class=\"bi bi-check\"></i></div></div><div><span class=\"fw-semibold\">All in one place</span></div></div></li></ul></div><div class=\"col-lg-6\"><div class=\"bg-primary overflow-hidden rounded-4 ps-8 pt-8\"><img alt=\"otoo trello exampple project\" src=\"./assets/img/marketing/trello2.png\" class=\"img-fluid\" alt=\"...\"></div></div></div></div><div id=\"custom\" class=\"section-step-lg\"><div class=\"row justify-content-between align-items-center\"><div class=\"col-lg-5 mb-7 mb-lg-0\"><h5 class=\"h5 mb-5 text-uppercase text-secondary fw-bolder\">Customization</h5><h1 class=\"font-display ls-tight fw-bolder mb-5\">We focus on making everything customizable</h1><p class=\"lead\">we prioritize customization in everything we do. We understand that every business is unique, with its own set of challenges and requirements. That's why our solutions are designed to be highly customizable, allowing you to tailor them to your specific needs and preferences. Whether it's adjusting the layout of your dashboard, fine-tuning data visualization settings, or integrating custom features, our platform puts you in full control. With a wide range of options and flexibility at your fingertips, you can create a solution that perfectly aligns with your business goals and enhances your workflow.</p><ul class=\"list-unstyled mt-6 mb-0\"><li class=\"py-2\"><div class=\"d-flex align-items-center\"><div><div class=\"icon icon-xs icon-shape bg-success text-white text-base rounded-circle me-3\"><i class=\"bi bi-check\"></i></div></div><div><span class=\"fw-semibold\">Highly Customizable Solutions</span></div></div></li><li class=\"py-2\"><div class=\"d-flex align-items-center\"><div><div class=\"icon icon-xs icon-shape bg-success text-white text-base rounded-circle me-3\"><i class=\"bi bi-check\"></i></div></div><div><span class=\"fw-semibold\">Tailored to Specific Needs</span></div></div></li><li class=\"py-2\"><div class=\"d-flex align-items-center\"><div><div class=\"icon icon-xs icon-shape bg-success text-white text-base rounded-circle me-3\"><i class=\"bi bi-check\"></i></div></div><div><span class=\"fw-semibold\">Empowering Control</span></div></div></li></ul></div><div class=\"col-lg-6\"><div class=\"bg-secondary overflow-hidden rounded-4 ps-8 pt-8\"><img alt=\"example of product page with staticts\" src=\"./assets/img/product.jpeg\" class=\"img-fluid\" alt=\"...\"></div></div></div></div></div></div><div class=\"py-20 py-lg-20\"><div class=\"container mw-screen-xl\"><div class=\"py-32 gradient-bottom-right start-gray middle-black end-gray rounded-5 px-lg-16 text-center text-md-start\"><div class=\"row justify-content-center\"><div class=\"col-12 col-md-10 col-lg-8 text-center\"><h1 class=\"ls-tight fw-bolder display-4 mb-5 text-white\">Ready to get started?</h1><p class=\"lead text-white opacity-8 mb-10\">Save time and money while getting more productive than ever before. Kickstart your development process now.</p><div class=\"mx-n2\"><a href=\"/dashboard\" class=\"btn btn-lg btn-white mx-2 px-lg-8\">Get started</a></div></div></div></div></div></div><div id=\"problem\" class=\"py-20 py-lg-20\"><div class=\"container mw-screen-xl\"><div class=\"row align-items-center mb-20\"><div class=\"col-lg-6 mb-4 mb-lg-0\"><h1 class=\"display-5 ls-tight mb-5\">Empowering Your Development Journey</h1><div class=\"row mx-n2 mt-12\"><div class=\"col-sm-4 mb-3 mb-sm-0\"><h2 class=\"text-secondary mb-1\"><span class=\"display-6 fw-semibold\">100+</span><!-- <span class=\"counter-extra\">+</span> --></h2><p class=\"text-muted\">Happy clients</p></div></div></div><div class=\"col-lg-5 ms-lg-auto\"><div class=\"vstack gap-8\"><div><div class=\"d-flex align-items-center gap-4 mb-4\"><div class=\"icon icon-shape text-bg-primary text-lg rounded-circle\"><i class=\"bi bi-door-open\"></i></div><h3 class=\"fw-semibold\">Efficiency and Time Savings</h3></div></div><hr class=\"my-0\"><div><div class=\"d-flex align-items-center gap-4 mb-4\"><div class=\"icon icon-shape text-bg-primary text-lg rounded-circle\"><i class=\"bi bi-rocket-takeoff\"></i></div><h3 class=\"fw-semibold\">Professional experience</h3></div></div></div></div></div><div class=\"row g-8 justify-content-center\"><div class=\"col-md-6 col-xl-4\"><div class=\"card shadow-soft-2 h-md-100\"><div class=\"card-body d-flex flex-column p-8\"><h5 class=\"h4 mb-4\">Problem Statement</h5><p class=\"text-sm\">Are you tired of spending hours manually transferring data between your e-commerce store and your ERP system? Do you find it challenging to keep your inventory, orders, and customer information synchronized across platforms?</p></div></div></div><div class=\"col-md-6 col-xl-4\"><div class=\"card shadow-soft-2 h-md-100\"><div class=\"card-body d-flex flex-column p-8\"><h5 class=\"h4 mb-4\">Manual Data Entry Overload</h5><p class=\"text-sm\">Businesses often struggle with the tedious and error-prone task of manually entering data from their e-commerce platforms into their ERP systems. Otoo automates this process, saving time and reducing the risk of inaccuracies.</p></div></div></div><div class=\"col-md-6 col-xl-4\"><div class=\"card shadow-soft-2 h-md-100\"><div class=\"card-body d-flex flex-column p-8\"><h5 class=\"h4 mb-4\">Lack of Real-Time Insights</h5><p class=\"text-sm\">Without access to real-time analytics, businesses may struggle to make informed decisions about inventory management, marketing strategies, and customer engagement. Otoo's dashboard provides actionable insights into key metrics, empowering businesses to optimize their operations and drive growth.</p></div></div></div><div class=\"col-md-6 col-xl-4\"><div class=\"card shadow-soft-2 h-md-100\"><div class=\"card-body d-flex flex-column p-8\"><h5 class=\"h4 mb-4\">Complex Integration Processes</h5><p class=\"text-sm\">Integrating e-commerce platforms with ERP systems can be complex and time-consuming, requiring technical expertise and custom development. Otoo simplifies the integration process with its user-friendly interface and pre-built connectors, allowing businesses to quickly set up and start benefiting from automated data transfer.</p></div></div></div><div class=\"col-md-6 col-xl-4\"><div class=\"card shadow-soft-2 h-md-100\"><div class=\"card-body d-flex flex-column p-8\"><h5 class=\"h4 mb-4\">Data Security Concernss</h5><p class=\"text-sm\">Manually transferring sensitive customer and financial data between systems can pose security risks, especially if done through insecure channels. Otoo prioritizes data security, employing encryption protocols and secure authentication methods to protect sensitive information throughout the integration process.</p></div></div></div><div class=\"col-md-6 col-xl-4\"><div class=\"card shadow-soft-2 h-md-100\"><div class=\"card-body d-flex flex-column p-8\"><h5 class=\"h4 mb-4\">Scalability Challenges</h5><p class=\"text-sm\">As e-commerce businesses grow and expand, their data integration needs become more complex and demanding. Otoo is designed to scale with businesses of all sizes, supporting high volumes of data and adapting to evolving requirements without sacrificing performance or reliability..</p></div></div></div></div></div></div></main><footer id=\"contact\" class=\"pt-24 pb-10\"><div class=\"container mw-screen-xl\"><div class=\"row\"><div class=\"col\"><div class=\"pe-5\"><h3 class=\"h2 text-heading fw-semibold lh-lg mb-3\">Let's talk about your project api integration, big data, streaming and much more at</h3><a href=\"mailto:hello@otoo.com\" class=\"h3 text-primary\">hello@otoo.com  <span class=\"svg-icon svg-align-baseline ms-3\"><i class=\"bi bi-arrow-right\"></i></span></a></div><div class=\"my-7\"><a href=\"/contact\" class=\"btn btn-dark\">Get In Touch</a></div></div></div><div class=\"row mt-5 mb-7\"><div class=\"col\"><ul class=\"nav mx-n4\"><!-- <li class=\"nav-item\"><a href=\"#\" class=\"nav-link text-lg text-muted text-primary-hover\"><i\n                                        class=\"bi bi-twitter\"></i></a></li>\n                            <li class=\"nav-item\"><a href=\"#\" class=\"nav-link text-lg text-muted text-primary-hover\"><i\n                                        class=\"bi bi-dribbble\"></i></a></li>\n                            <li class=\"nav-item\"><a href=\"#\" class=\"nav-link text-lg text-muted text-primary-hover\"><i\n                                        class=\"bi bi-github\"></i></a></li>\n                            <li class=\"nav-item\"><a href=\"#\" class=\"nav-link text-lg text-muted text-primary-hover\"><i\n                                        class=\"bi bi-youtube\"></i></a></li> --></ul></div></div><div class=\"row\"><div class=\"col-auto\"><p class=\"text-sm text-muted\">&copy; Copyright 2023 CloudDigital - Build like a PRO.</p></div></div></div></footer></div><script src=\"https://cdn.jsdelivr.net/npm/choices.js/public/assets/scripts/choices.min.js\"></script><script src=\"./assets/js/main.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
