package repository

import (
	"errors"
	"restAPI/CRUD/db"
	"restAPI/CRUD/internal/models"
	"restAPI/CRUD/internal/types"

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

func Delete(id string) (error, int) {
	var code int = 204
	result := db.DB.Delete(&models.Movie{}, "id = ?", id)
	if result.Error != nil {
		code = 400
	}
	return result.Error, code
}
