// routes/listUsersRoutes.go

package routes

import (
    "github.com/gin-gonic/gin"
    "backend-golang/controllers"
)

func SetupListUsersRoutes(router *gin.Engine) {
    // Example route to list users
    router.GET("/users", controllers.ListUsers)
    
    // Add more routes as needed
}
