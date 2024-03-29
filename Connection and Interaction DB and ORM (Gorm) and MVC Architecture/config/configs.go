package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}

	gonfig.GetConf("config/configs.json", &conf)
	return conf
}
