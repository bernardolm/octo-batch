package github

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/go-github/v50/github"
	"golang.org/x/sync/semaphore"

	"github.com/bernardolm/octo-batch/config"
)

func activityGetRepositorySubscription(ctx context.Context, repo *github.Repository) (*github.Subscription, error) {
	c, err := getClient()
	if err != nil {
		return nil, err
	}

	result, resp, err := c.Activity.GetRepositorySubscription(ctx, *repo.Owner.Login, *repo.Name)
	if err != nil {
		return nil, err
	}

	logRequestResponse(resp)
	return result, nil
}

func activitySetRepositorySubscriptionSubscribed(ctx context.Context, repo *github.Repository) error {
	return activitySetRepositorySubscription(ctx, repo, "subscribed")
}

func activitySetRepositorySubscriptionIgnored(ctx context.Context, repo *github.Repository) error {
	return activitySetRepositorySubscription(ctx, repo, "ignored")
}

func activitySetRepositorySubscription(ctx context.Context, repo *github.Repository, action string) error {
	body := strings.NewReader(fmt.Sprintf("{\"%s\":true}", action))
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("https://api.github.com/repos/%s/%s/subscription", *repo.Owner.Login, *repo.Name),
		body)
	if err != nil {
		return err
	}

	c, err := getClient()
	if err != nil {
		return err
	}

	resp, err := c.Do(ctx, req, body)
	if err != nil {
		return err
	}

	logRequestResponse(resp)
	return nil
}

func ActivitySetRepositoriesSubscription(ctx context.Context, repos []*github.Repository) error {
	sem := semaphore.NewWeighted(subscriptionMaxThreads)
	wg := new(sync.WaitGroup)

	for _, repo := range repos {
		wg.Add(1)

		go func(ctx context.Context, repo github.Repository) error {
			defer wg.Done()

			if err := sem.Acquire(ctx, 1); err != nil {
				return err
			}
			defer sem.Release(1)

			if subscriptionChangeConfirm {
				fmt.Printf("ToSubscribe Has %s? %v\n",
					repo.GetFullName(),
					config.Config.Subscriptions.ToSubscribe.Has(repo.GetFullName()))

				fmt.Printf("ToIgnore Has %s? %v\n",
					repo.GetFullName(),
					config.Config.Subscriptions.ToIgnore.Has(repo.GetFullName()))

				if config.Config.Subscriptions.ToSubscribe.Has(repo.GetFullName()) {
					if err := activitySetRepositorySubscriptionSubscribed(ctx, &repo); err != nil {
						return err
					}
				}

				if config.Config.Subscriptions.ToIgnore.Has(repo.GetFullName()) {
					if err := activitySetRepositorySubscriptionIgnored(ctx, &repo); err != nil {
						return err
					}
				}
			}

			return nil
		}(ctx, *repo)
	}

	wg.Wait()

	return nil
}
