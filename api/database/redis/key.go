package redis

import (
	"api/configs"
	"api/utils"
	"time"
)

func SetKey(key string, value string, ttlSeconds int) {
	configs.Redis.Set(utils.Ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

func GetKey(key string) (string, error) {
	return configs.Redis.Get(utils.Ctx, key).Result()
}

func DelKey(key string) {
	configs.Redis.Del(utils.Ctx, key).Err()
}
