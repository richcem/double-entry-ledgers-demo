package database

import (
	"log"

	"github.com/richcem/double-entry-ledgers-demo/models"
	"github.com/glebarez/sqlite" // 替换SQLite驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	// 使用glebarez/sqlite驱动连接数据库
	DB, err = gorm.Open(sqlite.Open("double_entry_ledgers.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 迁移数据表
	err = DB.AutoMigrate(
		&models.Account{},
		&models.Transaction{},
		// 注释掉未定义的TransactionItem模型迁移，避免编译错误
		// &models.TransactionItem{},
		&models.User{}, // 添加用户模型迁移
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 创建默认管理员账户
	createDefaultAdmin()
}

// createDefaultAdmin 创建默认管理员账户
func createDefaultAdmin() {
	var admin models.User
	// 检查管理员账户是否已存在
	result := DB.Where("username = ?", "admin").First(&admin)

	if result.Error == gorm.ErrRecordNotFound {
		// 创建默认管理员账户
		admin := models.User{
			Username: "admin",
			Password: "admin", // 将在BeforeSave钩子中自动加密
			Role:     "admin",
		}
		if err := DB.Create(&admin).Error; err != nil {
			log.Printf("Failed to create default admin: %v", err)
		} else {
			log.Println("Default admin user created successfully")
		}
	}
}
