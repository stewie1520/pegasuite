package initialize

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"pegasuite/global"
)

func Gorm() *gorm.DB {
	return gormPsql()
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate()

	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
