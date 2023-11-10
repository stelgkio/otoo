using Application.Handlers.Projects.Query;
using Infrastracture.Data;
using Infrastracture.Services.Repositories;
using Infrastracture.Services.WooCommerce;
using Konektor.Endpoints.Products;
using Konektor.Endpoints.ProjectDetails;
using Konektor.Endpoints.Projects;
using Konektor.Endpoints.WooCommerce;
using Konektor.Models;
using Microsoft.AspNetCore.ApiAuthorization.IdentityServer;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.EntityFrameworkCore;
using Microsoft.OpenApi.Models;
using System;
using System.Reflection;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
var connectionString = builder.Configuration.GetConnectionString("DefaultConnection") ?? throw new InvalidOperationException("Connection string 'DefaultConnection' not found.");
builder.Services.AddDbContext<ApplicationDbContext>(options =>
    options.UseNpgsql(connectionString));
builder.Services.AddDatabaseDeveloperPageExceptionFilter();

builder.Services.AddDefaultIdentity<ApplicationUser>(options => options.SignIn.RequireConfirmedAccount = false)
    .AddEntityFrameworkStores<ApplicationDbContext>();

builder.Services.AddIdentityServer()
    .AddApiAuthorization<ApplicationUser, ApplicationDbContext>();

builder.Services.AddAuthentication()
    .AddCookie(options =>
{
    // add an instance of the patched manager to the options:
    options.CookieManager = new ChunkingCookieManager();

    options.Cookie.HttpOnly = true;
    options.Cookie.SameSite = SameSiteMode.None;
    options.Cookie.SecurePolicy = CookieSecurePolicy.Always;
}).AddIdentityServerJwt();
builder.Services.Configure<JwtBearerOptions>(
    IdentityServerJwtConstants.IdentityServerJwtBearerScheme,
    options =>
    {
        options.TokenValidationParameters.ValidateIssuer = false;
    });

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddControllersWithViews();
builder.Services.AddRazorPages();
builder.Services.AddCors(options =>
{
    options.AddPolicy("api", policy =>
    {
        policy.AllowAnyOrigin().AllowAnyHeader().AllowAnyMethod();
    });
});

builder.Services.AddSwaggerGen(options =>
{
    options.SwaggerDoc("v1", new OpenApiInfo { Title = "Web API", Version = "v1" });

    //options.AddSecurityDefinition("JWT", new OpenApiSecurityScheme
    //{
    //    Type = SecuritySchemeType.ApiKey,
    //    Name = "Authorization",
    //    In = ParameterLocation.Header,
    //    Description = "JWT token"
    //});
    //options.OperationFilter<SecurityRequirementsOperationFilter>();
    options.AddSecurityDefinition("Bearer", new OpenApiSecurityScheme
    {
        In = ParameterLocation.Header,
        Description = "Please enter a valid token",
        Name = "Authorization",
        Type = SecuritySchemeType.Http,
        BearerFormat = "JWT",
        Scheme = "Bearer"
    });
    options.AddSecurityRequirement(new OpenApiSecurityRequirement
    {
        {
            new OpenApiSecurityScheme
            {
                Reference = new OpenApiReference
                {
                    Type=ReferenceType.SecurityScheme,
                    Id="Bearer"
                }
            },
            new string[]{}
        }
    });
    //var filename = Assembly.GetExecutingAssembly().GetName().Name + ".xml";
    //var filePath = Path.Combine(AppContext.BaseDirectory, filename);
    //options.IncludeXmlComments(filePath, true);
});


// Add services
builder.Services.AddTransient<IWooCommerceService, WooCommerceService>();
builder.Services.AddTransient<IProjectRepository, ProjectRepository>();
builder.Services.AddTransient<IProductRepository, ProductRepository>();
builder.Services.AddTransient<IExtentionRepository, ExtentionRepository>();


// Registers handlers and mediator types from the specified assemblies.
builder.Services.AddMediatR(cfg =>
{
    cfg.RegisterServicesFromAssemblies(typeof(GetProjectsQuery).Assembly);
});
var app = builder.Build();


// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseMigrationsEndPoint();
    //// Migrate database
    //using (IServiceScope scope = app.Services.CreateScope())
    //{
    //    IServiceProvider services = scope.ServiceProvider;
    //    ApplicationDbContext context = services.GetRequiredService<ApplicationDbContext>();
    //    context.Database.Migrate();

    //}
    app.UseSwagger(settings =>
    {
        settings.RouteTemplate = "swagger/{documentName}/openapi.json";
    });
    app.MapSwagger();
    app.UseSwaggerUI(settings =>
    {
        settings.DocumentTitle = "API Template description";
        settings.RoutePrefix = "swagger";
        settings.SwaggerEndpoint("/swagger/v1/openapi.json", "API description V1");
    });

    //// Migrate latest database changes during startup
    //using (var scope = app.Services.CreateScope())
    //{
    //    var dbContext = scope.ServiceProvider
    //        .GetRequiredService<ApplicationDbContext>();

    //    // Here is the migration executed
    //    dbContext.Database.Migrate();
    //}

    app.UseHttpsRedirection();
}
else
{
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

//app.UseStaticFiles();
app.UseRouting();

app.UseAuthentication();
//app.Use(async (context, next) => {
//    if (context.User != null && context.User.Identity.IsAuthenticated)
//    {
//        // add claims here 
//        context.User.Claims.Append(new Claim("type-x", "value-x"));
//    }
//    await next();
//});
app.UseIdentityServer();
app.UseAuthorization();
app.UseCors("api");

app.UseStaticFiles();



app.MapControllerRoute(
    name: "default",
    pattern: "{controller}/{action=Index}/{id?}");
app.MapRazorPages();

app.MapFallbackToFile("index.html");
app.MapWoocommerceEndpoints();
app.MapProductEndpoint();
app.MapProjectDetailsEndpoint();
// Map endpoints
app.GetProjects()
    .CreateProject()
    .GetProjectById()
    .DeleteProject()
    .UpdateProject();
    
    




app.Run();
