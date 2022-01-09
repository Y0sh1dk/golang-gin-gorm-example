package utils

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func StringToInt(s string) int {
	str, _ := strconv.Atoi(s)
	return str
}
