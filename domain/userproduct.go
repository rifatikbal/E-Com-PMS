package domain

type UserProduct struct {
	ID         uint64 `json:"id"`
	ProductId  uint64 `json:"productId"`
	UserID     uint64 `json:"userID"`
	ActionType string
}

type UserProductRepository interface {
	Store(m *UserProduct) error
}

type UserProductUseCase interface {
	Store(m *UserProduct) error
}
