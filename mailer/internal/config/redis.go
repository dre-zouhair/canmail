package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

type RedisConf struct {
	redisOptions *redis.Options
}

func NewRedisConf() *RedisConf {
	fmt.Println(os.Environ())
	redisHost := os.Getenv("REDIS_HOST")
	fmt.Println("REDIS_HOST ", redisHost)
	if redisHost == "" {
		fmt.Println("redisHost is not set")
		return nil
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")
	fmt.Println("redisPassword ", redisPassword)
	if redisPassword == "" {
		fmt.Println("REDIS_PASSWORD is not set")
		return nil
	}

	redisDB := os.Getenv("REDIS_DB")
	fmt.Println("REDIS_DB ", redisDB)
	if redisDB == "" {
		fmt.Println("redisPassword is not set")
		return nil
	}

	dbNumber, err := strconv.Atoi(redisDB)
	if err != nil {
		fmt.Println("redisDB does not have a valid value")
		return nil
	}

	return &RedisConf{
		redisOptions: &redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
			DB:       dbNumber,
		},
	}
}

func (dBConf *RedisConf) GetRedisOptions() *redis.Options {
	return dBConf.redisOptions
}
