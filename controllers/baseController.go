package controllers

import (
	"github.com/omenejoseph/go-crud/inititalizers"
	"gorm.io/gorm"
)

var db *gorm.DB


func init() {
	db = inititalizers.ConnectToDB()
}
