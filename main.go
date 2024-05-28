package main

import (
	"NGC_AVENGER/config"
	"NGC_AVENGER/handlers"
	"NGC_AVENGER/middleware"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	config.DB = database

	router := httprouter.New()

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.GET("/recipe", middleware.AuthMiddleware(handlers.GetAllRecipe))

	log.Fatal(http.ListenAndServe(":8080", router))

}
