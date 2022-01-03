package repositories

import (
	"context"
	"tutor-pet-api/src/models"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PetRepository struct {
	database   *dynamodb.Client
	table_name *string
}

func NewPetRepository(database *dynamodb.Client, table_name *string) *PetRepository {
	return &PetRepository{
		database:   database,
		table_name: table_name,
	}
}

func (p *PetRepository) FindOne(id string) (*dynamodb.QueryOutput, error) {
	cond := expression.Key("id").Equal(expression.Value(id))
	expr, err := expression.NewBuilder().WithKeyCondition(cond).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 p.table_name,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
	}

	res, err := p.database.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PetRepository) FindMany() (*dynamodb.ScanOutput, error) {
	input := &dynamodb.ScanInput{
		TableName: p.table_name,
	}

	res, err := p.database.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PetRepository) CreateOrUpdate(id string, pet models.Pet) (*dynamodb.UpdateItemOutput, error) {
	update := expression.Set(
		expression.Name("name"),
		expression.Value(pet.Name),
	).Set(
		expression.Name("age"),
		expression.Value(pet.Age),
	).Set(
		expression.Name("tutor_id"),
		expression.Value(pet.TutorID),
	)

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName:                 p.table_name,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	res, err := p.database.UpdateItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PetRepository) Delete(id string) (*dynamodb.DeleteItemOutput, error) {
	cond := expression.Name("id").Equal(expression.Value(id))

	expr, err := expression.NewBuilder().WithCondition(cond).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName:                 p.table_name,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	}

	res, err := p.database.DeleteItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
