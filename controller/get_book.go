package controller

import (
	"api_go_bwa/database"
	"api_go_bwa/dto"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetAllBooks(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT * FROM books")
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

	return c.Status(http.StatusOK).JSON(
		map[string]any{
			"status":  http.StatusOK,
			"message": "success",
			"data":    books,
		})
}
