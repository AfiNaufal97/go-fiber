package controller

import (
	"api_go_bwa/database"
	"api_go_bwa/dto"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func GetAllBooks(c *fiber.Ctx) error {

	var request dto.RequestBook
	err := c.QueryParser(&request)
	//
	query := "SELECT * FROM books"
	if err != nil {
		return err
	}
	//
	println(request.Search)
	//
	var arg []any
	if request.Search != "" {
		println("masuk sini")
		query += " WHERE LOWER(title) LIKE $1 OR LOWER(author) LIKE $1 OR LOWER(description) LIKE $1 OR LOWER(genre) LIKE $1 "
		filter := fmt.Sprintf("%%%s%%", strings.ToLower(request.Search))
		arg = append(arg, filter)
	}

	rows, err := database.DB.Query(query, arg...)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var books []dto.CreateBookRequest
	for rows.Next() {
		var book dto.CreateBookRequest
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Isbn, &book.Genre, &book.Author, &book.Stock, &book.PublishDate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if books != nil {
		return c.Status(http.StatusOK).JSON(
			map[string]any{
				"status":  http.StatusOK,
				"message": "success",
				"data":    books,
			})
	}

	return c.Status(http.StatusNotFound).JSON(
		map[string]any{
			"status":  http.StatusNotFound,
			"message": "Data tidak ditemukan",
		})

}
