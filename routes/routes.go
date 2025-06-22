package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/richcem/double-entry-ledgers-demo/controllers"
	"github.com/richcem/double-entry-ledgers-demo/middleware"
)

// InitRoutes 初始化路由
func InitRoutes(e *echo.Echo) {
	// 公开路由
	public := e.Group("/api")
	{
		// 认证路由
		auth := public.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
		}

		// 健康检查
		public.GET("/health", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"status": "ok"})
		})
	}

	// 需要认证的路由
	api := e.Group("/api")
	api.Use(middleware.AuthMiddleware)
	{
		// 用户信息
		auth := api.Group("/auth")
		{
			auth.GET("/me", controllers.GetCurrentUser)
		}

		// 账户路由
		accounts := api.Group("/accounts")
		{
			accounts.GET("", controllers.GetAllAccounts)
			accounts.GET("/:id", controllers.GetAccountByID)
			accounts.POST("", controllers.CreateAccount)
			accounts.PUT("/:id", controllers.UpdateAccount)
			accounts.DELETE("/:id", controllers.DeleteAccount)
		}

		// 交易路由
		transactions := api.Group("/transactions")
		{
			transactions.GET("", controllers.GetAllTransactions)
			transactions.GET("/:id", controllers.GetTransactionByID)
			transactions.POST("", controllers.CreateTransaction)
			transactions.POST("/deposit", controllers.CreateDeposit) // 存款
			// transactions.PUT("/:id", controllers.UpdateTransaction) // TODO: 实现controllers.UpdateTransaction函数以解决未定义错误

			// transactions.DELETE("/:id", controllers.DeleteTransaction) // TODO: 实现controllers.DeleteTransaction函数以解决未定义错误
		}
	}
}
