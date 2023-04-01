package routes

import (
	"github.com/gin-gonic/gin"
)

func HandleRouting() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	HandlePostRoutes(v1)
	HandleUsersRoute(v1)

	r.Run()
}
