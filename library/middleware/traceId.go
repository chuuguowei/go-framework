package middleware

import (
	"github.com/gin-gonic/gin"
	"go-framework/library/common/utils"
)

func traceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := utils.NewUuid()
		ctx.Set("traceId", traceId)
		ctx.Next()
	}
}
