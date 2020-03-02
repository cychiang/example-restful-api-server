package v1

import (
	"github.com/cychiang/restful-server/controllers/v1/slack"
	"github.com/gin-gonic/gin"
)

func SetSlackRoutes(router *gin.RouterGroup) {
	router.GET("/slack", slack.Get)
	router.POST("/slack", slack.Post)
}
