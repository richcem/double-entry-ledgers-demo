package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/richcem/double-entry-ledgers-demo/config"
)

// JWT密钥，从配置中获取
var jwtSecret []byte

// 初始化JWT密钥
func init() {
	cfg, err := config.Load()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	jwtSecret = []byte(cfg.JWTSecret)
}

// AuthMiddleware 验证JWT令牌的中间件
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 获取Authorization头
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "未提供认证令牌"})
		}

		// 检查令牌格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "认证令牌格式错误，应为Bearer {token}"})
		}

		// 解析JWT令牌
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return jwtSecret, nil
		})

		// 验证令牌有效性
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "无效的认证令牌"})
		}

		// 将用户信息存入上下文
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
			c.Set("username", claims["username"])
		}

		// 继续处理请求
		return next(c)
	}
}
