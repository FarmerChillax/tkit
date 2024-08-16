package tkit

import (
	"github.com/gin-gonic/gin"
)

// ResultError result failed
func resultError(ctx *gin.Context, err Error) {
	ctx.AbortWithStatusJSON(err.GetStatusCode(), err)
}
