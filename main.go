package main

import (
	"fmt"
	"log"
	"net/http"
	"restAPI/CRUD/routes"
)

var port string = ":8000"

// func getAllMovies(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	json.NewEncoder(w).Encode(repository.Movies)
// }

func main() {
	// repository.Movies = append(repository.Movies, factories.MoviesFactory("Matrix", "Sam", "Morfeo"))
	// repository.Movies = append(repository.Movies, factories.MoviesFactory("Africa mia", "Joe", "Doe"))
	r := routes.NewRouter()
	fmt.Printf("Starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
