package config

import (
	"os"
)

var (
	Port string
)

func init() {
	Port = getEnv("PORT", "4000")
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
