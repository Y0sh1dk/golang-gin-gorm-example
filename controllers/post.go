package controllers

import (
	"fmt"
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
	var post models.Post

	if err := models.GetPostByID(s.DB, &post, c.Param("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *Server) CreatePost(c *gin.Context) {
	var post models.Post

	// Validate input
	if err := c.BindJSON(&post); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	post.Prepare()

	// Create
	if err := models.CreatePost(s.DB, &post); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *Server) UpdatePost(c *gin.Context) {

}

// func (s *Server) DeletePost(c *gin.Context) {
// 	var post models.Post

// 	if err := models.DeletePost(s.DB, &post)
// }
