package utils

import (
	"os"
	"strconv"
	"strings"
)

func IsDev() bool {
	env := os.Getenv("ENV")
	return strings.Contains(strings.ToLower(env), "dev")
}

func GetEnvByDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvIntByDefault(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intVal
}
