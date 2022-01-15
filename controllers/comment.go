package controllers

import (
	"net/http"
	"strconv"

	"github.com/Y0sh1dk/golang-gin-gorm-example/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetComments(c *gin.Context) {
	var comments []models.Comment

	if err := models.GetComments(s.DB, &comments); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (s *Server) GetCommentByID(c *gin.Context) {
	var comment models.Comment

	if err := models.GetCommentByID(s.DB, &comment, c.Param("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (s *Server) CreateComment(c *gin.Context) {
	var comment models.Comment

	// Validate input
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	comment.Prepare()

	// Attach post
	if err := models.GetPostByID(s.DB, &comment.Post, strconv.Itoa(comment.PostID)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Create
	if err := models.CreateComment(s.DB, &comment); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (s *Server) UpdateCommentByID(c *gin.Context) {
	var comment models.Comment

	// Get the model if it exists
	if err := models.GetCommentByID(s.DB, &comment, c.Param("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Validate the input
	var input models.UpdateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	comment.Body = input.Body

	// Update
	if err := models.UpdateComment(s.DB, &comment); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (s *Server) DeleteCommentByID(c *gin.Context) {
	var comment models.Comment

	// Get comment if it exists
	if err := models.GetCommentByID(s.DB, &comment, c.Param("id")); err != nil {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	// Delete comment
	if err := models.DeleteComment(s.DB, &comment); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, comment)
}
