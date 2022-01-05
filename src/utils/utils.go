package utils

import (
	"tutor-pet-api/src/types"

	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const TableName = "pet-table"

func InitDBClient() (*DynamoDBClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return &DynamoDBClient{Client: dynamodb.NewFromConfig(cfg)}, nil
}

func BuildRes(statusCode int, body interface{}) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{}
	errorMsg := "{\"message\":\"Internal Server error!\"}"

	body_res, err := json.Marshal(body)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Body = errorMsg
	} else {
		res.StatusCode = statusCode
		res.Body = string(body_res)
	}

	res.Headers = map[string]string{
		"Content-Type": "application/json",
	}
	return res, nil
}

func BuildBadRequestErrorRes(err error) (events.APIGatewayProxyResponse, error) {
	return buildErrorRes(http.StatusBadRequest, "Wrong request body", err)
}

func BuildNotFoundErrorRes(err error) (events.APIGatewayProxyResponse, error) {
	return buildErrorRes(http.StatusNotFound, "Not found", err)
}

func BuildInternalServerErrorRes(err error) (events.APIGatewayProxyResponse, error) {
	return buildErrorRes(http.StatusInternalServerError, "Internal server error", err)
}

func buildErrorRes(statusCode int, msg string, err error) (events.APIGatewayProxyResponse, error) {
	fmt.Println(err)
	return BuildRes(statusCode, types.Response{Messsage: msg})
}
