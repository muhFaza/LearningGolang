package routes

import (
	"DailyChallenge6/controllers"
	"github.com/gorilla/mux"
)

func WarehouseRoutes(router *mux.Router) {
	router.HandleFunc("/warehouses", controllers.GetWarehouses).Methods("GET")
	router.HandleFunc("/warehouses", controllers.CreateWarehouse).Methods("POST")
}