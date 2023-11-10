using Domain.Models;
using Infrastracture.Data;
using Infrastracture.Services.Repositories.Base;

namespace Infrastracture.Services.Repositories
{
    public class ExtentionRepository : BaseRepository<Extention, ApplicationDbContext>, IExtentionRepository
    {

        public ExtentionRepository(ApplicationDbContext? dbContext)
            : base(dbContext)
    {
    }

    public async Task<IList<Extention>> GetAllByProjectIdAsync(string projectId, string userId, CancellationToken cancellationToken = default)
    {
        IList<Extention> result = await GetAsync<Extention>(predicate: p => p.ProjectId == projectId, 
            orderBy: cmp => cmp.OrderBy(std => std.Name),
         cancellationToken: cancellationToken).ConfigureAwait(false);
        return result;
    }
}
}
