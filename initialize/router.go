package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pegasuite/global"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	return Router
}
