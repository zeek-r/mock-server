package config

import (
	"os"
)

var (
	defaultConfig = map[string]string{
		"mockDirectory": "mocks",
	}
)

// GetConfig returns config from env or default
func GetConfig(key string) string {
	envConfig := os.Getenv(key)
	if envConfig == "" {
		return defaultConfig[key]
	}
	return envConfig
}
