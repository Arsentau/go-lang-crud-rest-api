package controllers

import (
	"encoding/json"
	"net/http"
	"restAPI/CRUD/pkg/repository"
	"restAPI/CRUD/pkg/services"
	"restAPI/CRUD/pkg/types"
	utils "restAPI/CRUD/pkg/utils/errors"

	"github.com/gorilla/mux"
)

func GetAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	bought := r.FormValue("bought")
	w.Header().Set("Content-type", "application/json")

	allMovies, err, code := repository.GetAllMoviesRepository(bought)
	if err != nil {
		utils.ErrorResponseHandler(w, err, code)
		return
	}
	json.NewEncoder(w).Encode(&allMovies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-type", "application/json")
	movie, err := repository.GetMovieRepository(params["id"])
	if err != nil {
		utils.ErrorResponseHandler(w, err, 404)
		return
	}
	json.NewEncoder(w).Encode(&movie)

}

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie types.Movie
	json.NewDecoder(r.Body).Decode(&newMovie)
	savedMovie, err, code := services.CreateMovieService(newMovie)
	if err != nil {
		utils.ErrorResponseHandler(w, err, code)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedMovie)
}
