package routes

import (
	"DailyChallenge6/controllers"
	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) {
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProducts).Methods("POST")
}