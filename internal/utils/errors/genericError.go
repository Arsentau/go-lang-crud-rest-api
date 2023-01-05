package utils

import (
	"encoding/json"
	"net/http"
	"restAPI/CRUD/internal/types"
)

func ErrorResponseHandler(w http.ResponseWriter, err error, code int) http.ResponseWriter {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	message := types.ErrorHttpResponse{
		Code:    code,
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(message)
	return w
}
