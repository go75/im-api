package router

import (
	"im-api/im-http/handler"

	"github.com/gin-gonic/gin"
)

func group(r *gin.RouterGroup) {
	r.GET("/head/:filename", handler.GroupHead)
	r.POST("/regist", handler.GroupRegist)
}