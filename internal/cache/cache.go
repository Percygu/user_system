package cache

import (
	"encoding/json"
	"golang.org/x/net/context"
	"time"
	"user_system/config"
	"user_system/internal/model"
	"user_system/pkg/constant"
	"user_system/utils"
)

func GetUserInfoFromCache(username string) (*model.User, error) {
	redisKey := constant.UserInfoPrefix + username
	val, err := utils.GetRedisCli().Get(context.Background(), redisKey).Result()
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(val), user)
	return user, err
}

func SetUserCacheInfo(user *model.User) error {
	redisKey := constant.UserInfoPrefix + user.Name
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}
	expired := time.Second * time.Duration(config.GetGlobalConf().Cache.UserExpired)
	_, err = utils.GetRedisCli().Set(context.Background(), redisKey, val, expired*time.Second).Result()
	return err
}

func UpdateCachedUserInfo(user *model.User) error {
	err := SetUserCacheInfo(user)
	if err != nil {
		redisKey := constant.UserInfoPrefix + user.Name
		utils.GetRedisCli().Del(context.Background(), redisKey).Result()
	}
	return err
}

func GetSessionInfo(session string) (*model.User, error) {
	redisKey := constant.SessionKeyPrefix + session
	val, err := utils.GetRedisCli().Get(context.Background(), redisKey).Result()
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(val), &user)
	return user, err
}

func SetSessionInfo(user *model.User, session string) error {
	redisKey := constant.SessionKeyPrefix + session
	val, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	expired := time.Second * time.Duration(config.GetGlobalConf().Cache.SessionExpired)
	_, err = utils.GetRedisCli().Set(context.Background(), redisKey, val, expired*time.Second).Result()
	return err
}

func DelSessionInfo(session string) error {
	redisKey := constant.SessionKeyPrefix + session
	_, err := utils.GetRedisCli().Del(context.Background(), redisKey).Result()
	return err
}
