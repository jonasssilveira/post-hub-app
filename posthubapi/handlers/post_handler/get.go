package handlers

import (
	"PostHubApp/domain/use_case/entity"
	"PostHubApp/domain/use_case/repository"
	"PostHubApp/posthubapi/handlers/error_handler"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiGet[T entity.Post] struct {
	db repository.DB[T]
}

func NewApiGet[T entity.Post](repo repository.DB[T]) *ApiGet[T] {
	return &ApiGet[T]{
		db: repo,
	}
}

func (api *ApiGet[T]) Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post, err := json.Marshal(api.db.Get(id))

	if err != nil {
		apiError := &error_handler.ApiError{
			Code:    http.StatusNotFound,
			Message: "Resource not found",
		}

		http.Error(w, apiError.Message, apiError.Code)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(post)

}
