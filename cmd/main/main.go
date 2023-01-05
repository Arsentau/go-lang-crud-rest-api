package main

import (
	"fmt"
	"log"
	"net/http"
	"restAPI/CRUD/db"
	"restAPI/CRUD/internal/models"
	"restAPI/CRUD/internal/routes"
	"restAPI/CRUD/internal/types"

	"github.com/spf13/viper"
)

func main() {
	// Get .env values with Viper
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	var config types.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// Connect to DB
	db.Database(config.DbHost, config.DbName, config.DbUser, config.DbPassword, config.DbPort)

	// Make migrations of movies
	db.DB.AutoMigrate(models.Movie{})

	r := routes.NewRouter()
	port := ":" + config.Port
	fmt.Printf("Starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
