package repository

import (
	"TestTask/model"

	"gorm.io/gorm"
)

type PersonRepository struct {
	DB *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{DB: db}
}

func (r *PersonRepository) Create(person *model.Person) error {
	return r.DB.Create(person).Error
}

func (r *PersonRepository) GetAll(limit, offset int) ([]model.Person, error) {
	var people []model.Person
	err := r.DB.Limit(limit).Offset(offset).Find(&people).Error
	return people, err
}

func (r *PersonRepository) GetByID(id uint) (*model.Person, error) {
	var person model.Person
	err := r.DB.First(&person, id).Error
	return &person, err
}

func (r *PersonRepository) Update(person *model.Person) error {
	return r.DB.Save(person).Error
}

func (r *PersonRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Person{}, id).Error
}
