package tag

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Tag, error)
	GetByID(id int) (Tag, error)
	Create(tag Tag) (Tag, error)
	Update(tag Tag) (Tag, error)
	Delete(tag Tag) (Tag, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Tag, error) {
	var tags []Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

func (r *repository) GetByID(id int) (Tag, error) {
	var tag Tag
	err := r.db.First(&tag, id).Error
	return tag, err
}

func (r *repository) Create(tag Tag) (Tag, error) {
	return tag, r.db.Create(&tag).Error
}

func (r *repository) Update(tag Tag) (Tag, error) {
	return tag, r.db.Save(tag).Error
}

func (r *repository) Delete(tag Tag) (Tag, error) {
	return tag, r.db.Delete(&tag).Error
}
