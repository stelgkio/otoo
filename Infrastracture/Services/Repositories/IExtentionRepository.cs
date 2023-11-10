using Domain.Models;
using Infrastracture.Services.Repositories.Base;


namespace Infrastracture.Services.Repositories
{
    public interface IExtentionRepository : IBaseRepository<Extention>
    {
        Task<IList<Extention>> GetAllByProjectIdAsync(string projectId,string userId,CancellationToken cancellationToken = default);
    }
}
