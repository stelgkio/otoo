using AutoMapper;
using MediatR;
using Microsoft.AspNetCore.Mvc;

namespace Konektor.Controllers
{
    
    public abstract class ApiController : ControllerBase
    {
        private IMediator _mediator;
        private IMapper _mapper;

        protected IMediator Mediator => _mediator ??= HttpContext.RequestServices.GetService<IMediator>();
        protected IMapper Mapper => _mapper ??= HttpContext.RequestServices.GetService<IMapper>();

        protected string? UserId
       => string.IsNullOrWhiteSpace(User?.Identity?.Name) ?
          null : User.Claims.Where(x => x.Type.Contains("nameidentifier")).Select(x => x.Value).FirstOrDefault();
    }
}
