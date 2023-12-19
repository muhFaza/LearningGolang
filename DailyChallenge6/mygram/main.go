package main

import (
	"mygram/config"
	"mygram/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbConfig := config.GetDBConfig()
	dsn := dbConfig.GetDBURL()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrate the User model
	err = db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})
	if err != nil {
		panic("failed to auto migrate")
	}

	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := db.Create(&user).Error
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to register",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success register",
			"user":    user,
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingUser model.User
		err := db.Where("email = ?", user.Email).First(&existingUser).Error
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to login",
			})
			return
		}

		if existingUser.Password != user.Password {
			c.JSON(500, gin.H{
				"message": "failed to login",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success login",
			"user":    existingUser,
		})
	})

	r.GET("/photos", func(c *gin.Context) {
		var photos []model.Photo
		err := db.Find(&photos).Error
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to get photos",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success get photos",
			"photos":  photos,
		})
	})

	// get all comments
	r.GET("/comments", func(c *gin.Context) {
		var comments []model.Comment
		err := db.Find(&comments).Error
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to get comments",
			})
			return
		}

		c.JSON(200, gin.H{
			"message":  "success get comments",
			"comments": comments,
		})
	})

	// get all social medias
	r.GET("/social-medias", func(c *gin.Context) {
		var socialMedias []model.SocialMedia
		err := db.Find(&socialMedias).Error
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to get social medias",
			})
			return
		}

		c.JSON(200, gin.H{
			"message":      "success get social medias",
			"socialMedias": socialMedias,
		})
	})

	r.Run(":8080")
}
