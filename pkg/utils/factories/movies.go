package factories

import (
	"restAPI/CRUD/pkg/models"

	"github.com/google/uuid"
)

func MoviesFactory(title string, directorName string, directorLastName string) models.Movie {
	uuid := uuid.New()
	return models.Movie{Id: uuid.String(), Title: title, Director: &models.Director{FirstName: directorName, LastName: directorLastName}}

}
