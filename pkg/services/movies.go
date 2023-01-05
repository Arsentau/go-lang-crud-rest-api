package services

import (
	"errors"
	"restAPI/CRUD/pkg/models"
	"restAPI/CRUD/pkg/repository"
	"restAPI/CRUD/pkg/types"

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
