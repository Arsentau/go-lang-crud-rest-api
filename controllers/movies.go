package controllers

import (
	"encoding/json"
	"net/http"
	"restAPI/CRUD/repository"
)

func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(repository.Movies)
}
