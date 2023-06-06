package middleware

import (
	"im-api/im-http/component"
	"time"

	"github.com/gin-gonic/gin"
)

var limiter = component.NewLimiter(time.Second, 100)

// 限流中间件
func RateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		c.Abort()
		return
	}
	c.Next()
}
