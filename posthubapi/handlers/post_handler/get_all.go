package handlers

import (
	"PostHubApp/domain/use_case/repository"
	"PostHubApp/posthubapi/handlers/errorshandle"
	"encoding/json"
	"net/http"
)

type ApiGetAll struct {
	db repository.DB
}

func NewApiGetAll(repo repository.DB) *ApiGetAll {
	return &ApiGetAll{
		db: repo,
	}
}

func (api *ApiGetAll) Handler(w http.ResponseWriter, r *http.Request) error {

	allPost, err := api.db.FindAll(r.Context(), nil)
	if err != nil {
		return err
	}
	post, err := json.Marshal(allPost)

	if err != nil {
		return errorshandler.ApiErrorNotFound{
			Code:    http.StatusNotFound,
			Message: "Resource not found",
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write(post)
	return nil
}
