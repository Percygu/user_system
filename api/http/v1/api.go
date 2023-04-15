package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"user_system/config"
	"user_system/internal/service"
	"user_system/pkg/constant"
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

// Login 登录
func Login(c *gin.Context) {
	req := &service.LoginRequest{}
	rsp := &HttpResponse{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}

	uuid := utils.Md5String(req.UserName + time.Now().GoString())
	ctx := context.WithValue(context.Background(), "uuid", uuid)
	log.Infof("loggin start,user:%s, password:%s", req.UserName, req.PassWord)
	session, err := service.Login(ctx, req)
	if err != nil {
		rsp.ResponseWithError(c, CodeLoginErr, err.Error())
		return
	}
	// 登陆成功，设置cookie
	c.SetCookie(constant.SessionKey, session, constant.CookieExpire, "/", "", false, true)
	rsp.ResponseSuccess(c)
}

// Logout 登出
func Logout(c *gin.Context) {
	session, _ := c.Cookie(constant.SessionKey)
	ctx := context.WithValue(context.Background(), constant.SessionKey, session)
	req := &service.LogoutRequest{}
	rsp := &HttpResponse{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Errorf("bind get logout request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}
	uuid := utils.Md5String(req.UserName + time.Now().GoString())
	ctx = context.WithValue(ctx, "uuid", uuid)
	if err := service.Logout(ctx, req); err != nil {
		rsp.ResponseWithError(c, CodeLogoutErr, err.Error())
		return
	}
	c.SetCookie(constant.SessionKey, session, -1, "/", "", false, true)
	rsp.ResponseSuccess(c)
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	userName := c.Query("username")
	session, _ := c.Cookie(constant.SessionKey)
	ctx := context.WithValue(context.Background(), constant.SessionKey, session)
	req := &service.GetUserInfoRequest{
		UserName: userName,
	}
	rsp := &HttpResponse{}
	uuid := utils.Md5String(req.UserName + time.Now().GoString())
	ctx = context.WithValue(ctx, "uuid", uuid)
	userInfo, err := service.GetUserInfo(ctx, req)
	if err != nil {
		rsp.ResponseWithError(c, CodeGetUserInfoErr, err.Error())
		return
	}
	rsp.ResponseWithData(c, userInfo)
}

// UpdateNickName 更新用户昵称
func UpdateNickName(c *gin.Context) {
	req := &service.UpdateNickNameRequest{}
	rsp := &HttpResponse{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Errorf("bind update user info request json err %v", err)
		rsp.ResponseWithError(c, CodeBodyBindErr, err.Error())
		return
	}
	session, _ := c.Cookie(constant.SessionKey)
	log.Infof("UpdateNickName|session=%s", session)
	ctx := context.WithValue(context.Background(), constant.SessionKey, session)
	uuid := utils.Md5String(req.UserName + time.Now().GoString())
	ctx = context.WithValue(ctx, "uuid", uuid)
	if err := service.UpdateUserNickName(ctx, req); err != nil {
		rsp.ResponseWithError(c, CodeUpdateUserInfoErr, err.Error())
		return
	}
	rsp.ResponseSuccess(c)
}
