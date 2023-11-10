using Application.Contracts.ProjectDetails;
using Application.Contracts.ProjectDto.CreateProject;
using Application.Handlers.ProjectDetails.Command;
using Konektor.Endpoints.Products;
using MediatR;
using Microsoft.AspNetCore.Mvc;
using System.Security.Claims;

namespace Konektor.Endpoints.ProjectDetails
{
    public static class ProjectDetailsEndPoint
    {
        public static void MapProjectDetailsEndpoint(this WebApplication app)
        {
            app.CreateProjectDetails();
            app.UpdateProjectDetails();
        }


        public static WebApplication CreateProjectDetails(this WebApplication app)
        {
            app.MapPost("projectdetails", async (IMediator mediator, ClaimsPrincipal user, [FromBody] AddProjectDetailsRequest req) =>
            {
               var project = await mediator.Send(new AddProjectDetailsCommand(req.ConsumerKey, req.ConsumerSecret, req.EndpointUrl, req.Id));

                if( project) return Results.Ok();
                return Results.NotFound();
            })
            .Produces(StatusCodes.Status200OK)
            .Produces(StatusCodes.Status404NotFound)
            .WithTags("ProjectDetailsEndpoint")
            .RequireAuthorization();

            return app;
        }
        public static WebApplication UpdateProjectDetails(this WebApplication app)
        {
            app.MapPut("projectdetails", async (IMediator mediator, ClaimsPrincipal user, [FromBody] UpdateProjectDetailsRequest req) =>
            {

                return Results.Ok();
            }).WithTags("ProjectDetailsEndpoint").RequireAuthorization();

            return app;
        }
    }
}
