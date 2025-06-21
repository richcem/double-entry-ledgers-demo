package migrations

import (
	"github.com/richcem/double-entry-ledgers-demo/models"
	"gorm.io/gorm"
)

// CreateTransactionsTable 创建交易表迁移
func CreateTransactionsTable(tx *gorm.DB) error {
	return tx.AutoMigrate(&models.Transaction{})
}

// DropTransactionsTable 删除交易表迁移
func DropTransactionsTable(tx *gorm.DB) error {
	return tx.Migrator().DropTable(&models.Transaction{})
}
