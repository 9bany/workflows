package utils

import (
	"os"
)

type Config struct {
	AppToken string `mapstructure:"SLACK_APP_TOKEN"`
	BotToken string `mapstructure:"SLACK_BOT_TOKEN"`
}

func LodConfig(path string) Config {
	return Config{
		AppToken: os.Getenv("SLACK_APP_TOKEN"),
		BotToken: os.Getenv("SLACK_BOT_TOKEN"),
	}
}
