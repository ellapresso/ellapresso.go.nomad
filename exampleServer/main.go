package main

import (
	v1 "exampleServer/routers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var routers *gin.Engine

func init() {
	routers = gin.New()
	temp := routers.Group("/")
	v1.InitRoutes(temp)
}

func main() {
	fmt.Println("Hello World")
	http.ListenAndServe(":4703", routers)
}
