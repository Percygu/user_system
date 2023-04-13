package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	api "user_system/api/http/v1"
	"user_system/config"
	"user_system/pkg/constant"
)

// InitRouterAndServe 路由配置、启动服务
func InitRouterAndServe() {

	setAppRunMode()
	r := gin.Default()

	// 健康检查
	r.GET("ping", api.Ping)
	//健康检查
	// 用户注册
	r.POST("/user/register", api.Register)
	// 用户登录
	r.POST("/user/login", api.Login)
	// 用户登出
	r.POST("/user/logout", AuthMiddleWare(), api.Logout)
	// 获取用户信息
	r.GET("/user/get_user_info", AuthMiddleWare(), api.GetUserInfo)
	// 更新用户信息
	r.POST("/user/update_nick_name", AuthMiddleWare(), api.UpdateNickName)

	r.Static("/static/", "./web/static/")
	r.Static("/upload/images/", "./web/upload/images/")

	// 启动server
	port := config.GetGlobalConf().AppConfig.Port
	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Error("start server err:" + err.Error())
	}
}

// setAppRunMode 设置运行模式
func setAppRunMode() {
	if config.GetGlobalConf().AppConfig.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session, err := c.Cookie(constant.SessionKey); err == nil {
			if session != "" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		c.Abort()
		return
	}
}
