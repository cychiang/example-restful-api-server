package slack

import (
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
