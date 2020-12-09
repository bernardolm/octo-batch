package github

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/go-github/v33/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/sync/semaphore"
)

func activityGetRepositorySubscription(ctx context.Context, repo *github.Repository) (*github.Subscription, error) {
	result, resp, err := getClient().Activity.GetRepositorySubscription(ctx, *repo.Owner.Login, *repo.Name)
	if err != nil {
		return nil, err
	}

	logRequestResponse(resp)
	return result, nil
}

func activitySetRepositorySubscription(ctx context.Context, repo *github.Repository) error {
	body := strings.NewReader("{\"subscribed\":true}")
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("https://api.github.com/repos/%s/%s/subscription", *repo.Owner.Login, *repo.Name),
		body)
	if err != nil {
		return err
	}

	resp, err := getClient().Do(ctx, req, body)
	if err != nil {
		return err
	}

	logRequestResponse(resp)
	return nil
}

func ActivitySetRepositoriesSubscription(ctx context.Context, repos []*github.Repository) error {
	sem := semaphore.NewWeighted(viper.GetInt64("SUBSCRIPTION_MAX_THREADS"))
	wg := new(sync.WaitGroup)

	for _, repo := range repos {
		wg.Add(1)

		go func(ctx context.Context, repo github.Repository) {
			defer wg.Done()

			if err := sem.Acquire(ctx, 1); err != nil {
				log.Panic(err)
			}
			defer sem.Release(1)

			sub, err := activityGetRepositorySubscription(ctx, &repo)
			if err != nil {
				log.Panic(err)
			}

			if sub.GetSubscribed() {
				log.WithField("repository", repo.GetFullName()).WithField("subscribed", sub.GetSubscribed()).
					Info("subscription")
			} else {
				log.WithField("repository", repo.GetFullName()).WithField("subscribed", sub.GetSubscribed()).
					Warn("subscription")
			}

			if viper.GetBool("SUBSCRIPTION_CHANGE_CONFIRM") && sub == nil {
				if err := activitySetRepositorySubscription(ctx, &repo); err != nil {
					log.Panic(err)
				}
			}
		}(ctx, *repo)
	}

	wg.Wait()

	return nil
}
