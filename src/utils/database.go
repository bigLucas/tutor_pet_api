package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	Client *dynamodb.Client
}

func (db *DynamoDBClient) Query(ctx context.Context, params *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return db.Client.Query(ctx, params)
}

func (db *DynamoDBClient) Scan(ctx context.Context, params *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return db.Client.Scan(ctx, params)
}

func (db *DynamoDBClient) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return db.Client.UpdateItem(ctx, params)
}

func (db *DynamoDBClient) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return db.Client.DeleteItem(ctx, params)
}
