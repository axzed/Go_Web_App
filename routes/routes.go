package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"web_app/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello web_app")
	})
	// 检测脚手架版本
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("版本号:%s", viper.GetString("app.version")))
	})

	return r
}
