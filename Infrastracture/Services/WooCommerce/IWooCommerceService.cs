using WooCommerceNET.WooCommerce.v3;

namespace Infrastracture.Services.WooCommerce
{
    public interface IWooCommerceService
    {
        Task<List<Product>> GetAllProducts(string siteUrl, string WooCommerceKey, string WooCommerceSecret);
        Task<List<Report>> GetAllReports(string siteUrl, string WooCommerceKey, string WooCommerceSecret);
    }
}