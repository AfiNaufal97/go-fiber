package controller

import (
	"api_go_bwa/database"
	"api_go_bwa/dto"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetDetailBook(c *fiber.Ctx) error {
	id := c.Params("id")
	row := database.DB.QueryRow("SELECT * FROM books WHERE id = $1", id)
	var book dto.DetailBook
	err := row.Scan(&book.ID, &book.Title, &book.Description, &book.Isbn, &book.Genre, &book.Author, &book.Stock, &book.PublishDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(
		map[string]any{
			"status":  http.StatusOK,
			"message": "success",
			"data":    book,
		})
}
