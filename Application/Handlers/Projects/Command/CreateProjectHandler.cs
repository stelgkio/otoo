using Application.Contracts.ProjectDto.GetProjects;
using Domain.Models;
using Infrastracture.Services.Repositories;
using MediatR;

namespace Application.Handlers.Projects.Command
{

    public class CreateProjectHandler : IRequestHandler<CreateProjectCommand, ProjectDto>
    {
        private readonly IProjectRepository projectRepository;
        public CreateProjectHandler(IProjectRepository projectRepository) {         
            this.projectRepository = projectRepository;
        }
        public async Task<ProjectDto> Handle(CreateProjectCommand request, CancellationToken cancellationToken)
        {
            var proejct = new Project(request.name, request.url, request.description, request.userId);
                
            var result =await projectRepository.AddAsync(proejct, cancellationToken);

            return new ProjectDto
            {
                Description = result.Description,
                Id = result.Id,
                Name = result.Name,
                Url = result.Url
            };
        }
    }
    public record CreateProjectCommand(string name, string url, string description, string userId) : IRequest<ProjectDto>;
}
