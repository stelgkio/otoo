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
                    "offcanvas-order_processing": "Order Processing #",
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
                    //Products Info
                    "off-canvas-customer-products-header": "Products",
                    "off-canvas-customer-product-name": "Product Name",
                    "off-canvas-customer-product-quantity": "Product Quantity",
                    "off-canvas-customer-product-price": "Product Price",
                    //Delivery Info
                    "off-canvas-delivery-info": "Delivery Info",
                    "off-canvas-special-instructions": "Special Instructions",
                    "off-canvas-customer-store-pickup": "Store pickup",
                    // Shipping Info
                    "off-canvas-shipping-information": "Shipping Information",
                    "off-canvas-shipping-order-number": "Order Number",
                    "off-canvas-shipping-parcel-weight": "Parcel Weight",
                    "off-canvas-shipping-company": "Shipping Company",
                    "off-canvas-shipping-provider": "Select Courier Provider",
                    //ACS
                    "off-canvas-delivery-options-acs": "Delivery Options (ACS)",
                    "off-canvas-acs-standard-delivery": "Standard Delivery",
                    "off-canvas-acs-saturday-delivery": "Saturday Delivery",
                    "off-canvas-acs-urgent-delivery": "Urgent Delivery",
                    "off-canvas-acs-pickup-delivery": "Pickup Delivery",
                    //Courier4U
                    "off-canvas-delivery-options-courier4u": "Delivery Options (Courier4U)",
                    "off-canvas-delivery-courier4u-standard-delivery": "Standard Delivery",
                    "off-canvas-delivery-courier4u-saturday-delivery": "Saturday Delivery",
                    "off-canvas-delivery-courier4u-urgent-delivery": "Urgent Delivery",
                    "off-canvas-delivery-courier4u-pickup-delivery": "Pickup Delivery",
                    "off-canvas-delivery-courier4u-sameday-delivery": "Same Day Delivery",
                    //button
                    "off-canvas-button-save": "Save Changes",






                }
            },
            el: {
                translation: {
                    // Index GR
                    "title": "Όλα σε μία Λύση, Ενσωμάτωση & Αναλυτικά Δεδομένα Πίνακα Ελέγχου Πιο Γρήγορα από Ποτέ.",
                    "Dashboard": "Πίνακας Ελέγχου",
                    // Off Canvas GR
                    "offcanvas-order_processing": "Επεξεργασία Παραγγελίας #",
                    "off-canvas-customer-tab": "Πελάτης",
                    "off-canvas-shipping-tab": "Αποστολή",
                    "off-canvas-customer-info-header": "Στοιχεία Πελάτη",
                    "off-canvas-customer-first_name": "Όνομα",
                    "off-canvas-customer-last_name": "Επώνυμο",
                    "off-canvas-customer-email": "Email",
                    "off-canvas-customer-phone": "Τηλέφωνο",
                    "off-canvas-customer-address": "Διεύθυνση",
                    "off-canvas-customer-city": "Πόλη",
                    "off-canvas-customer-postal": "Τ.K.",
                    //Products Info
                    "off-canvas-customer-products-header": "Προϊόντα",
                    "off-canvas-customer-product-name": "Όνομα Προϊόντος",
                    "off-canvas-customer-product-quantity": "Ποσότητα Προϊόντος",
                    "off-canvas-customer-product-price": "Τιμή Προϊόντος",
                    //Delivery Info
                    "off-canvas-delivery-info": "Πληροφορίες Αποστολής",
                    "off-canvas-special-instructions": "Σημειώσεις Παράδοσης",
                    "off-canvas-customer-store-pickup": "Παραλαβή Απο Κατάστημα",
                    //Shipping Info
                    "off-canvas-shipping-information": "Στοιχεία Αποστολής",
                    "off-canvas-shipping-order-number": "Αριθμός Παραγγελίας",
                    "off-canvas-shipping-parcel-weight": "Βάρος Πακέτου",
                    "off-canvas-shipping-company": "Εταιρεία Αποστολής",
                    "off-canvas-shipping-provider": "Επιλογή Ταχυδρομικής Εταιρείας",
                    //ACS
                    "off-canvas-delivery-options-acs": "Επιλογές Αποστολής (ACS)",
                    "off-canvas-acs-standard-delivery": "Κανονική Αποστολή",
                    "off-canvas-acs-saturday-delivery": "Αποστολή Σάββατο",
                    "off-canvas-acs-urgent-delivery": "Επείγον Αποστολή",
                    "off-canvas-acs-pickup-delivery": "Παραλαβή Απο Κατάστημα",
                    //Courier4U
                    "off-canvas-delivery-options-courier4u": "Επιλογές Αποστολής (Courier4U)",
                    "off-canvas-delivery-courier4u-standard-delivery": "Κανονική Αποστολή",
                    "off-canvas-delivery-courier4u-saturday-delivery": "Αποστολή Σάββατο",
                    "off-canvas-delivery-courier4u-urgent-delivery": "Επείγον Αποστολή",
                    "off-canvas-delivery-courier4u-pickup-delivery": "Παραλαβή Απο Κατάστημα",
                    "off-canvas-delivery-courier4u-sameday-delivery": "Παράδοση εντός ημέρας",
                    // Button  
                    "off-canvas-button-save": "Αποθήκευση",






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
