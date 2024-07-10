// routes/productsRoutes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"backend-golang/controllers"
	"backend-golang/middlewares"
)

func SetupProductsRoutes(router *gin.Engine) {
	productsRoutes := router.Group("/dashboard/products")
	{
		productsRoutes.POST("", middlewares.JWTMiddleware(), controllers.CreateProducts)
		productsRoutes.GET("", middlewares.JWTMiddleware(), controllers.GetProducts)
		productsRoutes.GET("/:id", middlewares.JWTMiddleware(), controllers.GetProductsById)
		productsRoutes.PUT("/:id", middlewares.JWTMiddleware(), controllers.UpdateProducts)
		productsRoutes.DELETE("/:id", middlewares.JWTMiddleware(), controllers.DeleteProducts)
	}
}
