package cache

import (
	"encoding/json"
	"time"
	"user_system/config"

	"user_system/internal/model"
	"user_system/pkg/constant"
	"user_system/utils"
)

func GetUserInfoFromCache(username string) (*model.User, error) {
	redisKey := constant.UserInfoPrefix + username
	val, err := utils.GetRedisCli().Get(redisKey).Result()
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
	_, err = utils.GetRedisCli().Set(redisKey, val, expired).Result()
	return err
}

func updateCachedUserinfo(user *model.User) error {
	err := SetUserCacheInfo(user)
	if err != nil {
		redisKey := constant.UserInfoPrefix + user.Name
		utils.GetRedisCli().Del(redisKey).Result()
	}
	return err
}

func GetTokenInfo(token string) (*model.User, error) {
	redisKey := constant.TokenKeyPrefix + token
	val, err := utils.GetRedisCli().Get(redisKey).Result()
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(val), &user)
	return user, err
}

func SetTokenInfo(user *model.User, token string) error {
	redisKey := constant.TokenKeyPrefix + token
	val, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	expired := time.Second * time.Duration(config.GetGlobalConf().Cache.TokenExpired)
	_, err = utils.GetRedisCli().Set(redisKey, val, expired).Result()
	return err
}

func DelTokenInfo(token string) error {
	redisKey := constant.TokenKeyPrefix + token
	_, err := utils.GetRedisCli().Del(redisKey).Result()
	return err
}
