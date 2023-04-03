package service

import (
	"fmt"
	"user_system/internal/utils"
)

// Register 用户注册
func Register(req *RegisterRequest) error {
	if req.Name == "" || req.Age <= 0 || !utils.Contains([]string{"male", "female"}, req.Gender) {
		return fmt.Errorf("register param invalid")
	}
	if err :=
}
