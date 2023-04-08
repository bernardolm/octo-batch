package sift

import "github.com/spf13/viper"

var tokenData, apiURL string

func config() {
	// org                       = viper.GetString("SIFT_ORG")
	// repositoryMaxThreads      = viper.GetInt64("SIFT_REPOSITORY_MAX_THREADS")
	// repositoryPerPage         = viper.GetInt("SIFT_REPOSITORY_PER_PAGE")
	// subscriptionChangeConfirm = viper.GetBool("SIFT_SUBSCRIPTION_CHANGE_CONFIRM")
	// subscriptionMaxThreads    = viper.GetInt64("SIFT_SUBSCRIPTION_MAX_THREADS")
	// username                  = viper.GetString("SIFT_USERNAME")
	tokenData = viper.GetString("SIFT_TOKEN_DATA")
	apiURL = viper.GetString("SIFT_API_URL")
	// tokenMedia = viper.GetString("SIFT_TOKEN_MEDIA")
}
