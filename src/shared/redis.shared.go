package shared

import (
	"context"
	configModule "golang-example-project/config"
	"time"
)

var ctx = context.Background()

func GetRedisByKey(key string) (string, error) {
	result, err := configModule.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func SetRedisByKey(key string, payload any, duration int) error {
	err := configModule.RedisClient.Set(ctx, key, payload, time.Duration(duration)).Err()
	if err != nil {
		return err
	}
	return nil
}
