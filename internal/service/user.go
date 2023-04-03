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
	fmt.Printf("1111111\n")
	if req.Name == "" || req.Password == "" || req.Age <= 0 || !utils.Contains([]string{"male", "female"}, req.Gender) {
		log.Errorf("register param invalid")
		return fmt.Errorf("register param invalid")
	}
	existedUser, err := dao.GetUserByName(req.Name)
	if err != nil {
		log.Errorf("Register|%v", err)
		return fmt.Errorf("register|%v", err)
	}
	fmt.Printf("22222222\n")
	if existedUser != nil {
		fmt.Printf("23r23r2r2r23r32r2")
		return fmt.Errorf("用户已经注册，不能重复注册")
	}
	fmt.Printf("333333333\n")

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
