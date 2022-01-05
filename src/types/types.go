package types

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Response struct {
	Messsage string `json:"message"`
	Key      string `json:"key,omitempty"`
}

type Database interface {
	Query(ctx context.Context, params *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	Scan(ctx context.Context, params *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
}
