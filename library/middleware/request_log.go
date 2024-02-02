package middleware

import (
	"fmt"
	"go-framework/library/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func requestLog() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}
		TimeStamp := time.Now()
		Cost := TimeStamp.Sub(start)
		if Cost > time.Minute {
			Cost = Cost.Truncate(time.Second)
		}

		requestMap := map[string]interface{}{
			"Path":      path,
			"Method":    c.Request.Method,
			"ClientIP":  c.ClientIP(),
			"Cost":      fmt.Sprintf("%s", Cost),
			"Status":    c.Writer.Status(),
			"Proto":     c.Request.Proto,
			"UserAgent": c.Request.UserAgent(),
			"Msg":       c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"Size":      c.Writer.Size(),
		}

		logger.WithContext(c).Serve(requestMap)
	}
}
