package main

import (
	"fmt"

	"github.com/Y0sh1dk/golang-gin-gorm-example/config"
	"github.com/Y0sh1dk/golang-gin-gorm-example/models"
	"github.com/Y0sh1dk/golang-gin-gorm-example/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var err error
	config.DB, err = gorm.Open("postgres", config.BuildDSN(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
		panic(err.Error())
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(models.Post{}, models.Comment{})

	r := routes.SetupRouter()

	r.Run()
}
