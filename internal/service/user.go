package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"user_system/internal/dao"
	"user_system/internal/model"
	"user_system/utils"
)

// Register 用户注册
func Register(req *RegisterRequest) error {
	if req.Name == "" || req.Password == "" || req.Age <= 0 || !utils.Contains([]string{"male", "female"}, req.Gender) {
		return fmt.Errorf("register param invalid")
	}
	existedUser, err := dao.GetUserByName(req.Name)
	if err != nil {
		log.Errorf("Register|%v", err)
		return fmt.Errorf("register|%v err", err)
	}
	if existedUser != nil {
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

// Register 用户登陆
func Login(req *LoginRequest) error {
	if req.Name == "" || req.Password == "" {
		return fmt.Errorf("login param invalid")
	}
	user, err := dao.GetUserByName(req.Name)
	if err != nil {
		log.Errorf("Login|%v", err)
		return fmt.Errorf("login|%v err", err)
	}
	if user == nil {
		return fmt.Errorf("用户尚未注册")
	}
	return nil
}
