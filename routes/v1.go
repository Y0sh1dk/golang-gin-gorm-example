package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func v1(r *gin.Engine) {
	v1 := r.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{"Status": "ok"})
		c.String(http.StatusOK, "pong")
	})
}
