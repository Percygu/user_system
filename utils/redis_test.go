package utils

import (
	"golang.org/x/net/context"
	"testing"
	"time"
	"user_system/config"
)

func TestRedis(t *testing.T) {
	ctx := context.Background()
	config.InitConfig()
	res, err := GetRedisCli().Set(ctx, "2222", "user2", 60*time.Second).Result()
	if err != nil {
		t.Errorf("redis set err:%v", err)
	}
	t.Logf("res=%s", res)

	val := GetRedisCli().Get(ctx, "2222").Val()
	if err != nil {
		t.Errorf("redis get err:%v", err)
	}
	t.Logf("val=%s", val)

}
