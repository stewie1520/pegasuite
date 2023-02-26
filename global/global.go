package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"pegasuite/config"
)

var (
	DB     *gorm.DB
	CONFIG config.Server

	VIPER *viper.Viper

	LOG *zap.Logger
)
