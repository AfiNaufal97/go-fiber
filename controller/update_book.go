package controller

import (
	"api_go_bwa/database"
	"api_go_bwa/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func UpdateBookController(c *fiber.Ctx) error {
	id := c.Params("id")

	var book dto.DetailBook

	err := c.BodyParser(&book)

	validate := validator.New()
	validationErrors := validate.Struct(book)

	if validationErrors != nil {
		valErrors := validationErrors.(validator.ValidationErrors)
		errRes := make(map[string]string)
		for _, err := range valErrors {
			message := "min 0"
			resMessage := err.Tag()
			if resMessage == "required" {
				message = "Tidak boleh kosong"
			} else {
				message = "Min nilai 0"
			}

			errRes[strings.ToLower(err.Field())] = message
		}
		//return c.Status(http.StatusBadRequest).JSON(map[string]any{
		//	"erors": errRes,
		//})

		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"status": http.StatusBadRequest,
			"error":  errRes,
		})

	}

	_, error := database.DB.Exec("UPDATE books SET title = $1, description = $2, isbn = $3, genre = $4, author = $5, stock = $6, publish_date = $7 WHERE id = $8",
		book.Title,
		book.Description,
		book.Isbn,
		book.Genre,
		book.Author,
		book.Stock,
		book.PublishDate,
		id)

	if err != nil {
		return err
	}
	if error != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"status": http.StatusBadRequest,
			"error":  "gagal update data",
		})
	}

	return c.Status(http.StatusCreated).JSON(map[string]any{
		"status":  http.StatusCreated,
		"message": "Book update successfully",
	})
}
