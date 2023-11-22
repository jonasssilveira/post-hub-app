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

func NewApiSaveComment(repo repository.DB) *ApiSave {
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
	var commentDTO dto.Comment

	if err = json.Unmarshal(readBody, &commentDTO); err != nil {
		apiError := &errorshandler.ApiErrorNotFound{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
		return apiError
	}
	comment := entity.NewComment(commentDTO.PostID, commentDTO.UserID, commentDTO.Message)
	err = api.db.MergeComment(c, comment)
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
