using Domain.Models;
using Infrastracture.Services.Repositories.Base;

namespace Infrastracture.Services.Repositories;

public interface IProjectRepository : IBaseRepository<Project>
{
    Task<IList<Project>> GetAllByUserIdAsync(string userId,CancellationToken cancellationToken = default);
    Task<IList<Project>> GetAllAsync(CancellationToken cancellationToken = default);

    Task<Project?> GetByIdAsync(string id, CancellationToken cancellationToken = default);

    Task<Project?> GetByNameAsync(string name, CancellationToken cancellationToken = default);

    Task SaveProjectAsync( CancellationToken cancellationToken = default);
}