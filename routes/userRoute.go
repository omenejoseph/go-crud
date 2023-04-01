package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/omenejoseph/go-crud/controllers"
)

func HandleUsersRoute(r *gin.RouterGroup) {
	r.POST("/login", controllers.UserLogin)
	userRoute := r.Group("/users")

	userRoute.POST("/", controllers.UserCreate)
}
