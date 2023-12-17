package main

import (
	"DailyChallenge6/controllers"
  "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/warehouses", controllers.GetWarehouses)
	r.POST("/warehouses", controllers.CreateWarehouse)


	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProducts)

	r.Run()
}