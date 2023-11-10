
using Infrastracture.Services.WooCommerce;
using Konektor.Models;
using Microsoft.AspNetCore.Identity;
using System.Security.Claims;

namespace Konektor.Endpoints.WooCommerce
{
    public static class WoocommerceEndpoint
    {
        public static void MapWoocommerceEndpoints(this WebApplication app)
        {
            app.MapGet("woocommerce/products", GetAllProducts);
            app.MapGet("woocommerce/reports", GetAllReports);
        }

        public static async Task<IResult> GetAllProducts(IWooCommerceService woocommerceService, ClaimsPrincipal user, UserManager<ApplicationUser> userManager)
        {            
            var userId = BaseEndpoint.GetUserId(user);          
            var products = await woocommerceService.GetAllProducts("http://test.local", "ck_16cfc3f9d2cb868ee6280bd602d91aad212710ae", "cs_8bfce11e60288885f293d5a75fea190532f5bd68");
            return Results.Ok(products);
        }
        public static async Task<IResult> GetAllReports(IWooCommerceService woocommerceService, ClaimsPrincipal user, UserManager<ApplicationUser> userManager)
        {
            var userId = BaseEndpoint.GetUserId(user);
            var products = await woocommerceService.GetAllReports("http://test.local", "ck_e253c807ff6c444bc1f6ce0c9ed799615a4ce4df", "cs_67de135b6ca22deef044e38568a06489607b8eb9");
            return Results.Ok(products);
        }
    }
}
