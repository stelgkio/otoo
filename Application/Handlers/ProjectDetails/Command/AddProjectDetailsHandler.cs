using Application.Contracts.ProjectDto.GetProjects;
using Domain.Models;
using Infrastracture.Services.Repositories;
using Infrastracture.Services.WooCommerce;
using MediatR;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Application.Handlers.ProjectDetails.Command
{

    public class AddProjectDetailsHandler : IRequestHandler<AddProjectDetailsCommand, bool>
    {
        private readonly IProjectRepository projectRepository;
        private readonly IWooCommerceService wooCommerce;
        public AddProjectDetailsHandler(IProjectRepository projectRepository, IWooCommerceService wooCommerce)
        {
            this.projectRepository = projectRepository;
            this.wooCommerce = wooCommerce;
        }
        public async Task<bool> Handle(AddProjectDetailsCommand request, CancellationToken cancellationToken)
        {
            Project project = await projectRepository.GetByIdAsync(request.Id, cancellationToken);
            if (project is null) return false;

            project.WithWoocommerceDetails(request.ConsumerKey, request.ConsumerSecret, request.ApiVersion);

            try
            {
                var data = await this.wooCommerce.GetAllReports(request.ApiVersion, request.ConsumerKey, request.ConsumerSecret);

            } catch (Exception ex)
            {
                throw new Exception("Invalid url or credetials", ex);
            }
            await projectRepository.SaveProjectAsync(cancellationToken);
           
            
            
            return true;
        }
    }
    public record AddProjectDetailsCommand(string ConsumerKey, string ConsumerSecret, string ApiVersion, string Id) : IRequest<bool>;

}
