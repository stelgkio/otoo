using System.Security.Claims;
using WooCommerce.NET.WordPress.v2;

namespace Konektor.Endpoints
{
    public static class BaseEndpoint
    {
        public static string? GetUserId(ClaimsPrincipal user) => user?.Claims.Where(x => x.Type.Contains("nameidentifier")).Select(x => x.Value).FirstOrDefault();

    }
}