using Application.Contracts.Extention;
using Application.Handlers.Extentions.Command;
using Application.Handlers.Extentions.Query;
using MediatR;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.ComponentModel.DataAnnotations;

namespace Konektor.Controllers
{
    
    [ApiController]
    [Route("[controller]")]
    //[Authorize]
    public class ExtentionController : ApiController
    {

        /// <summary>
        /// Get enebled extention for project
        /// </summary>
        /// <remarks>Get all e extentions</remarks>        
        /// <response code="201">Resource created</response>
        /// <response code="400">Bad Request</response>
        [HttpGet("{projectId}")]
        [ProducesResponseType(200)]
        [ProducesResponseType(400)]
        [ProducesResponseType(typeof(ExtentionResponse), 200)]
        public async Task<IActionResult> GetAllExtentions([FromRoute][Required] string projectId)
        {               
            var result = await Mediator.Send(new GetExtentionQuery(UserId, projectId));
            
            return Ok(new ExtentionResponse(result));
        }
        /// <summary>
        /// Add new extention
        /// </summary>
        /// <remarks>Adds an extention</remarks>
        /// <param name="body">extention</param>
        /// <param name="projectId">extention</param>
        /// <response code="201">Resource created</response>
        /// <response code="400">Bad Request</response>
        [HttpPost()]
        [ProducesResponseType(200)]
        [ProducesResponseType(400)]
        [ProducesResponseType(typeof(ExtentionResponse), 200)]
        public async Task<IActionResult> AddExtention([FromBody]ExtentionRequest body)
        {
            var result = await Mediator.Send(new AddExtentionCommand(UserId, body.ProjectId, body.Name,"",body.ProductType));

            return Ok(result);
        }
    }
}
