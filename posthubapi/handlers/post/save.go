package handlers

import (
	"PostHubApp/adapter/producer"
	"PostHubApp/domain/dto"
	"PostHubApp/domain/entity"
	"PostHubApp/domain/repository"
	"PostHubApp/domain/use_case"
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
	messagingProducer := producer.NewProducer()
	commentService := use_case.NewServicePost(api.db, messagingProducer)
	commentService.SavePost(c, post)
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
