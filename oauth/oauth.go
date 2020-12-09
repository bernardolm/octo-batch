package oauth

import (
	"github.com/bernardolm/octo-batch/cache"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func GetOAuth2Transport() (*oauth2.Transport, error) {
	httpcacheTransport, err := cache.GetRedisHTTPCacheTransport()
	if err != nil {
		return nil, err
	}

	tk := &oauth2.Token{
		AccessToken: viper.GetString("GITHUB_TOKEN"),
	}
	ts := oauth2.StaticTokenSource(tk)
	tsp := &oauth2.Transport{
		Base:   httpcacheTransport,
		Source: oauth2.ReuseTokenSource(nil, ts),
	}

	return tsp, nil
}
