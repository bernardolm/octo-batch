package github

import (
	"net/http"

	"github.com/bernardolm/octo-batch/oauth"
	githubSDK "github.com/google/go-github/v33/github"
	log "github.com/sirupsen/logrus"
)

var client *githubSDK.Client

func newClient() (*githubSDK.Client, error) {
	oauth2Transport, err := oauth.GetOAuth2Transport()
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Transport: oauth2Transport,
	}

	return githubSDK.NewClient(c), nil
}

func getClient() *githubSDK.Client {
	if client != nil {
		return client
	}

	c, err := newClient()
	if err != nil {
		log.Panic(err)
	}

	return c
}
