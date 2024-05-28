package main

import (
	"NGC_AVENGER/config"
	"NGC_AVENGER/handlers"
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

	log.Fatal(http.ListenAndServe(":8080", router))

}
