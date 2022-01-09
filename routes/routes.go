package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
	return r
}
