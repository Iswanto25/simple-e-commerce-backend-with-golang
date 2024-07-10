// routes/authRoutes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"backend-golang/controllers"
)

func SetupAuthRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/auth/register", controllers.CreateCustomer)
	router.POST("/auth/login", controllers.LoginCustomer)
}