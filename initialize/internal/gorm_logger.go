package internal

import (
	"fmt"
	"gorm.io/gorm/logger"
	"pegasuite/global"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{w}
}

func (w *writer) Printf(message string, args ...interface{}) {
	if global.CONFIG.Pgsql.LogZap {
		global.LOG.Info(fmt.Sprintf(message+"\n", args...))
		return
	}

	w.Writer.Printf(message, args...)
}
