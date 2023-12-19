package controllers

import (
	"DailyChallenge6/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// store data of warehouses
var productsData []models.Product

func GetProducts(c *gin.Context) {
	if len(productsData) == 0 {
		c.JSON(http.StatusNotFound, []models.Product{{ID: 1, Name: "Dummy Data", Category: "Dum Dum", Stock: 0, Price: 0}})
	} else {
		c.JSON(http.StatusOK, productsData)
	}
}

func CreateProducts(c *gin.Context) {
	var newProduct models.Product
	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		newProduct.ID = int64(len(productsData) + 1)
		newProduct.CreatedAt = time.Now()
		newProduct.UpdatedAt = time.Now()
		productsData = append(productsData, newProduct)

		c.JSON(http.StatusCreated, newProduct)
	}
}

func GetProductById (c *gin.Context) {
	ID := c.Param("id")
	for _, product := range productsData {
		if fmt.Sprintf("%v",product.ID) == ID {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	notFoundStr := fmt.Sprintf("No product with ID %v found", ID)
	c.JSON(http.StatusNotFound, gin.H{"message": notFoundStr})
}

func DeleteProductById(c *gin.Context){
	ID := c.Param("id")
	for index, product := range productsData {
		if fmt.Sprintf("%v", product.ID) == ID {
			productsData = append(productsData[:index], productsData[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "deleted": product})
			return
		}
	}
	
	notFoundStr := fmt.Sprintf("No product with ID %v found", ID)
	c.JSON(http.StatusNotFound, gin.H{"message": notFoundStr, "deleted": nil})
}