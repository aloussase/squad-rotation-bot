package config

import (
	"fmt"
	"os"
)

type Config struct {
	WebHookUrl  string
	DatabaseUrl string
}

const (
	WEB_HOOK_URL = "WEB_HOOK_URL"
	DATABASE_URL = "DATABASE_URL"
)

func ReadConfig() (*Config, error) {
	webHookUrl := os.Getenv(WEB_HOOK_URL)
	if webHookUrl == "" {
		return nil, fmt.Errorf("%s is not set", WEB_HOOK_URL)
	}

	databaseUrl := os.Getenv(DATABASE_URL)
	if databaseUrl == "" {
		return nil, fmt.Errorf("%s is not set", DATABASE_URL)
	}

	return &Config{
		WebHookUrl:  webHookUrl,
		DatabaseUrl: databaseUrl,
	}, nil
}
