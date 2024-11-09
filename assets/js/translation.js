// Initialize i18next
i18next
    // .use(i18nextHttpBackend) // Optional: Load translations from external files
    .use(i18nextBrowserLanguageDetector) // Detect user's language
    .init({
        fallbackLng: 'en', // Default language if detection fails
        debug: true, // Enable debug logs in the console

        resources: {
            en: {
                translation: {
                    // Index EN
                    "title": "All in one Solution,Integration & Dashboard Analytics	Faster than Ever.!",
                    "Dashboard": "Dashboard",
                    // Off Canvas EN
                    "offcanvas-order_processing": "Create Voucher#",
                    "off-canvas-customer-tab": "Customer",
                    "off-canvas-shipping-tab": "Shipping",
                    "off-canvas-customer-info-header": "Customer Info",
                    "off-canvas-customer-first_name": "First Name",
                    "off-canvas-customer-last_name": "Last Name",
                    "off-canvas-customer-email": "Email",
                    "off-canvas-customer-phone": "Phone",
                    "off-canvas-customer-address": "Address",
                    "off-canvas-customer-city": "City",
                    "off-canvas-customer-postal": "Postal",
                    //Off Canvas Products Info EN
                    "off-canvas-customer-products-header": "Products",
                    "off-canvas-customer-product-name": "Product Name",
                    "off-canvas-customer-product-quantity": "Product Quantity",
                    "off-canvas-customer-product-price": "Product Price",
                    //Off Canvas Delivery Info EN
                    "off-canvas-delivery-info": "Delivery Info",
                    "off-canvas-special-instructions": "Special Instructions",
                    "off-canvas-customer-store-pickup": "Store pickup",
                    // Off Canvas Shipping Info EN
                    "off-canvas-shipping-information": "Shipping Information",
                    "off-canvas-shipping-order-number": "Order Number",
                    "off-canvas-shipping-parcel-weight": "Parcel Weight",
                    "off-canvas-shipping-company": "Shipping Company",
                    "off-canvas-shipping-provider": "Select Courier Provider",
                    "off-canvas-parcel-cod": "COD",
                    //Off Canvas Billing Info EN
                    "off-canvas-modal-name": "Name",
                    "off-canvas-modal-last-name": "Last Name",
                    "off-canvas-email": "Email",
                    "off-canvas-modal-phone": "Phone",
                    "off-canvas-modal-address": "Address",
                    "off-canvas-modal-city": "City",
                    "off-canvas-modal-postal": "Postal",
                    //Off Canvas ACS EN
                    "off-canvas-delivery-options-acs": "Delivery Options (ACS)",
                    "off-canvas-acs-standard-delivery": "Standard Delivery",
                    "off-canvas-acs-saturday-delivery": "Saturday Delivery",
                    "off-canvas-acs-urgent-delivery": "Urgent Delivery",
                    "off-canvas-acs-pickup-delivery": "Pickup Delivery",
                    //Off Canvas Courier4U EN
                    "off-canvas-delivery-options-courier4u": "Delivery Options (Courier4U)",
                    "off-canvas-delivery-courier4u-standard-delivery": "Standard Delivery",
                    "off-canvas-delivery-courier4u-saturday-delivery": "Saturday Delivery",
                    "off-canvas-delivery-courier4u-urgent-delivery": "Urgent Delivery",
                    "off-canvas-delivery-courier4u-pickup-delivery": "Pickup Delivery",
                    "off-canvas-delivery-courier4u-sameday-delivery": "Same Day Delivery",
                    "off-canvas-delivery-courier4u-return-delivery": "Return Delivery",

                    "off-canvas-delivery-courier4u-ServiceProtocol-delivery": "Service Protocol Number",
                    //Off Canvas Parsel Weight/Height/Width Options EN
                    "off-canvas-shipping-parcel-courier4u-weight": "Package Weight",
                    "off-canvas-wh-options-courier4u": "Package Options",
                    "off-canvas-shipping-parcel-depth": "Parcel Depth",
                    "off-canvas-shipping-parcel-height": "Parcel Height",
                    "off-canvas-shipping-parcel-width": "Parcel Width",
                    "offcanvas-nav-customer": "Customer",
                    "offcanvas-nav-shipping": "Shipping",
                    //Off Canvas button
                    "off-canvas-save-btn": "Save Changes",
                    "off-canvas-close-btn": "Cancel",
                    //Voucher Header Table EN
                    "voucher-table-new-order": "New",
                    "voucher-table-processing-orders": "Processing",
                    "voucher-table-completed-orders": "Completed",
                    "voucher-table-cancelled-orders": "Cancelled",
                    "voucher-table-all-orders": "All",
                    "voucher-table-export-voucher-button": "Export Voucher",
                    "voucher-table-new-voucher-button": "New Voucher",
                    // Order Header Table EN
                    "order-table-all-orders": "All",
                    "order-table-completed-orders": "Completed",
                    "order-table-pending-orders": "Pending",
                    "order-table-new-orders": "New",
                    "order-table-processing-orders": "Processing",
                    "order-table-cancelled-orders": "Cancelled",
                    // Order Modal Billing EN
                    "order-modal-header-title": "Order Details",
                    "order-modal-header-billing": "Billing Information",
                    "order-modal-billing-name": "First Name",
                    "order-modal-billing-address": "Address",
                    "order-modal-billing-city": "City",
                    "order-modal-billing-postal": "Postal",
                    "order-modal-billing-email": "E-mail",
                    "order-modal-billing-phone": "Phone",
                    // Order Modal Shipping EN
                    "order-modal-header-shipping": "Shipping Information",
                    "order-modal-shipping-name": "First Name",
                    "order-modal-shipping-last-name": "Last Name",
                    "order-modal-shipping-address": "Address",
                    "order-modal-shipping-address_2": "Address 2",
                    "order-modal-shipping-city": "City",
                    "order-modal-shipping-postal": "Postal",
                    "order-modal-shipping-customer-note": "Customer Note",
                    "off-canvas-shipping-parclel-cod": "Cash on Delivery",
                    // Sorting Table EN
                    "sorting-table-voucher-id": "Voucher ID",
                    "sorting-table-order-id": "Order ID",
                    "sorting-table-order-date": "Order Date",
                    "sorting-table-order-cod": "COD",
                    "sorting-table-status-header": "Status",
                    "sorting-table-printed:": "Printed",
                    "sorting-table-action-header": "Action",
                    "sorting-no-results-msg": "No results found",
                    "sorting-create-voucher": "Create Voucher",
                    // Side nav
                    "side-nav-logout": "Logout",
                    "side-nav-projects": "Projects",
                    "side-nav-dashboard": "Dashboard",
                    "side-nav-default": "Default",
                    "side-nav-customers": "Customers",
                    "side-nav-products": "Products",
                    "side-nav-orders": "Orders",
                    "side-nav-resources": "Resources",
                    "side-nav-extensions": "Extensions",
                    "side-nav-add-extensions": "Add Extensions",
                    "side-nav-settings": "Settings",




                }
            },
            el: {
                translation: {
                    // Index GR
                    "title": "Όλα σε μία Λύση, Ενσωμάτωση & Αναλυτικά Δεδομένα Πίνακα Ελέγχου Πιο Γρήγορα από Ποτέ.",
                    "Dashboard": "Πίνακας Ελέγχου",
                    //Off Canvas GR
                    "offcanvas-order_processing": "Δημιουργία Voucher#",
                    "off-canvas-customer-tab": "Πελάτης",
                    "off-canvas-shipping-tab": "Αποστολή",
                    "off-canvas-customer-info-header": "Στοιχεία Πελάτη",
                    "off-canvas-modal-name": "Όνομα",
                    "off-canvas-modal-last-name": "Επώνυμο",
                    "off-canvas-email": "Email",
                    "off-canvas-modal-phone": "Τηλέφωνο",
                    "ooff-canvas-modal-address": "Διεύθυνση",
                    "off-canvas-modal-city": "Πόλη",
                    "off-canvas-modal-postal": "Τ.K.",
                    //Off Canvas Products Info
                    "off-canvas-customer-products-header": "Προϊόντα",
                    "off-canvas-customer-product-name": "Όνομα Προϊόντος",
                    "off-canvas-customer-product-quantity": "Ποσότητα Προϊόντος",
                    "off-canvas-customer-product-price": "Τιμή Προϊόντος",
                    //Off Canvas Delivery Info
                    "off-canvas-delivery-info": "Πληροφορίες Αποστολής",
                    "off-canvas-special-instructions": "Σημειώσεις Παράδοσης",
                    "off-canvas-customer-store-pickup": "Παραλαβή Απο Κατάστημα",
                    //Off Canvas Shipping Info
                    "off-canvas-shipping-information": "Στοιχεία Αποστολής",
                    "off-canvas-shipping-order-number": "Αριθμός Παραγγελίας",
                    "off-canvas-shipping-company": "Εταιρεία Αποστολής",
                    "off-canvas-shipping-provider": "Επιλογή Ταχυδρομικής Εταιρείας",
                    "off-canvas-parcel-cod": "Αντικαταβολή",
                    //Off Canvas ACS
                    "off-canvas-delivery-options-acs": "Επιλογές Αποστολής (ACS)",
                    "off-canvas-acs-standard-delivery": "Κανονική Αποστολή",
                    "off-canvas-acs-saturday-delivery": "Αποστολή Σάββατο",
                    "off-canvas-acs-urgent-delivery": "Επείγον Αποστολή",
                    "off-canvas-acs-pickup-delivery": "Παραλαβή Απο Κατάστημα",
                    //Off Canvas Courier4U
                    "off-canvas-delivery-options-courier4u": "Επιλογές Αποστολής (Courier4U)",
                    "off-canvas-delivery-courier4u-standard-delivery": "Κανονική Αποστολή",
                    "off-canvas-delivery-courier4u-saturday-delivery": "Αποστολή Σάββατο",
                    "off-canvas-delivery-courier4u-urgent-delivery": "Επείγον Αποστολή",
                    "off-canvas-delivery-courier4u-pickup-delivery": "Παραλαβή Απο Κατάστημα",
                    "off-canvas-delivery-courier4u-sameday-delivery": "Παράδοση εντός ημέρας",
                    "off-canvas-delivery-courier4u-ServiceProtocol-delivery": "Πρωτοκολλο Παράδοσης",
                    "off-canvas-delivery-courier4u-return-delivery": "Επιστροφή Προϊόντος",
                    //Off Canvas Parcel Weight/Height/Width Options GR
                    "off-canvas-wh-options-courier4u": "Επιλογές Δέματος",
                    "off-canvas-shipping-parcel-courier4u-weight": "Βάρος Δέματος",
                    "off-canvas-shipping-parcel-depth": "Βάθος Δέματος",
                    "off-canvas-shipping-parcel-height": "Ύψος Δέματος",
                    "off-canvas-shipping-parcel-width": "Πλάτος Δέματος",
                    "offcanvas-nav-customer": "Πελάτης",
                    "offcanvas-nav-shipping": "Αποστολή",
                    // Button  
                    "off-canvas-save-btn": "Αποθήκευση",
                    "off-canvas-close-btn": "Ακύρωση",
                    //Voucher Table
                    "voucher-table-new-order": "Νέες",
                    "voucher-table-processing-orders": "Σε Επεξεργασία",
                    "voucher-table-completed-orders": "Ολοκληρωμένες",
                    "voucher-table-cancelled-orders": "Ακυρωμένες",
                    "voucher-table-all-orders": "Όλες",
                    "voucher-table-export-voucher-button": "Εξαγωγή Voucher",
                    "voucher-table-new-voucher-button": "Νέο Voucher",
                    /// Order Header Table
                    "order-table-all-orders": "Όλες",
                    "order-table-completed-orders": "Ολοκληρωμένες",
                    "order-table-pending-orders": "Σε Εκκρεμότητα",
                    "order-table-processing-orders": "Σε Επεξεργασία",
                    "order-table-new-orders": "Νέες",
                    "order-table-cancelled-orders": "Ακυρωμένες",
                    // Order Modal Shipping GR
                    "order-modal-header-title": "Πληροφορίες Παραγγελίας",
                    "order-modal-header-billing": "Πληροφορίες Πληρωμής",
                    "order-modal-billing-name": "Όνομα",
                    "order-modal-billing-last-name": "Επώνυμο",
                    "order-modal-billing-address": "Διεύθυνση",
                    "order-modal-billing-city": "Πόλη",
                    "order-modal-billing-postal": "Τ.Κ.",
                    "order-modal-billing-email": "E-mail",
                    "order-modal-billing-phone": "Τηλέφωνο",
                    // Order Modal Shipping EN
                    "order-modal-header-shipping": "Πληροφορίες Αποστολής",
                    "order-modal-shipping-name": "Όνομα",
                    "order-modal-shipping-last-name": "Επώνυμο",
                    "order-modal-shipping-address": "Διεύθυνση",
                    "order-modal-shipping-address_2": "Διεύθυνση 2",
                    "order-modal-shipping-city": "Πόλη",
                    "order-modal-shipping-postal": "Τ.Κ.",
                    "order-modal-shipping-customer-note": "Σημειώσεις Πελάτη",
                    "off-canvas-shipping-parclel-cod": "Αντικαταβολή",
                    // Sorting Table GR
                    "sorting-table-voucher-id": "VOUCHER ID",
                    "sorting-table-order-id": "ID ΠΑΡΑΓΓΕΛΙΑΣ",
                    "sorting-table-order-date": "ΗΜ/ΝΙΑ",
                    "sorting-table-order-cod": "ΑΝΤΙΚΑΤΑΒΟΛΗ",
                    "sorting-table-status-header": "ΚΑΤΑΣΤΑΣΗ",
                    "sorting-table-printed": "ΕΚΤΥΠΩΜΕΝΟ",
                    "sorting-table-action-header": "ΕΝΕΡΓΕΙΕΣ",
                    "sorting-no-results-msg": "Δεν υπάρχουν αποτελέσματα",
                    "sorting-create-voucher": "Δημιουργία Voucher",
                    // Side nav
                    "side-nav-logout": "Αποσύνδεση",
                    "side-nav-projects": "Έργα",
                    "side-nav-dashboard": "Πίνακας Ελέγχου",
                    "side-nav-default": "Προεπιλογή",
                    "side-nav-customers": "Πελάτες",
                    "side-nav-products": "Προϊόντα",
                    "side-nav-orders": "Παραγγελίες",
                    "side-nav-resources": "Πόροι",
                    "side-nav-extensions": "Επεκτάσεις",
                    "side-nav-add-extensions": "Προσθήκη Επεκτάσεων",
                    "side-nav-settings": "Ρυθμίσεις",





                }
            }
        }
    }, function (err, t) {
        // Initialize the content with the current language
        updateContent();
    });

// Function to change language
function changeLanguage(lng) {
    i18next.changeLanguage(lng, updateContent);
}

// Function to update the content  / support placeholders etc.
function updateContent() {
    document.querySelectorAll('[data-i18n]').forEach(function (element) {
        var dataI18n = element.getAttribute('data-i18n');

        if (dataI18n.includes('[')) {
            // This is an attribute translation
            var parts = dataI18n.match(/\[([^\]]+)\](.+)/);
            var attribute = parts[1]; // e.g., 'placeholder'
            var key = parts[2]; // e.g., 'off-canvas-special-instructions'
            element.setAttribute(attribute, i18next.t(key.trim()));
        } else {
            // This is a text content translation
            var key = dataI18n;
            element.textContent = i18next.t(key);
        }
    });
}


// Listen to HTMX event when content is dynamically updated
document.body.addEventListener('htmx:afterSettle', function (evt) {
    // Update i18n content in the dynamically updated section
    updateContent();
});
