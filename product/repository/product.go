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

	q := p.GormDB.
		Table("products")
	if ctr.Name != nil {
		q = q.Where("name = ?", *ctr.Name)
	}
	if ctr.Genr != nil {
		q = q.Where("genr = ?", *ctr.Genr)
	}
	if ctr.PageSize != nil {
		q = q.Limit(*ctr.PageSize)
	}
	if ctr.Offset != nil {
		q = q.Offset(*ctr.Offset)
	}

	if err := q.Scan(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *product) Delete(ctr *domain.ProductCriteria) error {
	m := &domain.Product{}

	if err := p.GormDB.
		Delete(m).
		Where(`
			hashed_id = ?
			name = ?`,
			*ctr.Name,
			*ctr.HashedId).Error; err != nil {
		return err
	}
	return nil
}

func (p *product) Update(m *domain.Product) error {
	if err := p.GormDB.
		Save(m).
		Where(`
			hashed_id = ?`,
			m.HashedId).Error; err != nil {
		return err
	}
	return nil
}

func (p *product) FetchProductCount(ctr *domain.ProductCriteria) (*uint64, error) {
	var productCount *uint64
	q := p.GormDB.
		Table("products")
	if ctr.Genr != nil {
		q = q.Where("genr = ?", *ctr.Genr)
	}
	if err := q.Scan(&productCount).Error; err != nil {
		return nil, err
	}

	return productCount, nil
}
