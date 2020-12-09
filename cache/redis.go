package cache

import (
	redigoRedis "github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

func getRedisClient() (redigoRedis.Conn, error) {
	return redigoRedis.DialURL(viper.GetString("REDIS_URL"))
}
