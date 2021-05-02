package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gofiber/fiber/v2"
	"go-crud-dynamodb-gofiber/api/routes"
	"go-crud-dynamodb-gofiber/pkg/book"
	"log"
)

func main() {
	dyn := DatabaseConnection()
	bookRepo := book.NewRepo(dyn, "books")
	bookService := book.NewService(bookRepo)
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	_ = app.Listen(":8080")
}

func DatabaseConnection() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	return dynamodb.NewFromConfig(cfg)
}
