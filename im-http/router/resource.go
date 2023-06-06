package router

import "github.com/gin-gonic/gin"

func resource(r *gin.Engine) {
	r.LoadHTMLFiles("./front/dist/index.html", "./front/view/login.html", "./front/view/regist.html")
	r.Static("/assets", "./front/dist/assets")
}