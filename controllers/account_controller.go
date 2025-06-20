package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/richcem/double-entry-ledgers-demo/database"
	"github.com/richcem/double-entry-ledgers-demo/models"
)

// CreateAccount creates a new account
func CreateAccount(c echo.Context) error {
	account := new(models.Account)
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.DB.Create(account).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create account"})
	}

	return c.JSON(http.StatusCreated, account)
}

// GetAllAccounts returns all accounts
func GetAllAccounts(c echo.Context) error {
	var accounts []models.Account
	if err := database.DB.Find(&accounts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch accounts"})
	}

	return c.JSON(http.StatusOK, accounts)
}

// GetAccountByID returns an account by ID
func GetAccountByID(c echo.Context) error {
	id := c.Param("id")
	var account models.Account

	if err := database.DB.First(&account, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Account not found"})
	}

	return c.JSON(http.StatusOK, account)
}

// UpdateAccount updates an existing account
func UpdateAccount(c echo.Context) error {
	id := c.Param("id")
	var account models.Account

	if err := database.DB.First(&account, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Account not found"})
	}

	updatedAccount := new(models.Account)
	if err := c.Bind(updatedAccount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	account.Name = updatedAccount.Name
	account.Type = updatedAccount.Type

	if err := database.DB.Save(&account).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update account"})
	}

	return c.JSON(http.StatusOK, account)
}

// DeleteAccount deletes an account
func DeleteAccount(c echo.Context) error {
	id := c.Param("id")
	var account models.Account

	if err := database.DB.First(&account, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Account not found"})
	}

	if err := database.DB.Delete(&account).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete account"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Account deleted successfully"})
}

// GetAccountBalance returns the balance of an account
func GetAccountBalance(c echo.Context) error {
	id := c.Param("id")
	var account models.Account

	if err := database.DB.First(&account, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Account not found"})
	}

	return c.JSON(http.StatusOK, map[string]float64{"balance": account.Balance})
}
