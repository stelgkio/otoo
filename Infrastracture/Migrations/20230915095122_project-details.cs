using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Infrastracture.Migrations
{
    /// <inheritdoc />
    public partial class projectdetails : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "cceaddd5-df0d-4e76-894c-56a9fdfeb13e");

            migrationBuilder.AddColumn<string>(
                name: "WoocommerceDetails_ConsumerKey",
                table: "Projects",
                type: "text",
                nullable: true);

            migrationBuilder.AddColumn<string>(
                name: "WoocommerceDetails_ConsumerSecret",
                table: "Projects",
                type: "text",
                nullable: true);

            migrationBuilder.AlterColumn<bool>(
                name: "IsVisible",
                table: "Extentions",
                type: "boolean",
                nullable: false,
                defaultValue: true,
                oldClrType: typeof(bool),
                oldType: "boolean",
                oldNullable: true,
                oldDefaultValue: true);

            migrationBuilder.InsertData(
                table: "Products",
                columns: new[] { "Id", "CreatedDate", "Description", "IsUpdatedDate", "LastModifiedDate", "Name", "ProductType", "UpdatedDate" },
                values: new object[] { "edb7735d-d3d1-485f-b866-dfcc0689c298", new DateTime(2023, 9, 15, 9, 51, 22, 791, DateTimeKind.Utc).AddTicks(7376), "Softone description", false, new DateTime(2023, 9, 15, 9, 51, 22, 791, DateTimeKind.Utc).AddTicks(7378), "SoftOne", 0, null });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "edb7735d-d3d1-485f-b866-dfcc0689c298");

            migrationBuilder.DropColumn(
                name: "WoocommerceDetails_ConsumerKey",
                table: "Projects");

            migrationBuilder.DropColumn(
                name: "WoocommerceDetails_ConsumerSecret",
                table: "Projects");

            migrationBuilder.AlterColumn<bool>(
                name: "IsVisible",
                table: "Extentions",
                type: "boolean",
                nullable: true,
                defaultValue: true,
                oldClrType: typeof(bool),
                oldType: "boolean",
                oldDefaultValue: true);

            migrationBuilder.InsertData(
                table: "Products",
                columns: new[] { "Id", "CreatedDate", "Description", "IsUpdatedDate", "LastModifiedDate", "Name", "ProductType", "UpdatedDate" },
                values: new object[] { "cceaddd5-df0d-4e76-894c-56a9fdfeb13e", new DateTime(2023, 7, 18, 15, 21, 11, 935, DateTimeKind.Utc).AddTicks(5445), "Softone description", false, new DateTime(2023, 7, 18, 15, 21, 11, 935, DateTimeKind.Utc).AddTicks(5449), "SoftOne", 0, null });
        }
    }
}
