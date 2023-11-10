using Application.Contracts.Extention;
using Domain.Models;
using Infrastracture.Services.Repositories;
using MediatR;
using Microsoft.Extensions.Logging;


namespace Application.Handlers.Extentions.Query
{
    public class GetExtentionHandler : IRequestHandler<GetExtentionQuery, IList<ExtentionDto>>
    {
        private readonly ILogger _logger;
        private readonly IExtentionRepository _extentionRepository;
        private readonly IProductRepository _productRepository;
        public GetExtentionHandler(IExtentionRepository extentionRepository, IProductRepository productRepository, ILogger<GetExtentionHandler> logger)
        {
            _logger = logger;
            _extentionRepository= extentionRepository;
            _productRepository= productRepository;
        }      

        public async Task<IList<ExtentionDto>> Handle(GetExtentionQuery request, CancellationToken cancellationToken)
        {
            var productList = await _productRepository.GetAllAsync(cancellationToken);


            var extentionList = await _extentionRepository.GetAllByProjectIdAsync(request.projectId, request.userId, cancellationToken);

            if(extentionList.Count==0) {

                IList<ExtentionDto> productData = productList.Select(x => new ExtentionDto("", x.Name, x.Description, request.projectId, false, x.IsCommingSoon, x.ProductType)).ToList();


                return productData;

            }
            var result = (from p in productList
                          join e in extentionList
                          on p.Name equals e.Name
                          select new
                          {
                              Name = p.Name,
                              Description = p.Description,
                              ProductType = p.ProductType,
                              ExtentionId = e.Id,
                              ProjectId = e.ProjectId,
                              isAdded = true,
                              isCommingSoon = p.IsCommingSoon

                          }).ToList();

            var result2 = productList.Where(a => !result.Select(b => b.Name).Contains(a.Name)) .Select(p =>
                           new
                           {
                               Name = p.Name,
                               Description = p.Description,
                               ProductType = p.ProductType,
                               ExtentionId = "",
                               ProjectId = request.projectId,
                               isAdded = false,
                               isCommingSoon = p.IsCommingSoon
                           }).ToList();



            IList<ExtentionDto> data=  result.Concat(result2).Select(x => new ExtentionDto(x.ExtentionId, x.Name, x.Description, x.ProjectId, x.isAdded,x.isCommingSoon,x.ProductType)).ToList();

            
            return data;
        }
    }
    public record GetExtentionQuery(string userId, string projectId) : IRequest<IList<ExtentionDto>>;
}
