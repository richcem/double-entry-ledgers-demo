package migrations

import (
	"github.com/richcem/double-entry-ledgers-demo/models"
	"gorm.io/gorm"
)

// CreateUsersTable 创建用户表迁移
func CreateUsersTable(tx *gorm.DB) error {
	return tx.AutoMigrate(&models.User{})
}

// DropUsersTable 删除用户表迁移
func DropUsersTable(tx *gorm.DB) error {
	return tx.Migrator().DropTable(&models.User{})
}