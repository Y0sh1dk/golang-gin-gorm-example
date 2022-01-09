package routes

import (
	"net/http"

	"github.com/Y0sh1dk/golang-gin-gorm-example/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(server *controllers.Server) {
	server.Router = gin.Default()

	v1(server)

	server.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
}
