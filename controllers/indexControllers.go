// controllers/indexControllers.go

package controllers

import (
	"log"

	"github.com/matoous/go-nanoid/v2"
	"github.com/gin-gonic/gin"
)


func Index(c *gin.Context) {
	id, err := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890", 54)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(id)
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Welcome to the API",
		"data": id,
	})
}
