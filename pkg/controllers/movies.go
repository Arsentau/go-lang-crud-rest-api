package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restAPI/CRUD/pkg/repository"
	"restAPI/CRUD/pkg/services"
	"restAPI/CRUD/pkg/types"
	utils "restAPI/CRUD/pkg/utils/errors"

	"github.com/gorilla/mux"
)

func GetAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	allMovies, err := repository.GetAllMoviesRepository()
	if err != nil {
		utils.ErrorResponseHandler(w, err)
		return
	}
	json.NewEncoder(w).Encode(&allMovies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-type", "application/json")
	movie, err := repository.GetMovieRepository(params["id"])
	if err != nil {
		utils.ErrorResponseHandler(w, err)
		return
	}
	json.NewEncoder(w).Encode(&movie)

}

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("bought")
	fmt.Print("BOUGHT", key)
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
