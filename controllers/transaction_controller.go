package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/richcem/double-entry-ledgers-demo/database"
	"github.com/richcem/double-entry-ledgers-demo/models"
	"gorm.io/gorm"
)

// GetAllTransactions returns all transactions
func GetAllTransactions(c echo.Context) error {
	var transactions []models.Transaction
	if err := database.DB.Preload("Entries.Account").Find(&transactions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch transactions"})
	}

	return c.JSON(http.StatusOK, transactions)
}

// GetTransactionByID returns a transaction by ID
func GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	var transaction models.Transaction

	if err := database.DB.Preload("Entries.Account").First(&transaction, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Transaction not found"})
	}

	return c.JSON(http.StatusOK, transaction)
}

// CreateTransfer handles a transfer between two accounts
func CreateTransfer(c echo.Context) error {
	// Define transfer request structure
	var transfer struct {
		FromAccountID uint    `json:"from_account_id"`
		ToAccountID   uint    `json:"to_account_id"`
		Amount        float64 `json:"amount"`
		Description   string  `json:"description"`
	}

	if err := c.Bind(&transfer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Validate transfer data
	if transfer.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Amount must be greater than zero"})
	}

	if transfer.FromAccountID == transfer.ToAccountID {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "From and to account cannot be the same"})
	}

	// Verify accounts exist
	var fromAccount models.Account
	if err := database.DB.First(&fromAccount, transfer.FromAccountID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "From account not found"})
	}

	var toAccount models.Account
	if err := database.DB.First(&toAccount, transfer.ToAccountID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "To account not found"})
	}

	// Create transaction with two entries (debit and credit)
	transaction := models.Transaction{
		Description: transfer.Description,
		Amount:      transfer.Amount,
		Date:        time.Now(),
		Status:      "completed",
		Entries: []models.Entry{
			{
				AccountID: transfer.FromAccountID,
				Amount:    transfer.Amount,
				Type:      "credit",
			},
			{
				AccountID: transfer.ToAccountID,
				Amount:    transfer.Amount,
				Type:      "debit",
			},
		},
	}

	// Use transaction to ensure atomicity
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		// Create transaction
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// Process entries and update account balances
		for _, entry := range transaction.Entries {
			var account models.Account
			if err := tx.First(&account, entry.AccountID).Error; err != nil {
				return err
			}

			if entry.Type == "debit" {
				account.Balance += entry.Amount
			} else if entry.Type == "credit" {
				account.Balance -= entry.Amount
				// Check if account has sufficient balance
				if account.Balance < 0 {
					return echo.NewHTTPError(http.StatusBadRequest, "Insufficient balance in source account")
				}
			}

			if err := tx.Save(&account).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Transfer failed: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, transaction)
}

// @Title CreateTransaction
// @Description 创建新交易
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   transaction body models.Transaction true "交易信息"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/transactions [post]
func CreateTransaction(c echo.Context) error {
	transaction := new(models.Transaction)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Set default values
	transaction.Date = time.Now()
	transaction.Status = "completed"

	// Use transaction to ensure data consistency
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		// Create transaction
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		// Create entries and update account balances
		for i := range transaction.Entries {
			entry := &transaction.Entries[i]
			entry.TransactionID = transaction.ID

			// Create entry
			if err := tx.Create(entry).Error; err != nil {
				return err
			}

			// Update account balance
			var account models.Account
			if err := tx.First(&account, entry.AccountID).Error; err != nil {
				return err
			}

			if entry.Type == "debit" {
				account.Balance += entry.Amount
			} else if entry.Type == "credit" {
				account.Balance -= entry.Amount
			}

			if err := tx.Save(&account).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create transaction: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, transaction)
}
