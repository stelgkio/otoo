using Application.Contracts.Product.GetProducts;
using Infrastracture.Services.Repositories;
using MediatR;


namespace Application.Handlers.Products.Query
{
    public class GetProductHandler : IRequestHandler<GetProductsQuery, IList<ProductDto>>
    {
        private readonly IProductRepository _productRepository;
        public GetProductHandler(IProductRepository productRepository)
        {
            this._productRepository = productRepository;
        }

        public async Task<IList<ProductDto>> Handle(GetProductsQuery request, CancellationToken cancellationToken)
        {
            var retult = await _productRepository.GetAllAsync(cancellationToken);
            IList<ProductDto> model = retult.Select(x => new ProductDto { Id = x.Id, Description = x.Description, Name = x.Name }).ToList();

            return model;         
        }
    }
    public record GetProductsQuery(string UserId) : IRequest<IList<ProductDto>>;
}



