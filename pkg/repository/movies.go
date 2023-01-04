package repository

import (
	"restAPI/CRUD/db"
	"restAPI/CRUD/pkg/models"
)

func CreateMovieRepository(newMovie *models.Movie) error {
	NewMovie := db.DB.Save(newMovie)
	return NewMovie.Error
}

func GetAllMoviesRepository() ([]models.Movie, error) {
	var movies []models.Movie
	result := db.DB.Find(&movies)
	return movies, result.Error
}

func GetMovieRepository(id string) (*models.Movie, error) {
	var movie models.Movie
	result := db.DB.First(&movie, "id = ?", id)
	return &movie, result.Error
}
