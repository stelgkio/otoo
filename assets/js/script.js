document.addEventListener('DOMContentLoaded', function () {
    const searchInput = document.getElementById('project-search');
    const projectList = document.getElementById('project-list');
    const projects = projectList.querySelectorAll('[data-project-type="pro"]');

    searchInput.addEventListener('input', function () {
        const filter = searchInput.value.toLowerCase();

        for (let i = 0; i < projects.length; i++) {
            const project = projects[i];
            const projectName = project.getElementsByTagName('a')[0].innerText.toLowerCase();

            if (projectName.includes(filter)) {
                project.setAttribute('style', 'display:inline !important');
            } else {
                project.setAttribute('style', 'display:none !important');
            }
        }
    });
});


 // Bootstrap validation script
 (function () {
    'use strict';

    window.addEventListener('load', function () {
        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        var forms = document.getElementsByClassName('needs-validation');

        // Loop over them and prevent submission
        var validation = Array.prototype.filter.call(forms, function (form) {
            form.addEventListener('submit', function (event) {
                if (form.checkValidity() === false) {
                    event.preventDefault();
                    event.stopPropagation();
                }
                form.classList.add('was-validated');
            }, false);
        });
    }, false);
})();