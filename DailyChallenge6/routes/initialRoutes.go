package routes

import (
	"github.com/gorilla/mux"
)

func InitializeRoutes () *mux.Router {
	router := mux.NewRouter()
	WarehouseRoutes(router)
	ProductRoutes(router)

	return router
}