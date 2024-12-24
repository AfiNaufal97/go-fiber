package controller

import (
	"api_go_bwa/database"
	"github.com/gofiber/fiber/v2"
)

func DeleteBookController(c *fiber.Ctx) error {
	id := c.Params("id")

	// Delete the book from the database
	result, err := database.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if rowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})

}
