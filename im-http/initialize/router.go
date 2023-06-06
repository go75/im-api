package initialize

import (
	"im-api/im-http/middleware"
	"im-api/im-http/router"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors)
	r.Use(middleware.RateLimiter)
	r.Use(middleware.Jwt)
	router.Init(r)
	return r
}
