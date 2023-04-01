package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omenejoseph/go-crud/models"
	token_utils "github.com/omenejoseph/go-crud/utils"
	"golang.org/x/crypto/bcrypt"
)

var UserRequest struct {
	Name     string
	Email    string
	Password string
}

var LoginRequest struct {
	Email    string
	Password string
}

func UserCreate(c *gin.Context) {
	c.Bind(&UserRequest)

	hash, err := bcrypt.GenerateFromPassword([]byte(UserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}

	user := models.User{Name: UserRequest.Name, Email: UserRequest.Email, Password: string(hash)}

	result := db.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

func UserLogin(c *gin.Context) {
	c.Bind(&LoginRequest)
	var user models.User

	result := db.Where(&models.User{Email: LoginRequest.Email}).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Username or Password does not match our records",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginRequest.Password)); err != nil {
		// TODO: Properly handle error
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Username or Password does not match our records",
		})

		return
	}

	//generate JWT
	token, err := token_utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Username or Password does not match our records",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
