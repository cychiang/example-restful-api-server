package main

import (
	"fmt"
	"github.com/cychiang/restful-server/routers/v1"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.New()
	routerGroup := router.Group("/v1")
	v1.InitRoutes(routerGroup)
}

func main() {
	fmt.Println("Server Running on Port: ", 8080)
	err := router.Run(":8080")
	if err != nil {
		fmt.Print(err)
	}
}
