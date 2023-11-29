package handlers

import (
	"PostHubApp/domain/repository"
	"PostHubApp/posthubapi/handlers/errorshandle"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiGetAll struct {
	db repository.DB
}

func NewApiGetAllComment(repo repository.DB) *ApiGetAll {
	return &ApiGetAll{
		db: repo,
	}
}

func (api *ApiGetAll) Handler(c *gin.Context) error {

	allPost, err := api.db.FindAllComment(c, nil)
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
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(post)
	return nil
}
