package health

import (
	"net/http"
)

func IsHealth(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("value: UP"))
	return nil
}
