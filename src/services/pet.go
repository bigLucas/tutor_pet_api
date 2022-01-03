package services

import (
	"fmt"
	"tutor-pet-api/src/models"
	"tutor-pet-api/src/repositories"

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
	return pet, nil
}

func (p *PetService) FindOne(id string) (models.Pet, error) {
	//todo: remove this var
	mock_res := models.Pet{
		Name:    "Lily",
		Age:     7,
		TutorID: 1,
	}

	res, err := p.repo.FindOne(id)
	if err != nil {
		return mock_res, err
	}
	fmt.Println("debug 1", res.Items)
	// todo: map dynamodb response to model
	return mock_res, nil
}

func (p *PetService) FindMany() ([]models.Pet, error) {
	//todo: remove this var
	mock_res := []models.Pet{}

	res, err := p.repo.FindMany()
	if err != nil {
		return mock_res, err
	}
	fmt.Println("debug 2", res.Items)
	// todo: map dynamodb response to model
	return mock_res, nil
}

func (p *PetService) Update(id string, pet models.Pet) (models.Pet, error) {
	//todo: remove this var
	mock_res := models.Pet{}

	res, err := p.repo.CreateOrUpdate(id, pet)
	if err != nil {
		return mock_res, err
	}
	fmt.Println("debug 3", res)
	// todo: map dynamodb response to model
	return mock_res, nil
}

func (p *PetService) Delete(id string) error {
	_, err := p.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
