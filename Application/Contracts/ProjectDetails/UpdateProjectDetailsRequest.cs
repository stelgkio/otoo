using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Application.Contracts.ProjectDetails
{
    public class UpdateProjectDetailsRequest
    {
        public string Id { get; set; }
        public string ConsumerKey { get; set; }
        public string ConsumerSecret { get; set; }
        public string ApiVersion { get; set; }
    }
}
