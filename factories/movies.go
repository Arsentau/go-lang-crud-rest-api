package factories

import (
	"restAPI/CRUD/models"

	"github.com/google/uuid"
)

func MoviesFactory(title string, directorName string, directorLastName string) models.Movie {
	uuid := uuid.New()
	return models.Movie{Id: uuid.String(), Title: title, Director: &models.Director{FirstName: directorName, LastName: directorLastName}}

}

// func BulkMovieFactory(qty int) {
// 	for i := 1; i <= qty; i++ {
// 		MoviesFactory("Matrix", "Sam", "Morfeo")
// 	}
// 	return
// }
