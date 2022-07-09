package repository

import (
	"github.com/rifatikbal/E-Com-PMS/domain"
	"github.com/rifatikbal/E-Com-PMS/internal/conn"
)

type product struct {
	*conn.DB
}

func New(conn *conn.DB) domain.ProductRepository {
	return &product{
		conn,
	}

}

func (p *product) Store(m *domain.Product) error {
	if err := p.GormDB.Create(m).Error; err != nil {
		return err
	}

	return nil
}

func (p *product) Fetch(ctr *domain.ProductCriteria) (*domain.Product, error) {
	product := &domain.Product{}

	if err := p.GormDB.
		Table("products").
		Where("name = ?",
			*ctr.Name).
		Scan(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
