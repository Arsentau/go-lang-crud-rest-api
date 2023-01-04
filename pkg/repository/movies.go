package repository

import (
	"restAPI/CRUD/db"
	"restAPI/CRUD/pkg/models"
	"restAPI/CRUD/pkg/types"
)

func CreateMovieRepository(newMovie *models.Movie) error {
	NewMovie := db.DB.Save(newMovie)
	return NewMovie.Error
}

func GetAllMoviesRepository() ([]types.Movie, error) {
	var movies []types.Movie
	result := db.DB.Find(&movies)
	return movies, result.Error
}

func GetMovieRepository(id string) (*types.Movie, error) {
	var movie types.Movie
	result := db.DB.First(&movie, "id = ?", id)
	return &movie, result.Error
}
