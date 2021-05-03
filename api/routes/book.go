package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-crud-dynamodb-gofiber/pkg/book"
	"go-crud-dynamodb-gofiber/pkg/entities"
)

func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", getBooks(service))
	app.Post("/books", addBook(service))
	app.Put("/books", updateBook(service))
	app.Delete("/books", removeBook(service))
}

func removeBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.BookKey
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.RemoveBook(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func updateBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.UpdateBook(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func addBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.InsertBook(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func getBooks(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchBooks()
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"books":  fetched,
		})
	}
}
