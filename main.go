package main

import (
	"flag"
	"fmt"
	"go-framework/bootstrap"
	"go-framework/library/config"
	"go-framework/library/logger"
	"go-framework/library/middleware"
	"go-framework/library/route"

	"github.com/gin-gonic/gin"
	"go-framework/config"
)

var configFile = flag.String("f", "./.env.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	// 加载配置
	library.MustLoad(*configFile, &c)
	if c.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 框架初始化
	server := gin.New()
	server.HandleMethodNotAllowed = true

	// 依赖初始化
	bootstrap.MustInit(c)
	// 注册中间件
	middleware.RegisterGlobalMiddlewares(server)
	// 注册路由
	route.RegisterHandlers(server)

	fmt.Printf("Server Started at localhost%s...\n", c.App.Port)

	if err := server.Run(c.App.Port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
		logger.ErrorF("Start server error:%v", err)
	}

}
