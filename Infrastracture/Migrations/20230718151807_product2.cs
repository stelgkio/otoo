using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Infrastracture.Migrations
{
    /// <inheritdoc />
    public partial class product2 : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "1db1da22-3872-4a6e-808b-05559f426419");

            migrationBuilder.DropColumn(
                name: "Description",
                table: "Projects");

            migrationBuilder.DropColumn(
                name: "Name",
                table: "Projects");

            migrationBuilder.DropColumn(
                name: "Url",
                table: "Projects");

            migrationBuilder.DropColumn(
                name: "UserId",
                table: "Projects");

            migrationBuilder.DropColumn(
                name: "Name",
                table: "Products");

            migrationBuilder.DropColumn(
                name: "Description",
                table: "Extentions");

            migrationBuilder.DropColumn(
                name: "Name",
                table: "Extentions");

            migrationBuilder.InsertData(
                table: "Products",
                columns: new[] { "Id", "CreatedDate", "Description", "IsUpdatedDate", "LastModifiedDate", "UpdatedDate" },
                values: new object[] { "aa37289b-d373-4509-acdf-ea380eb91a8e", new DateTime(2023, 7, 18, 15, 18, 7, 16, DateTimeKind.Utc).AddTicks(3445), "Softone description", false, new DateTime(2023, 7, 18, 15, 18, 7, 16, DateTimeKind.Utc).AddTicks(3448), null });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "Products",
                keyColumn: "Id",
                keyValue: "aa37289b-d373-4509-acdf-ea380eb91a8e");

            migrationBuilder.AddColumn<string>(
                name: "Description",
                table: "Projects",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<string>(
                name: "Name",
                table: "Projects",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<string>(
                name: "Url",
                table: "Projects",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<string>(
                name: "UserId",
                table: "Projects",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<string>(
                name: "Name",
                table: "Products",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<string>(
                name: "Description",
                table: "Extentions",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<string>(
                name: "Name",
                table: "Extentions",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.InsertData(
                table: "Products",
                columns: new[] { "Id", "CreatedDate", "Description", "IsUpdatedDate", "LastModifiedDate", "Name", "UpdatedDate" },
                values: new object[] { "1db1da22-3872-4a6e-808b-05559f426419", new DateTime(2023, 7, 18, 14, 21, 59, 469, DateTimeKind.Utc).AddTicks(5028), "Softone integration date", false, new DateTime(2023, 7, 18, 14, 21, 59, 469, DateTimeKind.Utc).AddTicks(5031), "SoftOne", null });
        }
    }
}
