package cache

import (
	"cstore/pkg/logging"
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

func init() {
	db, _ := strconv.ParseUint("1", 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379",
		DB:   int(db),
	})
	pong, err := client.Ping().Result()
	logging.Info(pong)
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
