using Domain.Enums;
using Domain.Models;
using Duende.IdentityServer.EntityFramework.Options;
using Konektor.Models;
using Microsoft.AspNetCore.ApiAuthorization.IdentityServer;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Options;
using System.Reflection.Metadata;

namespace Infrastracture.Data
{
    public class ApplicationDbContext : ApiAuthorizationDbContext<ApplicationUser>
    {
       
        public ApplicationDbContext(DbContextOptions options, IOptions<OperationalStoreOptions> operationalStoreOptions)
            : base(options, operationalStoreOptions)
        {
           
        }
        public DbSet<Project> Projects { get; set; }
        public DbSet<Project> Extentions { get; set; }
        public DbSet<Product> Products { get; set; }
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Project>().ToTable("Projects").HasKey(p => p.Id); 
            modelBuilder.Entity<Extention>().ToTable("Extentions").HasKey(p => p.Id);
            modelBuilder.Entity<Project>()              
                .HasMany(e => e.Extentions)
                .WithOne(e => e.Project)
                .HasForeignKey(e => e.ProjectId)
                .IsRequired(true);

            modelBuilder.Entity<Project>().OwnsOne(x => x.WoocommerceDetails, woo =>
            {
                woo.Property(z => z.ConsumerKey).IsRequired(false);
                woo.Property(z => z.ConsumerSecret).IsRequired(false);
                woo.Property(z => z.ApiVersion).IsRequired(false);

            });

            modelBuilder.Entity<Extention>().Property(e => e.IsVisible).HasDefaultValue(true);

            modelBuilder.Entity<Product>().Property(e => e.Description).HasMaxLength(300);            

            modelBuilder.Entity<Product>().HasData(new Product("SoftOne","Softone description",ProductType.SoftOne,true));
            modelBuilder.Entity<Product>().HasData(new Product("MyData", "Mydata", ProductType.MyData,false));
            modelBuilder.Entity<Product>().HasData(new Product("EpsilonNet", "EpsilonNet description", ProductType.EpsilonNet, true));



            base.OnModelCreating(modelBuilder);

        }
    }
}