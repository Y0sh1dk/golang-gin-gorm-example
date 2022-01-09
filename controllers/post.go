package controllers

import (
	"net/http"

	"github.com/Y0sh1dk/golang-gin-gorm-example/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := models.GetPosts(s.DB, &posts); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (s *Server) GetPostByID(c *gin.Context) {

}

func (s *Server) CreatePost(c *gin.Context) {

}

func (s *Server) UpdatePost(c *gin.Context) {

}

func (s *Server) DeletePost(c *gin.Context) {

}
