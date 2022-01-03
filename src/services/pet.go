package services

import (
	"errors"
	"tutor-pet-api/src/models"
	"tutor-pet-api/src/repositories"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/google/uuid"
)

type PetService struct {
	repo *repositories.PetRepository
}

func NewPetService(repository *repositories.PetRepository) *PetService {
	return &PetService{repo: repository}
}

func (p *PetService) Create(pet models.Pet) (models.Pet, error) {
	id := uuid.New()
	_, err := p.repo.CreateOrUpdate(id.String(), pet)
	if err != nil {
		return pet, err
	}
	pet.ID = id.String()
	return pet, nil
}

func (p *PetService) FindOne(id string) (models.Pet, error) {
	pet := models.Pet{}

	res, err := p.repo.FindOne(id)
	if err != nil {
		return pet, err
	}

	if len(res.Items) == 0 || len(res.Items) > 1 {
		return pet, errors.New("not found")
	}

	// mapping
	if attributevalue.UnmarshalMap(res.Items[0], &pet); err != nil {
		return pet, err
	}

	return pet, nil
}

func (p *PetService) FindMany() ([]models.Pet, error) {
	pets := []models.Pet{}

	res, err := p.repo.FindMany()
	if err != nil {
		return pets, err
	}

	// mapping
	if attributevalue.UnmarshalListOfMaps(res.Items, &pets); err != nil {
		return pets, err
	}

	return pets, nil
}

func (p *PetService) Update(id string, pet models.Pet) (models.Pet, error) {
	_, err := p.repo.CreateOrUpdate(id, pet)
	if err != nil {
		return pet, err
	}

	return pet, nil
}

func (p *PetService) Delete(id string) error {
	_, err := p.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
