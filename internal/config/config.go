package config

import "os"

const (
	SERVER_ADDRESS string = "SERVER_ADDRESS"
	ADMIN_ADDRESS  string = "ADMIN_ADDRESS"
)

type Config struct {
	AdminAddress  string
	ServerAddress string
}

// NewFromEnv creates a Config from the environment variables with sensible fallbacks
func NewFromEnv() *Config {
	serverAddress := envWithDefaultValue(SERVER_ADDRESS, ":8080")
	adminAddress := envWithDefaultValue(ADMIN_ADDRESS, ":8081")

	return &Config{adminAddress, serverAddress}
}

func envWithDefaultValue(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	return value
}
