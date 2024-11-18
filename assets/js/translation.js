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
                            "submit": "Submit",
                            "apply": "Apply"
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
                        "overview": {
                            "orders": "Orders",
                            "customers": "Customers",
                            "products": "Products",
                            "total_orders": "Total Orders",
                            "day_orders": "Orders (24H)",
                            "week_orders": "Orders (7D)",
                            "month_orders": "Orders (30D)",
                            "default": "Default",
                            "quick_stats": "Quick Stats",
                            // Tooltips
                            "tooltip": {
                                "day_orders": "The total completed order the last 24 hours.",
                                "week_orders": "The total completed order the last 7 days.",
                                "month_orders": "The total completed order the last month."
                            },
                            "weekly_balance": {
                                "title": "Weekly Balance",
                                "active_rate": "Active Order Rate",
                                "total_revenue": "Total Revenue"
                            },
                            "stats": {
                                "orders": {
                                    "title": "Orders",
                                    "icon": "bag-plus"
                                },
                                "customers": {
                                    "title": "Customers",
                                    "icon": "file-earmark-person"
                                },
                                "products": {
                                    "title": "Products",
                                    "icon": "shop"
                                }

                            }

                        },
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
                                "type": "Type",
                                "search_placeholder": "Search all products",
                                "search_aria_label": "Search products"
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
                            }
                        }
                    },
                    // Modal section
                    "modal": {
                        "order": {
                            "title": "Order",
                            "payment_method": "Payment Method",
                            "details": {
                                "header": "Order Details",
                                "payment_info": "Payment Information",
                                "products_list": "Products List"
                            },
                            "billing": {
                                "header": "Billing Information",
                                "address_2": "Address 2"
                            },
                            "shipping": {
                                "header": "Shipping Information",
                                "address_2": "Address 2",
                                "customer_note": "Customer Notes"
                            },
                            "products": {
                                "header": "Products",
                                "table": {
                                    "name": "Product",
                                    "quantity": "Qty",
                                    "price": "Price"
                                }
                            },
                            "buttons": {
                                "save": "Save Changes",
                                "close": "Close",
                                "saving": "Saving..."
                            }
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
                    },
                    "voucher": {
                        "types": {
                            "new": "New",
                            "processing": "Processing",
                            "completed": "Completed",
                            "cancelled": "Cancelled",
                            "all": "All"
                        },
                        "actions": {
                            "bulk": "Bulk Action",
                            "download_multiple": "Download Multiple Vouchers",
                            "new": "New Voucher",
                            "create": "Create Voucher",
                            "export": "Export"
                        },
                        "table": {
                            "id": "Voucher ID",
                            "order_id": "Order ID",
                            "date": "Date",
                            "cod": "COD",
                            "status": "Status",
                            "printed": "Printed",
                            "action": "Action"
                        },
                        "messages": {
                            "no_results": "No vouchers found",
                            "error": "An error occurred",
                            "success": "Operation completed successfully"
                        }
                    },
                    "extensions": {
                        "header": "Extensions",
                        "pricing": {
                            "per_month": "€/mo",
                            "per_year": "€/yr"
                        },
                        "descriptions": {
                            "acs": "Streamline your shipping process with ACS Courier integration. Automate order fulfillment, track shipments in real-time, and provide excellent customer service.",
                            "courier4u": "Optimize your delivery operations with Courier4U. Manage multiple shipments, automate tracking updates, and enhance your logistics efficiency.",
                            "wallet": "Take control of your finances with comprehensive expense tracking, automated tax calculations, and integrated campaign management.",
                            "data_sync": "Keep your business data synchronized across all platforms. Seamlessly integrate customer information, product data, and analytics in real-time.",
                            "team": "Expand your team's capabilities with collaborative tools, role-based access control, and streamlined workflow management."
                        },
                        "couriers": {
                            "acs": {
                                "title": "ACS Courier Integration",
                                "cta": "Buy Extension",
                                "features": {
                                    "status_updates": "Automatic order status updates",
                                    "vouchers": "Download shipping vouchers",
                                    "notifications": "Send customer notifications via email"
                                }
                            },
                            "courier4u": {
                                "title": "Courier4U Integration",
                                "cta": "Buy Extension",
                                "features": {
                                    "status_updates": "Automatic order status updates",
                                    "vouchers": "Download shipping vouchers",
                                    "notifications": "Send customer notifications via email"
                                }
                            }
                        },
                        "wallet": {
                            "title": "Wallet & Expenses",
                            "cta": "Buy Extension",
                            "features": {
                                "facebook": "Integrate with Facebook campaigns",
                                "courier_charges": "Monitor courier charges",
                                "tax": "Automatic tax calculations",
                                "support": "Priority support for KonektorX users"
                            }
                        },
                        "data_sync": {
                            "title": "Data Synchronizer",
                            "cta": "Buy Extension",
                            "features": {
                                "customer_data": "Retrieve customer data seamlessly",
                                "product_sync": "Sync product information and variations",
                                "order_history": "Access full order history",
                                "analytics": "Build custom analytics and reports"
                            }
                        },
                        "team": {
                            "title": "Team Management",
                            "cta": "Buy Extension",
                            "features": {
                                "unlimited": "Add as many users you want to your project"
                            }
                        }
                    },
                    "settings": {
                        "header": "Project Settings",
                        "navigation": {
                            "general": "General",
                            "secrets": "Secrets",
                            "webhooks": "Webhooks",
                            "notifications": "Notifications",
                            "payment": "Payment & Subscription",
                            "team": "Team"
                        },
                        "general": {
                            "title": "General",
                            "description": "Update your project data",
                            "fields": {
                                "name": "Name",
                                "description": "Description",
                                "domain": "Domain",
                                "domain_tooltip": "Contact our support team to change your domain."
                            },
                            "buttons": {
                                "update": "Update"
                            },
                            "delete": {
                                "title": "Delete Project",
                                "warning": "By deleting this project you are losing all of your data!",
                                "extension_warning": "Make sure you have deactivate all of your extensions!",
                                "confirmation": "Are you sure you want to delete this project?",
                                "button": "Delete"
                            }
                        },
                        "secrets": {
                            "title": "Project Secrets",
                            "reset": {
                                "title": "Secrets Reset",
                                "description": "By updating your project secrets you will reset all the secrets for this project.",
                                "consumer_key": "ConsumerKey",
                                "consumer_secret": "ConsumerSecret",
                                "error": "Unable to connect to the server, check your secrets and try again."
                            }
                        },
                        "couriers": {
                            "acs": {
                                "title": "ACS Courier",
                                "description": "DeActivate extension and remove subscription",
                                "fields": {
                                    "user_id": "User Id",
                                    "user_password": "User Password",
                                    "company_id": "Company Id",
                                    "company_password": "Company Password",
                                    "billing_code": "Billing Code",
                                    "acs_api_key": "AcsAPIKey"
                                },
                                "printer": {
                                    "label": "Select Printer Type",
                                    "laser": "Laser",
                                    "thermal": "Thermal"
                                }
                            },
                            "courier4u": {
                                "title": "Courier4u",
                                "description": "DeActivate extension and remove subscription",
                                "fields": {
                                    "api_key": "APIKey"
                                }
                            }
                        },
                        "team": {
                            "title": "Team",
                            "description": "The new user can not add new extension to project",
                            "add_member": {
                                "title": "Add team member",
                                "fields": {
                                    "email": "Email",
                                    "name": "Name",
                                    "last_name": "Last Name",
                                    "password": "Password",
                                    "confirm_password": "Confirmation Password",
                                    "role": {
                                        "label": "Role",
                                        "user": "User",
                                        "admin": "Admin"
                                    },
                                    "notifications": "Receive Notification"
                                }
                            },
                            "members": "Members"
                        },
                        "payment": {
                            "title": "Payments",
                            "table": {
                                "extension_name": "Extension Name",
                                "amount": "Amount",
                                "status": "Status",
                                "created": "Created",
                                "status_values": {
                                    "paid": "Paid",
                                    "failed": "Failed",
                                    "pending": "Pending"
                                }
                            }
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
                            "submit": "Υποβολή",
                            "apply": "Εφαρμογή"
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
                        "overview": {
                            "orders": "Παραγγελίες",
                            "customers": "Πελάτες",
                            "products": "Προϊόντα",
                            "total_orders": "Σύνολο Παραγγελιών",
                            "day_orders": "Παραγγελίες",
                            "week_orders": "Παραγγελίες",
                            "month_orders": "Παραγγελίες",
                            "default": "Προεπιλογή",
                            "quick_stats": "Γρήγορα Στατιστικά",
                            // Tooltips
                            "tooltip": {
                                "day_orders": "Οι συνολικές ολοκληρωμένες παραγγελίες τις τελευταίες 24 ώρες.",
                                "week_orders": "Οι συνολικές ολοκληρωμένες παραγγελίες τις τελευταίες 7 ημέρες.",
                                "month_orders": "Οι συνολικές ολοκληρωμένες παραγγελίες τον τελευταίο μήνα."
                            },
                            "weekly_balance": {
                                "title": "Εβδομαδιαίο Υπόλοιπο",
                                "active_rate": "Ποσοστό Ενεργών Παραγγελιών",
                                "total_revenue": "Συνολικά Έσοδα"
                            },
                            "stats": {
                                "orders": {
                                    "title": "Παραγγελίες",
                                    "icon": "bag-plus"
                                },
                                "customers": {
                                    "title": "Πελάτες",
                                    "icon": "file-earmark-person"
                                },
                                "products": {
                                    "title": "Προϊόντα",
                                    "icon": "shop"
                                }
                            }

                        },

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
                                "id": "ID Προϊόντος",
                                "category": "Κατηγορία",
                                "type": "Τύπος",
                                "search_placeholder": "Αναζήτηση σε όλα τα προϊόντα",
                                "search_aria_label": "Αναζήτηση προϊόντων"
                            }
                        },
                        "order": {
                            "header": "Λεπτομέρειες Παραγγελίας",
                            "types": {
                                "all": "Όλες",
                                "new": "Νέες",
                                "completed": "Ολοκληρωμένες",
                                "pending": "Εκκρεμείς",
                                "processing": "Επεξεργασία",
                                "cancelled": "Ακυρωμένες"
                            },
                            "actions": {
                                "bulk": "Ομαδικές Ενέργειες"
                            },
                            "bulk_actions": {
                                "change_status_completed": "Αλλαγή κατάστασης σε ολοκληρωμένη",
                                "change_status_pending": "Αλλαγή κατάστασης σε εκκρεμή",
                                "change_status_processing": "Αλλαγή κατάστασης σε επεξεργασία",
                                "change_status_cancelled": "Αλλαγή κατάστασης σε ακυρωμένη"
                            }
                        },
                        "tables": {
                            "order": {
                                "id": "ID Παραγγελίας",
                                "order_id": "Αριθμός Παραγγελίας",
                                "date": "Ημ/νία Παραγγελίας",
                                "status": "Κατάσταση Παραγγελίας",
                                "action": "Ενέργεια",
                                "created": "Ημ/νία Δημιουργίας",
                                "total": "Συνολικό Ποσό"
                            },
                            "customer": {
                                "id": "ID Πελάτη",
                                "name": "Όνομα Πελάτη",
                                "email": "Email",
                                "phone": "Τηλέφωνο",
                                "total_spent": "Σύνολο Δαπανών",
                                "total_orders": "Σύνολο Παραγγελιών"
                            },
                            "product": {
                                "id": "ID Προϊόντος",
                                "name": "Όνομα Προϊόντος",
                                "quantity": "Ποσότητα",
                                "price": "Τιμή",
                                "category": "Κατηγορία",
                                "type": "Τύπος Προϊόντος",
                                "total_orders": "Σύνολο Παραγγελιών"
                            }
                        }
                    },

                    // Modal section
                    "modal": {
                        "order": {
                            "title": "Παραγγελία",
                            "payment_method": "Μέθοδος Πληρωμής",
                            "details": {
                                "header": "Λεπτομέρειες Παραγγελίας",
                                "payment_info": "Πληροφορίες Πληρωμής",
                                "products_list": "Λίστα Προϊόντων"
                            },
                            "billing": {
                                "header": "Πληροφορίες Τιμολόγησης",
                                "address_2": "Διεύθυνση 2"
                            },
                            "shipping": {
                                "header": "Πληροφορίες Αποστολής",
                                "address_2": "Διεύθυνση 2",
                                "customer_note": "Σημειώσεις Πελάτη"
                            },
                            "products": {
                                "header": "Προϊόντα",
                                "table": {
                                    "name": "Προϊόν",
                                    "quantity": "Ποσ",
                                    "price": "Τιμή"
                                }
                            },
                            "buttons": {
                                "save": "Αποθήκευση Αλλαγών",
                                "close": "Κλείσιμο",
                                "saving": "Αποθήκευση..."
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
                                "create": "Δημιουργία Voucher",
                                "close": "Κλείσιμο",
                                "processing": "Επεξεργασία..."
                            }
                        }
                    },

                    // Voucher Μanagement
                    "voucher": {
                        "header": "Διαχείριση Voucher",
                        "actions": {
                            "create": "Δημιουργία Voucher",
                            "update": "Ενημέρωση Voucher",
                            "new": "Νέο Voucher",
                            "export": "Εξαγωγή Voucher"
                        },
                        "table": {
                            "id": "ID Voucher",
                            "code": "Κωδικός Voucher",
                            "discount": "Έκπτωση",
                            "expiry_date": "Ημ/νία Λήξης",
                            "status": "Κατάσταση"
                        }
                    },

                    // Project Settings
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
                            "courier4u": "Courier4U Extension",
                            "acs_courier": "ACS Courier Extension"
                        }
                    },
                    // Greek version
                    "voucher": {
                        "types": {
                            "new": "Νέο",
                            "processing": "Σε Επεξεργασία",
                            "completed": "Ολοκληρωμένο",
                            "cancelled": "Ακυρωμένο",
                            "all": "Όλα"
                        },
                        "actions": {
                            "bulk": "Μαζική Ενέργεια",
                            "download_multiple": "Λήψη Πολλαπλών Vouchers",
                            "new": "Νέο Voucher",
                            "create": "Δημιουργία Voucher",
                            "export": "Εξαγωγή"
                        },
                        "table": {
                            "id": "ID Voucher",
                            "order_id": "ID Παραγγελίας",
                            "date": "Ημ/νία",
                            "cod": "Αντικαταβολή",
                            "status": "Κατάσταση",
                            "printed": "Εκτυπωμένο",
                            "action": "Ενέργεια"
                        },
                        "messages": {
                            "no_results": "Δεν βρέθηκαν vouchers",
                            "error": "Παρουσιάστηκε σφάλμα",
                            "success": "Η ενέργεια ολοκληρώθηκε με επιτυχία"
                        }
                    },
                    // Greek translations
                    "extensions": {
                        "header": "Επεκτάσεις",
                        "pricing": {
                            "per_month": "€/μήνα",
                            "per_year": "€/έτος"
                        },
                        "descriptions": {
                            "acs": "Απλοποιήστε τη διαδικασία αποστολής με την ενσωμάτωση ACS Courier. Αυτοματοποιήστε την εκπλήρωση παραγγελιών, παρακολουθήστε τις αποστολές σε πραγματικό χρόνο και παρέχετε εξαιρετική εξυπηρέτηση πελατών.",
                            "courier4u": "Βελτιστοποιήστε τις λειτουργίες παράδοσης με το Courier4U. Διαχειριστείτε πολλαπλές αποστολές, αυτοματοποιήστε τις ενημερώσεις παρακολούθησης και βελτιώστε την αποδοτικότητα των logistics σας.",
                            "wallet": "Αναλάβετε τον έλεγχο των οικονομικών σας με ολοκληρωμένη παρακολούθηση εξόδων, αυτοματοποιημένους υπολογισμούς φόρων και ολοκληρωμένη διαχείριση καμπάνιας.",
                            "data_sync": "Διατηρήστε τα επιχειρηματικά σας δεδομένα συγχρονισμένα σε όλες τις πλατφόρμες. Ενσωματώστε απρόσκοπτα πληροφορίες πελατών, δεδομένα προϊόντων και αναλύσεις σε πραγματικό χρόνο.",
                            "team": "Επεκτείνετε τις δυνατότητες της ομάδας σας με εργαλεία συνεργασίας, έλεγχο πρόσβασης βάσει ρόλων και εξορθολογισμένη ροή εργασίας."
                        },
                        "couriers": {
                            "acs": {
                                "title": "ACS Courier",
                                "cta": "Αγορά Extension",
                                "features": {
                                    "status_updates": "Αυτόματες ενημερώσεις κατάστασης παραγγελίας",
                                    "vouchers": "Λήψη vouchers αποστολής",
                                    "notifications": "Αποστολή ειδοποιήσεων στον πελάτη μέσω email"
                                }
                            },
                            "courier4u": {
                                "title": "Courier4U",
                                "cta": "Αγορά Extension",
                                "features": {
                                    "status_updates": "Αυτόματες ενημερώσεις κατάστασης παραγγελίας",
                                    "vouchers": "Λήψη vouchers αποστολής",
                                    "notifications": "Αποστολή ειδοποιήσεων στον πελάτη μέσω email"
                                }
                            }
                        },
                        "wallet": {
                            "title": "Πορτοφόλι & Έξοδα",
                            "cta": "Αγορά Extension",
                            "features": {
                                "facebook": "Ενσωμάτωση με καμπάνιες Facebook",
                                "courier_charges": "Παρακολούθηση χρεώσεων courier",
                                "tax": "Αυτόματοι υπολογισμοί φόρων",
                                "support": "Προτεραιότητα υποστήριξης για χρήστες KonektorX"
                            }
                        },
                        "data_sync": {
                            "title": "Συγχρονισμός Δεδομένων",
                            "cta": "Αγορά Extension",
                            "features": {
                                "customer_data": "Λήψη δεδομένων πελατών χωρίς διακοπή",
                                "product_sync": "Συγχρονισμός πληροφοριών προϊόντων και παραλλαγών",
                                "order_history": "Πρόσβαση σε πλήρες ιστορικό παραγγελιών",
                                "analytics": "Δημιουργία προσαρμοσμένων αναλύσεων και αναφορών"
                            }
                        },
                        "team": {
                            "title": "Διαχείριση Ομάδας",
                            "cta": "Προσθέστε Μέλος Ομάδας στο project σας!",
                            "features": {
                                "unlimited": "Προσθέστε όσους χρήστες θέλετε στο project σας"
                            }
                        }
                    },
                    "settings": {
                        "header": "Ρυθμίσεις Project",
                        "navigation": {
                            "general": "Γενικά",
                            "secrets": "Μυστικά",
                            "webhooks": "Webhooks",
                            "notifications": "Ειδοποιήσεις",
                            "payment": "Πληρωμές & Συνδρομές",
                            "team": "Ομάδα"
                        },
                        "general": {
                            "title": "Γενικά",
                            "description": "Ενημέρωση των δεδομένων του project",
                            "fields": {
                                "name": "Όνομα",
                                "description": "Περιγραφή",
                                "domain": "Domain",
                                "domain_tooltip": "Επικοινωνήστε με την ομάδα υποστήριξης για να αλλάξετε το domain σας."
                            },
                            "buttons": {
                                "update": "Ενημέρωση"
                            },
                            "delete": {
                                "title": "Διαγραφή Project",
                                "warning": "Διαγράφοντας αυτό το project χάνετε όλα σας τα δεδομένα!",
                                "extension_warning": "Βεβαιωθείτε ότι έχετε απενεργοποιήσει όλες τις επεκτάσεις!",
                                "confirmation": "Είστε σίγουροι ότι θέλετε να διαγράψετε αυτό το project;",
                                "button": "Διαγραφή"
                            }
                        },
                        "secrets": {
                            "title": "Μυστικά Project",
                            "reset": {
                                "title": "Επαναφορά Μυστικών",
                                "description": "Ενημερώνοντας τα μυστικά του project θα γίνει επαναφορά όλων των μυστικών.",
                                "consumer_key": "Consumer Key",
                                "consumer_secret": "Consumer Secret",
                                "error": "Αδυναμία σύνδεσης στον server, ελέγξτε τα μυστικά σας και δοκιμάστε ξανά."
                            }
                        },
                        "couriers": {
                            "acs": {
                                "title": "ACS Courier",
                                "description": "Απενεργοποίηση επέκτασης και διαγραφή συνδρομής",
                                "fields": {
                                    "user_id": "ID Χρήστη",
                                    "user_password": "Κωδικός Χρήστη",
                                    "company_id": "ID Εταιρίας",
                                    "company_password": "Κωδικός Εταιρίας",
                                    "billing_code": "Κωδικός Χρέωσης",
                                    "acs_api_key": "ACS API Key"
                                },
                                "printer": {
                                    "label": "Επιλέξτε Τύπο Εκτυπωτή",
                                    "laser": "Laser",
                                    "thermal": "Thermal"
                                }
                            },
                            "courier4u": {
                                "title": "Courier4u",
                                "description": "Απενεργοποίηση επέκτασης και διαγραφή συνδρομής",
                                "fields": {
                                    "api_key": "API Key"
                                }
                            }
                        },
                        "team": {
                            "title": "Ομάδα",
                            "description": "Ο νέος χρήστης δεν μπορεί να προσθέσει νέες επεκτάσεις στο project",
                            "add_member": {
                                "title": "Προσθήκη μέλους ομάδας",
                                "fields": {
                                    "email": "Email",
                                    "name": "Όνομα",
                                    "last_name": "Επώνυμο",
                                    "password": "Κωδικός",
                                    "confirm_password": "Επιβεβαίωση Κωδικού",
                                    "role": {
                                        "label": "Ρόλος",
                                        "user": "Χρήστης",
                                        "admin": "Διαχειριστής"
                                    },
                                    "notifications": "Λήψη Ειδοποιήσεων"
                                }
                            },
                            "members": "Μέλη"
                        },
                        "payment": {
                            "title": "Πληρωμές",
                            "table": {
                                "extension_name": "Όνομα Επέκτασης",
                                "amount": "Ποσό",
                                "status": "Κατάσταση",
                                "created": "Ημ/νία Δημιουργίας",
                                "status_values": {
                                    "paid": "Πληρώθηκε",
                                    "failed": "Απέτυχε",
                                    "pending": "Σε Εκκρεμότητα"
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