
using Domain.Enums;

namespace Domain.Models
{
    public class Extention : Entity
    {
        public string Name { get; private set; }        
        public string Description { get; private set; }
        public virtual Project? Project { get; private set; }
        public string ProjectId { get; private set; }
        public ProductType ProductType { get; private set;  }
        public bool IsVisible { get; private set; } = true;

        public Extention(string name, string description,  string projectId, ProductType type)
        {
            Name = name;
            Description = description;            
            ProjectId = projectId;
            ProductType = type;
        }
        public Extention()
        {
                
        }
    }
}
