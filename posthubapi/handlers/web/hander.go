package web

import (
	errorshandler "PostHubApp/posthubapi/handlers/errorshandle"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"reflect"
)

const (
	errInternalServer = "unexpected server error"
	contentType       = "application/json; charset=utf-8"
)

var jsonMarshalContent = json.Marshal

var requestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "requests_total",
		Help: "Total number of requests.",
	},
	[]string{"request_id"},
)

func HandleFunc(fn func(c *gin.Context) error) func(c *gin.Context) {

	return func(c *gin.Context) {
		err := fn(c)

		if err != nil {
			status, errEncoder := errorEncoder(c.Copy(), err, c.Copy())
			c.JSON(status, errEncoder)
		}
		return
	}
}

func errorEncoder(ctx context.Context, err error, c *gin.Context) (int, []byte) {
	requestID := generateRequestID()
	mappedError := mapperError(ctx, requestID, err)
	requestCounter.WithLabelValues(requestID).Inc()

	c.Header("Content-Type", contentType)

	if mappedError.Status != http.StatusInternalServerError {
		res, marshalErr := jsonMarshalContent(mappedError)
		if marshalErr != nil {
			log.Fatal(ctx, "error could not be marshaled.", err)
			mappedError.Status = http.StatusInternalServerError
			mappedError.Message = fmt.Sprintf("error could not be marshaled. %s", marshalErr.Error())
			mappedError.Type = getType(marshalErr)
			mappedError.Err = marshalErr.Error()
			res, _ = json.Marshal(mappedError)
		}

		return mappedError.Status, res
	}

	c.Header("CustomAttribute", requestID)
	return mappedError.Status, nil
}

func mapperError(ctx context.Context, requestID string, err error) errorshandler.CustomError {
	switch err.(type) {
	case errorshandler.ApiErrorNotFound:
		return errorshandler.NewError(http.StatusBadRequest,
			requestID,
			"object not found",
			getType(err),
			err.Error())
	default:
		log.Fatal(ctx, "logging InternalServerError", fmt.Sprintln("error_description", err.Error()))
		return errorshandler.NewError(http.StatusInternalServerError,
			requestID,
			errInternalServer,
			getType(err),
			err.Error())
	}
}

func getType(myvar interface{}) string {
	return reflect.TypeOf(myvar).String()
}

func generateRequestID() string {
	return uuid.NewString()
}
