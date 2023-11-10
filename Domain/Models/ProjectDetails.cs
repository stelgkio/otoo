using Domain.Common;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Domain.Models
{
    public class ProjectDetails  : ValueObject
    {
        public ProjectDetails() { }
        public ProjectDetails(string consumerKey, string consumerSecret, string apiVersion)
        {
            ConsumerKey = consumerKey;
            ConsumerSecret = consumerSecret;
            ApiVersion = apiVersion;
        }

        public string ConsumerKey { get; private set; }
        public string ConsumerSecret { get; private set; }
        public string ApiVersion { get; private set; }

      

        protected override IEnumerable<object> GetAtomicValues()
        {
            yield return ConsumerKey;
            yield return ConsumerSecret;
            yield return ApiVersion;
        }
    }
}
