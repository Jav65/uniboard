package thread

import (
	"sort"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	GetSorted(sortBy string, search string) ([]Thread, error)
	GetByID(id int) (Thread, error)
	Create(thread Thread) (Thread, error)
	Update(thread Thread) (Thread, error)
	Delete(thread Thread) (Thread, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetSorted(sortBy string, search string) ([]Thread, error) {
	var threads []Thread
	err := r.db.Find(&threads).Error

	if search != "" {
		search = strings.ToLower(search) 
		filteredThreads := []Thread{}
		for _, thread := range threads {
			if strings.Contains(strings.ToLower(thread.Title), search) ||
				strings.Contains(strings.ToLower(thread.Content), search) {
				filteredThreads = append(filteredThreads, thread)
			}
		}
		threads = filteredThreads
	}

	switch sortBy {
	case "Most Likes":
		sort.Slice(threads, func(i, j int) bool { return threads[i].Likes > threads[j].Likes })
	case "Most Comments":
		sort.Slice(threads, func(i, j int) bool { return threads[i].Comments > threads[j].Comments })
	case "Most Views":
		sort.Slice(threads, func(i, j int) bool { return threads[i].Views > threads[j].Views })
	case "Most Recent":
		sort.Slice(threads, func(i, j int) bool { return threads[i].CreatedAt.After(threads[j].CreatedAt) })
	default:
		sort.Slice(threads, func(i, j int) bool { return threads[i].CreatedAt.After(threads[j].CreatedAt) })
	}
	return threads, err
}

func (r *repository) GetByID(id int) (Thread, error) {
	var thread Thread
	err := r.db.First(&thread, id).Error
	return thread, err
}

func (r *repository) Create(thread Thread) (Thread, error) {
	return thread, r.db.Create(&thread).Error
}

func (r *repository) Update(thread Thread) (Thread, error) {
	return thread, r.db.Save(thread).Error
}

func (r *repository) Delete(thread Thread) (Thread, error) {
	return thread, r.db.Delete(&thread).Error
}
