package utils

import (
	"encoding/json"
	"net/http"
	"restAPI/CRUD/pkg/types"
)

func ErrorResponseHandler(w http.ResponseWriter, err error) http.ResponseWriter {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	message := types.ErrorHttpResponse{Message: err.Error()}
	json.NewEncoder(w).Encode(message)
	return w
}
