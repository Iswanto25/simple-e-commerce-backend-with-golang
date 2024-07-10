// routes/indexRoutes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"backend-golang/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Example route
	router.GET("/", controllers.Index)

	// Add more routes here as needed
}