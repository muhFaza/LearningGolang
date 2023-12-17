package controllers

import (
	"DailyChallenge6/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// store data of warehouses
var warehouses []models.Warehouse

func GetWarehouses(c *gin.Context) {
	if len(warehouses) == 0 {
		c.JSON(http.StatusNotFound, []models.Warehouse{{ID: 1, Name: "Dummy Warehouse", Address: "Dummy Address"}})
	} else {
		c.JSON(http.StatusOK, warehouses)
	}
}

func CreateWarehouse(c *gin.Context) {
	var newWarehouse models.Warehouse
	err := c.ShouldBindJSON(&newWarehouse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		newWarehouse.ID = int64(len(warehouses) + 1)
		newWarehouse.CreatedAt = time.Now()
		newWarehouse.UpdatedAt = time.Now()
		warehouses = append(warehouses, newWarehouse)

		c.JSON(http.StatusCreated, newWarehouse)
	}
}

func GetWarehouseById(c *gin.Context) {
	ID := c.Param("id")
	for _, warehouse := range warehouses {
		if fmt.Sprintf("%v", warehouse.ID) == ID {
			c.JSON(http.StatusOK, warehouse)
			return
		}
	}

	notFoundStr := fmt.Sprintf("Warehouse with ID %v not found", ID)
	c.JSON(http.StatusNotFound, gin.H{"message": notFoundStr})
}