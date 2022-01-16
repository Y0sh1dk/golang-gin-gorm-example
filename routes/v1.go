package routes

import (
	"net/http"

	"github.com/Y0sh1dk/golang-gin-gorm-example/controllers"
	"github.com/gin-gonic/gin"
)

func v1(server *controllers.Server) {
	v1 := server.Router.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1.GET("/posts", server.GetPosts)
	v1.POST("/posts", server.CreatePost)
	v1.GET("/posts/:id", server.GetPostByID)
	v1.DELETE("/posts/:id", server.DeletePostByID)
	v1.PUT("/posts/:id", server.UpdatePostByID)

	v1.GET("/posts/:id/comments", server.GetCommentsByPostID)

	v1.GET("/comments", server.GetComments)
	v1.POST("/comments", server.CreateComment)
	v1.GET("/comments/:id", server.GetCommentByID)
	v1.DELETE("/comments/:id", server.DeleteCommentByID)
	v1.PUT("/comments/:id", server.UpdateCommentByID)
}
