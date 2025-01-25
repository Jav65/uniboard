package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	GetByEmail(email string) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetByID(id int) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *repository) GetByEmail(email string) (User, error) {
	var user User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *repository) Create(user User) (User, error) {
	return user, r.db.Create(&user).Error
}

func (r *repository) Update(user User) (User, error) {
	return user, r.db.Save(user).Error
}

func (r *repository) Delete(user User) (User, error) {
	return user, r.db.Delete(&user).Error
}
