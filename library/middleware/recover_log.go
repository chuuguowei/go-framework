package middleware

import (
	"fmt"
	"go-framework/library/common/response"
	"go-framework/library/logger"

	"github.com/gin-gonic/gin"
)

func recoverLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.WithContext(ctx).Error(fmt.Sprintf("%+v", err))
				response.ServiceUnavailable(ctx)
				ctx.Abort()
				return
			}
		}()

		ctx.Next()

	}
}
