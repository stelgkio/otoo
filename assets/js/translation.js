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
                    "subtitle": "With Otoo you can connect any e-commerce platform quicker than ever.",
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
                    "off-canvas-modal-desc": "Description",
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
                    // "off-canvas-delivery-courier4u-standard-delivery": "Standard Delivery",
                    "off-canvas-delivery-courier4u-saturday-delivery": "Delivery on Saturday",
                    "off-canvas-delivery-courier4u-urgent-delivery": "Urgent Delivery",
                    "off-canvas-delivery-courier4u-pickup-delivery": "Pickup from the store",
                    "off-canvas-delivery-courier4u-sameday-delivery": "Same-Day",
                    "off-canvas-delivery-courier4u-return-delivery": "Delivery - Pickup",
                    "off-canvas-delivery-courier4u-ServiceProtocol-delivery": "Protocol Number",
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
                    //Voucher Header Section Table EN
                    "voucher-table-new-order": "New",
                    "voucher-table-processing-orders": "Processing",
                    "voucher-table-completed-orders": "Completed",
                    "voucher-table-cancelled-orders": "Cancelled",
                    "voucher-table-all-orders": "All",
                    "voucher-table-export-voucher-button": "Export Voucher",
                    "voucher-table-new-voucher-button": "New Voucher",
                    // Order Header Section Table EN
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
                    // Dashboard overview EN ( Shared with Order/Customer/Product Overview )
                    "dashboard-overview_orders": "Orders",
                    "dashboard-overview_products": "Products",
                    "dashboard-overview_customers": "Customers",
                    "dashboard-overview-total_orders": "Total Orders",
                    "dashboard-overview_orders": "Orders",
                    "dashboard-overview-week_orders": "Orders",
                    "dashboard-overview-month_orders": "Orders",
                    // Dashboard Shared Tables EN
                    "dashboard-table-order_id": "Order ID",
                    "dashboard-table-order-created": "Created",
                    "dashboard-table-order_date": "Order Date",
                    "dashboard-table-order_status": "Order Status",
                    "dashboard-table-order_total": "Order Total",
                    "dashboard-table-order_action": "Action",
                    "dashboard-table-customer_name": "Name",
                    "dashboard-table-customer_email": "Email",
                    "dashboard-table-customer-total_orders": "Total Orders",
                    "dashboard-table-customer-total_spent": "Total Spent",
                    "dashboard-table-customer_id": "product ID",
                    "dashboard-table-customer_email": "Product Name",
                    "dashboard-table-customer-total_spent": "Πελάτης δεν βρέθηκε",
                    "dashboard-table-products_id": "product ID",
                    "dashboard-table-products_name": "Product Name",
                    "dashboard-table-products_price": "Pricing",
                    "dashboard-table-products_category": "Category",
                    "dashboard-table-products_type": "Type",
                    "dashboard-table-products-total_orders": "Total Orders",
                    "Bulk_Action": "Bulk Action",
                    "Download_Mutliple_Vouchers": "Download Multiple Vouchers",
                    "Courier": "Courier",
                    "General": "Γενικά",
                    "Update_your_project_data": "Update your project data.",
                    "Secrets": "Secrets",
                    "Notifications": "Notifications",
                    "Payment_Subscription": "Payment & Subscription",
                    "Team": "Team",
                    'update': "Update",
                    "Delete_Project": "Delete Project",
                    "deleteconfim": "Are you sure you want to delete this project?",
                    "deleteWarning": "By deleting this project you are lossing all of your data!",
                    "deleteWarning2": "Make sure you have deactivated all of your extensions!",
                    "orders": "Orders",
                    "weeklyBalance": "Weekly Balance",
                    "quickStats": "Quick Stats",
                    "bestSeller": "Best Seller",
                    "latestOrderHistory": "Latest Order History",
                    "changeStatusToCompleted": "Change status to completed",
                    "changeStatusToPending": "Change status to pending",
                    "changeStatusToProcessing": "Change status to processing",
                    "changeStatusToCancelled": "Change status to cancelled",
                    "apply": "Apply",
                    "letsTalkAboutYourProject": "Let's talk about your project",
                    "getInTouch": "Get in Touch",
                    "off-canvas-modal-fullname": "Full Name",
                    "off-canvas-modal-message": "Message",
                    "vouchers": "Voucher",
                }
            },
            el: {
                translation: {
                    // Index GR
                    "title": "Μια Ολοκληρωμένη Λύση για την επιχείρηση σας, Πιο Γρήγορα από Ποτέ!",
                    "subtitle": "Με το Otoo, μπορείτε να συνδέσετε οποιαδήποτε πλατφόρμα e-commerce πιο γρήγορα από ποτέ.",
                    "Dashboard": "Πίνακας Ελέγχου",
                    //Off Canvas GR
                    "offcanvas-order_processing": "Δημιουργία Voucher#",
                    "off-canvas-customer-tab": "Πελάτης",
                    "off-canvas-shipping-tab": "Αποστολή",
                    "off-canvas-customer-info-header": "Στοιχεία Πελάτη",
                    "off-canvas-modal-name": "Όνομα",
                    "off-canvas-modal-fullname": "Όνοματεπώνυμο",
                    "off-canvas-modal-last-name": "Επώνυμο",
                    "off-canvas-email": "Email",
                    "off-canvas-modal-phone": "Τηλέφωνο",
                    "ooff-canvas-modal-address": "Διεύθυνση",
                    "off-canvas-modal-city": "Πόλη",
                    "off-canvas-modal-postal": "Τ.K.",
                    "off-canvas-modal-message": "Μήνυμα",
                    //Off Canvas Products Info GR
                    "off-canvas-customer-products-header": "Προϊόντα",
                    "off-canvas-customer-product-name": "Όνομα Προϊόντος",
                    "off-canvas-customer-product-quantity": "Ποσότητα Προϊόντος",
                    "off-canvas-customer-product-price": "Τιμή Προϊόντος",
                    //Off Canvas Delivery Info GR
                    "off-canvas-delivery-info": "Πληροφορίες Αποστολής",
                    "off-canvas-special-instructions": "Σημειώσεις Παράδοσης",
                    "off-canvas-customer-store-pickup": "Παραλαβή Απο Κατάστημα",
                    //Off Canvas Shipping Info GR
                    "off-canvas-shipping-information": "Στοιχεία Αποστολής",
                    "off-canvas-shipping-order-number": "Αριθμός Παραγγελίας",
                    "off-canvas-shipping-company": "Εταιρεία Αποστολής",
                    "off-canvas-shipping-provider": "Επιλογή Ταχυδρομικής Εταιρείας",
                    "off-canvas-parcel-cod": "Αντικαταβολή",
                    //Off Canvas ACS GR
                    "off-canvas-delivery-options-acs": "Επιλογές Αποστολής (ACS)",
                    "off-canvas-acs-standard-delivery": "Κανονική Αποστολή",
                    "off-canvas-acs-saturday-delivery": "Αποστολή Σάββατο",
                    "off-canvas-acs-urgent-delivery": "Επείγον Αποστολή",
                    "off-canvas-acs-pickup-delivery": "Παραλαβή Απο Κατάστημα",
                    //Off Canvas Courier4U GR
                    "off-canvas-delivery-options-courier4u": "Επιλογές Αποστολής (Courier4U)",
                    "off-canvas-delivery-courier4u-standard-delivery": "Κανονική Αποστολή",
                    "off-canvas-delivery-courier4u-saturday-delivery": "Παράδοση Σάββατο",
                    "off-canvas-delivery-courier4u-urgent-delivery": "Επείγουσα Παράδοση",
                    "off-canvas-delivery-courier4u-pickup-delivery": "Παραλαβή από το κατάστημα",
                    "off-canvas-delivery-courier4u-sameday-delivery": "Αυθημερόν",
                    "off-canvas-delivery-courier4u-ServiceProtocol-delivery": "Αριθμός Πρωτοκόλλου",
                    "off-canvas-delivery-courier4u-return-delivery": "Παράδοση - Παραλαβή",
                    //Off Canvas Parcel Weight/Height/Width Options GR
                    "off-canvas-wh-options-courier4u": "Επιλογές Δέματος",
                    "off-canvas-shipping-parcel-courier4u-weight": "Βάρος Δέματος",
                    "off-canvas-shipping-parcel-depth": "Βάθος Δέματος",
                    "off-canvas-shipping-parcel-height": "Ύψος Δέματος",
                    "off-canvas-shipping-parcel-width": "Πλάτος Δέματος",
                    "offcanvas-nav-customer": "Πελάτης",
                    "offcanvas-nav-shipping": "Αποστολή",
                    //Button GR
                    "off-canvas-save-btn": "Αποθήκευση",
                    "off-canvas-close-btn": "Ακύρωση",
                    //Voucher Table GR
                    "voucher-table-new-order": "Νέες",
                    "voucher-table-processing-orders": "Σε Επεξεργασία",
                    "voucher-table-completed-orders": "Ολοκληρωμένες",
                    "voucher-table-cancelled-orders": "Ακυρωμένες",
                    "voucher-table-all-orders": "Όλες",
                    "voucher-table-export-voucher-button": "Εξαγωγή Voucher",
                    "voucher-table-new-voucher-button": "Νέο Voucher",
                    //Order Header Section Table GR
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
                    "off-canvas-modal-desc": "Περιγραφή",
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
                    "sorting-table-order-id": "ΑΡΘ. ΠΑΡΑΓΓΕΛΙΑΣ",
                    "sorting-table-order-date": "ΗΜ/ΝΙΑ",
                    "sorting-table-order-cod": "ΑΝΤΙΚΑΤΑΒΟΛΗ",
                    "sorting-table-status-header": "ΚΑΤΑΣΤΑΣΗ",
                    "sorting-table-printed": "ΕΚΤΥΠΩΜΕΝΟ",
                    "sorting-table-action-header": "ΕΝΕΡΓΕΙΕΣ",
                    "sorting-no-results-msg": "Δεν υπάρχουν αποτελέσματα",
                    "sorting-create-voucher": "Δημιουργία Voucher",
                    // Side nav GR
                    "side-nav-logout": "Αποσύνδεση",
                    "side-nav-projects": "Έργα",
                    "side-nav-dashboard": "Πίνακας Ελέγχου",
                    "side-nav-default": "Στατιστικά",
                    "side-nav-customers": "Πελάτες",
                    "side-nav-products": "Προϊόντα",
                    "side-nav-orders": "Παραγγελίες",
                    "side-nav-resources": "Πόροι",
                    "side-nav-extensions": "Επεκτάσεις",
                    "side-nav-add-extensions": "Προσθήκη Επεκτάσεων",
                    "side-nav-settings": "Ρυθμίσεις",
                    // Dashboard overview EN ( Shared with Order/Customer/Product Overview )
                    "dashboard-overview_orders": "Παραγγελίες",
                    "dashboard-overview_products": "Προϊόντα",
                    "dashboard-overview_customers": "Πελάτες",
                    "dashboard-overview-total_orders": "Συνολικές Παραγγελίες",
                    "dashboard-overview_orders": "Παραγγελίες",
                    "dashboard-overview-week_orders": "Παραγγελίες",
                    "dashboard-overview-month_orders": "Παραγγελίες",
                    // Dashboard Shared Tables EN
                    "dashboard-table-order_id": "ID Παραγγελίας",
                    "dashboard-table-order-created": "Ημ/νία Δημιουργίας",
                    "dashboard-table-order_date": "Ημ/νία Παραγγελίας",
                    "dashboard-table-order_status": "Κατάσταση Παραγγελίας",
                    "dashboard-table-order_total": "Συνολικό Ποσό",
                    "dashboard-table-order_action": "Ενέργειες",
                    "dashboard-table-customer_name": "Όνομα Πελάτη",
                    "dashboard-table-customer_email": "Email",
                    "dashboard-table-customer-total_orders": "Συνολικές Παραγγελίες",
                    "dashboard-table-customer-total_spent": "Συνολικό Ποσό",
                    "dashboard-table-products_id": "ID Προϊόντος",
                    "dashboard-table-products_name": "Όνομα Προϊόντος",
                    "dashboard-table-products_price": "Τιμή Προϊόντος",
                    "dashboard-table-products_category": "Κατηγορία Προϊόντος",
                    "dashboard-table-products_type": "Τύπος Προϊόντος",
                    "dashboard-table-products-total_orders": "Συνολικές Παραγγελίες",
                    "Bulk_Action": "Μαζικές Ενέργειες",
                    "Download_Mutliple_Vouchers": "Λήψη Πολλαπλών Voucher",
                    "Courier": "Μεταφορική",
                    "General": "Γενικά",
                    "Update_your_project_data": "Ενημέρωση δεδομένων έργου",
                    "Secrets": "Συνθηματικά",
                    "Notifications": "Ειδοποιήσεις",
                    "Payment_Subscription": "Πληρωμές & Εγγραφές",
                    "Team": "Ομάδα",
                    "Delete_Project": "Διαγραφή Έργου",
                    'update': "Ενημέρωση",
                    "deleteconfim": "Είστε σίγουροι ότι θέλετε να διαγράψετε αυτό το έργο;",
                    "deleteWarning": "Ανακαλώντας αυτό το έργο, θα χάσετε όλα τα δεδομένα σας!",
                    "deleteWarning2": "Βεβαιωθείτε ότι έχετε απενεργοποιήσει όλες τις επεκτάσεις σας!",
                    "orders": "Παραγγελίες",
                    "weeklyBalance": "Εβδομαδιαίος Τζίρος",
                    "quickStats": "Στατιστικά Στοιχεία",
                    "bestSeller": "Δημοφιλής Προϊόντα",
                    "latestOrderHistory": "Ιστορικό Παραγγελιών",
                    "changeStatusToCompleted": "Αλλαγή κατάστασης σε ολοκληρωμένο",
                    "changeStatusToPending": "Αλλαγή κατάστασης σε εκκρεμότητα",
                    "changeStatusToProcessing": "Αλλαγή κατάστασης σε επεξεργασία",
                    "changeStatusToCancelled": "Αλλαγή κατάστασης σε ακυρωμένο",
                    "apply": "Εφαρμογή",
                    "letsTalkAboutYourProject": "Ας μιλήσουμε για το έργο σας",
                    "getInTouch": "Επικοινωνήστε μαζί μας",
                    "vouchers": "Παραστατικά",
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
