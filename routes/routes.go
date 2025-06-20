package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/richcem/double-entry-ledgers-demo/controllers"
)

// InitRoutes initializes all application routes
func InitRoutes(e *echo.Echo) {
	// API group
	api := e.Group("/api")

	// Account routes
	account := api.Group("/accounts")
	account.POST("/", controllers.CreateAccount)
	account.GET("/", controllers.GetAllAccounts)
	account.GET("/:id", controllers.GetAccountByID)
	account.PUT("/:id", controllers.UpdateAccount)
	account.DELETE("/:id", controllers.DeleteAccount)

	// Transaction routes
	transaction := api.Group("/transactions")
	transaction.POST("/", controllers.CreateTransaction)
	transaction.GET("/", controllers.GetAllTransactions)
	transaction.GET("/:id", controllers.GetTransactionByID)

	// Transfer route
	api.POST("/transfers", controllers.CreateTransfer)

	// Balance routes
	api.GET("/accounts/:id/balance", controllers.GetAccountBalance)
}
