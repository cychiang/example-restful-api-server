package slack

import (
	models "github.com/cychiang/restful-server/models/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Post(ctx *gin.Context) {
	var payload models.Payload
	if ctx.ContentType() != "application/json" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Content-Type": ctx.ContentType()})
		return
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, payload)
}
