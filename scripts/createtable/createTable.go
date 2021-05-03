package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

func main() {
	//TODO: remove custom resolver to work with real instance of dynamodb
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: "http://localhost:8000",
		}, nil
	})

	defaultConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ca-central-1"), config.WithEndpointResolver(customResolver))

	if err != nil {
		return
	}
	client := dynamodb.NewFromConfig(defaultConfig)

	param := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Title"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("Author"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Title"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("Author"),
				KeyType:       types.KeyTypeRange,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("Books"),
	}

	_, err = client.CreateTable(context.TODO(), param)
	if err != nil {
		log.Fatalf("Fail to create table %v \n", err)
	}

	log.Printf("Table Book is created")
}
