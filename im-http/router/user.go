package router

import (
	"im-api/im-http/handler"

	"github.com/gin-gonic/gin"
)

func user(r *gin.Engine) {
	r.POST("login", handler.UserLogin)
	r.POST("regist", handler.UserRegist)
	r.GET("head/:filename", handler.UserHead)
}
