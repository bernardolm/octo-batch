package github

import (
	"net/http"

	githubSDK "github.com/google/go-github/v50/github"

	"github.com/bernardolm/octo-batch/oauth"
)

var client *githubSDK.Client

func newClient() (*githubSDK.Client, error) {
	oauth2Transport, err := oauth.GetOAuth2Transport(token)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Transport: oauth2Transport,
	}

	return githubSDK.NewClient(c), nil
}

func getClient() (*githubSDK.Client, error) {
	if client == nil {
		c, err := newClient()
		if err != nil {
			return nil, err
		}

		client = c
	}

	return client, nil
}
