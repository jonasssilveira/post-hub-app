package error_handler

import (
	"encoding/json"
)

type ApiError struct {
	Code    int
	Message string
}

func (error *ApiError) ReturnJson() ([]byte, error) {
	return json.Marshal(error)
}
