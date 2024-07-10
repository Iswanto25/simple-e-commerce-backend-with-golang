package routes

import (
	"github.com/gin-gonic/gin"
	"backend-golang/controllers"
	"backend-golang/middlewares"
)

func SetupOrdersRoutes(router *gin.Engine) {
    router.GET("/orders", middlewares.JWTMiddleware(), controllers.ListProducts)
    router.POST("/orders", middlewares.JWTMiddleware(), controllers.CreateOrders)
    // router.GET("/orders/:id", middlewares.JWTMiddleware(), controllers.GetOrder)
    // router.PUT("/orders/:id", middlewares.JWTMiddleware(), controllers.UpdateOrder)
    // router.DELETE("/orders/:id", middlewares.JWTMiddleware(), controllers.DeleteOrder)
}
