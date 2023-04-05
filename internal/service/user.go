package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"user_system/internal/cache"
	"user_system/internal/dao"
	"user_system/internal/model"
	"user_system/pkg/constant"
	"user_system/utils"
)

// Register 用户注册
func Register(req *RegisterRequest) error {
	if req.Name == "" || req.Password == "" || req.Age <= 0 || !utils.Contains([]string{"male", "female"}, req.Gender) {
		log.Errorf("register param invalid")
		return fmt.Errorf("register param invalid")
	}
	existedUser, err := dao.GetUserByName(req.Name)
	if err != nil {
		log.Errorf("Register|%v", err)
		return fmt.Errorf("register|%v", err)
	}
	if existedUser != nil {
		fmt.Printf("23r23r2r2r23r32r2")
		return fmt.Errorf("用户已经注册，不能重复注册")
	}

	user := &model.User{
		Name:     req.Name,
		Age:      req.Age,
		Gender:   req.Gender,
		PassWord: req.Password,
	}

	if err := dao.CreateUser(user); err != nil {
		log.Errorf("Register|%v", err)
		return fmt.Errorf("register|%v", err)
	}
	return nil
}

func getUserInfo(userName string) (*model.User, error) {
	user, err := cache.GetUserInfoFromCache(userName)
	if err == nil && user.Name == userName {
		return user, nil
	}

	user, err = dao.GetUserByName(userName)
	if err != nil {
		return user, err
	}

	if user == nil {
		return nil, fmt.Errorf("用户尚未注册")
	}

	err = cache.SetUserCacheInfo(user)
	if err != nil {
		log.Error("cache userinfo failed for user:", user.Name, " with err:", err.Error())
	}
	return user, nil
}

// Login 用户登陆
func Login(ctx context.Context, req *LoginRequest) error {
	uuid := ctx.Value(constant.ReqUuid)
	log.Debugf(" %s| Login access from:%s,@,%s", uuid, req.UserName, req.PassWord)

	user, err := getUserInfo(req.UserName)
	if err != nil {
		log.Errorf("Login|%v", err)
		return fmt.Errorf("login|%v", err)
	}

	// 用户存在
	if req.PassWord != user.PassWord {
		log.Errorf("Login|password err: req.password=%s|user.password=%s", req.PassWord, user.PassWord)
	}

	token := utils.GenerateToken(user.Name)
	err = cache.SetTokenInfo(user, token)

	if err != nil {
		log.Errorf(" Login|Failed to SetTokenInfo, uuid=%s|user_name=%s|token=%s|err=%v", uuid, user.Name, token, err)
		return fmt.Errorf("login|SetTokenInfo fail:%v", err)
	}

	log.Infof("Login successfully, %s@%s with token %s", req.UserName, req.PassWord, token)
	return nil
}

// Logout 退出登陆
func Logout(ctx context.Context, req *LogoutRequest) error {
	uuid := ctx.Value(constant.ReqUuid)
	log.Infof("%s|Logout access from,user_name=%s|token=%s", uuid, req.UserName, req.Token)
	token := utils.GenerateToken(req.UserName)
	if token != req.Token {
		log.Errorf("user_name:%s and token:%s is not matched", req.UserName, req.Token)
		return fmt.Errorf("user_name and token is not matched")
	}
	err := cache.DelTokenInfo(req.Token)
	if err != nil {
		log.Errorf("%s|Failed to delTokenInfo :%s", uuid, req.Token)
		return fmt.Errorf("del token err:%v", err)
	}
	log.Errorf("%s|Failed to delTokenInfo :%s", uuid, req.Token)
	return nil
}

// GetUserInfo 获取用户信息请求，只能在用户登陆的情况下使用
func GetUserInfo(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	uuid := ctx.Value(constant.ReqUuid)
	log.Infof("%s|GetUserInfo access from,user_name=%s|token=%s", uuid, req.UserName, req.Token)

	if req.Token == "" || req.UserName == "" {
		return nil, fmt.Errorf("GetUserInfo|request params invalid")
	}

	user, err := cache.GetTokenInfo(req.Token)
	if err != nil {
		log.Errorf("%s|Failed to get with token=%s|err =%v", uuid, req.Token, err)
		return nil, fmt.Errorf("getUserInfo|GetTokenInfo err:%v", err)
	}

	if user.Name != req.UserName {
		log.Errorf("%s|token info not match with username=%s", uuid, req.UserName)
	}
	log.Infof("%s|Succ to GetUserInfo|user_name=%s|token=%s", uuid, req.UserName, req.Token)
	return &GetUserInfoResponse{
		UserName: user.Name,
		Age:      user.Age,
		Gender:   user.Gender,
		PassWord: user.PassWord,
	}, nil
}

func UpdateUserInfo(ctx context.Context, req *UpdateUserInfoRequest) {
	uuid := ctx.Value(constant.ReqUuid)
	log.Infof("%s|UpdateUserInfo access from,user_name=%s|token=%s", uuid, req.UserName, req.Token)
}
