// @title Double-Entry Ledgers API
// @version 1.0
// @description 双记账系统API文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/richcem/double-entry-ledgers-demo/database"
	_ "github.com/richcem/double-entry-ledgers-demo/docs" // 导入生成的docs包
	"github.com/richcem/double-entry-ledgers-demo/middleware"
	"github.com/richcem/double-entry-ledgers-demo/routes"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// Initialize database connection
	database.InitDB()

	// Create Echo instance
	e := echo.New()
	e.Use(emiddleware.Logger()) // 注册日志中间件..

	// 注册验证器 middleware
	e.Validator = middleware.NewValidator()

	// Configure routes
	routes.InitRoutes(e)

	// Health check route
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// 添加Swagger UI路由
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	log.Println("Starting server on :8080")
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
