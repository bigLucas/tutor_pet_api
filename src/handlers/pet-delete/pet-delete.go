package main

import (
	"tutor-pet-api/src/repositories"
	"tutor-pet-api/src/services"
	"tutor-pet-api/src/types"
	"tutor-pet-api/src/utils"

	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db_client, err := utils.InitDBClient()
	if err != nil {
		return utils.BuildInternalServerErrorRes(err)
	}
	repository := repositories.NewPetRepository(db_client, aws.String(utils.TableName))
	service := services.NewPetService(repository)
	id := event.PathParameters["id"]

	// calling the service
	err = service.Delete(id)
	if err != nil {
		if err.Error() == "not found" {
			return utils.BuildNotFoundErrorRes(err)
		}
		return utils.BuildInternalServerErrorRes(err)
	}

	return utils.BuildRes(http.StatusOK, types.Response{Messsage: "Successfully deleted"})
}

func main() {
	lambda.Start(handler)
}
