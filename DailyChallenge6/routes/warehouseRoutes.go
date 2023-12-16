package routes

import (
	"DailyChallenge6/controllers"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/warehouses", controllers.GetWarehouses).Methods("GET")
	router.HandleFunc("/warehouses", controllers.CreateWarehouse).Methods("POST")
	return router
}