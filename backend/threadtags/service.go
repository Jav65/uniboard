package threadtags

import "backend/tag"

type Service interface {
	GetAll() ([]ThreadTags, error)
	GetByID(id int) (ThreadTags, error)
	GetByThreadID(thread_id int) []string
	GetByTagID(tag_id int) (ThreadTags, error)
	Associate(tags_id []int, thread_id int) (ThreadTags, error)
	Create(threadtagsRequest ThreadTagsRequest) (ThreadTags, error)
	Update(id int, threadtagsRequest ThreadTagsRequest) (ThreadTags, error)
	Delete(id int) (ThreadTags, error)
}

type service struct {
	repository Repository
	tagSvc     tag.Service
}

func NewService(repository Repository, tagSvc tag.Service) *service {
	return &service{repository, tagSvc}
}

func (s *service) GetAll() ([]ThreadTags, error) {
	return s.repository.GetAll()
}

func (s *service) GetByID(id int) (ThreadTags, error) {
	return s.repository.GetByID(id)
}

func (s *service) GetByThreadID(thread_id int) []string {
	tag_names := []string{}
	tags, _ := s.repository.GetByThreadID(thread_id)
	for _, tag := range tags {
		name, _ := s.tagSvc.GetByID(tag.Tag_ID)
		tag_names = append(tag_names, name.Name)
	}
	return tag_names
}

func (s *service) GetByTagID(tag_id int) (ThreadTags, error) {
	return s.repository.GetByTagID(tag_id)
}

func (s *service) Associate(tags_id []int, thread_id int) (ThreadTags, error) {
	println(tags_id)
	res := ThreadTags{}
	for i := 0; i < len(tags_id); i++ {
		threadtags := ThreadTags{
			Thread_ID: thread_id,
			Tag_ID:    tags_id[i],
		}
		res, _ = s.repository.Create(threadtags)
	}
	return res, nil
}

func (s *service) Create(threadtagsRequest ThreadTagsRequest) (ThreadTags, error) {
	threadtags := ThreadTags{
		Thread_ID: threadtagsRequest.Thread_ID,
		Tag_ID:    threadtagsRequest.Tag_ID,
	}
	return s.repository.Create(threadtags)
}

func (s *service) Update(id int, threadtagsRequest ThreadTagsRequest) (ThreadTags, error) {
	threadtags, _ := s.repository.GetByID(id)

	threadtags.Thread_ID = threadtagsRequest.Thread_ID
	threadtags.Tag_ID = threadtagsRequest.Tag_ID

	return s.repository.Update(threadtags)
}

func (s *service) Delete(id int) (ThreadTags, error) {
	threadtags, _ := s.repository.GetByID(id)
	return s.repository.Delete(threadtags)
}
