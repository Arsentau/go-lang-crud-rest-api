package main

import (
	"fmt"
	"log"
	"net/http"
	"restAPI/CRUD/db"
	"restAPI/CRUD/pkg/routes"
	"restAPI/CRUD/pkg/types"

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
	port := ":" + config.Port
	db.Database(config.DbHost, config.DbName, config.DbUser, config.DbPassword, config.DbPort)

	r := routes.NewRouter()
	fmt.Printf("Starting server at port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
