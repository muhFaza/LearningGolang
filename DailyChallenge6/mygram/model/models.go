package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `validate:"required"`
	Email        string `gorm:"unique" validate:"required,email"`
	Password     string `validate:"required,min=6"`
	Age          int    `validate:"required,gte=8"`
	Photos       []Photo
	SocialMedias []SocialMedia
}

type Photo struct {
	gorm.Model
	Title    string `validate:"required"`
	Caption  string
	PhotoURL string `gorm:"column:photo_url" validate:"required,url"`
	UserID   uint   `gorm:"column:user_id"`
	Comments []Comment
}

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id"`
	PhotoID uint   `gorm:"column:photo_id"`
	Message string `validate:"required"`
}

type SocialMedia struct {
	gorm.Model
	Name           string `validate:"required"`
	SocialMediaURL string `gorm:"column:social_media_url" validate:"required,url"`
	UserID         uint   `gorm:"column:user_id"`
}
