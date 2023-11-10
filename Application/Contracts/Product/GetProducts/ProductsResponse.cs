namespace Application.Contracts.Product.GetProducts
{
    public class ProductsResponse
    {
        public ProductsResponse(IList<ProductDto> products)
        {
            Products = products;
        }
        public IList<ProductDto> Products { get; set; }


    }
}
