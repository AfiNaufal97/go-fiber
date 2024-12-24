package main

import (
	"api_go_bwa/controller"
	"api_go_bwa/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	database.InitDb()
	app.Get("/", controller.GetAllBooks)
	app.Get("/:id", controller.GetDetailBook)
	app.Post("/post", controller.CreateBookController)
	app.Put("/:id", controller.UpdateBookController)
	app.Delete("/:id", controller.DeleteBookController)

	log.Fatal(app.Listen(":3000"))
}
