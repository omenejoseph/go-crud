package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omenejoseph/go-crud/controllers"
	"github.com/omenejoseph/go-crud/middlewares"
)

func HandlePostRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	posts.Use(middlewares.JwtAuthMiddleware())
	posts.POST("/", controllers.PostCreate)
	posts.GET("/", controllers.PostIndex)
	posts.GET("/:id", controllers.PostFind)
	posts.PATCH("/:id", controllers.PostUpdate)
	posts.DELETE("/:id", controllers.PostDelete)
}
