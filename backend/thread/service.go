package thread

import (
	"backend/threadtags"
	"backend/user"
)

type Service interface {
	GetSorted(sortBy string, search string) ([]ThreadRequest, error)
	GetByID(id int) (ThreadRequest, error)
	Create(threadRequest ThreadRequest) (Thread, error)
	Update(id int, threadRequest ThreadRequest) (Thread, error)
	Delete(id int) (Thread, error)
}

type service struct {
	repository   Repository
	threadTagSvc threadtags.Service
	userSvc      user.Service
}

func NewService(repository Repository, threadTagSvc threadtags.Service, userSvc user.Service) *service {
	return &service{repository, threadTagSvc, userSvc}
}

func (s *service) GetSorted(sortBy string, search string) ([]ThreadRequest, error) {
	thread := []ThreadRequest{}
	threads, err := s.repository.GetSorted(sortBy, search)
	if err != nil {
		return nil, err
	}
	for _, t := range threads {
		Author, _ := s.userSvc.GetByID(t.AuthorID)
		Tags_name := s.threadTagSvc.GetByThreadID(t.ID)

		thread = append(thread, ThreadRequest{
			ID:          t.ID,
			Author_id:   t.AuthorID,
			Author_name: Author.Username,
			Title:       t.Title,
			Content:     t.Content,
			Tags_name:   Tags_name,
			CreatedAt:   t.CreatedAt,
			Likes:       t.Likes,
			Comments:    t.Comments,
			Views:       t.Views,
		})
	}
	return thread, nil
}

func (s *service) GetByID(id int) (ThreadRequest, error) {
	t, _ := s.repository.GetByID(id)
	Author, _ := s.userSvc.GetByID(t.AuthorID)
	Tags_name := s.threadTagSvc.GetByThreadID(t.ID)
	thread := ThreadRequest{
		ID:          t.ID,
		Author_id:   t.AuthorID,
		Author_name: Author.Username,
		Title:       t.Title,
		Content:     t.Content,
		Tags_name:   Tags_name,
	}
	return thread, nil
}

func (s *service) Create(threadRequest ThreadRequest) (Thread, error) {
	thread := Thread{
		Title:    threadRequest.Title,
		Content:  threadRequest.Content,
		AuthorID: threadRequest.Author_id,
		Likes:    0,
		Comments: 0,
		Views:    0,
	}

	thread, _ = s.repository.Create(thread)
	s.threadTagSvc.Associate(threadRequest.Tags, thread.ID)

	return thread, nil
}

func (s *service) Update(id int, threadRequest ThreadRequest) (Thread, error) {
	thread, _ := s.repository.GetByID(id)

	thread.Title = threadRequest.Title
	thread.Content = threadRequest.Content

	return s.repository.Update(thread)
}

func (s *service) Delete(id int) (Thread, error) {
	thread, _ := s.repository.GetByID(id)
	return s.repository.Delete(thread)
}
