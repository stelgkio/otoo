using Domain.Enums;
using System;
using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;

namespace Application.Contracts.Extention
{
    public class ExtentionRequest
    {
        /// <summary>
        /// Name of the Extention
        /// </summary>
        ///<example>Extention name</example>
        [Required]
        public string Name { get; set; }     
        /// <summary>
        /// Name of the project
        /// </summary>
        public string ProjectId { get; set; }       

        [Required]
      //  [JsonConverter(typeof(JsonStringEnumConverter))]
        /// <summary>
        ///     ProductType
        /// </summary>
        public ProductType ProductType { get; set; }


    }
}
