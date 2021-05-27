package DBHandler

import (
	"github.com/go-redis/redis"
	"github.com/northcity0406/delayTask/db"
)

func GetValue(key string) (interface{}, error) {
	client := db.RedisClient
	stringCmd := client.Get(key)
	err, value := stringCmd.Err(), stringCmd.Val()
	if err != nil {
		return nil, err
	}
	return value, err
}

func SetValue(key string, value interface{}) error {
	statusCmd := db.RedisClient.Set(key, value, 0)
	err, value := statusCmd.Err(), statusCmd.Val()
	if err != nil {
		return err
	}
	return nil
}

func ZAddValue(key string, value interface{}, score float64) error {
	var zAttrs []redis.Z
	z := redis.Z{
		Score:  score,
		Member: value,
	}
	zAttrs = append(zAttrs, z)
	zCmd := db.RedisClient.ZAdd(key, zAttrs...)
	err := zCmd.Err()
	return err
}

func ZRangeValue(key string, start, stop int64) ([]string, error) {
	zCmd := db.RedisClient.ZRange(key, start, stop)
	return zCmd.Val(), zCmd.Err()
}

func ZRemValueByMembers(key string, members []interface{}) error {
	zRem := db.RedisClient.ZRem(key, members)
	return zRem.Err()
}
