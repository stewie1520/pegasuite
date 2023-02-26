package main

import (
	"go.uber.org/zap"
	"pegasuite/core"
	"pegasuite/global"
	"pegasuite/initialize"
)

func main() {
	global.VIPER = core.Viper()
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)
	global.DB = initialize.Gorm()

	if global.DB != nil {
		initialize.RegisterTables()
		db, _ := global.DB.DB()
		defer db.Close()
	}

	core.RunServer()
}
