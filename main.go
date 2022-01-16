package main

import (
	"fmt"
	"os"

	"github.com/Y0sh1dk/golang-gin-gorm-example/command"
	"github.com/Y0sh1dk/golang-gin-gorm-example/config"
	"github.com/Y0sh1dk/golang-gin-gorm-example/controllers"
	"github.com/Y0sh1dk/golang-gin-gorm-example/routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	if err := command.ParseCommands(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server := controllers.Server{}
	server.Init(*config.BuildDBConfig())

	defer server.DB.Close()

	routes.SetupRouter(&server)

	server.Router.Run()
}
