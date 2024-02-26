package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"go-framework/bootstrap"
	"go-framework/library/configs"
	"go-framework/library/logger"
	"go-framework/library/middleware"
	"go-framework/library/route"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go-framework/config"
)

var (
	configFile = flag.String("f", "./.env.yaml", "the configs file")

	//go:embed web/html/*.html
	htmlFs embed.FS

	//go:embed web/static/* web/static/fonts/*
	staticFs embed.FS
)

func main() {
	// 解析命令行参数
	flag.Parse()
	var c config.Config

	// 加载配置
	configs.MustLoad(*configFile, &c)
	if c.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Web模块
	module, err := bootstrap.NewWebModule("/", "web", htmlFs, staticFs)
	if err != nil {
		logger.ErrorF("create web module error, {}", err)
		return
	}

	// 关闭模块前进行资源清理
	defer module.Close()

	// 框架初始化
	server := gin.Default()
	server.HandleMethodNotAllowed = true

	// 处理静态文件请求
	module.ServStaticFiles(server)
	// 处理HTML文件请求
	module.ServHtmlFiles(server)

	// 依赖初始化
	bootstrap.MustInit(c)
	// 注册中间件
	middleware.RegisterGlobalMiddlewares(server)
	// 注册路由
	route.RegisterHandlers(server)

	// 监听端口
	listener, err := net.Listen("tcp", c.App.Port)
	if err != nil {
		return
	}
	// 初始化服务
	srv := &http.Server{
		Handler: server,
	}
	go func() {
		// 服务连接
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start server error, listen: %s\n", err)
		}
	}()

	fmt.Printf("Server Started at localhost%s...\n", c.App.Port)

	gracefulShutdown(srv)
}

func gracefulShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shut down the server with a timeout of seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown error:", err)
	}
}
