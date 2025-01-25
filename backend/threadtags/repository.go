package threadtags

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]ThreadTags, error)
	GetByID(id int) (ThreadTags, error)
	GetByThreadID(thread_id int) ([]ThreadTags, error)
	GetByTagID(tag_id int) (ThreadTags, error)
	Create(threadtags ThreadTags) (ThreadTags, error)
	Update(threadtags ThreadTags) (ThreadTags, error)
	Delete(threadtags ThreadTags) (ThreadTags, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]ThreadTags, error) {
	var threadtagss []ThreadTags
	err := r.db.Find(&threadtagss).Error
	return threadtagss, err
}

func (r *repository) GetByID(id int) (ThreadTags, error) {
	var threadtags ThreadTags
	err := r.db.First(&threadtags, id).Error
	return threadtags, err
}

func (r *repository) GetByThreadID(thread_id int) ([]ThreadTags, error) {
	var threadtags []ThreadTags
	err := r.db.Find(&threadtags, "thread_id = ?", thread_id).Error
	return threadtags, err
}

func (r *repository) GetByTagID(tag_id int) (ThreadTags, error) {
	var threadtags ThreadTags
	err := r.db.First(&threadtags, "tag_id = ?", tag_id).Error
	return threadtags, err
}

func (r *repository) Create(threadtags ThreadTags) (ThreadTags, error) {
	return threadtags, r.db.Create(&threadtags).Error
}

func (r *repository) Update(threadtags ThreadTags) (ThreadTags, error) {
	return threadtags, r.db.Save(threadtags).Error
}

func (r *repository) Delete(threadtags ThreadTags) (ThreadTags, error) {
	return threadtags, r.db.Delete(&threadtags).Error
}
