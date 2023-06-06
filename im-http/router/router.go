package router

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	resource(r)
	swagger(r)
	page(r)
	user(r)
	group(r.Group("group"))
}