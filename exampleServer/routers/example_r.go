package v1

import (
	"exampleServer/controllers"

	"github.com/gin-gonic/gin"
)

func SetExampleRoutes(router *gin.RouterGroup) {
	router.GET("/ping", controllers.Pong)
}
