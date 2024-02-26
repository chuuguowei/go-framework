package route

import (
	"github.com/gin-gonic/gin"
	"go-framework/controller"
	"go-framework/library/common/response"
)

func RegisterHandlers(r *gin.Engine) {

	r.NoMethod(func(ctx *gin.Context) {
		response.NoMethod(ctx)
		return
	})
	r.NoRoute(func(ctx *gin.Context) {
		response.NoRoute(ctx)
		return
	})

	// 路由注册
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "hello world")
	//})

	// 用户
	r.GET("/user", controller.IndexHandler())
	r.GET("/userinfo", controller.InfoHandler())
	r.GET("/adduser", controller.AddHandler())

}
