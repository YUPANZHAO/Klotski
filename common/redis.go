package common

import (
	"KlotskiWeb/config"

	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func InitRedisDB() error {
	conf := config.NewRedisConfig()

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.DB,
	})

	_, err := RedisDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func CloseRedisDB() error {
	if RedisDB != nil {
		return RedisDB.Close()
	}
	return nil
}
