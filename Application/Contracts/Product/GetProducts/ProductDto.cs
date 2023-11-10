using Domain.Enums;

namespace Application.Contracts.Product.GetProducts
{
    public class ProductDto
    {
        /// <summary>
        /// Name of the project
        /// </summary>
        public string Id { get; set; }

        /// <summary>
        /// Name of the project
        /// </summary>
        public string Name { get; set; }

        /// <summary>
        /// Description of the project
        /// </summary>
        public string Description { get; set; }

        /// <summary>
        /// ProductType of the project
        /// </summary>
        public ProductType ProductType { get; set; }

    }
}
