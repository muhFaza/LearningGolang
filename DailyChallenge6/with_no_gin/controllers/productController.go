package controllers

import (
	"DailyChallenge6/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// store data of warehouses
var productsData []models.Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(productsData) == 0 {
		temp := []models.Product{{ID: 0, Name: "Dummy Data", Category: "Dum Dum", Stock: 0, Price: 0}}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(temp)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(productsData)
	}
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reqBody models.Product
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", 400)
	} else {
		reqBody.ID = int64(len(productsData) + 1)
		reqBody.CreatedAt = time.Now()
		reqBody.UpdatedAt = time.Now()
		productsData = append(productsData, reqBody)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(reqBody)
	}
}
