package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	p, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
