package main

import (
	"restApi/config"
	"restApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to DB
	config.ConnectDB()

	// Initialize router
	r := gin.Default()

	// Register routes
	routes.UserRoutes(r)

	// Start server
	r.Run(":8080")
}
