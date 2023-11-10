using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Infrastracture.Migrations
{
    /// <inheritdoc />
    public partial class projectdetailsapiversion : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "edb7735d-d3d1-485f-b866-dfcc0689c298");

            migrationBuilder.AddColumn<string>(
                name: "WoocommerceDetails_ApiVersion",
                table: "Projects",
                type: "text",
                nullable: true);

            migrationBuilder.InsertData(
                table: "Products",
                columns: new[] { "Id", "CreatedDate", "Description", "IsUpdatedDate", "LastModifiedDate", "Name", "ProductType", "UpdatedDate" },
                values: new object[] { "96af749b-77b1-4d41-a21b-8361324e6e51", new DateTime(2023, 9, 15, 10, 12, 3, 933, DateTimeKind.Utc).AddTicks(507), "Softone description", false, new DateTime(2023, 9, 15, 10, 12, 3, 933, DateTimeKind.Utc).AddTicks(510), "SoftOne", 0, null });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "96af749b-77b1-4d41-a21b-8361324e6e51");

            migrationBuilder.DropColumn(
                name: "WoocommerceDetails_ApiVersion",
                table: "Projects");

            migrationBuilder.InsertData(
                table: "Products",
                columns: new[] { "Id", "CreatedDate", "Description", "IsUpdatedDate", "LastModifiedDate", "Name", "ProductType", "UpdatedDate" },
                values: new object[] { "edb7735d-d3d1-485f-b866-dfcc0689c298", new DateTime(2023, 9, 15, 9, 51, 22, 791, DateTimeKind.Utc).AddTicks(7376), "Softone description", false, new DateTime(2023, 9, 15, 9, 51, 22, 791, DateTimeKind.Utc).AddTicks(7378), "SoftOne", 0, null });
        }
    }
}
