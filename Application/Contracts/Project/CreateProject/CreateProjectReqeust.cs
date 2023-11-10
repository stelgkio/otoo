namespace Application.Contracts.ProjectDto.CreateProject
{
    /// <summary>
    /// Create new project
    /// </summary>
    public class CreateProjectReqeust
    {

        /// <summary>
        /// Name of the project
        /// </summary>
        public string Name { get; set; }

        /// <summary>
        /// Description of the project
        /// </summary>
        public string Description { get; set; }

        /// <summary>
        /// Add the web site of the project you want to connect
        /// </summary>
        public string Url { get; set; }

    }
}
