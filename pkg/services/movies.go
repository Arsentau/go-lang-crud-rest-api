package services

import (
	"errors"
	"restAPI/CRUD/pkg/models"
	"restAPI/CRUD/pkg/repository"
	"restAPI/CRUD/pkg/types"

	"github.com/google/uuid"
)

func CreateMovieService(movie types.Movie) (types.Movie, error) {
	var err error
	movie.ID = uuid.NewString()
	NewMovie := models.Movie(movie)
	e := repository.CreateMovieRepository(&NewMovie)
	if e != nil {
		err = errors.New("Ups! Something went wrong :(")
	}
	return movie, err
}
