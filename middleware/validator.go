package middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator 自定义验证器
type CustomValidator struct {
	validator *validator.Validate
}

// Validate 实现Echo的Validator接口
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// NewValidator 创建新的验证器实例
func NewValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}
