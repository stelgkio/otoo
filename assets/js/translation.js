// Initialize i18next
i18next
    .use(i18nextBrowserLanguageDetector)
    .init({
        fallbackLng: localStorage.getItem('preferredLanguage') || 'en',
        debug: true,
        resources: {
            "en": {
                "translation": {
                    // Application-wide shared components and text
                    "common": {
                        "form_fields": {
                            "first_name": "First Name",
                            "last_name": "Last Name",
                            "email": "Email",
                            "phone": "Phone",
                            "address": "Address",
                            "city": "City",
                            "postal": "Postal Code",
                            "order_number": "Order Number"
                        },
                        "buttons": {
                            "save": "Save",
                            "close": "Close",
                            "cancel": "Cancel",
                            "create": "Create",
                            "update": "Update",
                            "export": "Export",
                            "submit": "Submit"
                        },
                        "status": {
                            "processing": "Processing...",
                            "loading": "Loading...",
                            "completed": "Completed",
                            "pending": "Pending",
                            "cancelled": "Cancelled"
                        },
                        "messages": {
                            "success": "Operation completed successfully",
                            "error": "An error occurred",
                            "no_results": "No results found"
                        },
                        "validation": {
                            "required": "This field is required",
                            "email": "Please enter a valid email address",
                            "phone": "Please enter a valid phone number",
                            "postal": "Please enter a valid postal code",
                            "amount": "Please enter a valid amount"
                        }
                    },

                    // Homepage content
                    "index": {
                        "title": "All-in-One Solution, Integration & Dashboard Analytics Faster than Ever!",
                        "subtitle": "With KonektorX, you can connect any e-commerce platform quicker than ever."
                    },

                    // Main navigation
                    "navigation": {
                        "side_nav": {
                            "dashboard": "Dashboard",
                            "projects": "Projects",
                            "customers": "Customers",
                            "products": "Products",
                            "orders": "Orders",
                            "resources": "Resources",
                            "extensions": "Extensions",
                            "add_extensions": "Add Extensions",
                            "settings": "Settings",
                            "default": "Default",
                            "logout": "Logout"
                        }
                    },

                    // Dashboard and main components
                    "dashboard": {
                        "header": "Dashboard",
                        "customer": {
                            "header": "Customer Information",
                            "search": {
                                "placeholder": "Search all customers",
                                "aria_label": "Search customers"
                            },
                            "metrics": {
                                "total_orders": "Total Orders",
                                "total_spent": "Total Spent"
                            }
                        },
                        "product": {
                            "header": "Products",
                            "fields": {
                                "name": "Product Name",
                                "quantity": "Quantity",
                                "price": "Price",
                                "id": "Product ID",
                                "category": "Category",
                                "type": "Type"
                            }
                        },
                        "order": {
                            "header": "Order Details",
                            "types": {
                                "all": "All Orders",
                                "new": "New Orders",
                                "completed": "Completed Orders",
                                "pending": "Pending Orders",
                                "processing": "Processing Orders",
                                "cancelled": "Cancelled Orders"
                            },
                            "actions": {
                                "bulk": "Bulk Actions"
                            },
                            "bulk_actions": {
                                "change_status_completed": "Change status to completed",
                                "change_status_pending": "Change status to pending",
                                "change_status_processing": "Change status to processing",
                                "change_status_cancelled": "Change status to cancelled"
                            }
                        },
                        "tables": {
                            "order": {
                                "id": "Order ID",
                                "order_id": "Order Number",
                                "date": "Order Date",
                                "status": "Order Status",
                                "action": "Action",
                                "created": "Order Created",
                                "total": "Total Amount"
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
                                "category": "Category",
                                "type": "Product Type",
                                "total_orders": "Total Orders"
                            },

                        }
                    },

                    // Offcanvas components
                    "offcanvas": {
                        "order_processing": {
                            "header": {
                                "title": "Order Processing #",
                                "close": "Close"
                            },
                            "tabs": {
                                "customer": "Customer",
                                "shipping": "Shipping"
                            },
                            "customer": {
                                "header": "Customer Info",
                                "products": {
                                    "header": "Products",
                                    "table": {
                                        "product": "Product",
                                        "quantity": "Quantity",
                                        "price": "Price"
                                    }
                                },
                                "delivery": {
                                    "header": "Delivery Info",
                                    "instructions": "Delivery Instructions"
                                }
                            },
                            "shipping": {
                                "header": "Shipping Information",
                                "form": {
                                    "company": "Shipping Company",
                                    "provider": "Select Courier Provider",
                                    "order_number": "Order Number",
                                    "cod": "COD (€)"
                                },
                                "dimensions": {
                                    "weight": "Parcel Weight (kg)",
                                    "depth": "Parcel Depth (cm)",
                                    "width": "Parcel Width (cm)",
                                    "height": "Parcel Height (cm)"
                                },
                                "courier_options": {
                                    "acs": {
                                        "header": "Delivery Options (ACS)",
                                        "standard": "Standard Delivery",
                                        "saturday": "Saturday Delivery",
                                        "urgent": "Urgent Delivery",
                                        "pickup": "Pickup"
                                    },
                                    "courier4u": {
                                        "header": "W/D/W/H Options (Courier4U)",
                                        "delivery_header": "Delivery Options (Courier4U)",
                                        "saturday": "Delivery on Saturday",
                                        "urgent": "Urgent Delivery",
                                        "return": "Delivery - Pickup",
                                        "pickup": "Pickup from the store",
                                        "same_day": "Same-Day",
                                        "protocol": "Protocol Number"
                                    }
                                }
                            },
                            "actions": {
                                "create": "Create Voucher",
                                "close": "Close",
                                "processing": "Processing..."
                            }
                        }
                    },

                    // Voucher management
                    "voucher": {
                        "header": "Voucher Management",
                        "actions": {
                            "create": "Create Voucher",
                            "update": "Update Voucher",
                            "new": "New Voucher",
                            "export": "Export Vouchers"
                        },
                        "table": {
                            "id": "Voucher ID",
                            "code": "Voucher Code",
                            "discount": "Discount",
                            "expiry_date": "Expiry Date",
                            "status": "Status"
                        }
                    },

                    // Project settings
                    "project_settings": {
                        "header": "Project Settings",
                        "sections": {
                            "general": "General",
                            "secrets": "Secrets",
                            "webhooks": "Webhooks",
                            "notifications": "Notifications",
                            "payment_subscription": "Payment & Subscriptions",
                            "team": "Team Management"
                        },
                        "integrations": {
                            "courier4u": "Courier4U Integration",
                            "acs_courier": "ACS Courier Integration"
                        }
                    }
                }
            },

            "el": {
                "translation": {
                    // Application-wide shared components and text
                    "common": {
                        "form_fields": {
                            "first_name": "Όνομα",
                            "last_name": "Επώνυμο",
                            "email": "Email",
                            "phone": "Τηλέφωνο",
                            "address": "Διεύθυνση",
                            "city": "Πόλη",
                            "postal": "Ταχυδρομικός Κώδικας",
                            "order_number": "Αριθμός Παραγγελίας"
                        },
                        "buttons": {
                            "save": "Αποθήκευση",
                            "close": "Κλείσιμο",
                            "cancel": "Ακύρωση",
                            "create": "Δημιουργία",
                            "update": "Ενημέρωση",
                            "export": "Εξαγωγή",
                            "submit": "Υποβολή"
                        },
                        "status": {
                            "processing": "Επεξεργασία...",
                            "loading": "Φόρτωση...",
                            "completed": "Ολοκληρώθηκε",
                            "pending": "Εκκρεμεί",
                            "cancelled": "Ακυρώθηκε"
                        },
                        "messages": {
                            "success": "Η ενέργεια ολοκληρώθηκε με επιτυχία",
                            "error": "Παρουσιάστηκε σφάλμα",
                            "no_results": "Δεν βρέθηκαν αποτελέσματα"
                        },
                        "validation": {
                            "required": "Το πεδίο είναι υποχρεωτικό",
                            "email": "Παρακαλώ εισάγετε έγκυρη διεύθυνση email",
                            "phone": "Παρακαλώ εισάγετε έγκυρο αριθμό τηλεφώνου",
                            "postal": "Παρακαλώ εισάγετε έγκυρο ταχυδρομικό κώδικα",
                            "amount": "Παρακαλώ εισάγετε έγκυρο ποσό"
                        }
                    },

                    // Homepage content
                    "index": {
                        "title": "Μια Ολοκληρωμένη Λύση, Πιο Γρήγορα από Ποτέ!",
                        "subtitle": "Με το KonektorX, μπορείτε να συνδέσετε οποιαδήποτε πλατφόρμα e-commerce πιο γρήγορα από ποτέ."
                    },

                    // Main navigation
                    "navigation": {
                        "side_nav": {
                            "dashboard": "Πίνακας Ελέγχου",
                            "projects": "Projects",
                            "customers": "Πελάτες",
                            "products": "Προϊόντα",
                            "orders": "Παραγγελίες",
                            "resources": "Πόροι",
                            "extensions": "Επεκτάσεις",
                            "add_extensions": "Προσθήκη Επεκτάσεων",
                            "settings": "Ρυθμίσεις",
                            "default": "Προεπιλογή",
                            "logout": "Αποσύνδεση"
                        }
                    },

                    // Dashboard and main components
                    "dashboard": {
                        "header": "Πίνακας Ελέγχου",
                        "customer": {
                            "header": "Πληροφορίες Πελάτη",
                            "search": {
                                "placeholder": "Αναζήτηση σε όλους τους πελάτες",
                                "aria_label": "Αναζήτηση πελατών"
                            },
                            "metrics": {
                                "total_orders": "Σύνολο Παραγγελιών",
                                "total_spent": "Σύνολο Δαπανών"
                            }
                        },
                        "product": {
                            "header": "Προϊόντα",
                            "fields": {
                                "name": "Όνομα Προϊόντος",
                                "quantity": "Ποσότητα",
                                "price": "Τιμή",
                                "id": "Αναγνωριστικό Προϊόντος",
                                "category": "Κατηγορία",
                                "type": "Τύπος"
                            }
                        },
                        "order": {
                            "header": "Λεπτομέρειες Παραγγελίας",
                            "types": {
                                "all": "Όλες οι Παραγγελίες",
                                "new": "Νέες Παραγγελίες",
                                "completed": "Ολοκληρωμένες Παραγγελίες",
                                "pending": "Εκκρεμείς Παραγγελίες",
                                "processing": "Παραγγελίες σε Επεξεργασία",
                                "cancelled": "Ακυρωμένες Παραγγελίες"
                            },
                            "actions": {
                                "bulk": "Ομαδικές Ενέργειες"
                            },
                            "bulk_actions": {
                                "change_status_completed": "Αλλαγή κατάστασης σε ολοκληρωμένη",
                                "change_status_pending": "Αλλαγή κατάστασης σε εκκρεμή",
                                "change_status_processing": "Αλλαγή κατάστασης σε επεξεργασία",
                                "change_status_cancelled": "Αλλαγή κατάστασης σε ακυρωμένη"
                            },

                        },

                        "tables": {
                            "order": {
                                "id": "Αναγνωριστικό Παραγγελίας",
                                "order_id": "Αριθμός Παραγγελίας",
                                "date": "Ημερομηνία Παραγγελίας",
                                "status": "Κατάσταση Παραγγελίας",
                                "action": "Ενέργεια",
                                "created": "Ημερομηνία Δημιουργίας",
                                "total": "Συνολικό Ποσό"
                            },
                            "customer": {
                                "id": "Αναγνωριστικό Πελάτη",
                                "name": "Όνομα Πελάτη",
                                "email": "Email",
                                "phone": "Τηλέφωνο",
                                "total_spent": "Σύνολο Δαπανών",
                                "total_orders": "Σύνολο Παραγγελιών"
                            },
                            "product": {
                                "id": "Αναγνωριστικό Προϊόντος",
                                "name": "Όνομα Προϊόντος",
                                "quantity": "Ποσότητα",
                                "price": "Τιμή",
                                "category": "Κατηγορία",
                                "type": "Τύπος Προϊόντος",
                                "total_orders": "Σύνολο Παραγγελιών"
                            }
                        }
                    },

                    // Offcanvas components
                    "offcanvas": {
                        "order_processing": {
                            "header": {
                                "title": "Επεξεργασία Παραγγελίας #",
                                "close": "Κλείσιμο"
                            },
                            "tabs": {
                                "customer": "Πελάτης",
                                "shipping": "Αποστολή"
                            },
                            "customer": {
                                "header": "Στοιχεία Πελάτη",
                                "products": {
                                    "header": "Προϊόντα",
                                    "table": {
                                        "product": "Προϊόν",
                                        "quantity": "Ποσότητα",
                                        "price": "Τιμή"
                                    }
                                },
                                "delivery": {
                                    "header": "Πληροφορίες Παράδοσης",
                                    "instructions": "Οδηγίες Παράδοσης"
                                }
                            },
                            "shipping": {
                                "header": "Πληροφορίες Αποστολής",
                                "form": {
                                    "company": "Εταιρεία Μεταφοράς",
                                    "provider": "Επιλέξτε Πάροχο Κούριερ",
                                    "order_number": "Αριθμός Παραγγελίας",
                                    "cod": "Αντικαταβολή (€)"
                                },
                                "dimensions": {
                                    "weight": "Βάρος Πακέτου (kg)",
                                    "depth": "Βάθος Πακέτου (cm)",
                                    "width": "Πλάτος Πακέτου (cm)",
                                    "height": "Ύψος Πακέτου (cm)"
                                },
                                "courier_options": {
                                    "acs": {
                                        "header": "Επιλογές Παράδοσης (ACS)",
                                        "standard": "Κανονική Παράδοση",
                                        "saturday": "Παράδοση Σάββατο",
                                        "urgent": "Επείγουσα Παράδοση",
                                        "pickup": "Παραλαβή"
                                    },
                                    "courier4u": {
                                        "header": "Επιλογές Β/Π/Υ/Μ (Courier4U)",
                                        "delivery_header": "Επιλογές Παράδοσης (Courier4U)",
                                        "saturday": "Παράδοση Σάββατο",
                                        "urgent": "Επείγουσα Παράδοση",
                                        "return": "Παράδοση - Παραλαβή",
                                        "pickup": "Παραλαβή από το κατάστημα",
                                        "same_day": "Αυθημερόν",
                                        "protocol": "Αριθμός Πρωτοκόλλου"
                                    }
                                }
                            },
                            "actions": {
                                "create": "Δημιουργία Κουπονιού",
                                "close": "Κλείσιμο",
                                "processing": "Επεξεργασία..."
                            }
                        }
                    },

                    // Voucher management
                    "voucher": {
                        "header": "Διαχείριση Voucher",
                        "actions": {
                            "create": "Δημιουργία Voucher",
                            "update": "Ενημέρωση Voucher",
                            "new": "Νέο Κουπόνι",
                            "export": "Εξαγωγή Voucher"
                        },
                        "table": {
                            "id": "Αναγνωριστικό Voucher",
                            "code": "Κωδικός Voucher",
                            "discount": "Έκπτωση",
                            "expiry_date": "Ημερομηνία Λήξης",
                            "status": "Κατάσταση"
                        }
                    },

                    // Project settings
                    "project_settings": {
                        "header": "Ρυθμίσεις Έργου",
                        "sections": {
                            "general": "Γενικές Ρυθμίσεις",
                            "secrets": "Μυστικά",
                            "webhooks": "Webhooks",
                            "notifications": "Ειδοποιήσεις",
                            "payment_subscription": "Πληρωμές & Συνδρομές",
                            "team": "Διαχείριση Ομάδας"
                        },
                        "integrations": {
                            "courier4u": "Ενοποίηση Courier4U",
                            "acs_courier": "Ενοποίηση ACS Courier"
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
        if (!dataI18n) return;

        if (dataI18n.includes('[')) {
            const parts = dataI18n.match(/\[([^\]]+)\](.+)/);
            if (parts && parts[1] && parts[2]) {
                const attribute = parts[1];
                const key = parts[2].trim();
                element.setAttribute(attribute, i18next.t(key) || '');
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