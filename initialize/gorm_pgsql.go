package initialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pegasuite/global"
	"pegasuite/initialize/internal"
)

func gormPsql() *gorm.DB {
	conf := global.CONFIG.Pgsql
	if conf.Dbname == "" {
		return nil
	}

	pgsqlConfig := postgres.Config{
		DSN:                  conf.Dsn(),
		PreferSimpleProtocol: false,
	}

	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(conf.Prefix, conf.Singular)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		return db
	}
}
