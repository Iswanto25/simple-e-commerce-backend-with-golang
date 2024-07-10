// routes/dashboardRoutes.go

package routes

import (
	"github.com/gin-gonic/gin"

	// 
	"backend-golang/controllers"
	"backend-golang/middlewares"
)

func SetupDashboardRoutes(router *gin.Engine) {
	router.GET("/dashboard", middlewares.JWTMiddleware(), controllers.Dashboard)
	router.GET("/dashboard/orders", middlewares.JWTMiddleware(), controllers.ListOrders)
	router.PUT("/dashboard/orders/:id", middlewares.JWTMiddleware(), controllers.UpdateOrders)
	router.DELETE("/dashboard/orders/:id", middlewares.JWTMiddleware(), controllers.DeleteOrders)
}