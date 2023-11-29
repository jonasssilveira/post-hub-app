package handlers

import (
	"PostHubApp/domain/repository"
	errorshandler "PostHubApp/posthubapi/handlers/errorshandle"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiGet struct {
	db repository.DB
}

func NewApiGetPost(repo repository.DB) *ApiGet {
	return &ApiGet{
		db: repo,
	}
}

func (api *ApiGet) Handler(c *gin.Context) error {

	id := c.Param("id")

	getPost, err := api.db.GetPost(c, id)

	if err != nil {
		return err
	}

	post, err := json.Marshal(getPost)

	if err != nil {
		return &errorshandler.ApiErrorNotFound{
			Code:    http.StatusInternalServerError,
			Message: "Resource not found",
		}
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(post)
	return nil
}
