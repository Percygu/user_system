package dao

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"user_system/internal/model"
	"user_system/utils"
)

// GetUserByName 根据姓名获取用户
func GetUserByName(name string) (*model.User, error) {
	user := &model.User{}
	if err := utils.GetDB().Model(model.User{}).Where("name=?", name).First(user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		log.Errorf("GetUserByName fail:%v", err)
		return nil, fmt.Errorf("GetUserByName fail:%v", err)
	}
	return user, nil
}

// CreateUser 创建一个用户
func CreateUser(user *model.User) error {
	if err := utils.GetDB().Model(&model.User{}).Create(user).Error; err != nil {
		log.Errorf("CreateUser fail: %v", err)
		return fmt.Errorf("CreateUser fail: %v", err)
	}
	log.Infof("insert success")
	return nil
}

// UpdateUserInfo 更新昵称
func UpdateUserInfo(userName string, user *model.User) int64 {
	return utils.GetDB().Model(&model.User{}).Where("`name` = ?", userName).Updates(user).RowsAffected
}
