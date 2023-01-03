package controllers

import (
	"encoding/json"
	"net/http"
	"restAPI/CRUD/pkg/services"
	"restAPI/CRUD/pkg/types"
	utils "restAPI/CRUD/pkg/utils/errors"
)

// func GetAllMoviesHandler(w http.ResponseWriter, _ *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	json.NewEncoder(w).Encode(repository.Movies)
// }

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie types.Movie
	json.NewDecoder(r.Body).Decode(&newMovie)
	savedMovie, err := services.CreateMovieService(newMovie)
	if err != nil {
		utils.ErrorResponseHandler(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedMovie)
}
