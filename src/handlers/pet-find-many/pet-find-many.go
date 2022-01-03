package main

import (
	"tutor-pet-api/src/repositories"
	"tutor-pet-api/src/services"

	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	const TABLE_NAME = "pet-table"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "{\"message\":\"Internal Server error!\"}", Headers: map[string]string{"Content-Type": "application/json"}}, nil
	}
	db_client := dynamodb.NewFromConfig(cfg)
	repository := repositories.NewPetRepository(db_client, aws.String(TABLE_NAME))
	service := services.NewPetService(repository)

	// calling the service
	_, err = service.FindMany()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "{\"message\":\"Internal Server error!\"}", Headers: map[string]string{"Content-Type": "application/json"}}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "{\"message\":\"Everything is good!\"}", Headers: map[string]string{"Content-Type": "application/json"}}, nil
}

func main() {
	lambda.Start(handler)
}