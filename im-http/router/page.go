package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func page(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/regist", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "regist.html", nil)
	})
}
