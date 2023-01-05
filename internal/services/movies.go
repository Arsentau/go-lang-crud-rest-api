package services

import (
	"errors"
	"restAPI/CRUD/internal/models"
	"restAPI/CRUD/internal/repository"
	"restAPI/CRUD/internal/types"

	"github.com/google/uuid"
)

func CreateMovieService(movie types.Movie) (types.Movie, error, int) {
	var err error
	movie.ID = uuid.NewString()
	NewMovie := models.Movie(movie)
	var code int = 201
	e := repository.CreateMovieRepository(&NewMovie)
	if e != nil {
		err = errors.New("Ups! Something went wrong :(")
		code = 400
	}
	return movie, err, code
}
