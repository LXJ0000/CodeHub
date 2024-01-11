package redis

import (
	"bluebell/conf"
	"github.com/go-redis/redis"
	"strings"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

func Init() {
	redisConfig := conf.Conf.RedisConfig
	client = redis.NewClient(&redis.Options{
		Addr:         strings.Join([]string{redisConfig.Host + ":" + redisConfig.Port}, ""),
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		PoolSize:     redisConfig.PoolSize,
		MinIdleConns: redisConfig.MinIdleConns,
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
}
