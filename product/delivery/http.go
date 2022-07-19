package delivery

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rifatikbal/E-Com-PMS/domain"
	"github.com/rifatikbal/E-Com-PMS/internal/utils"
	"log"
	"net/http"
	"strconv"
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
		r.Get("/", handler.GetProduct)
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

	var ctr *domain.ProductCriteria
	name := r.URL.Query().Get("name")
	if name != "" {
		ctr.Name = &name
	}

	genr := r.URL.Query().Get("genr")
	if genr != "" {
		ctr.Genr = &genr
	}

	if r.URL.Query().Get("page") != "" {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			log.Println("could not parsee page")
			return
		}
		ctr.Page = &page
	}

	if r.URL.Query().Get("page_size") != "" {
		pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
		if err != nil {
			log.Println("could not parsee page_size")
			return
		}
		ctr.PageSize = &pageSize
	}

	var count *uint64
	if ctr.PageSize != nil {
		var err error
		count, err = p.productUseCase.FetchProductCount(ctr)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{
				"error": err.Error(),
			})
		}
	}

	var offset int
	if count != nil {
		offset = int(*ctr.PageSize) * (*ctr.Page - 1)

		ctr.Offset = &offset
	}

	product, err := p.productUseCase.Fetch(ctr)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}

	pagination := utils.Pagination{
		Total: int(*count),
	}

	resp := utils.Response{
		Data:       product,
		Pagination: pagination,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}
