package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gofiber/fiber/v2"
	"go-crud-dynamodb-gofiber/api/routes"
	"go-crud-dynamodb-gofiber/pkg/book"
	"log"
)

func main() {
	dyn := DatabaseConnection()
	bookRepo := book.NewRepo(dyn, "Books")
	bookService := book.NewService(bookRepo)
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture dynamoDB (go aws sdk V2) book shop!"))
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	_ = app.Listen(":3000")
}

func DatabaseConnection() *dynamodb.Client {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: "http://localhost:8000",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ca-central-1"), config.WithEndpointResolver(customResolver))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	return dynamodb.NewFromConfig(cfg)
}
