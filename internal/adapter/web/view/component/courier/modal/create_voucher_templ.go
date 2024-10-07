// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	v "github.com/stelgkio/otoo/internal/core/domain/courier"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
)

func CreateVoucher(projectId string, voucher *v.Voucher, order *w.OrderRecord) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Offcanvas for Courier Information --><div class=\"offcanvas-header\"><h5 class=\"offcanvas-title\" id=\"orderProcessingOffcanvasLabel\">Επεξεργασία Παραγγελίας #<span id=\"offcanvas-order-id\"></span></h5><button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"offcanvas\" aria-label=\"Close\"></button></div><div class=\"offcanvas-body\"><div class=\"accordion\" id=\"orderProcessingAccordion\"><!-- Shipping Company Section --><div class=\"accordion-item\"><h2 class=\"accordion-header\"><button class=\"accordion-button\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#shippingCompanyCollapse\">Εταιρεία Αποστολής</button></h2><div id=\"shippingCompanyCollapse\" class=\"accordion-collapse collapse show\" data-bs-parent=\"#orderProcessingAccordion\"><div class=\"accordion-body\"><select class=\"form-select mb-3\" id=\"shippingCompany\"><option value=\"\">Επιλέξτε εταιρεία...</option> <option value=\"acs\">ACS Courier</option> <option value=\"dhl\">DHL</option> <option value=\"fedex\">FedEx</option></select> <select class=\"form-select\" id=\"shippingMethod\"><option value=\"\">Επιλέξτε μέθοδο...</option> <option value=\"standard\">Standard</option> <option value=\"express\">Επείγουσα Παράδοση</option> <option value=\"express\">Παράδοση - Παραλαβή</option> <option value=\"express\">Αυθημερόν</option> <option value=\"express\">Παραλαβή από το κατάστημα</option></select></div></div></div><!-- Order Info Section --><div class=\"accordion-item\"><h2 class=\"accordion-header\"><button class=\"accordion-button collapsed\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#orderInfoCollapse\">Πληροφορίες Παραγγελίας</button></h2><div id=\"orderInfoCollapse\" class=\"accordion-collapse collapse\" data-bs-parent=\"#orderProcessingAccordion\"><div class=\"accordion-body\"><div class=\"mb-3\"><label class=\"form-label\">Ποσό Αντικαταβολής</label><div class=\"input-group\"><span class=\"input-group-text\">€</span> <input type=\"number\" class=\"form-control\" id=\"codAmount\" value=\"72.50\"></div></div><div class=\"mb-3\"><label class=\"form-label\">Βάρος Δέματος (kg)</label> <input type=\"number\" class=\"form-control\" id=\"packageWeight\" step=\"0.1\" value=\"0.5\"></div></div></div></div><!-- Special Instructions Section --><div class=\"accordion-item\"><h2 class=\"accordion-header\"><button class=\"accordion-button collapsed\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#specialInstructionsCollapse\">Ειδικές Οδηγίες</button></h2><div id=\"specialInstructionsCollapse\" class=\"accordion-collapse collapse\" data-bs-parent=\"#orderProcessingAccordion\"><div class=\"accordion-body\"><textarea class=\"form-control\" id=\"specialInstructions\" rows=\"3\" placeholder=\"Προσθέστε ειδικές οδηγίες προς την εταιρεία αποστολής...\"></textarea></div></div></div><!-- Special Instructions Section --><div class=\"accordion-item\"><h2 class=\"accordion-header\"><button class=\"accordion-button collapsed\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#recieptInstructionsCollapse\">Εντολή Παράδοσης</button></h2><div id=\"recieptInstructionsCollapse\" class=\"accordion-collapse collapse\" data-bs-parent=\"#orderProcessingAccordion\"><div class=\"accordion-body\"><select class=\"form-select\" id=\"recieptMethod\"><option value=\"\">Επιλογή Οχήματος</option> <option value=\"standard\">Μηχανή</option> <option value=\"express\">Αυτοκίνητο</option></select></div><div class=\"accordion-body\"><textarea class=\"form-control\" id=\"recieptInstructions\" rows=\"2\" placeholder=\"Γράψτε οδηγίες για τον διανομέα πχ όροφος ή το τι θα παραλάβουμε..\"></textarea></div></div></div></div></div><div class=\"offcanvas-footer p-3\"><button type=\"button\" class=\"btn btn-secondary\" data-bs-dismiss=\"offcanvas\">Ακύρωση</button> <button type=\"button\" class=\"btn btn-primary\" onclick=\"sendToShipping()\">Αποστολή στην Εταιρεία</button></div><!-- Toast Container --><div class=\"toast-container\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
