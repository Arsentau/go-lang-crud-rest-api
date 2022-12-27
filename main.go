package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restAPI/CRUD/factories"
	"restAPI/CRUD/repository"

	"github.com/gorilla/mux"
)

var port string = ":8000"

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(repository.Movies)
}

func main() {
	r := mux.NewRouter()
	repository.Movies = append(repository.Movies, factories.MoviesFactory("Matrix", "Sam", "Morfeo"))
	repository.Movies = append(repository.Movies, factories.MoviesFactory("Africa mia", "Joe", "Doe"))

	r.HandleFunc("/movies", getAllMovies).Methods("GET")

	fmt.Printf("Starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
