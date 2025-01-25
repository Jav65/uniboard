package comment

import "backend/user"

type Service interface {
	GetAll() ([]Comment, error)
	GetAllByThreadID(id int) ([]CommentRequest, error)
	Create(commentRequest CommentRequest) (Comment, error)
}

type service struct {
	repository Repository
	userSvc    user.Service
}

func NewService(repository Repository, userSvc user.Service) *service {
	return &service{repository, userSvc}
}

func (s *service) GetAll() ([]Comment, error) {
	return s.repository.GetAll()
}

func (s *service) GetAllByThreadID(id int) ([]CommentRequest, error) {
	comment := []CommentRequest{}
	comments, err := s.repository.GetAllByThreadID(id)
	if err != nil {
		return nil, err
	}
	for _, c := range comments {
		Author, _ := s.userSvc.GetByID(c.Author_id)
		comment = append(comment, CommentRequest{
			ID:          c.ID,
			Thread_id:   c.Thread_id,
			Content:     c.Content,
			Author_id:   c.Author_id,
			Author_name: Author.Username,
			CreatedAt:   c.CreatedAt,
		})
	}
	return comment, nil
}

func (s *service) Create(commentRequest CommentRequest) (Comment, error) {
	comment := Comment{
		Thread_id: commentRequest.Thread_id,
		Content:   commentRequest.Content,
		Author_id: commentRequest.Author_id,
		CreatedAt: commentRequest.CreatedAt,
	}
	return s.repository.Create(comment)
}
