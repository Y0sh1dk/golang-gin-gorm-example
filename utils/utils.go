package utils

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func StringToInt(s string) int {
	str, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return str
}
