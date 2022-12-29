package controllers

import (
	"encoding/json"
	"net/http"
	"restAPI/CRUD/pkg/repository"
)

func GetAllMoviesController(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(repository.Movies)
}
