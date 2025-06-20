package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/richcem/double-entry-ledgers-demo/database"
	"github.com/richcem/double-entry-ledgers-demo/routes"
)

func main() {
	// Initialize database connection
	database.InitDB()

	// Create Echo instance
	e := echo.New()

	// Configure routes
	routes.InitRoutes(e)

	// Health check route
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// Start server
	log.Println("Starting server on :8080")
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
