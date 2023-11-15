package errorshandler

import (
	"encoding/json"
	"fmt"
)

var jsonMarshalContent = json.Marshal

type ApiErrorNotFound struct {
	Code    int
	Message string
}

func (error ApiErrorNotFound) Error() string {
	marshal, err := jsonMarshalContent(error)
	if err != nil {
		return err.Error()
	}
	return string(marshal)
}

type CustomError struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Type      string `json:"type"`
	Err       string `json:"error"`
}

func NewError(status int, requestID, cusMsg, typeError, orgMsg string) CustomError {
	return CustomError{
		Status:    status,
		Message:   cusMsg,
		RequestID: requestID,
		Type:      typeError,
		Err:       orgMsg,
	}
}

func (e CustomError) Error() string {
	var answer string

	res, marshalErr := jsonMarshalContent(e)
	if marshalErr != nil {
		answer = fmt.Sprintf("error could not be marshaled. %s", marshalErr.Error())
	} else {
		answer = string(res)
	}

	return answer
}

func (e CustomError) GetStatus() string {
	return e.GetStatus()
}
