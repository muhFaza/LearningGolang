package controllers

import (
	"DailyChallenge6/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// store data of warehouses
var warehouses []models.Warehouse

func GetWarehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// dummy data

	// if theres data in warehouses, return it
	// else return dummy data
	if len(warehouses) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(warehouses)
		return
	} else {
		warehouse := models.Warehouse{ID: 1, Name: "Warehouse 1", Address: "Address 1"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(warehouse)
	}
}

func CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var newWarehouse models.Warehouse
	err := json.NewDecoder(r.Body).Decode(&newWarehouse)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newWarehouse.ID = int64(len(warehouses) + 1)
	newWarehouse.CreatedAt = time.Now()
	newWarehouse.UpdatedAt = time.Now()
	warehouses = append(warehouses, newWarehouse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newWarehouse)
}
