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

func (p *product) Fetch(ctr *domain.ProductCriteria) (*domain.Product, error) {

	product, err := p.productRepo.Fetch(ctr)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *product) FetchProductCount(ctr *domain.ProductCriteria) (*uint64, error) {
	count, err := p.productRepo.FetchProductCount(ctr)
	if err != nil {
		return nil, err
	}

	return count, nil
}
