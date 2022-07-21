package repository

import (
	"github.com/rifatikbal/E-Com-PMS/domain"
	"github.com/rifatikbal/E-Com-PMS/internal/conn"
)

type userProduct struct {
	*conn.DB
}

func New(conn *conn.DB) domain.UserProductRepository {
	return &userProduct{
		conn,
	}

}

func (p *userProduct) Store(m *domain.UserProduct) error {
	if err := p.GormDB.Create(m).Error; err != nil {
		return err
	}

	return nil
}
