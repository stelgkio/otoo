// js/order_table.js

function orderTable(projectId) {
    return {
        projectID: projectId,
        currentTab: 'all',
        orders: [],
        selectedOrders: [],
        selectAllCheckbox: false,
        sortKey: 'orderId',
        sortAsc: false,
        currentPage: 1,
        itemsPerPage: 10,
        totalItems: 0,
        totalPages: 0,
        loading: false,
        selectedStatus: '',
        errorMessage: '',
        showModal: false,
        modalOrder: {},

        async init() {
            await this.fetchOrders(this.currentPage);
        },

        async fetchOrders(page = 1) {
            this.loading = true;
            try {
                const url = this.getUrlForTab(this.currentTab, page);
                const response = await fetch(url);
                const result = await response.json();
                if (response.ok) {
                    this.orders = result.data || [];
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
            const baseUrl = `${window.location.origin}/order/table/${this.projectID}`;
            const sortDirection = this.sortAsc ? 'asc' : 'desc';
            switch (tab) {
                case 'all':
                    return `${baseUrl}/all/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'completed':
                    return `${baseUrl}/completed/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'processing':
                    return `${baseUrl}/processing/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'pending':
                    return `${baseUrl}/pending/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                case 'cancelled':
                    return `${baseUrl}/cancelled/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
                default:
                    return `${baseUrl}/completed/${page}?sort=${this.sortKey}&direction=${sortDirection}`;
            }
        },

        loadModalData(order) {
            this.modalOrder = order;
            this.showModal = true;
        },

        closeModal() {
            this.showModal = false;
            this.modalOrder = {};
        },

        selectTab(tab) {
            this.currentTab = tab;
            this.currentPage = 1;
            this.fetchOrders(this.currentPage);
        },

        selectAll() {
            this.selectedOrders = this.selectAllCheckbox ? this.orders.map(order => order.orderId) : [];
        },

        sortTable(key) {
            if (this.sortKey === key) {
                this.sortAsc = !this.sortAsc;
            } else {
                this.sortKey = key;
                this.sortAsc = true;
            }
            this.fetchOrders(this.currentPage);
        },

        getSortIcon(key) {
            if (this.sortKey !== key) return '';
            return this.sortAsc ? 'bi bi-chevron-up' : 'bi bi-chevron-down';
        },

        changePage(page) {
            if (page < 1 || page > this.totalPages) return;
            this.fetchOrders(page);
        },

        get paginatedOrders() {
            return this.orders;
        }
    };
}
