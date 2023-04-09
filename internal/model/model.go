package model

import "time"

// CreateModel 内嵌model
type CreateModel struct {
	Creator    string    `gorm:"type:varchar(100);not null;default ''"`
	CreateTime time.Time `gorm:"autoCreateTime"` // 在创建记录时自动生成时间
}

// ModifyModel 内嵌model
type ModifyModel struct {
	Modifier   string    `gorm:"type:varchar(100);not null;default ''"`
	ModifyTime time.Time `gorm:"autoUpdateTime"` // 在更新记录时自动生成时间
}

// User 用户
type User struct {
	CreateModel
	ModifyModel
	ID       int    `gorm:"column:id"`
	Name     string `gorm:"column:name"`     //姓名
	Gender   string `gorm:"column:gender"`   //性别
	Age      int    `gorm:"column:age"`      //年龄
	PassWord string `gorm:"column:password"` //密码
	NickName string `gorm:"column:nickname"` //昵称
}

// TableName 表名
func (t *User) TableName() string {
	return "t_user"
}
