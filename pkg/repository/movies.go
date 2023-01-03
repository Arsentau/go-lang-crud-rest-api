package repository

import (
	"restAPI/CRUD/db"
	"restAPI/CRUD/pkg/models"
)

func CreateMovieRepository(newMovie *models.Movie) error {
	NewMovie := db.DB.Save(newMovie)
	return NewMovie.Error
}
