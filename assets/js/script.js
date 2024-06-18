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


