package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"user_system/config"
	"user_system/internal/service"
)

// Ping 健康检查
func Ping(c *gin.Context) {
	appConfig := config.GetGlobalConf().App
	confInfo, _ := json.MarshalIndent(appConfig, "", "  ")
	appInfo := fmt.Sprintf("app_name: %s\nversion: %s\n\n%s", appConfig.AppName, appConfig.Version,
		string(confInfo))
	c.String(http.StatusOK, appInfo)
}

// Register 注册
func Register(c *gin.Context) {
	req := &service.RegisterRequest{}
	rsp := &HttpResponse{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}
	if err := service.Register(req); err != nil {
		rsp.ResponseWithError(c, CodeRegisterErr, err.Error())
		return
	}
	rsp.ResponseSuccess(c)
}

// Login 注册
func Login(c *gin.Context) {
	req := &service.LoginRequest{}
	rsp := &HttpResponse{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}
	if err := service.Login(req); err != nil {
		rsp.ResponseWithError(c, CodeLoginErr, "login err")
		return
	}
	rsp.ResponseSuccess(c)
}
