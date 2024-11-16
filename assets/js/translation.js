// Initialize i18next
i18next
    .use(i18nextBrowserLanguageDetector) // Detect user's language
    .init({
        // Determine fallback language and load from localStorage
        fallbackLng: localStorage.getItem('preferredLanguage') || 'en',
        debug: true, // Enable debug logs for easier troubleshooting

        resources: {
            // English Language Resources
            "en": {
                "translation": {
                    // Homepage Text
                    "index": {
                        "title": "All-in-One Solution, Integration & Dashboard Analytics Faster than Ever!",
                        "subtitle": "With KonektorX, you can connect any e-commerce platform quicker than ever."
                    },
                    // General Phrases
                    "general": {
                        "dashboard": "Dashboard",
                        "no_results": "No results found"
                    },
                    // Dashboard Section
                    "dashboard": {
                        // Customer Subsection
                        "customer": {
                            "header": "Customer Information",
                            "first_name": "First Name",
                            "last_name": "Last Name",
                            "email": "Email",
                            "phone": "Phone",
                            "address": "Address",
                            "city": "City",
                            "postal": "Postal Code",
                            "total_orders": "Total Orders",
                            "total_spent": "Total Spent"
                        },
                        // Product Subsection
                        "product": {
                            "header": "Products",
                            "name": "Product Name",
                            "quantity": "Quantity",
                            "price": "Price",
                            "id": "Product ID",
                            "category": "Category",
                            "type": "Type"
                        },
                        // Order Subsection
                        "order": {
                            "header": "Order Details",
                            "all": "All Orders",
                            "new": "New Orders",
                            "completed": "Completed Orders",
                            "pending": "Pending Orders",
                            "processing": "Processing Orders",
                            "cancelled": "Cancelled Orders",
                            // Voucher Subsection under Orders
                            "voucher": {
                                "header": "Voucher Management",
                                "create": "Create Voucher",
                                "update": "Update Voucher",
                                "new_voucher": "New Voucher",
                                "export": "Export Vouchers"
                            },
                            // Tables Subsection for Orders, Customers, Products, and Vouchers
                            "tables": {
                                "order": {
                                    "id": "Order ID",
                                    "order_id": "Order Number",
                                    "date": "Order Date",
                                    "status": "Order Status",
                                    "action": "Actions"
                                },
                                "customer": {
                                    "id": "Customer ID",
                                    "name": "Customer Name",
                                    "email": "Email",
                                    "phone": "Phone",
                                    "total_spent": "Total Spent",
                                    "total_orders": "Total Orders"
                                },
                                "product": {
                                    "id": "Product ID",
                                    "name": "Product Name",
                                    "quantity": "Stock Quantity",
                                    "price": "Price",
                                    "category": "Category"
                                },
                                "voucher": {
                                    "id": "Voucher ID",
                                    "code": "Voucher Code",
                                    "discount": "Discount",
                                    "expiry_date": "Expiry Date",
                                    "status": "Status"
                                }
                            }
                        }
                    },
                    // Shipping Information
                    "shipping": {
                        "header": "Shipping Information",
                        "order_number": "Order Number",
                        "parcel_weight": "Parcel Weight (kg)",
                        "parcel_options": "Parcel Options",
                        "parcel_dimensions": "Dimensions (L×W×H)",
                        "provider": "Courier Provider",
                        "cod": "Cash on Delivery"
                    },
                    // Billing Information
                    "billing": {
                        "header": "Billing Information",
                        "first_name": "First Name",
                        "last_name": "Last Name",
                        "address": "Address",
                        "city": "City",
                        "postal": "Postal Code",
                        "email": "Email",
                        "phone": "Phone"
                    },
                    // Sidebar Navigation
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
                    // Project Settings
                    "project_settings": {
                        "header": "Project Settings",
                        "general": "General",
                        "secrets": "Secrets",
                        "webhooks": "Webhooks",
                        "notifications": "Notifications",
                        "payment_subscription": "Payment & Subscriptions",
                        "team": "Team Management",
                        "courier4u": "Courier4U Integration",
                        "acs_courier": "ACS Courier Integration"
                    }
                }
            },
            // Greek Language Resources
            "el": {
                "translation": {
                    // Homepage Text
                    "index": {
                        "title": "Μια Ολοκληρωμένη Λύση, Πιο Γρήγορα από Ποτέ!",
                        "subtitle": "Με το KonektorX, μπορείτε να συνδέσετε οποιαδήποτε πλατφόρμα e-commerce πιο γρήγορα από ποτέ."
                    },
                    // General Phrases
                    "general": {
                        "dashboard": "Πίνακας Ελέγχου",
                        "no_results": "Δεν υπάρχουν αποτελέσματα"
                    },
                    // Dashboard Section
                    "dashboard": {
                        // Customer Subsection
                        "customer": {
                            "header": "Πληροφορίες Πελάτη",
                            "first_name": "Όνομα",
                            "last_name": "Επώνυμο",
                            "email": "Email",
                            "phone": "Τηλέφωνο",
                            "address": "Διεύθυνση",
                            "city": "Πόλη",
                            "postal": "Ταχ. Κώδικας",
                            "total_orders": "Συνολικές Παραγγελίες",
                            "total_spent": "Συνολικές Αγορές"
                        },
                        // Product Subsection
                        "product": {
                            "header": "Προϊόντα",
                            "name": "Όνομα Προϊόντος",
                            "quantity": "Ποσότητα",
                            "price": "Τιμή",
                            "id": "ID Προϊόντος",
                            "category": "Κατηγορία",
                            "type": "Τύπος"
                        },
                        // Order Subsection
                        "order": {
                            "header": "Πληροφορίες Παραγγελίας",
                            "all": "Όλες οι Παραγγελίες",
                            "new": "Νέες Παραγγελίες",
                            "completed": "Ολοκληρωμένες",
                            "pending": "Σε Εκκρεμότητα",
                            "processing": "Υπό Επεξεργασία",
                            "cancelled": "Ακυρωμένες",
                            // Voucher Subsection
                            "voucher": {
                                "header": "Διαχείριση Voucher",
                                "create": "Δημιουργία Voucher",
                                "update": "Ενημέρωση Voucher",
                                "new_voucher": "Νέο Voucher",
                                "export": "Εξαγωγή Voucher"
                            },
                            // Tables Subsection
                            "tables": {
                                "order": {
                                    "id": "ID Παραγγελίας",
                                    "order_id": "Αριθμός Παραγγελίας",
                                    "date": "Ημερομηνία",
                                    "status": "Κατάσταση",
                                    "action": "Ενέργειες"
                                },
                                "customer": {
                                    "id": "ID Πελάτη",
                                    "name": "Όνομα Πελάτη",
                                    "email": "Email",
                                    "phone": "Τηλέφωνο",
                                    "total_spent": "Σύνολο Αγορών",
                                    "total_orders": "Σύνολο Παραγγελιών"
                                }
                            }
                        }
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
    i18next.changeLanguage(lng, () => {
        localStorage.setItem('preferredLanguage', lng);
        updateContent();
    });
}

// Function to update content and support placeholders
function updateContent() {
    document.querySelectorAll('[data-i18n]').forEach(function (element) {
        const dataI18n = element.getAttribute('data-i18n');
        if (!dataI18n) return; // Skip if no valid key

        if (dataI18n.includes('[')) {
            const parts = dataI18n.match(/\[([^\]]+)\](.+)/);
            if (parts && parts[1] && parts[2]) {
                const attribute = parts[1];
                const key = parts[2];
                element.setAttribute(attribute, i18next.t(key.trim()) || '');
            }
        } else {
            element.textContent = i18next.t(dataI18n) || '';
        }
    });
}

// Listen for HTMX dynamic updates and reapply translations
document.body.addEventListener('htmx:afterSettle', () => {
    updateContent();
});
