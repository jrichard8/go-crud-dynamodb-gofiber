package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
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

	param := &dynamodb.ListTablesInput{
		Limit: aws.Int32(10),
	}

	tables, err := client.ListTables(context.TODO(), param)
	if err != nil {
		return
	}
	for i, s := range tables.TableNames {
		println(i, s)
	}
}
