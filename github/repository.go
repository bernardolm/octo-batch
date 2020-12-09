package github

import (
	"context"
	"sync"

	"github.com/google/go-github/v33/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/sync/semaphore"
)

type repositoriesCh chan []*github.Repository
type repositoriesListByPageType func(context.Context, int) ([]*github.Repository, int, error)
type repositoriesListType func(context.Context, repositoriesCh, *sync.WaitGroup) error

func repositoriesListPaginate(ctx context.Context, fn repositoriesListByPageType, ch repositoriesCh, wg *sync.WaitGroup) error {
	r, lastPage, err := fn(ctx, 1)
	if err != nil {
		return err
	}
	ch <- r

	if lastPage == 1 {
		return nil
	}

	sem := semaphore.NewWeighted(viper.GetInt64("REPOSITORY_MAX_THREADS"))

	for i := 2; i <= lastPage; i++ {
		wg.Add(1)

		go func(ctx context.Context, i int, ch repositoriesCh) {
			defer wg.Done()

			if err := sem.Acquire(ctx, 1); err != nil {
				log.Panic(err)
			}
			defer sem.Release(1)

			r, _, err := fn(ctx, i)
			if err != nil {
				log.Panic(err)
			}

			ch <- r
		}(ctx, i, ch)
	}

	return nil
}

func repositoriesListByPage(ctx context.Context, page int) ([]*github.Repository, int, error) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: viper.GetInt("REPOSITORY_PER_PAGE"),
		},
	}

	result, resp, err := getClient().Repositories.List(ctx, viper.GetString("GITHUB_USERNAME"), opt)
	if err != nil {
		return nil, 0, err
	}

	logRequestResponse(resp)
	return result, resp.LastPage, nil
}

func repositoriesList(ctx context.Context, ch repositoriesCh, wg *sync.WaitGroup) error {
	if err := repositoriesListPaginate(ctx, repositoriesListByPage, ch, wg); err != nil {
		return err
	}
	return nil
}

func RepositoriesList(ctx context.Context) ([]*github.Repository, error) {
	ch := make(repositoriesCh)
	wg := new(sync.WaitGroup)

	if err := repositoriesList(ctx, ch, wg); err != nil {
		return nil, err
	}

	wg.Wait()

	var repos []*github.Repository

	for r := range ch {
		repos = append(repos, r...)
	}

	return repos, nil
}

func repositoriesListByOrgByPage(ctx context.Context, page int) ([]*github.Repository, int, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: viper.GetInt("REPOSITORY_PER_PAGE"),
		},
	}

	result, resp, err := getClient().Repositories.ListByOrg(ctx, viper.GetString("GITHUB_ORG"), opt)
	if err != nil {
		return nil, 0, err
	}

	logRequestResponse(resp)
	return result, resp.LastPage, nil
}

func repositoriesListByOrg(ctx context.Context, ch repositoriesCh, wg *sync.WaitGroup) error {
	if err := repositoriesListPaginate(ctx, repositoriesListByOrgByPage, ch, wg); err != nil {
		return err
	}
	return nil
}

func RepositoriesListByOrg(ctx context.Context) ([]*github.Repository, error) {
	ch := make(repositoriesCh)
	wg := new(sync.WaitGroup)

	if err := repositoriesListByOrg(ctx, ch, wg); err != nil {
		return nil, err
	}

	wg.Wait()

	var repos []*github.Repository

	for r := range ch {
		repos = append(repos, r...)
	}

	return repos, nil
}

func RepositoriesListAll(ctx context.Context) ([]*github.Repository, error) {
	ch := make(repositoriesCh)
	wg := new(sync.WaitGroup)

	fns := []repositoriesListType{
		repositoriesList,
		repositoriesListByOrg,
	}

	var repos []*github.Repository

	go func() {
		for r := range ch {
			repos = append(repos, r...)
		}
	}()

	for _, fn := range fns {
		if err := fn(ctx, ch, wg); err != nil {
			log.Panic(err)
		}
	}

	wg.Wait()

	return repos, nil
}
