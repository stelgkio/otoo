
using System.ComponentModel.DataAnnotations;

namespace Domain.Models
{
    public class Entity
    {
        protected Entity()
        {
            Id= Guid.NewGuid().ToString();
            CreatedDate = DateTime.UtcNow;
            IsUpdatedDate = false;
            UpdatedDate = null;
            LastModifiedDate = DateTime.UtcNow;
        }
        [Required]
        [Key]
        public virtual string Id { get ; protected set; }
        public virtual DateTime CreatedDate { get; protected set; }
        public virtual bool IsUpdatedDate { get; protected set; }  
        public virtual DateTime? UpdatedDate { get; protected set; }
        public virtual DateTime LastModifiedDate { get; protected set; }

        public void UpdatedEntity()
        {
            UpdatedDate = DateTime.UtcNow;
            IsUpdatedDate = true;
            LastModifiedDate = DateTime.UtcNow;
        }

    }
}
