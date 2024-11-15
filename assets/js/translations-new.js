// Initialize i18next
i18next
    // .use(i18nextHttpBackend) // Optional: Load translations from external files
    .use(i18nextBrowserLanguageDetector) // Detect user's language
    .init({
        // fallbackLng: 'en',
        fallbackLng: localStorage.getItem('preferredLanguage') || 'en',
        debug: true, // Enable debug logs in the console

        resources: {

            "en": {
                "translation": {
                    "index": {
                        "title": "All in one Solution, Integration & Dashboard Analytics Faster than Ever.!",
                        "subtitle": "With KonektorX you can connect any e-commerce platform quicker than ever."
                    },
                    "general": {
                        "dashboard": "Dashboard",
                        "no_results": "No results found"
                    },

                    // Dashboard - Customer - Product - Order 
                    "dashboard": {
                        "customer": {
                            "header": "Customer Info",
                            "first_name": "First Name",
                            "last_name": "Last Name",
                            "email": "Email",
                            "phone": "Phone",
                            "address": "Address",
                            "city": "City",
                            "postal": "Postal",
                            "total_orders": "Total Orders",
                            "total_spent": "Total Spent"
                        },
                        "product": {
                            "header": "Products",
                            "name": "Product Name",
                            "quantity": "Product Quantity",
                            "price": "Product Price",
                            "id": "Product ID",
                            "category": "Category",
                            "type": "Type"
                        },
                        "order": {
                            "header": "Order Details",
                            "all": "All",
                            "new": "New",
                            "completed": "Completed",
                            "pending": "Pending",
                            "processing": "Processing",
                            "cancelled": "Cancelled",
                            // Voucher //
                            "voucher": {
                                "header": "Voucher",
                                "create": "Create Voucher",
                                "update": "Update Voucher",
                                "New Voucher": "New Voucher",
                                "export": "Export Voucher",
                                "table": {
                                    "id": "Voucher ID",
                                    "order_id": "Order ID",
                                    "date": "Order Date",
                                    "status": "Status",
                                    "action": "Action"
                                },
                                // Modal //
                                "modal": {

                                },

                            }
                        },

                    },
                    //End Dashboard - Customer - Product - Order


                    "shipping": {
                        "header": "Shipping Information",
                        "order_number": "Order Number",
                        "parcel_weight": "Parcel Weight",
                        "parcel_options": "Package Options",
                        "parcel_depth": "Parcel Depth",
                        "parcel_height": "Parcel Height",
                        "parcel_width": "Parcel Width",
                        "provider": "Select Courier Provider",
                        "cod": "COD"
                    },
                    "billing": {
                        "header": "Billing Information",
                        "name": "First Name",
                        "last_name": "Last Name",
                        "address": "Address",
                        "city": "City",
                        "postal": "Postal",
                        "email": "E-mail",
                        "phone": "Phone"
                    },
                    "side_nav": {
                        "logout": "Logout",
                        "projects": "Projects",
                        "dashboard": "Dashboard",
                        "customers": "Customers",
                        "products": "Products",
                        "orders": "Orders",
                        "resources": "Resources",
                        "extensions": "Extensions",
                        "settings": "Settings"
                    },
                    "project_settings": {
                        "header": "Project Settings",
                        "general": "General",
                        "secrets": "Secrets",
                        "webhooks": "Webhooks",
                        "notifications": "Notifications",
                        "payment_subscription": "Payment & Subscription",
                        "team": "Team",
                        "courier4u": "Courier4u",
                        "acs_courier": "ACS Courier"
                    }
                }
            },
            "el": {
                "translation": {
                    "index": {
                        "title": "Μια Ολοκληρωμένη Λύση για την επιχείρηση σας, Πιο Γρήγορα από Ποτέ!",
                        "subtitle": "Με το KonektorX, μπορείτε να συνδέσετε οποιαδήποτε πλατφόρμα e-commerce πιο γρήγορα από ποτέ."
                    },
                    "general": {
                        "dashboard": "Πίνακας Ελέγχου",
                        "no_results": "Δεν υπάρχουν αποτελέσματα"
                    },
                    "dashboard": {
                        "customer": {
                            "header": "Στοιχεία Πελάτη",
                            "first_name": "Όνομα",
                            "last_name": "Επώνυμο",
                            "email": "Email",
                            "phone": "Τηλέφωνο",
                            "address": "Διεύθυνση",
                            "city": "Πόλη",
                            "postal": "Τ.K.",
                            "total_orders": "Συνολικές Παραγγελίες",
                            "total_spent": "Συνολικό Ποσό"
                        },
                        "product": {
                            "header": "Προϊόντα",
                            "name": "Όνομα Προϊόντος",
                            "quantity": "Ποσότητα Προϊόντος",
                            "price": "Τιμή Προϊόντος",
                            "id": "ID Προϊόντος",
                            "category": "Κατηγορία Προϊόντος",
                            "type": "Τύπος Προϊόντος"
                        },
                        "order": {
                            "header": "Πληροφορίες Παραγγελίας",
                            "all": "Όλες",
                            "new": "Νέες",
                            "completed": "Ολοκληρωμένες",
                            "pending": "Σε Εκκρεμότητα",
                            "processing": "Σε Επεξεργασία",
                            "cancelled": "Ακυρωμένες",
                            // Voucher //
                            "voucher": {
                                "header": "Voucher",
                                "create": "Δημιουργία Voucher",
                                "update": "Επεξεργασία Voucher",
                                "New Voucher": "Νέο Voucher",
                                "export": "Εξαγωγή Voucher",
                                "table": {
                                    "id": "Voucher ID",
                                    "order_id": "ID Παραγγελίας",
                                    "date": "Ημ/νία Παραγγελίας",
                                    "status": "Κατάσταση",
                                    "action": "Ενέργειες"
                                },
                                "Modal": {

                                },
                            }
                        },

                    },
                    "shipping": {
                        "header": "Στοιχεία Αποστολής",
                        "order_number": "Αριθμός Παραγγελίας",
                        "parcel_weight": "Βάρος Δέματος",
                        "parcel_options": "Επιλογές Δέματος",
                        "parcel_depth": "Βάθος Δέματος",
                        "parcel_height": "Ύψος Δέματος",
                        "parcel_width": "Πλάτος Δέματος",
                        "provider": "Επιλογή Ταχυδρομικής Εταιρείας",
                        "cod": "Αντικαταβολή"
                    },
                    "billing": {
                        "header": "Πληροφορίες Πληρωμής",
                        "name": "Όνομα",
                        "last_name": "Επώνυμο",
                        "address": "Διεύθυνση",
                        "city": "Πόλη",
                        "postal": "Τ.Κ.",
                        "email": "E-mail",
                        "phone": "Τηλέφωνο"
                    },
                    "side_nav": {
                        "logout": "Αποσύνδεση",
                        "projects": "Έργα",
                        "dashboard": "Πίνακας Ελέγχου",
                        "customers": "Πελάτες",
                        "products": "Προϊόντα",
                        "orders": "Παραγγελίες",
                        "resources": "Πόροι",
                        "extensions": "Επεκτάσεις",
                        "settings": "Ρυθμίσεις"
                    },
                    "project_settings": {
                        "header": "Ρυθμίσεις Έργου",
                        "general": "Γενικά",
                        "secrets": "Μυστικά",
                        "webhooks": "Webhooks",
                        "notifications": "Ειδοποιήσεις",
                        "payment_subscription": "Πληρωμές & Συνδρομές",
                        "team": "Ομάδα",
                        "courier4u": "Courier4u",
                        "acs_courier": "ACS Courier"
                    }
                }
            }
        }


    }, function (err, t) {
        if (err) console.error("i18next initialization error:", err);
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
            // Match and extract attribute and key parts
            var parts = dataI18n.match(/\[([^\]]+)\](.+)/);
            if (parts) { // Add a check to avoid errors
                var attribute = parts[1];
                var key = parts[2];
                element.setAttribute(attribute, i18next.t(key.trim()));
            }
        } else {
            // Text content translation
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
