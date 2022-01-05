//go:build unit

package repositories

import (
	"context"
	"testing"
	"tutor-pet-api/src/models"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoDBClientMock struct {
	updateItemWasCalled bool
	queryWasCalled      bool
	scanWasCalled       bool
	deleteItemWasCalled bool
}

func (db *dynamoDBClientMock) Query(ctx context.Context, params *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	db.queryWasCalled = true
	return &dynamodb.QueryOutput{}, nil
}

func (db *dynamoDBClientMock) Scan(ctx context.Context, params *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	db.scanWasCalled = true
	return &dynamodb.ScanOutput{}, nil
}

func (db *dynamoDBClientMock) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	db.updateItemWasCalled = true
	return &dynamodb.UpdateItemOutput{}, nil
}

func (db *dynamoDBClientMock) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	db.deleteItemWasCalled = true
	return &dynamodb.DeleteItemOutput{}, nil
}

var mock_table_name = "dummy_table_name"
var mock_id = "mock_id"
var mock_pet = models.Pet{
	Name:    "luke",
	Age:     18,
	TutorID: 1818,
}

func TestCreateOrUpdate(t *testing.T) {
	// mock-spy
	dynamodb_client_mock := &dynamoDBClientMock{}

	// execution
	repo := NewPetRepository(dynamodb_client_mock, &mock_table_name)
	got, _ := repo.CreateOrUpdate(mock_id, mock_pet)

	// assertions
	if got == nil {
		t.Errorf("Want 'defined', got '%v'", got)
	}
	if !dynamodb_client_mock.updateItemWasCalled {
		t.Errorf("Expected to call 'UpdateItem' of 'Database' in 'CreateOrUpdate', but it was not")
	}
}

func TestFindOne(t *testing.T) {
	// mock-spy
	dynamodb_client_mock := &dynamoDBClientMock{}

	// execution
	repo := NewPetRepository(dynamodb_client_mock, &mock_table_name)
	got, _ := repo.FindOne(mock_id)

	// assertions
	if got == nil {
		t.Errorf("Want 'defined', got '%v'", got)
	}
	if !dynamodb_client_mock.queryWasCalled {
		t.Errorf("Expected to call 'Query' of 'Database' in 'FindOne', but it was not")
	}
}

func TestFindMany(t *testing.T) {
	// mock-spy
	dynamodb_client_mock := &dynamoDBClientMock{}

	// execution
	repo := NewPetRepository(dynamodb_client_mock, &mock_table_name)
	got, _ := repo.FindMany()

	// assertions
	if got == nil {
		t.Errorf("Want 'defined', got '%v'", got)
	}
	if !dynamodb_client_mock.scanWasCalled {
		t.Errorf("Expected to call 'Scan' of 'Database' in 'FindMany', but it was not")
	}
}

func TestDelete(t *testing.T) {
	// mock-spy
	dynamodb_client_mock := &dynamoDBClientMock{}

	// execution
	repo := NewPetRepository(dynamodb_client_mock, &mock_table_name)
	got, _ := repo.Delete(mock_id)

	// assertions
	if got == nil {
		t.Errorf("Want 'defined', got '%v'", got)
	}
	if !dynamodb_client_mock.deleteItemWasCalled {
		t.Errorf("Expected to call 'Delete' of 'Database' in 'Delete', but it was not")
	}
}
