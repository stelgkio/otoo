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
                    "title": "All in one Solution,Integration & Dashboard Analytics	Faster than Ever.!",
                    "Dashboard": "Dashboard",

                }
            },
            el: {
                translation: {
                    "title": "Όλα σε μία Λύση, Ενσωμάτωση & Αναλυτικά Δεδομένα Πίνακα Ελέγχου Πιο Γρήγορα από Ποτέ.",
                    "Dashboard": "Πίνακας Ελέγχου",

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

// Function to update the content
function updateContent() {
    document.querySelectorAll('[data-i18n]').forEach(function (element) {
        var key = element.getAttribute('data-i18n');
        element.textContent = i18next.t(key);
    });
}

// Listen to HTMX event when content is dynamically updated
document.body.addEventListener('htmx:afterSettle', function (evt) {
    // Update i18n content in the dynamically updated section
    updateContent();
});
