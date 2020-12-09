package main

import (
	"context"

	"github.com/bernardolm/octo-batch/config"
	"github.com/bernardolm/octo-batch/github"
	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	config.Defaults()

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	if viper.GetBool("DEBUG") {
		log.SetLevel(log.DebugLevel)
	}

	ctx := context.Background()

	repos, err := github.RepositoriesListAll(ctx)
	if err != nil {
		log.Panic(err)
	}

	for _, r := range repos {
		pp.Println(r.GetFullName())
	}

	if err := github.ActivitySetRepositoriesSubscription(ctx, repos); err != nil {
		log.Panic(err)
	}
}
