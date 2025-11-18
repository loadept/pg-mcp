package config

import (
	"os"
	"sync"
)

var (
	once sync.Once
	envs map[string]string
)

func LoadEnvs() {
	once.Do(func() {
		envs = map[string]string{
			"POSTGRES_URI": getEnvOrDefault("POSTGRES_URI", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		}
	})
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func GetEnv(key string) string {
	variable, ok := envs[key]
	if !ok {
		return ""
	}

	return variable
}
