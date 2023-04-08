package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Debugging bool

func Init() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if Debugging {
		log.SetLevel(log.DebugLevel)
	}

	Debugging = viper.GetBool("DEBUG")

	if Debugging {
		viper.Debug()
	}

	// if err := loadYAML(); err != nil {
	// 	return err
	// }

	return nil
}
