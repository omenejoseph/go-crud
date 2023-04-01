package main

import (
	"github.com/omenejoseph/go-crud/inititalizers"
	"github.com/omenejoseph/go-crud/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = inititalizers.ConnectToDB()
}

func main() {
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.User{})
}
