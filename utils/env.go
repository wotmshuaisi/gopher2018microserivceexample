package utils

import (
	"os"
)

// GetSysEnv get env value, if doesn't exists return default
func GetSysEnv(key, value string) string {
	val := os.Getenv(key)
	if val == "" {
		return value
	}
	return val
}
