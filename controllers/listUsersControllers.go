// controllers/listUsersControllers.go

package controllers

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"

    // project modules
    "backend-golang/config"	
    "backend-golang/models"
)

func ListUsers(c *gin.Context) {
    var users []models.Users
    result := config.DB.Find(&users)
    if result.Error != nil {
        log.Printf("Failed to fetch users: %v", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
	c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "List of users",
        "data":    users,
	})
}
