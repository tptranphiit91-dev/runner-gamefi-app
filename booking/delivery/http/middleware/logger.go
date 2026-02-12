package middleware

import (
	"time"
	
	"github.com/gin-gonic/gin"
)

// Logger middleware for logging HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		
		c.Next()
		
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		
		println("📊 [HTTP]", method, path, "- Status:", statusCode, "- Latency:", latency.String())
	}
}

