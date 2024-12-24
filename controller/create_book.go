package controller

import (
	"api_go_bwa/database"
	"api_go_bwa/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func CreateBookController(c *fiber.Ctx) error {
	var req dto.CreateBookRequest

	err := c.BodyParser(&req)

	validate := validator.New()
	validationErrors := validate.Struct(req)

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
		return c.Status(http.StatusBadRequest).JSON(map[string]any{
			"erors": errRes,
		})
	}

	ID := uuid.New()
	if err != nil {
		return err
	}

	row, err := database.DB.Exec("INSERT INTO books( title, description, isbn, genre, author, stock, publish_date, id) values( $1, $2, $3,$4, $5, $6, $7, $8) ",
		req.Title,
		req.Description,
		req.Isbn,
		req.Genre,
		req.Author,
		req.Stock,
		req.PublishDate,
		ID,
	)

	println(row)

	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(map[string]any{
		"status":  http.StatusCreated,
		"message": "Book created successfully",
	})
}
