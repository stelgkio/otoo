using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

#pragma warning disable CA1814 // Prefer jagged arrays over multidimensional

namespace Infrastracture.Migrations
{
    /// <inheritdoc />
    public partial class extentions : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "96af749b-77b1-4d41-a21b-8361324e6e51");

            migrationBuilder.AddColumn<bool>(
                name: "IsCommingSoon",
                table: "Products",
                type: "boolean",
                nullable: false,
                defaultValue: false);

            migrationBuilder.AddColumn<int>(
                name: "ProductType",
                table: "Extentions",
                type: "integer",
                nullable: false,
                defaultValue: 0);

            migrationBuilder.InsertData(
                table: "Products",
                columns: new[] { "Id", "CreatedDate", "Description", "IsCommingSoon", "IsUpdatedDate", "LastModifiedDate", "Name", "ProductType", "UpdatedDate" },
                values: new object[,]
                {
                    { "44fca1ff-7c08-45fa-a2ee-97e4a356ab10", new DateTime(2023, 9, 25, 16, 16, 32, 551, DateTimeKind.Utc).AddTicks(1275), "Softone description", false, false, new DateTime(2023, 9, 25, 16, 16, 32, 551, DateTimeKind.Utc).AddTicks(1281), "SoftOne", 0, null },
                    { "798de812-c86b-4b1e-bf9c-18c2c0a87866", new DateTime(2023, 9, 25, 16, 16, 32, 551, DateTimeKind.Utc).AddTicks(1399), "EpsilonNet description", false, false, new DateTime(2023, 9, 25, 16, 16, 32, 551, DateTimeKind.Utc).AddTicks(1400), "EpsilonNet", 2, null },
                    { "7b1fb59f-d084-46c5-a00e-76ff5a85140a", new DateTime(2023, 9, 25, 16, 16, 32, 551, DateTimeKind.Utc).AddTicks(1358), "Mydata", true, false, new DateTime(2023, 9, 25, 16, 16, 32, 551, DateTimeKind.Utc).AddTicks(1359), "MyData", 1, null }
                });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "44fca1ff-7c08-45fa-a2ee-97e4a356ab10");

            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "798de812-c86b-4b1e-bf9c-18c2c0a87866");

            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "7b1fb59f-d084-46c5-a00e-76ff5a85140a");

            migrationBuilder.DropColumn(
                name: "IsCommingSoon",
                table: "Products");

            migrationBuilder.DropColumn(
                name: "ProductType",
                table: "Extentions");

            migrationBuilder.InsertData(
                table: "Products",
                columns: new[] { "Id", "CreatedDate", "Description", "IsUpdatedDate", "LastModifiedDate", "Name", "ProductType", "UpdatedDate" },
                values: new object[] { "96af749b-77b1-4d41-a21b-8361324e6e51", new DateTime(2023, 9, 15, 10, 12, 3, 933, DateTimeKind.Utc).AddTicks(507), "Softone description", false, new DateTime(2023, 9, 15, 10, 12, 3, 933, DateTimeKind.Utc).AddTicks(510), "SoftOne", 0, null });
        }
    }
}
