package handlers

import (
	"PostHubApp/domain/use_case/dto"
	"PostHubApp/domain/use_case/entity"
	"PostHubApp/domain/use_case/repository"
	"PostHubApp/posthubapi/handlers/errorshandle"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ApiSave struct {
	db repository.DB
}

func NewApiSavePost(repo repository.DB) *ApiSave {
	return &ApiSave{
		db: repo,
	}
}

func (api *ApiSave) Handler(c *gin.Context) error {

	readBody, err := io.ReadAll(c.Request.Body)

	if err != nil {
		apiError := &errorshandler.ApiErrorNotFound{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}

		return apiError
	}
	var postDTO dto.Post

	if err = json.Unmarshal(readBody, &postDTO); err != nil {
		apiError := &errorshandler.ApiErrorNotFound{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
		return apiError
	}
	post := entity.NewPost(postDTO.Title, postDTO.Message, postDTO.UserID)
	err = api.db.MergePost(c, post)
	if err != nil {
		apiError := &errorshandler.ApiErrorNotFound{
			Code:    http.StatusNotFound,
			Message: "it was not possible save or update",
		}

		http.Error(c.Writer, apiError.Message, apiError.Code)
	}

	c.Writer.WriteHeader(http.StatusCreated)
	return nil
}
