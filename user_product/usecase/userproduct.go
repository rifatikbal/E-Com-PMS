package usecase

import "github.com/rifatikbal/E-Com-PMS/domain"

type userProduct struct {
	userProductRepo domain.UserProductRepository
}

func New(userProductRepo domain.UserProductRepository) domain.UserProductUseCase {
	return &userProduct{
		userProductRepo: userProductRepo,
	}
}

func (p *userProduct) Store(m *domain.UserProduct) error {
	if err := p.userProductRepo.Store(m); err != nil {
		return err
	}

	return nil
}
