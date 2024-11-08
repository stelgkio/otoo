function createVoucher(projectId) {
    return {
        // Core UI state
        projectId: projectId,  // Store the projectId
        showOffcanvas: false,
        activeTab: 'customer',
        errors: {},
        isSubmitting: false,
        toastMessage: '',
        toastType: 'bg-success',
        selectedCourier: '',

        // Generic voucher data
        voucher_object: {
            orderId: '',
            billing: {
                first_name: '',
                last_name: '',
                email: '',
                phone: '',
                address_1: '',
                city: '',
                postcode: ''
            },
            shipping: {
                first_name: '',
                last_name: '',
                address_1: '',
                city: '',
                postcode: '',
                courier: '',
                deliveryOption: ''
            },
            products: [],
            note: "",
            cod: '',
            hermes_settings: {
                ServiceSavvato: '',
                ServiceEpigon: '',
                ServiceEpistrofi: '',
                ServiceSameday: '',
                ServiceProtocol: '',
                ServiceReception: '',
                ParcelWeight: "1.00",
                ParcelDepth: '',
                ParcelWidth: '',
                ParcelHeight: '',
            },
        },

        // Hermes-specific data


        // Initialize form
        init() {
            console.log('Initializing component with projectId:', this.projectId);
            this.setupValidationWatchers();
            this.initializeBootstrapComponents();
        },

        // Prepare payload for Hermes API
        prepareHermesPayload() {
            return {
                ReceiverName: `${this.voucher_object.shipping.first_name} ${this.voucher_object.shipping.last_name}`,
                ReceiverAddress: this.voucher_object.shipping.address_1,
                ReceiverCity: this.voucher_object.shipping.city,
                ReceiverPostal: parseInt(this.voucher_object.shipping.postcode, 10),
                ReceiverTelephone: this.voucher_object.billing.phone,
                Notes: this.voucher_object.note,
                OrderID: String(this.voucher_object.orderId),
                Cod: parseFloat(this.voucher_object.cod),

                // Hermes specific services
                ServiceSavvato: this.voucher_object.hermes_settings.ServiceSavvato === true ? 1 : null,
                ServiceEpigon: this.voucher_object.hermes_settings.ServiceEpigon === true ? 1 : null,
                ServiceEpistrofi: this.voucher_object.hermes_settings.ServiceEpistrofi === true ? 1 : null,
                ServiceSameday: this.voucher_object.hermes_settings.ServiceSameday === true ? 1 : null,
                ServiceProtocol: this.voucher_object.hermes_settings.ServiceProtocol === true ? 1 : null,
                ServiceReception: this.voucher_object.hermes_settings.ServiceReception === true ? 1 : null,

                // Parcel details
                ParcelWeight: parseFloat(this.voucher_object.hermes_settings.ParcelWeight),
                ParcelDepth: parseFloat(this.voucher_object.hermes_settings.ParcelDepth),
                ParcelWidth: parseFloat(this.voucher_object.hermes_settings.ParcelWidth),
                ParcelHeight: parseFloat(this.voucher_object.hermes_settings.ParcelHeight),

            };
        },

        // Handle form submission
        async handleSubmit() {
            if (!this.validateForm()) {
                this.showToast('Please check the form for errors', 'bg-danger');
                return;
            }

            // Start the submission process
            this.isSubmitting = true;
            try {
                if (this.selectedCourier === 'courier4u') {
                    const payload = this.prepareHermesPayload();

                    // Log the payload for debugging
                    console.log('Full Payload:', payload);


                    const response = await fetch(`/voucher/courier4u/create/${this.projectID}`, {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(payload) // Use the payload directly
                    });

                    if (!response.ok) {
                        const errorData = await response.json();
                        console.error('Response Error:', errorData);
                        throw new Error('Failed to create voucher: ' + errorData.message || response.statusText);
                    }



                    this.closeOffcanvas();
                }
            } catch (error) {
                console.error('Error creating voucher:', error);
                this.showToast('Failed to create voucher', 'bg-danger');
            } finally {
                this.isSubmitting = false;
            }
        },

        // Show the order offcanvas and populate data if voucher exists
        openOffcanvas(voucher) {
            console.log('Opening offcanvas with voucher:', voucher);
            console.log('Project ID:', this.projectId);

            // Ensure products is an array in voucher data
            if (!Array.isArray(voucher.products)) {
                voucher.products = [];
            }

            // Merge voucher data into voucher_object
            this.voucher_object = {
                ...this.voucher_object,
                orderId: voucher.orderId,
                billing: { ...this.voucher_object.billing, ...voucher.billing },
                shipping: { ...this.voucher_object.shipping, ...voucher.shipping },
                products: [...voucher.products],
                note: voucher.note,
                cod: voucher.total_amount,
                hermes_settings: {
                    ServiceSavvato: "",
                    ServiceEpigon: "",
                    ServiceEpistrofi: "",
                    ServiceSameday: "",
                    ServiceProtocol: "",
                    ServiceReception: "",
                    ParcelWeight: "1.00",
                    ParcelDepth: "",
                    ParcelWidth: "",
                    ParcelHeight: "",
                },
            };

            this.selectedCourier = voucher.shipping?.courier || '';
            this.showOffcanvas = true;
            this.activeTab = 'customer';
            this.errors = {};

            // Initialize Bootstrap offcanvas component if it doesn't exist
            if (!this.offcanvas) {
                this.initializeBootstrapComponents();
            }

            // Show the offcanvas component
            if (this.offcanvas) {
                this.offcanvas.show();
            }
        },

        // Initialize and configure Bootstrap components
        initializeBootstrapComponents() {
            try {
                const offcanvasElement = document.getElementById('orderProcessingOffcanvas');
                if (offcanvasElement && typeof bootstrap !== 'undefined') {
                    // Initialize the offcanvas only if it hasn't been initialized
                    if (!this.offcanvas) {
                        this.offcanvas = new bootstrap.Offcanvas(offcanvasElement, {
                            backdrop: true,
                            keyboard: true
                        });

                        // Set up an event listener to reset the form when offcanvas is closed
                        offcanvasElement.addEventListener('hidden.bs.offcanvas', () => {
                            this.showOffcanvas = false;
                            this.resetForm();
                        });
                    }
                }
            } catch (error) {
                console.error('Error initializing Bootstrap components:', error);
            }
        },

        // Set active tab in the form
        setActiveTab(tab) {
            console.log('Setting active tab:', tab);
            this.activeTab = tab;
        },



        // Set up watchers for validation of form fields
        setupValidationWatchers() {
            try {
                // Watch changes in billing fields
                ['first_name', 'last_name', 'email', 'phone', 'address_1', 'city', 'postcode'].forEach(field => {
                    this.$watch(`voucher_object.billing.${field}`, () => {
                        this.validateField(`billing.${field}`);
                    });
                });

                // Watch changes in shipping fields
                ['first_name', 'last_name', 'address_1', 'city', 'postcode'].forEach(field => {
                    this.$watch(`voucher_object.shipping.${field}`, () => {
                        this.validateField(`shipping.${field}`);
                    });
                });

                // Watch changes in shipping-specific fields
                this.$watch('selectedCourier', () => {
                    this.voucher_object.shipping.courier = this.selectedCourier;
                    this.validateField('shipping.courier');
                });

                this.$watch('voucher_object.hermes_settings.ParcelWeight', () => {
                    this.validateField('voucher_object.hermes_settings.ParcelWeight');
                });

            } catch (error) {
                console.error('Error setting up validation watchers:', error);
            }
        },

        // Validate individual fields with custom rules
        validateField(field) {
            console.log('Validating field:', field);

            // Get the value of the field to validate
            var value = field.split('.').reduce((obj, key) => obj?.[key], this.voucher_object);

            console.log('value :', value);
            delete this.errors[field];  // Clear previous errors

            const optionalFields = ['shipping.deliveryOption'];
            if (!value && optionalFields.includes(field)) return true;  // Skip validation for optional fields


            // Basic required field validation
            if (!value?.toString().trim()) {
                this.errors[field] = 'This field is required';
                return false;
            }

            // Specific validations based on field type
            switch (true) {
                case field.endsWith('email'):
                    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
                    if (!emailRegex.test(value)) {
                        this.errors[field] = 'Please enter a valid email address';
                        return false;
                    }
                    break;

                case field.endsWith('phone'):
                    const phoneRegex = /^\+?[\d\s-]{10,}$/;
                    if (!phoneRegex.test(value)) {
                        this.errors[field] = 'Please enter a valid phone number';
                        return false;
                    }
                    break;

                case field.endsWith('postcode'):
                    if (value.length < 5) {
                        this.errors[field] = 'Please enter a valid postcode';
                        return false;
                    }
                    break;

                case field.endsWith('ParcelWeight'):
                    console.log('Validating field tttt:', field);
                    console.log('Validatinvalue:', value);
                    const weight = parseFloat(value);
                    if (isNaN(weight) || weight < 1) {
                        this.errors[field] = 'Please enter a valid weight greater than 0';
                        return false;
                    } else {
                        this.errors[field] = null; // Clear the error explicitly by setting to null
                    }
                    break;

                case field === 'cod':
                    const codAmount = parseFloat(value);
                    if (isNaN(codAmount)) {
                        this.errors[field] = 'Please enter a valid COD amount';
                        return false;
                    }
                    if (codAmount <= 0.00) {
                        this.errors[field] = 'COD amount cannot be negative';
                        return false;
                    }
                    if (codAmount > 499.99) {
                        this.errors[field] = 'COD amount must be less than or equal to 499.99';
                        return false;
                    }
                    break;
            }
            return true;
        },
        validateParcelWeight(weightValue) {
            // Ensure that 'errors' is initialized as an empty object if not defined
            this.errors = this.errors || {};

            // Parse the input value as a float
            const weight = parseFloat(weightValue);

            // Check if the weight is valid
            if (isNaN(weight) || weight <= 0) {
                // Add an error message to the errors object if invalid
                this.errors['hermes_settings.ParcelWeight'] = 'Please enter a valid weight greater than 0';
                return false;
            } else {
                // Clear the error message explicitly by setting it to an empty string
                this.errors['hermes_settings.ParcelWeight'] = '';
                return true;
            }
        },
        // Validate the entire form before submission
        validateForm() {
            const fieldsToValidate = [
                'billing.first_name', 'billing.last_name', 'billing.email', 'billing.phone',
                'billing.address_1', 'billing.city', 'billing.postcode',
                'shipping.first_name', 'shipping.last_name', 'shipping.address_1',
                'shipping.city', 'shipping.postcode', 'shipping.courier',
                'shipping.deliveryOption', 'cod', 'hermes_settings.ParcelWeight'
            ];
            let valid = true;

            fieldsToValidate.forEach(field => {
                valid = this.validateField(field) && valid;  // Combine validation results
            });

            return valid;  // Return overall validation status
        },

        // Reset the form fields and errors
        resetForm() {
            console.log('Resetting form fields');
            this.voucher_object = {
                ...this.voucher_object,
                orderId: '',
                billing: {
                    first_name: '',
                    last_name: '',
                    email: '',
                    phone: '',
                    address_1: '',
                    city: '',
                    postcode: ''
                },
                shipping: {
                    first_name: '',
                    last_name: '',
                    address_1: '',
                    city: '',
                    postcode: '',
                    courier: '',
                    deliveryOption: ''
                },
                products: [],
                note: "",
                cod: ''
            };
            this.errors = {};  // Clear all errors
            this.selectedCourier = '';  // Reset selected courier
        },

        // Display toast notifications
        showToast(message, type) {
            this.toastMessage = message;
            this.toastType = type;

            const toastElement = document.getElementById('toast');
            if (toastElement) {
                const toast = new bootstrap.Toast(toastElement);
                toast.show();  // Show the toast
            }
        },

        // Close the order offcanvas
        closeOffcanvas() {
            if (this.offcanvas) {
                this.offcanvas.hide();
            }
        },
    };
}

