using WooCommerceNET;
using WooCommerceNET.WooCommerce.v3;
using WooCommerceNET.WooCommerce.v3.Extension;

namespace Infrastracture.Services.WooCommerce
{
    /// <summary>
    /// 
    /// </summary>
    public class WooCommerceService : IWooCommerceService
    {
        public WooCommerceService()
        {

        }

        public async Task<List<Product>> GetAllProducts(string siteUrl, string WooCommerceKey, string WooCommerceSecret)
        {
            RestAPI rest2 = new RestAPI($"{siteUrl}/wp-json/wc/v3/", WooCommerceKey, WooCommerceSecret);
            WCObject wc2 = new WCObject(rest2);
            return await wc2.Product.GetAll();
        }

        public async Task<List<Report>> GetAllReports(string siteUrl, string WooCommerceKey, string WooCommerceSecret)
        {
            RestAPI rest2 = new RestAPI($"{siteUrl}/wp-json/wc/v3/", WooCommerceKey, WooCommerceSecret);
            WCObject wc2 = new WCObject(rest2);

            return await wc2.Report.GetAll();
        }
       

    }
     
}
