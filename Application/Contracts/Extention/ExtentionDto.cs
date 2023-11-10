using Domain.Enums;
using Domain.Models;
namespace Application.Contracts.Extention
{
    public class ExtentionDto
    {
        public string Id { get; }
        public string Name { get; }
        public string Description { get; }
        public string ProjectId { get; }
        public ProductType ProductType { get; }
        public bool IsAdded { get; }
        public bool IsCommingSoon { get; }

        public ExtentionDto(string id, string name, string description, string projectId, bool isAdded, bool isCommingSoon, ProductType type)
        {
            Id = id;
            Name = name;
            Description = description;
            ProjectId = projectId;
            IsAdded = isAdded;
            IsCommingSoon = isCommingSoon;
            ProductType = type;
        }
    }
}
