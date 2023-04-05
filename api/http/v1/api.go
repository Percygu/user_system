package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
	"user_system/config"
	"user_system/internal/service"
	"user_system/utils"
)

// Ping 健康检查
func Ping(c *gin.Context) {
	appConfig := config.GetGlobalConf().AppConfig
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

// Login 登陆
func Login(c *gin.Context) {
	req := &service.LoginRequest{}
	rsp := &HttpResponse{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}

	uuid := utils.GenerateToken(req.UserName)
	ctx := context.WithValue(context.Background(), "uuid", uuid)
	log.Infof("loggin start,user:%s, password:%s", req.UserName, req.PassWord)
	if err := service.Login(ctx, req); err != nil {
		rsp.ResponseWithError(c, CodeLoginErr, "login err")
		return
	}
	rsp.ResponseSuccess(c)
}

func GetUserInfo(c *gin.Context) {
	token := c.Header["token"]
	req := &service.GetUserInfoRequest{}
	rsp := &HttpResponse{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Errorf("bind get user info request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}
	uuid := utils.GenerateToken(req.UserName)
	ctx := context.WithValue(context.Background(), "uuid", uuid)
	userInfo, err := service.GetUserInfo(ctx, req)
	if err != nil {
		rsp.ResponseWithError(c, CodeGetUserInfoErr, err.Error())
		return
	}
	rsp.ResponseWithData(c, userInfo)
}

func UpdateUserInfo(c *gin.Context) {
	req := &service.UpdateUserInfoRequest{}
	rsp := &HttpResponse{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Errorf("bind update user info request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}
	uuid := utils.GenerateToken(req.UserName)
	ctx := context.WithValue(context.Background(), "uuid", uuid)

}
