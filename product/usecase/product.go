package usecase

import "github.com/rifatikbal/E-Com-PMS/domain"

type product struct {
	productRepo domain.ProductRepository
}

func New(productRepo domain.ProductRepository) domain.ProductUseCase {
	return &product{
		productRepo: productRepo,
	}
}

func (p *product) Store(m *domain.Product) error {
	if err := p.productRepo.Store(m); err != nil {
		return err
	}

	return nil
}
