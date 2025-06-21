package migrations

import (
	"github.com/richcem/double-entry-ledgers-demo/models"
	"gorm.io/gorm"
)

// CreateAccountsTable 创建账户表迁移
func CreateAccountsTable(tx *gorm.DB) error {
	return tx.AutoMigrate(&models.Account{})
}

// DropAccountsTable 删除账户表迁移
func DropAccountsTable(tx *gorm.DB) error {
	return tx.Migrator().DropTable(&models.Account{})
}
