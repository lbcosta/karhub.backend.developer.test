package config

import "os"

const (
	DevelopmentEnvironment = "development"
	DefaultPort            = "8080"
)

func GetEnvironment() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = DevelopmentEnvironment
	}
	return env
}

func GetPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = DefaultPort
	}
	return port
}
