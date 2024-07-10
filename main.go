package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// modules project
	"backend-golang/config"
	"backend-golang/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the database
	config.InitDB()

	// Set Gin to release mode when deploying to production
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin router
	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Setup routes
	routes.SetupRoutes(router)
	routes.SetupListUsersRoutes(router)
	routes.SetupAuthRoutes(router)
	routes.SetupDashboardRoutes(router)
	routes.SetupProductsRoutes(router)
	routes.SetupOrdersRoutes(router)

	// Start the server
	err = router.Run(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
