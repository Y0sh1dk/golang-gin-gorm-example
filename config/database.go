package config

import (
	"fmt"

	"github.com/Y0sh1dk/golang-gin-gorm-example/utils"

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
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.StringToInt(utils.GetEnv("DB_PORT", "5432")),
		User:     utils.GetEnv("DB_USER", "pg"),
		Password: utils.GetEnv("DB_PASSWORD", "password"),
		DBName:   utils.GetEnv("DB_NAME", "golang-gin-gorm-example"),
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
