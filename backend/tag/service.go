package tag

type Service interface {
	GetAllName() ([]string, error)
	GetAll() ([]Tag, error)
	GetByID(id int) (Tag, error)
	Create(tagRequest TagRequest) (Tag, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllName() ([]string, error) {
	tags_name := []string{}
	tags, _ := s.repository.GetAll()
	for _, tag := range tags {
		tags_name = append(tags_name, tag.Name)
	}
	return tags_name, nil
}

func (s *service) GetAll() ([]Tag, error) {
	tags := []Tag{}
	t, _ := s.repository.GetAll()
	for _, tag := range t {
		tags = append(tags, Tag{
			Name: tag.Name,
			ID:   tag.ID,
			})
		}
	return tags, nil
}

func (s *service) GetByID(id int) (Tag, error) {
	return s.repository.GetByID(id)
}

func (s *service) Create(tagRequest TagRequest) (Tag, error) {
	tag := Tag{
		Name: tagRequest.Name,
	}
	return s.repository.Create(tag)
}
