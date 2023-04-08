package github

import (
	"net/http"

	githubSDK "github.com/google/go-github/v50/github"
	log "github.com/sirupsen/logrus"
)

func logRequestResponse(r *githubSDK.Response) {
	fields := log.Fields{
		"from_cache":     r.Header.Get("X-From-Cache"),
		"method":         r.Request.Method,
		"rate_limit":     r.Rate.Limit,
		"rate_remaining": r.Rate.Remaining,
		"status":         r.Status,
		"url":            r.Request.URL.String(),
	}

	switch status := r.StatusCode; {
	case status >= http.StatusInternalServerError:
		log.WithFields(fields).Errorf("github API request done")
	case status >= http.StatusBadRequest:
		log.WithFields(fields).Warnf("github API request done")
	default:
		log.WithFields(fields).Infof("github API request done")
	}
}
