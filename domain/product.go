package domain

import "time"

type Product struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Genr        string    `json:"genr"`
	Type        uint      `json:"type"`
	Price       float64   `json:"price"`
	Discount    float64   `json:"discount"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	Count       int       `json:"count"`
	Power       int       `json:"power"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProductCriteria struct {
	Name     *string `json:"product_name"`
	Genr     *string `json:"genr"`
	PageSize *int    `json:"page_size"`
	Page     *int    `json:"page"`
	Offset   *int    `json:"offset"`
}

type ProductRepository interface {
	Store(m *Product) error
	Fetch(ctr *ProductCriteria) (*Product, error)
	FetchProductCount(ctr *ProductCriteria) (*uint64, error)
}

type ProductUseCase interface {
	Store(m *Product) error
	Fetch(ctr *ProductCriteria) (*Product, error)
	FetchProductCount(ctr *ProductCriteria) (*uint64, error)
}
