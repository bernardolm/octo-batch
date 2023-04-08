package oauth

import (
	"github.com/bernardolm/octo-batch/cache"
	"golang.org/x/oauth2"
)

func GetOAuth2Transport(token string) (*oauth2.Transport, error) {
	httpcacheTransport, err := cache.GetRedisHTTPCacheTransport()
	if err != nil {
		return nil, err
	}

	tk := &oauth2.Token{
		AccessToken: token,
	}
	ts := oauth2.StaticTokenSource(tk)
	tsp := &oauth2.Transport{
		Base:   httpcacheTransport,
		Source: oauth2.ReuseTokenSource(nil, ts),
	}

	return tsp, nil
}
