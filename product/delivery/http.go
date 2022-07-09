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
		r.Get("/{name}", handler.GetProduct)
	})

}

func (p *ProductHandler) StoreProduct(w http.ResponseWriter, r *http.Request) {
	product := &domain.Product{}
	log.Println("redirected")
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "couldn't parse body",
		})
		return
	}

	if err := p.productUseCase.Store(product); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
	}
}

func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "name") == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "name required",
		})
		return
	}

	productName := chi.URLParam(r, "name")

	ctr := &domain.ProductCriteria{
		Name: &productName,
	}

	product, err := p.productUseCase.Fetch(ctr)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, product)
}
