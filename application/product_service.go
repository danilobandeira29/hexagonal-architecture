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

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	p := NewProduct()
	p.Name = name
	p.Price = price
	_, err := p.IsValid()
	if err != nil {
		return &Product{}, err
	}
	product, err := s.Persistence.Save(p)
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func (s *ProductService) Enable(p ProductInterface) (ProductInterface, error) {
	err := p.Enable()
	if err != nil {
		return &Product{}, err
	}
	product, err := s.Persistence.Save(p)
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}
