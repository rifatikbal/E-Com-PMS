package domain

type ProductUser struct {
	ID        uint64 `json:"id"`
	ProductId uint64 `json:"productId"`
	UserID    uint64 `json:"userID"`
}

type ProductUserRepository interface {
}

type ProductUserUseCase interface {
}
