package handlers

import (
	"PostHubApp/domain/use_case/entity"
	"PostHubApp/domain/use_case/repository"
	"PostHubApp/posthubapi/handlers/error_handler"
	"encoding/json"
	"net/http"
)

type ApiGetAll struct {
	db repository.DB[entity.Post]
}

func NewApiGetAll(repo repository.DB[entity.Post]) *ApiGetAll {
	return &ApiGetAll{
		db: repo,
	}
}

func (api *ApiGetAll) Handler(w http.ResponseWriter, r *http.Request) {

	post, err := json.Marshal(api.db.FindAll())

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
