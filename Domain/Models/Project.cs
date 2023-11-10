using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Domain.Models
{
    public class Project : Entity
    {

        public string UserId { get; private set; }
        public string Name { get; private set; }
        public string Url { get; private set; }
        public string Description { get; private set; }
        public ICollection<Extention> Extentions { get; set; }

        public ProjectDetails? WoocommerceDetails { get; private set; }

        public Project()
        {

        }
        public Project(string name, string url, string description, string userId)
        {

            Name = name ?? throw new ArgumentException("Name is null", nameof(name));
            Url = url ?? throw new ArgumentException("Url is null", nameof(url));
            Description = description;
            UserId = userId ?? throw new ArgumentException("UserId is null", nameof(userId));
            WoocommerceDetails = null;
        }

        public Project WithWoocommerceDetails(string consumerKey, string consumerSecret, string apiVersion)
        {
            WoocommerceDetails = new ProjectDetails(consumerKey, consumerSecret, apiVersion);
            return this;
        }
    }
}
