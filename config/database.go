package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     GetEnv("DB_HOST", "localhost"),
		Port:     StringToInt(GetEnv("DB_PORT", "5432")),
		User:     GetEnv("DB_USER", "pg"),
		Password: GetEnv("DB_PASSWORD", "password"),
		DBName:   GetEnv("DB_NAME", "golang-gin-gorm-example"),
	}
	return &dbConfig
}

func BuildDSN(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Australia/Sydney application_name=golang-gin-gorm-api",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port)
}

// TODO(yoshi): factor this out
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
