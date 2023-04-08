package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	Config configYAML
)

type list []string

func (l list) Has(a string) bool {
	for i := range l {
		if l[i] == a {
			return true
		}
	}
	return false
}

type configYAML struct {
	Subscriptions struct {
		ToSubscribe list `yaml:"to_subscribe"`
		ToIgnore    list `yaml:"to_ignore"`
	} `yaml:"subscriptions"`
}

func loadYAML() error {
	f, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(f, &Config); err != nil {
		return err
	}

	return nil
}
