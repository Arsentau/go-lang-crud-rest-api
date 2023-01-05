package repository

import (
	"errors"
	"restAPI/CRUD/db"
	"restAPI/CRUD/pkg/models"
	"restAPI/CRUD/pkg/types"

	"gorm.io/gorm"
)

func CreateMovieRepository(newMovie *models.Movie) error {
	NewMovie := db.DB.Save(newMovie)
	return NewMovie.Error
}

func GetAllMoviesRepository(bought string) ([]types.Movie, error, int) {
	var movies []types.Movie
	var code int = 200
	var result *gorm.DB
	if bought == "true" || bought == "false" {
		result = db.DB.Find(&movies, "bought = ?", bought)
		return movies, result.Error, code
	}
	if bought == "" {
		result = db.DB.Find(&movies)
		return movies, result.Error, code
	}
	code = 400
	return movies, errors.New("Invalid query params"), code

}

func GetMovieRepository(id string) (*types.Movie, error) {
	var movie types.Movie
	result := db.DB.First(&movie, "id = ?", id)
	return &movie, result.Error
}
