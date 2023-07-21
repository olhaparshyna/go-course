package config

import "os"

type Config struct {
	ApiKey string
}

func New() *Config {
	return &Config{ApiKey: getEnv("API_KEY", "")}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
