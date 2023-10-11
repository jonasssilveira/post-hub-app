package handlers

import (
	"PostHubApp/domain/use_case/entity"
	"PostHubApp/domain/use_case/repository"
	"PostHubApp/posthubapi/handlers/error_handler"
	"encoding/json"
	"io"
	"net/http"
)

type ApiSave struct {
	db repository.DB[entity.Post]
}

func NewApiSave(repo repository.DB[entity.Post]) *ApiSave {
	return &ApiSave{
		db: repo,
	}
}

func (api *ApiSave) Handler(w http.ResponseWriter, r *http.Request) {

	readBody, err := io.ReadAll(r.Body)

	if err != nil {
		apiError := &error_handler.ApiError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}

		http.Error(w, apiError.Message, apiError.Code)
		return
	}
	var post entity.Post

	if err = json.Unmarshal(readBody, &post); err != nil {
		apiError := &error_handler.ApiError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}

		http.Error(w, apiError.Message, apiError.Code)
		return
	}

	api.db.Merge(post)

	w.WriteHeader(http.StatusCreated)

}
