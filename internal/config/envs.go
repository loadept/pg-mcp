package config

import (
	"os"
	"sync"
)

var (
	once sync.Once
	envs map[string]string
)

// LoadEnvs initializes the environment variables map.
// It loads environment variables from the system with fallback default values.
// This function is safe for concurrent use and will only execute once.
func LoadEnvs() {
	once.Do(func() {
		envs = map[string]string{
			"POSTGRES_URI": getEnvOrDefault("POSTGRES_URI", "postgres://user:password@localhost:5432/mydb?sslmode=disable"),
		}
	})
}

// getEnvOrDefault retrieves an environment variable or returns a default value.
// It checks if the environment variable exists and returns its value,
// otherwise returns the provided default value.
//
// Parameters:
//   - key: Environment variable name to look up
//   - defaultValue: Default value to return if the variable is not set
//
// Returns:
//   - string: The environment variable value or the default value
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

// GetEnv retrieves a previously loaded environment variable.
// This function should be called after LoadEnvs() has been invoked.
//
// Parameters:
//   - key: Environment variable name to retrieve
//
// Returns:
//   - string: The environment variable value, or empty string if not found
func GetEnv(key string) string {
	variable, ok := envs[key]
	if !ok {
		return ""
	}

	return variable
}
