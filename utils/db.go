package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
	"user_system/config"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// openDB 连接db
func openDB() {
	mysqlConf := config.GetGlobalConf().DbConfig
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.User,
		mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Dbname)
	log.Info("mdb addr:" + connArgs)

	var err error
	db, err = gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("fetch db connection err:" + err.Error())
	}

	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConn)                                        //设置最大空闲连接
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConn)                                        //设置最大打开的连接
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.MaxIdleTime * int64(time.Second))) //设置空闲时间为(s)
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	dbOnce.Do(openDB)
	return db
}
