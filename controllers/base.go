package controllers

import (
	"fmt"

	"github.com/Y0sh1dk/golang-gin-gorm-example/config"
	"github.com/Y0sh1dk/golang-gin-gorm-example/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (s *Server) Init(dbconf config.DBConfig) {
	var err error
	s.DB, err = gorm.Open("postgres", config.BuildDSN(&dbconf))
	if err != nil {
		fmt.Println("Status:", err)
		panic(err.Error())
	}

	s.DB.AutoMigrate(models.Post{}, models.Comment{})
}
