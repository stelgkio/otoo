using Application.Contracts.ProjectDto.GetProjects;
using Infrastracture.Services.Repositories;
using MediatR;

namespace Application.Handlers.Projects.Query
{
    public class GetProjectsHandler : IRequestHandler<GetProjectsQuery, IList<ProjectDto>>
    {
        private readonly IProjectRepository projectRepository;
        public GetProjectsHandler(IProjectRepository projectRepository)
        {
            this.projectRepository = projectRepository;
        }



        public async Task<IList<ProjectDto>> Handle(GetProjectsQuery request, CancellationToken cancellationToken)
        {
            var retult = await projectRepository.GetAllByUserIdAsync(request.UserId, cancellationToken);
            IList<ProjectDto> model = retult.Select(x => new ProjectDto {Id=x.Id,  Description = x.Description, Url = x.Url, Name = x.Name }).ToList();

            return model;
        }
    }

    public record GetProjectsQuery(string UserId) : IRequest<IList<ProjectDto>>;
}
