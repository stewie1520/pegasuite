package initialize

import (
	"github.com/gin-gonic/gin"
	"pegasuite/middleware"
)

func InstallPlugin(Router *gin.Engine) {
	Router.Use(middleware.DefaultLogger())
}
