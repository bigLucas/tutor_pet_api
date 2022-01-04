package main

import (
	"tutor-pet-api/src/models"
	"tutor-pet-api/src/repositories"
	"tutor-pet-api/src/services"
	"tutor-pet-api/src/utils"

	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pet := models.Pet{}

	if err := json.Unmarshal([]byte(event.Body), &pet); err != nil {
		return utils.BuildBadRequestErrorRes(err)
	}

	db_client, err := utils.InitDBClient()
	if err != nil {
		return utils.BuildInternalServerErrorRes(err)
	}

	repository := repositories.NewPetRepository(db_client, aws.String(utils.TableName))
	service := services.NewPetService(repository)

	// calling the service
	res, err := service.Create(pet)
	if err != nil {
		return utils.BuildInternalServerErrorRes(err)
	}
	return utils.BuildRes(http.StatusOK, res)
}

func main() {
	lambda.Start(handler)
}
