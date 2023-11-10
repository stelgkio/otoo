using Domain.Enums;

namespace Domain.Models
{
    public class Product : Entity
    {         
        public string Name { get; private set; }          
        public string Description { get; private set; }
        public bool IsCommingSoon { get; private set; }
        public ProductType ProductType { get; private set;  }

        public Product(string name, string description, ProductType productType, bool isCommingSoon)
        {
            Name = name;
            Description = description;
            ProductType = productType;
            IsCommingSoon = isCommingSoon;
        }

        public Product()
        {
                
        }
    }
}
