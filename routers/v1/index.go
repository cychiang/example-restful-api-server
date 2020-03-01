package v1

import "github.com/gin-gonic/gin"

func InitRoutes(routerGroup *gin.RouterGroup) {
	SetSlackRoutes(routerGroup)
}
