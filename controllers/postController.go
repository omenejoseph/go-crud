package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omenejoseph/go-crud/models"
	token_utils "github.com/omenejoseph/go-crud/utils"
)

var PostRequest struct {
	Body  string
	Title string
}

func PostCreate(c *gin.Context) {

	c.Bind(&PostRequest)

	user_id, err := token_utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := models.GetUserByID(user_id)

	fmt.Println(user.Email) // user email

	post := models.Post{Title: PostRequest.Title, Body: PostRequest.Body}

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

	c.Bind(&PostRequest)

	db.Model(&post).Updates(models.Post{Title: PostRequest.Title, Body: PostRequest.Body})

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
