package github

import "github.com/spf13/viper"

var (
	org                       = viper.GetString("GITHUB_ORG")
	repositoryMaxThreads      = viper.GetInt64("GITHUB_REPOSITORY_MAX_THREADS")
	repositoryPerPage         = viper.GetInt("GITHUB_REPOSITORY_PER_PAGE")
	subscriptionChangeConfirm = viper.GetBool("GITHUB_SUBSCRIPTION_CHANGE_CONFIRM")
	subscriptionMaxThreads    = viper.GetInt64("GITHUB_SUBSCRIPTION_MAX_THREADS")
	token                     = viper.GetString("GITHUB_TOKEN")
	username                  = viper.GetString("GITHUB_USERNAME")
)

func init() {
	viper.SetDefault("GITHUB_REPOSITORY_MAX_THREADS", 5)
	viper.SetDefault("GITHUB_REPOSITORY_PER_PAGE", 5)
	viper.SetDefault("GITHUB_SUBSCRIPTION_CHANGE_CONFIRM", false)
	viper.SetDefault("GITHUB_SUBSCRIPTION_MAX_THREADS", 5)
}
