using Domain.Models;
using Infrastracture.Services.Repositories.Base;

namespace Infrastracture.Services.Repositories
{
    public interface IProductRepository : IBaseRepository<Product>
    {
        Task<IList<Product>> GetAllAsync(CancellationToken cancellationToken = default);
    }
}
