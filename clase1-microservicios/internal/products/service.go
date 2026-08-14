package products

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetAll() []Product {
	return s.repository.GetAll()
}

func (s *Service) GetByID(id string) (Product, error) {
	return s.repository.GetByID(id)
}
