package Repositories

import (
	"enCodeTest/src/models"
	"gorm.io/gorm"
	"log"
)

type GormPersonRepository struct {
	db *gorm.DB
}

func NewGormPersonRepository(db *gorm.DB) *GormPersonRepository {
	return &GormPersonRepository{db}
}

func (r *GormPersonRepository) Create(person *models.Person) error {
	return r.db.Create(person).Error
}

func (r *GormPersonRepository) GetByID(id int) (*models.Person, error) {
	var person models.Person
	if err := r.db.First(&person, id).Error; err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *GormPersonRepository) GetAll(limit, offset int, search string) ([]models.Person, error) {
	var persons []models.Person
	query := r.db.Limit(limit).Offset(offset)
	if search != "" {
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if err := query.Find(&persons).Error; err != nil {
		return nil, err
	}
	log.Println(persons)

	return persons, nil
}

func (r *GormPersonRepository) Update(person *models.Person) error {
	return r.db.Save(person).Error
}

func (r *GormPersonRepository) Delete(id int) error {
	return r.db.Delete(&models.Person{}, id).Error
}
