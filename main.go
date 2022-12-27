package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restAPI/CRUD/factories"
	"restAPI/CRUD/models"

	"github.com/gorilla/mux"
)

var movies []models.Movie
var port string = ":8000"

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, factories.MoviesFactory("Matrix", "Sam", "Morfeo"))
	movies = append(movies, factories.MoviesFactory("Africa mia", "Joe", "Doe"))

	r.HandleFunc("/movies", getAllMovies).Methods("GET")

	fmt.Printf("Starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
