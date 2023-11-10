using Domain.Models;
using Infrastracture.Data;
using Infrastracture.Services.Repositories.Base;

namespace Infrastracture.Services.Repositories;

public class ProjectRepository : BaseRepository<Project, ApplicationDbContext>, IProjectRepository
{

    public ProjectRepository(ApplicationDbContext? dbContext)
        : base(dbContext)
    {
    }


    public async Task<Project?> CreateProjectAsync(string id, CancellationToken cancellationToken = default)
    {
        return await GetSingleOrDefaultAsync<Project>(std => std.Id == id, cancellationToken: cancellationToken).ConfigureAwait(false);
    }
    public async Task<IList<Project>> GetAllByUserIdAsync(string userId, CancellationToken cancellationToken = default)
    {
        IList<Project> result = await GetAsync<Project>(predicate: p=> p.UserId == userId,  orderBy: cmp => cmp.OrderBy(std => std.Name),
            cancellationToken: cancellationToken).ConfigureAwait(false);
        return result;
    }
    public async Task<IList<Project>> GetAllAsync(CancellationToken cancellationToken = default)
    {
        IList<Project> result = await GetAsync<Project>(orderBy: cmp => cmp.OrderBy(std => std.Name),
            cancellationToken: cancellationToken).ConfigureAwait(false);
        return result;
    }

    public async Task<Project?> GetByIdAsync(string id, CancellationToken cancellationToken = default)
    {
        return await GetSingleOrDefaultAsync<Project>(std => std.Id == id, cancellationToken: cancellationToken).ConfigureAwait(false);
    }


    public async Task<Project?> GetByNameAsync(string name, CancellationToken cancellationToken = default)
    {
        return await GetSingleOrDefaultAsync<Project>(std => std.Name == name, cancellationToken: cancellationToken).ConfigureAwait(false);
    }

    public async Task SaveProjectAsync( CancellationToken cancellationToken = default)
    {
     await DatabaseContext.SaveChangesAsync(cancellationToken);
     
    }
}