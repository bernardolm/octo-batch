package cache

import (
	redigoRedis "github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var redisURL string

func getRedisClient() (redigoRedis.Conn, error) {
	redisURL = viper.GetString("REDIS_URL")
	// debug.Print("cache.getRedisClient.redisURL", redisURL)
	return redigoRedis.DialURL(redisURL)
}
