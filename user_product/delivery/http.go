package delivery

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rifatikbal/E-Com-PMS/domain"
	"net/http"
)

type UserProductHandler struct {
	userProductUseCase domain.UserProductUseCase
}

func New(r *chi.Mux, uc domain.UserProductUseCase) {

	handler := &UserProductHandler{
		userProductUseCase: uc,
	}

	r.Route("/product-user", func(r chi.Router) {
		r.Post("/", handler.userAction)

	})
}

func (up *UserProductHandler) userAction(w http.ResponseWriter, r *http.Request) {
	userProduct := &domain.UserProduct{}

	if err := json.NewDecoder(r.Body).Decode(userProduct); err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
	}

	if err := up.userProductUseCase.Store(userProduct); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
	}
}
