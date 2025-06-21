package migrations

import (
	"github.com/richcem/double-entry-ledgers-demo/models"
	"gorm.io/gorm"
)

// CreateEntriesTable 创建分录表迁移
func CreateEntriesTable(tx *gorm.DB) error {
	return tx.AutoMigrate(&models.Entry{})
}

// DropEntriesTable 删除分录表迁移
func DropEntriesTable(tx *gorm.DB) error {
	return tx.Migrator().DropTable(&models.Entry{})
}
