package routes

import (
	"restAPI/CRUD/pkg/controllers"

	"github.com/gorilla/mux"
)

func addMainHandlers(r *mux.Router) {
	r.HandleFunc("/movies", controllers.GetAllMoviesController).Methods("GET")
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	addMainHandlers(r)
	return r
}
