package comment

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Comment, error)
	GetAllByThreadID(id int) ([]Comment, error)
	Create(comment Comment) (Comment, error)
	Update(comment Comment) (Comment, error)
	Delete(comment Comment) (Comment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Comment, error) {
	var comments []Comment
	err := r.db.Find(&comments).Error
	return comments, err
}

func (r *repository) GetAllByThreadID(id int) ([]Comment, error) {
	var comment []Comment
	err := r.db.Find(&comment, "thread_id = ?", id).Error
	return comment, err
}

func (r *repository) Create(comment Comment) (Comment, error) {
	return comment, r.db.Create(&comment).Error
}

func (r *repository) Update(comment Comment) (Comment, error) {
	return comment, r.db.Save(comment).Error
}

func (r *repository) Delete(comment Comment) (Comment, error) {
	return comment, r.db.Delete(&comment).Error
}
