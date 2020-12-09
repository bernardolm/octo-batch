package cache

import (
	"github.com/gregjones/httpcache"
	httpcacheRedis "github.com/gregjones/httpcache/redis"
)

func GetRedisHTTPCacheTransport() (*httpcache.Transport, error) {
	c, err := getRedisClient()
	if err != nil {
		return nil, err
	}
	cch := httpcacheRedis.NewWithClient(c)
	return httpcache.NewTransport(cch), nil
}
