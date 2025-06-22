package controllers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/richcem/double-entry-ledgers-demo/config"
	"github.com/richcem/double-entry-ledgers-demo/database"
	"github.com/richcem/double-entry-ledgers-demo/models"
)

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expires_at"`
	User      *models.User `json:"user"`
}

// Login 用户登录
// @Title Login
// @Description 用户登录并获取JWT令牌
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   login body LoginRequest true "登录信息"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
func Login(c echo.Context) error {

	// 绑定请求数据
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无效的请求数据"})
	}

	// 验证请求数据
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 查询用户
	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户名或密码错误"})
	}

	// 验证密码
	if !user.CheckPassword(req.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户名或密码错误"})
	}

	// 生成JWT令牌
	cfg, _ := config.Load()
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      expiresAt,
	})

	// 签名令牌
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "生成令牌失败"})
	}

	// 返回响应
	return c.JSON(http.StatusOK, LoginResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt,
		User:      &user,
	})
}

// GetCurrentUser 获取当前登录用户信息
// @Title GetCurrentUser
// @Description 获取当前登录用户信息
// @Tags auth
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]string
// @Router /api/auth/me [get]
func GetCurrentUser(c echo.Context) error {
	// 从上下文获取用户ID
	userID := c.Get("userID")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "未授权访问"})
	}

	// 查询用户
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "用户不存在"})
	}

	return c.JSON(http.StatusOK, user)
}
