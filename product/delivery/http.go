package delivery

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rifatikbal/E-Com-PMS/domain"
	"log"
	"net/http"
)

type ProductHandler struct {
	productUseCase domain.ProductUseCase
}

func New(r *chi.Mux, uc domain.ProductUseCase) {

	handler := &ProductHandler{
		productUseCase: uc,
	}

	r.Route("/product", func(r chi.Router) {
		r.Post("/", handler.StoreProduct)

	})

}

func (p *ProductHandler) StoreProduct(w http.ResponseWriter, r *http.Request) {
	product := domain.Product{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "couldn't parse body",
		})
		return
	}

}
