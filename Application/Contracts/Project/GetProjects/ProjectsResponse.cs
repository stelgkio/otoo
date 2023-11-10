namespace Application.Contracts.ProjectDto.GetProjects
{
    /// <summary>
    /// Create new Project 
    /// </summary>
    public class ProjectsResponse
    {
        public ProjectsResponse(IList<ProjectDto> projects)
        {
            Projects = projects;
        }
        public IList<ProjectDto> Projects { get; set; }



    }
}
