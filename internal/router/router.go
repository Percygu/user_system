package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
	api "user_system/api/http/v1"
	"user_system/config"
)

// InitRouterAndServe 路由配置、启动服务
func InitRouterAndServe() {

	setAppRunMode()
	r := gin.Default()

	//健康检查
	r.GET("ping", api.Ping)
	//健康检查
	r.POST("/user/register", api.Register)
	r.POST("/user/login", api.Login)

	// 启动server
	port := config.GetGlobalConf().App.Port
	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Error("start server err:" + err.Error())
	}
}

// setAppRunMode 设置运行模式
func setAppRunMode() {
	if config.GetGlobalConf().App.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}
