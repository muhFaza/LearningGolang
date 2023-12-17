package main

import (
	"fmt"
	"net/http"
	"DailyChallenge6/routes"
  "github.com/gin-gonic/gin"
)

func main() {
	router := routes.InitializeRoutes()
	fmt.Println("Server is running on port 8080")

	// start the server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}