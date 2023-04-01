package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omenejoseph/go-crud/inititalizers"
	"github.com/omenejoseph/go-crud/models"
	"gorm.io/gorm"
)

var db *gorm.DB

var Request struct {
	Body  string
	Title string
}

func init() {
	db = inititalizers.ConnectToDB()
}

func PostCreate(c *gin.Context) {

	c.Bind(&Request)

	post := models.Post{Title: Request.Title, Body: Request.Body}

	result := db.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post

	results := db.Find(&posts)

	if results.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func PostFind(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := db.Find(&post, id)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func PostUpdate(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := db.Find(&post, id)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Bind(&Request)

	db.Model(&post).Updates(models.Post{Title: Request.Title, Body: Request.Body})

	c.JSON(http.StatusAccepted, gin.H{
		"data": post,
	})
}

func PostDelete(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := db.Find(&post, id)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	db.Delete(&post)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "deleted",
	})
}
