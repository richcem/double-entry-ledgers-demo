package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Migrate 执行所有数据库迁移
func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:       "20231015000000",
			Migrate:  CreateAccountsTable,
			Rollback: DropAccountsTable,
		},
		{
			ID:       "20231015000001",
			Migrate:  CreateTransactionsTable,
			Rollback: DropTransactionsTable,
		},
		{
			ID:       "20231015000002",
			Migrate:  CreateEntriesTable,
			Rollback: DropEntriesTable,
		},
	})

	return m.Migrate()
}
