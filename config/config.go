package config

import (
	"github.com/spf13/viper"
)

func Defaults() {
	viper.SetDefault("REPOSITORY_MAX_THREADS", 5)
	viper.SetDefault("REPOSITORY_PER_PAGE", 5)
	viper.SetDefault("SUBSCRIPTION_CHANGE_CONFIRM", false)
	viper.SetDefault("SUBSCRIPTION_MAX_THREADS", 5)
}
