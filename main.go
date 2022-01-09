package main

import (
	"github.com/Y0sh1dk/golang-gin-gorm-example/config"
	"github.com/Y0sh1dk/golang-gin-gorm-example/controllers"
	"github.com/Y0sh1dk/golang-gin-gorm-example/routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	server := controllers.Server{}
	server.Init(*config.BuildDBConfig())

	defer server.DB.Close()

	routes.SetupRouter(&server)

	server.Router.Run()
}
