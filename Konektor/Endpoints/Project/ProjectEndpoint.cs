using Application.Contracts.ProjectDto.CreateProject;
using Application.Contracts.ProjectDto.GetProjects;
using Application.Handlers.Projects.Command;
using Application.Handlers.Projects.Query;
using Konektor.Endpoints.Products;
using MediatR;
using Microsoft.AspNetCore.Mvc;
using System.Security.Claims;


namespace Konektor.Endpoints.Projects
{

    public static class ProjectEndpoint
    {

        public static WebApplication GetProjects(this WebApplication app)
        {
            app.MapGet("projects", async (IMediator mediator, ClaimsPrincipal user) =>
            {
                var userId = BaseEndpoint.GetUserId(user);
                IList<ProjectDto> project = await mediator.Send(new GetProjectsQuery(userId));
                return Results.Ok(new ProjectsResponse(project));
            })
            .Produces<ProjectsResponse>(StatusCodes.Status200OK)
            .Produces(StatusCodes.Status404NotFound)
            .WithTags("ProjectEndpoint")
            .RequireAuthorization();

            return app;
        }

        //
        public static WebApplication CreateProject(this WebApplication app)
        {
            app.MapPost("project", async (IMediator mediator, ClaimsPrincipal user, [FromBody] CreateProjectReqeust req) =>
            {

                var userId = BaseEndpoint.GetUserId(user);
                var project = await mediator.Send(new CreateProjectCommand(req.Name, req.Url, req.Description, userId));
                return Results.Ok(project);
            })
            .Produces(StatusCodes.Status200OK)
            .Produces(StatusCodes.Status404NotFound)
            .WithTags("ProjectEndpoint")
            .RequireAuthorization();

            return app;
        }
        public static WebApplication UpdateProject(this WebApplication app)
        {
            app.MapPut("project", async (IMediator mediator, ClaimsPrincipal user, [FromBody] CreateProjectReqeust req) =>
            {

                return Results.Ok();
            }).WithTags("ProjectEndpoint").RequireAuthorization();

            return app;
        }
        public static WebApplication GetProjectById(this WebApplication app)
        {
            app.MapGet("project/{id}", async (int id, IMediator mediator, ClaimsPrincipal user) =>
            {
                return Results.Ok();
            }).WithTags("ProjectEndpoint").RequireAuthorization();

            return app;
        }
        public static WebApplication DeleteProject(this WebApplication app)
        {
            app.MapDelete("project{id}", async (int id, IMediator mediator, ClaimsPrincipal user) =>
            {
                try
                {

                    return Results.Ok();
                } catch (Exception ex)
                {
                    return Results.BadRequest(new { message = ex.Message });
                }
            }).WithTags("ProjectEndpoint").RequireAuthorization();

            return app;
        }

    }
}
