package models

// 建议添加到 models/transaction.go 或创建 models/requests.go
type DepositRequest struct {
	AccountID   uint    `json:"account_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Description string  `json:"description"`
}
