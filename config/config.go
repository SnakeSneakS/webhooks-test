package config

import (
	"log"

	env "github.com/netflix/go-env"
)

type Config struct {
	SLACK_WEBHOOK_URL   string `env:"SLACK_WEBHOOK_URL"`
	DISCORD_WEBHOOK_URL string `env:"DISCORD_WEBHOOK_URL"`
	Extras              env.EnvSet
}

func GetConfig() Config {
	var conf Config
	envset, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		log.Fatalln(err)
	}
	conf.Extras = envset
	return conf
}
