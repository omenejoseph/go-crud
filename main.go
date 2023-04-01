package main

import (
	"github.com/gin-gonic/gin"
	"github.com/omenejoseph/go-crud/controllers"
	"github.com/omenejoseph/go-crud/inititalizers"
)

func init() {
	inititalizers.LoadEnvVariables()
	inititalizers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/", controllers.PostIndex)
	r.GET("/:id", controllers.PostFind)
	r.PATCH("/:id", controllers.PostUpdate)
	r.DELETE("/:id", controllers.PostDelete)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
