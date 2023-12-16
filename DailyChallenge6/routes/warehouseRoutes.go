package routes

import (
	"DailyChallenge6/controllers"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/warehouses", controllers.GetWarehouses).Methods("GET")
	router.HandleFunc("/warehouses", controllers.CreateWarehouse).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProducts).Methods("POST")
	return router
}