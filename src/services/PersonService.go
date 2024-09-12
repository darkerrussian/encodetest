package services

import (
	"enCodeTest/src/Repositories"
	"enCodeTest/src/models"
)

type PersonService struct {
	repo Repositories.PersonRepository
}

func NewPersonService(repo Repositories.PersonRepository) *PersonService {
	return &PersonService{repo}
}

func (s *PersonService) CreatePerson(person *models.Person) error {
	return s.repo.Create(person)
}

func (s *PersonService) GetPersonByID(id int) (*models.Person, error) {
	return s.repo.GetByID(id)
}

func (s *PersonService) GetAllPersons(limit, offset int, search string) ([]models.Person, error) {
	return s.repo.GetAll(limit, offset, search)
}

func (s *PersonService) UpdatePerson(person *models.Person) error {
	return s.repo.Update(person)
}

func (s *PersonService) DeletePerson(id int) error {
	return s.repo.Delete(id)
}
