package config

import "os"

type ConfigService struct {
}

func (*ConfigService) Getenv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func NewService() *ConfigService {
	return &ConfigService{}
}
