var CurrentTab = "new";
var currentPage = 1;




function voucherTable(projectId) {
    return {
        projectID: projectId,
        currentTab: 'new',
        vouchers: [],
        selectedVouchers: [],
        selectAllCheckbox: false,
        sortKey: 'created_at',
        sortAsc: false,
        currentPage: 1,
        itemsPerPage: 10,
        totalItems: 0,
        totalPages: 0,
        loading: false,
        selectedMultipleOption: '',
        errorMessage: '',
        isNewTab: true,
        isPrinted: false,
        isDownloading: false,


        async init() {
            await this.fetchVouchers(this.currentPage);
        },
        async fetchVouchers(page = 1) {
            this.loading = true;
            try {
                const url = this.getUrlForTab(this.currentTab, page);
                const response = await fetch(url);
                const result = await response.json();
                if (response.ok) {
                    this.vouchers = result.data || [];
                    this.totalItems = result.meta.totalItems || 0;
                    this.currentPage = result.meta.currentPage || 1;
                    this.itemsPerPage = result.meta.itemsPerPage || 10;
                    this.totalPages = result.meta.totalPages || 0;
                } else {
                    console.error('Error fetching data:', result.message);
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            } finally {
                this.loading = false;
            }
        },
        getUrlForTab(tab, page) {
            const baseUrl = `${window.location.origin}/voucher/table/${this.projectID}`;
            const sortDirection = this.sortAsc ? 'asc' : 'desc';
            switch (tab) {
                case 'all':
                    return `${baseUrl}/all/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'pending':
                    return `${baseUrl}/pending/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'completed':
                    return `${baseUrl}/completed/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'processing':
                    return `${baseUrl}/processing/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'cancelled':
                    return `${baseUrl}/cancelled/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'new':
                    return `${baseUrl}/new/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                default:
                    return `${baseUrl}/all/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
            }
        },

        selectTab(tab) {
            if (tab != 'new') {
                this.isNewTab = false;
                this.isPrinted = true;
            } else {
                this.isNewTab = true;
                this.isPrinted = false;
            }
            this.currentTab = tab;
            this.currentPage = 1;
            CurrentTab = this.currentTab
            CurrentPage = this.currentPage
            this.fetchVouchers(this.currentPage);
        },

        selectAll() {
            this.selectedVouchers = this.selectAllCheckbox ? this.vouchers.map(voucher => voucher.Id) : [];
        },

        sortTable(key) {
            if (this.sortKey === key) {
                this.sortAsc = !this.sortAsc;
            } else {
                this.sortKey = key;
                this.sortAsc = true;
            }
            this.fetchVouchers(this.currentPage);
        },
        getSortIcon(key) {
            if (this.sortKey !== key) return '';
            return this.sortAsc ? 'bi bi-chevron-up' : 'bi bi-chevron-down';
        },

        changePage(page) {
            if (page < 1 || page > this.totalPages) return;
            this.fetchVouchers(page);
        },

        async downloadVoucher(voucherId, courier_provider) {
            // Ask for confirmation before proceeding
            const userConfirmed = window.confirm("Are you sure you want to download this voucher? If you download this voucher, it will be marked as printed and cannot be edited.");
            if (!userConfirmed) {
                console.log("Voucher download cancelled by user.");
                return; // Exit the function if user cancels
            }
            isDownloading = true
            var url = '';
            if (courier_provider == 'courier4u') {
                url = `/voucher/courier4u/donwload/${voucherId}/${this.projectID}`;
            } else if (courier_provider == 'acs-courier') {
                url = `/voucher/acscourier/download/${voucherId}/${this.projectID}`;
            } else {
                url = `/voucher/redcourier/download/${voucherId}/${this.projectID}`;

            }
            console.log(`Downloading voucher with ID: ${voucherId}`);
            // Assuming you have a function to fetch the data from your endpoint
            await fetch(url)
                .then(response => {
                    if (response.ok) {
                        return response.json(); // Assuming response has { filename, base64 }
                    } else {
                        throw new Error("Network response was not OK");
                    }
                })
                .then(data => {
                    // Extract filename and base64 content
                    const { filename, base64 } = data;

                    // Ensure there are no data URI prefixes
                    const cleanedBase64 = base64.split(',').pop();

                    // Decode base64 to binary string
                    let binaryString;
                    try {
                        binaryString = atob(cleanedBase64);
                    } catch (error) {
                        console.error("Failed to decode base64 data:", error);
                        return;
                    }

                    // Convert binary string to Uint8Array
                    const byteArray = new Uint8Array(binaryString.length);
                    for (let i = 0; i < binaryString.length; i++) {
                        byteArray[i] = binaryString.charCodeAt(i);
                    }

                    // Create a Blob with 'application/pdf' MIME type
                    const blob = new Blob([byteArray], { type: 'application/pdf' });

                    // Create an anchor element to trigger download
                    const link = document.createElement('a');
                    link.href = window.URL.createObjectURL(blob);
                    link.download = filename || "downloaded-file.pdf";
                    link.click();

                    // Clean up URL object
                    window.URL.revokeObjectURL(link.href);
                })
                .catch(error => {
                    console.error("There was a problem with the fetch operation:", error);
                    isDownloading = false
                });
            isDownloading = false
            await this.fetchVouchers(page);

        },

        get paginatedVouchers() {
            return this.vouchers;
        },
        get currentPageStart() {
            return (this.currentPage - 1) * this.itemsPerPage + 1;
        },

        get currentPageEnd() {
            return Math.min(this.currentPage * this.itemsPerPage, this.totalItems);
        },
        get pageNumbers() {
            const range = 2; // Number of pages to show around the current page
            let start = Math.max(1, this.currentPage - range);
            let end = Math.min(this.totalPages, this.currentPage + range);

            // Adjust range if there are not enough pages on one side
            if (this.totalPages - end < range) {
                end = this.totalPages;
                start = Math.max(1, end - 2 * range);
            } else if (start <= range) {
                start = 1;
                end = Math.min(this.totalPages, start + 2 * range);
            }

            return Array.from({ length: end - start + 1 }, (_, i) => start + i);
        },

        badgeClass(status) {
            const baseClass = 'badge bg-body-secondary badge-custom';
            switch (status) {
                case 'pending':
                    return `${baseClass} text-info`;
                case 'new':
                    return `${baseClass} text-primary`;
                case 'completed':
                    return `${baseClass} text-success`;
                case 'cancelled':
                    return `${baseClass} text-danger`;
                case 'processing':
                    return `${baseClass} text-warning`;
                default:
                    return baseClass;
            }
        },

        async bulkAction() {
            if (!this.selectedMultipleOption || !this.selectedVouchers.length) {
                this.errorMessage = 'Please select a option and at least one voucher.';
                return;
            }
            this.loading = true;
            try {

                const selectedVoucherObjects = this.selectedVouchers.map(selectedId =>
                    this.vouchers.find(voucher => voucher.Id === selectedId)
                );

                // Filter vouchers based on the courier provider
                const courier4uVouchers = selectedVoucherObjects.filter(voucher => voucher.courier_provider === 'courier4u');
                const redcourierVouchers = selectedVoucherObjects.filter(voucher => voucher.courier_provider === 'redcourier');
                const acsCourierVouchers = selectedVoucherObjects.filter(voucher => voucher.courier_provider === 'acs-courier');

                // Function to handle the request for each provider and download the file
                const downloadFile = async (voucherIds, provider) => {
                    const url = `${window.location.origin}/voucher/${provider}/download-multiple/${this.projectID}`;
                    await fetch(url, {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify({
                            voucherIds,
                        })
                    })
                        .then(response => {
                            if (response.ok) {
                                return response.json(); // Assuming response has { filename, base64 }
                            } else {
                                throw new Error("Network response was not OK");
                            }
                        })
                        .then(data => {
                            // Extract filename and base64 content
                            const { filename, base64 } = data;

                            // Ensure there are no data URI prefixes
                            const cleanedBase64 = base64.split(',').pop();

                            // Decode base64 to binary string
                            let binaryString;
                            try {
                                binaryString = atob(cleanedBase64);
                            } catch (error) {
                                console.error("Failed to decode base64 data:", error);
                                return;
                            }

                            // Convert binary string to Uint8Array
                            const byteArray = new Uint8Array(binaryString.length);
                            for (let i = 0; i < binaryString.length; i++) {
                                byteArray[i] = binaryString.charCodeAt(i);
                            }

                            // Create a Blob with 'application/pdf' MIME type
                            const blob = new Blob([byteArray], { type: 'application/pdf' });

                            // Create an anchor element to trigger download
                            const link = document.createElement('a');
                            link.href = window.URL.createObjectURL(blob);
                            link.download = filename || "downloaded-file.pdf";
                            link.click();

                            // Clean up URL object
                            window.URL.revokeObjectURL(link.href);
                        })
                        .catch(error => {
                            console.error("There was a problem with the fetch operation:", error);
                            isDownloading = false
                        })
                        .finally(() => {
                            isDownloading = false;
                        });
                };




                // Send download requests for each provider
                // Prepare downloadPromises array only for non-empty voucher lists
                const downloadPromises = [];

                if (courier4uVouchers.length > 0) {
                    downloadPromises.push(downloadFile(courier4uVouchers, 'courier4u'));
                }
                if (redcourierVouchers.length > 0) {
                    downloadPromises.push(downloadFile(redcourierVouchers, 'redcourier'));
                }
                if (acsCourierVouchers.length > 0) {
                    downloadPromises.push(downloadFile(acsCourierVouchers, 'acs-courier'));
                }
                // Wait for all downloads to finish
                await Promise.all(downloadPromises);

                // Reset form after downloads
                this.selectedVouchers = [];
                this.selectedMultipleOption = '';
                this.selectAllCheckbox = false;
                this.fetchVouchers(this.currentPage);

            } catch (error) {
                this.errorMessage = `An error occurred: ${error.message}`;
                console.error(error);
            } finally {
                this.loading = false;
            }
        },



    };
}


function createVoucher(projectId) {
    return {
        // Core UI state
        projectId: projectId,  // Store the projectId
        showOffcanvas: false,
        activeTab: 'customer',
        errors: {},
        isSubmitting: false,
        toastMessage: '',
        toastMessageSuuccess: '',
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
        createVoucherInit() {
            console.log('Initializing component with projectId:', this.projectId);
            this.setupValidationWatchers();
            this.initializeBootstrapComponents();
            this.hideToast();
        },
        triggerUpdateVoucherEvent() {
            // Dispatch the event to the parent component
            this.$dispatch('update-voucher');
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
                ParcelWeight: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelWeight)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelWeight),

                ParcelDepth: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelDepth)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelDepth),
                ParcelWidth: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelWidth)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelWidth),
                ParcelHeight: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelHeight)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelHeight),
                CustomOrderId: false
            };
        },

        // Handle form submission
        async handleSubmit() {
            this.hideToast()
            if (!this.validateForm()) {
                this.showToast('Please check the form for errors', 'bg-danger');
                return;
            }

            // Start the submission process
            this.isSubmitting = true;
            try {
                if (this.selectedCourier === 'courier4u') {
                    await this.createCourier4uVoucher()
                }
                if (this.selectedCourier === 'redcourier') {
                    await this.createRedCourierVoucher()
                }
            } catch (error) {
                console.error('Error creating voucher:', error);
                this.showToast(error, 'bg-danger');
            } finally {
                this.isSubmitting = false;
            }
        },
        hideToast() {
            this.toastMessage = '';
            this.toastMessageSuccess = '';
        },
        async createCourier4uVoucher() {
            const payload = this.prepareHermesPayload();

            // Log the payload for debugging
            console.log('Full Payload (Courier4u):', payload);


            const response = await fetch(`/voucher/courier4u/create/${this.projectID}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload) // Use the payload directly
            });

            if (!response.ok) {
                const errorData = await response.json();
                console.error('Response Error (Courier4u):', errorData);
                this.showToast(errorData.error, 'bg-danger');
                throw new Error('Failed to create voucher: ' + errorData.error);
            }

            this.toastMessageSuccess = 'Courier4u voucher created successfully!';
            this.triggerUpdateVoucherEvent()
            this.closeOffcanvas();


        },

        async createRedcourierVoucher() {
            const payload = this.prepareHermesPayload();

            // Log the payload for debugging
            console.log('Full Payload (Redcourier):', payload);


            const response = await fetch(`/voucher/redcourier/create/${this.projectID}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload) // Use the payload directly
            });

            if (!response.ok) {
                const errorData = await response.json();
                console.error('Response Error (Redcourier):', errorData);
                this.showToast(errorData.error, 'bg-danger');
                throw new Error('Failed to create voucher: ' + errorData.error);
            }

            this.toastMessageSuccess = 'Redcourier voucher created successfully!';
            this.triggerUpdateVoucherEvent()
            this.closeOffcanvas();


        },
        // Show the order offcanvas and populate data if voucher exists
        openOffcanvas(voucher) {
            console.log('Opening offcanvas with voucher:', voucher);
            console.log('Project ID:', this.projectId);
            this.hideToast()
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
                    if (value.length < 5 || value.length > 5) {
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
                }
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


function newVoucher(projectId) {
    return {
        // Core UI state
        projectId: projectId,  // Store the projectId
        showOffcanvas: false,
        activeTab: 'customer',
        errors: {},
        isSubmitting: false,
        toastMessage: '',
        toastMessageSuuccess: '',
        toastType: 'bg-success',
        selectedCourier: '',
        offcanvasInstances: {},
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
        hasInteracted: false,
        // Hermes-specific data

        markInteracted() {
            if (!this.hasInteracted) this.hasInteracted = true;
        },
        // Initialize form
        newVoucherInit() {
            console.log('Initializing component with projectId:', this.projectId);
            this.setupValidationWatchers();
            this.initializeBootstrapComponents('newVoucherInit');
            this.hideToast();
        },
        triggerUpdateVoucherEvent() {
            // Dispatch the event to the parent component
            this.$dispatch('update-voucher');
        },
        formatOrderIdWithLeadingZeros() {
            // Convert orderId to a string and pad with five leading zeros
            const orderIdString = String(this.voucher_object.orderId);
            this.voucher_object.orderId = orderIdString + '000000';
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
                ParcelWeight: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelWeight)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelWeight),

                ParcelDepth: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelDepth)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelDepth),
                ParcelWidth: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelWidth)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelWidth),
                ParcelHeight: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelHeight)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelHeight),
                CustomOrderId: true,
                ReceiverEmail: this.voucher_object.billing.email,
            };
        },

        // Handle form submission
        async handleSubmit() {
            this.hideToast()
            this.hasInteracted = true;
            if (!this.validateForm()) {
                this.showToast('Please check the form for errors', 'bg-danger');
                return;
            }

            // Start the submission process
            this.isSubmitting = true;
            try {
                if (this.selectedCourier === 'courier4u') {
                    await this.createCourier4uVoucher()
                }
                if (this.selectedCourier === 'redcourier') {
                    await this.createRedCourierVoucher()
                }
            } catch (error) {
                console.error('Error creating voucher:', error);
                this.showToast(error, 'bg-danger');
            } finally {
                this.isSubmitting = false;
            }
        },
        hideToast() {
            this.toastMessage = '';
            this.toastMessageSuccess = '';
        },
        async createCourier4uVoucher() {
            const payload = this.prepareHermesPayload();

            // Log the payload for debugging
            console.log('Full Payload (Courier4u):', payload);


            const response = await fetch(`/voucher/courier4u/create/${this.projectID}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload) // Use the payload directly
            });

            if (!response.ok) {
                const errorData = await response.json();
                console.error('Response Error (Courier4u):', errorData);
                this.showToast(errorData.error, 'bg-danger');
                throw new Error('Failed to create voucher: ' + errorData.error);
            }

            this.toastMessageSuccess = 'Courier4u voucher created successfully!';
            this.triggerUpdateVoucherEvent()
            this.closeOffcanvas();


        },

        async createRedcourierVoucher() {
            const payload = this.prepareHermesPayload();

            // Log the payload for debugging
            console.log('Full Payload (Redcourier):', payload);


            const response = await fetch(`/voucher/redcourier/create/${this.projectID}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload) // Use the payload directly
            });

            if (!response.ok) {
                const errorData = await response.json();
                console.error('Response Error (Redcourier):', errorData);
                this.showToast(errorData.error, 'bg-danger');
                throw new Error('Failed to create voucher: ' + errorData.error);
            }

            this.toastMessageSuccess = 'Redcourier voucher created successfully!';
            this.triggerUpdateVoucherEvent()
            this.closeOffcanvas();


        },
        // Show the order offcanvas and populate data if voucher exists
        openNewVoucherOffcanva(voucher) {
            console.log('Opening offcanvas with voucher:', voucher);
            console.log('Project ID:', this.projectId);
            this.hideToast()
            // Ensure products is an array in voucher dif (!Array.isArray(vata


            // Merge voucher data into voucher_object
            this.voucher_object = mapToNewVoucherObject(null)

            this.selectedCourier = voucher?.shipping?.courier || '';
            this.showOffcanvas = true;
            this.activeTab = 'customer';
            this.errors = {};

            // Initialize Bootstrap offcanvas component if it doesn't exist
            if (!this.offcanvasInstances["newVoucherInit"]) {
                this.initializeBootstrapComponents('newVoucherInit');
            }

            // Show the offcanvas component
            if (this.offcanvasInstances["newVoucherInit"]) {
                this.offcanvasInstances["newVoucherInit"].show();
            }
        },

        // Initialize and configure Bootstrap components

        initializeBootstrapComponents(offcanvasId) {
            try {
                const offcanvasElement = document.getElementById("newVoucherOffcanvas");

                // Initialize the offcanvas only if it hasn't been initialized
                if (!this.offcanvasInstances[offcanvasId]) {
                    this.offcanvasInstances[offcanvasId] = new bootstrap.Offcanvas(offcanvasElement, {
                        backdrop: true,
                        keyboard: true
                    });

                    // Set up an event listener to reset the form when offcanvas is closed
                    offcanvasElement.addEventListener('hidden.bs.offcanvas', () => {
                        this.offcanvasInstances["newVoucherInit"] = false;
                        this.resetForm();
                    });
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
                        if (this.hasInteracted) {
                            this.validateField(`billing.${field}`);
                        }
                    });
                });

                // Watch changes in shipping fields
                ['first_name', 'last_name', 'address_1', 'city', 'postcode'].forEach(field => {
                    this.$watch(`voucher_object.shipping.${field}`, () => {
                        if (this.hasInteracted) {
                            this.validateField(`shipping.${field}`);
                        }
                    });
                });

                // Watch changes in shipping-specific fields
                this.$watch('selectedCourier', () => {
                    this.voucher_object.shipping.courier = this.selectedCourier;
                    if (this.hasInteracted) {
                        this.validateField('shipping.courier');
                    }
                });

                this.$watch('voucher_object.hermes_settings.ParcelWeight', () => {
                    if (this.hasInteracted) {
                        this.validateField('voucher_object.hermes_settings.ParcelWeight');
                    }
                });

            } catch (error) {
                console.error('Error setting up validation watchers:', error);
            }
        },

        // Validate individual fields with custom rules
        validateField(field) {
            if (!this.hasInteracted) return true;
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
                    if (value.length < 5 || value.length > 5) {
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
            if (!this.hasInteracted) return true;
            const fieldsToValidate = [
                'billing.email', 'billing.phone',

                'shipping.first_name', 'shipping.last_name', 'shipping.address_1',
                'shipping.city', 'shipping.postcode', 'shipping.courier',
                'cod', 'hermes_settings.ParcelWeight'
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
                }
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
            this.hasInteracted = false
            if (this.offcanvasInstances["newVoucherInit"]) {
                this.offcanvasInstances["newVoucherInit"].hide();
            }
        },
    };
}


function updateHermeVoucher(projectId) {
    return {
        // Core UI state
        projectId: projectId,  // Store the projectId
        showOffcanvas: false,
        activeTab: 'customer',
        errors: {},
        isSubmitting: false,
        toastMessage: '',
        toastMessageSuuccess: '',
        toastType: 'bg-success',
        selectedCourier: '',
        offcanvasInstances: {},
        isPrinted: false,
        // Generic voucher data
        voucher_object: {
            orderId: '',
            voucherId: '',
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
        UpdateHermeVoucherInit() {
            console.log('Initializing component with projectId:', this.projectId);
            this.setupValidationWatchers();
            this.initializeBootstrapComponents('UpdateHermeVoucherInit');
            this.hideToast();
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
                ParcelWeight: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelWeight)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelWeight),

                ParcelDepth: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelDepth)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelDepth),
                ParcelWidth: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelWidth)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelWidth),
                ParcelHeight: isNaN(parseFloat(this.voucher_object.hermes_settings.ParcelHeight)) ? 1.00 : parseFloat(this.voucher_object.hermes_settings.ParcelHeight),
                CustomOrderId: true
            };
        },

        // Handle form submission
        async handleSubmit() {
            this.hideToast()
            if (!this.validateForm()) {
                this.showToast('Please check the form for errors', 'bg-danger');
                return;
            }

            // Start the submission process
            this.isSubmitting = true;
            try {
                if (this.selectedCourier === 'courier4u') {
                    await this.updateCourier4uVoucher()

                }
                if (this.selectedCourier === 'redcourier') {
                    await this.updateRedcourierVoucher()
                }
            } catch (error) {
                console.error('Error creating voucher:', error);
                this.showToast(error, 'bg-danger');
            } finally {
                this.isSubmitting = false;
            }
        },
        hideToast() {
            this.toastMessage = '';
            this.toastMessageSuccess = '';
        },
        triggerUpdateVoucherEvent() {
            // Dispatch the event to the parent component
            this.$dispatch('update-voucher');
        },
        async updateCourier4uVoucher() {
            const payload = this.prepareHermesPayload();

            // Log the payload for debugging
            console.log('Full Payload (Courier4u):', payload);


            const response = await fetch(`/voucher/courier4u/update/${this.voucher_object.voucherId}/${this.projectID}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload) // Use the payload directly
            });

            if (!response.ok) {
                const errorData = await response.json();
                console.error('Response Error (Courier4u):', errorData);
                this.showToast(errorData.error, 'bg-danger');
                throw new Error('Failed to update voucher: ' + errorData.error);
            }

            this.toastMessageSuccess = 'Courier4u voucher updated successfully!';
            // Check if $parent is defined before calling fetchVouchers

            this.triggerUpdateVoucherEvent()
            this.closeOffcanvas("UpdateHermeVoucherInit");


        },

        async updateRedcourierVoucher() {
            const payload = this.prepareHermesPayload();

            // Log the payload for debugging
            console.log('Full Payload (Redcourier):', payload);


            const response = await fetch(`/voucher/redcourier/update/${this.voucher_object.voucherId}/${this.projectID}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload) // Use the payload directly
            });

            if (!response.ok) {
                const errorData = await response.json();
                console.error('Response Error (Redcourier):', errorData);
                this.showToast(errorData.error, 'bg-danger');
                throw new Error('Failed to update voucher: ' + errorData.error);
            }

            this.toastMessageSuccess = 'Redcourier voucher updated successfully!';
            triggerUpdateVoucherEvent()
            this.closeOffcanvas("UpdateHermeVoucherInit");


        },
        // Show the order offcanvas and populate data if voucher exists
        openUpdateHemerOffCamva(voucher) {
            console.log('Opening Hermer Update offcanvas with voucher:', voucher);
            console.log('Project ID:', this.projectId);
            this.hideToast()
            // Ensure products is an array in voucher data
            if (!Array.isArray(voucher.products)) {
                voucher.products = [];
            }

            // Merge voucher data into voucher_object
            this.voucher_object = mapToUpdateVoucherObject(voucher);


            this.selectedCourier = voucher.courier_provider || '';
            this.showOffcanvas = true;
            this.activeTab = 'customer';
            this.errors = {};
            this.isPrinted = voucher.is_printed || false;

            // Initialize Bootstrap offcanvas component if it doesn't exist
            if (!this.offcanvasInstances["UpdateHermeVoucherInit"]) {
                this.initializeBootstrapComponents("UpdateHermeVoucherInit");
            }

            // Show the offcanvas component
            if (this.offcanvasInstances["UpdateHermeVoucherInit"]) {
                this.offcanvasInstances["UpdateHermeVoucherInit"].show();
            }
        },



        initializeBootstrapComponents(offcanvasId) {
            try {
                const offcanvasElement = document.getElementById("openUpdateHemerOffCamva");

                // Initialize the offcanvas only if it hasn't been initialized
                if (!this.offcanvasInstances[offcanvasId]) {
                    this.offcanvasInstances[offcanvasId] = new bootstrap.Offcanvas(offcanvasElement, {
                        backdrop: true,
                        keyboard: true
                    });

                    // Set up an event listener to reset the form when offcanvas is closed
                    offcanvasElement.addEventListener('hidden.bs.offcanvas', () => {
                        this.showOffcanvas = false;
                        this.resetForm();
                    });
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
                    // Check if the value is exactly 5 characters long and only contains digits
                    if (!/^\d{5}$/.test(value)) {
                        this.errors[field] = 'Please enter a valid 5-digit postcode containing only numbers';
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
                }

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
        closeOffcanvas(offcanvasId) {
            try {
                const offcanvasInstance = this.offcanvasInstances[offcanvasId];
                if (offcanvasInstance) {
                    // Close the offcanvas programmatically
                    offcanvasInstance.hide();
                } else {
                    console.error(`Offcanvas with ID ${offcanvasId} is not initialized.`);
                }
            } catch (error) {
                console.error('Error closing offcanvas:', error);
            }
        },


    };
}


function mapToUpdateVoucherObject(data) {
    return {
        orderId: data.orderId || '',
        voucherId: data.voucherId || '',
        billing: {
            first_name: data.billing?.first_name || '',
            last_name: data.billing?.last_name || '',
            email: data.billing?.email || '',
            phone: data.billing?.phone || '',
            address_1: data.billing?.address_1 || '',
            city: data.billing?.city || '',
            postcode: data.billing?.postcode || ''
        },
        shipping: {
            first_name: data.shipping?.first_name || '',
            last_name: data.shipping?.last_name || '',
            address_1: data.shipping?.address_1 || '',
            city: data.shipping?.city || '',
            postcode: data.shipping?.postcode || '',

        },
        products: data.products?.map(product => ({
            id: product.id || '',
            name: product.name || '',
            product_id: product.product_id || '',
            quantity: product.quantity || 0,
            subtotal: product.subtotal || '',
            subtotal_tax: product.subtotal_tax || '',
            total: product.total || '',
            total_tax: product.total_tax || '',
            price: product.price || '',
            meta_data: product.meta_data || []
        })) || [],
        note: data.note || '',
        cod: data.total_amount || '',
        hermes_settings: {
            ServiceSavvato: data.hermes_courier?.ServiceSavvato === 1 || false,
            ServiceEpigon: data.hermes_courier?.ServiceEpigon === 1 || false,
            ServiceEpistrofi: data.hermes_courier?.ServiceEpistrofi === 1 || false,
            ServiceSameday: data.hermes_courier?.ServiceSameday === 1 || false,
            ServiceProtocol: data.hermes_courier?.ServiceProtocol === 1 || false,
            ServiceReception: data.hermes_courier?.ServiceReception === 1 || false,
            ParcelWeight: data.hermes_courier?.ParcelWeight?.toFixed(2) || "1.00",
            ParcelDepth: data.hermes_courier?.ParcelDepth || '',
            ParcelWidth: data.hermes_courier?.ParcelWidth || '',
            ParcelHeight: data.hermes_courier?.ParcelHeight || ''
        }
    };
}

function mapToNewVoucherObject(data) {
    return {
        orderId: '',
        voucherId: '',
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

        },
        products: [],
        note: '',
        cod: '',
        hermes_settings: {
            ServiceSavvato: false,
            ServiceEpigon: false,
            ServiceEpistrofi: false,
            ServiceSameday: false,
            ServiceProtocol: false,
            ServiceReception: false,
            ParcelWeight: "1.00",
            ParcelDepth: '',
            ParcelWidth: '',
            ParcelHeight: ''
        }
    };
}

