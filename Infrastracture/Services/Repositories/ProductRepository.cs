using Domain.Models;
using Infrastracture.Data;
using Infrastracture.Services.Repositories.Base;
using Microsoft.EntityFrameworkCore.Query;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using System.Threading.Tasks;

namespace Infrastracture.Services.Repositories
{
    public class ProductRepository : BaseRepository<Product, ApplicationDbContext>, IProductRepository
    {

        public ProductRepository(ApplicationDbContext? dbContext)
            : base(dbContext)
        {
        }

        public async Task<IList<Product>> GetAllAsync(CancellationToken cancellationToken = default)
        {
            IList<Product> result = await GetAsync<Product>(orderBy: cmp => cmp.OrderBy(std => std.Name),
             cancellationToken: cancellationToken).ConfigureAwait(false);
            return result;
        }
    }
}
