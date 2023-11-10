using Application.Contracts.Extention;
using Domain.Enums;
using Domain.Models;
using Infrastracture.Services.Repositories;
using MediatR;
using Microsoft.Extensions.Logging;


namespace Application.Handlers.Extentions.Command
{
    public class AddNewExtentionHandler : IRequestHandler<AddExtentionCommand, ExtentionDto>
    {
        private readonly ILogger _logger;
        private readonly IExtentionRepository _extentionRepository;
        private readonly IProjectRepository _projectRepository;
        private readonly IProductRepository _productRepository;

        public AddNewExtentionHandler(IExtentionRepository extentionRepository, IProductRepository productRepository, IProjectRepository projectRepository, ILogger<AddNewExtentionHandler> logger)
        {
            _logger = logger;
            _extentionRepository = extentionRepository;
            _projectRepository = projectRepository;
            _productRepository = productRepository;
        }



        public async Task<ExtentionDto> Handle(AddExtentionCommand request, CancellationToken cancellationToken)
        {
            var projectExist = await _projectRepository.GetByIdAsync(request.projectId, cancellationToken);
            if (projectExist is null)
            {
                _logger.LogError("Invalid project id");
                throw new ArgumentException();
            }

            var product = await _productRepository.GetAllAsync();

            var productData = product.Where(x => x.Name == request.name)
                             .Where(x => x.ProductType == request.type).FirstOrDefault();

            if (productData is null)
            {
                _logger.LogError("Invalid product data");
                throw new ArgumentException();
            }
            if (productData.IsCommingSoon)
            {
                _logger.LogError("Invalid product data");
                throw new ArgumentException();
            }

            var extention = new Extention(request.name, request.description, request.projectId, request.type);
            var result = await _extentionRepository.AddAsync(extention);



            return new ExtentionDto(result.Id, result.Name, result.Description, result.ProjectId, true, productData.IsCommingSoon, request.type);
        }
    }
    public record AddExtentionCommand(string userId, string projectId, string name, string description, ProductType type) : IRequest<ExtentionDto>;
}


;