using Application.Contracts.Product.GetProducts;
using Application.Handlers.Products.Query;

using MediatR;
using System.Security.Claims;


namespace Konektor.Endpoints. Products
{
    /// <summary>
    /// Product Endpoint
    /// </summary>
    public static class ProductEndpoint 
    {
        public static void MapProductEndpoint(this WebApplication app)
        {
            app.GetProduct();
        }

        public static WebApplication GetProduct(this WebApplication app)
        {
            app.MapGet("products", async (IMediator mediator, ClaimsPrincipal user) =>
            {
                var userId = BaseEndpoint.GetUserId(user);
                IList<ProductDto> project = await mediator.Send(new GetProductsQuery(userId));
                return Results.Ok(new ProductsResponse(project));
            })
            .Produces<ProductsResponse>(StatusCodes.Status200OK)
            .Produces(StatusCodes.Status500InternalServerError)            
            .WithTags("ProductEndpoint")             
            .RequireAuthorization(); 

            return app;
        }

       
        
    }
}

