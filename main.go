package main

import (
	""
	"context"
	"github.com/Jacky-88/go-base-framework"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"k8s-demo-test/config"
	"k8s-demo-test/controller"
	"k8s-demo-test/service"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// test

	// 初始化K8s clientset
	service.K8s.Init()
	// 初始化路由配置
	r := gin.Default()
	// 初始化路由
	controller.Router.InitApiRouter(r)

	//gin server启动
	srv := &http.Server{
		Addr:    config.ListenAddr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen:" + err.Error())
		}
	}()
	//等待中断信号，优雅关闭所有的server及DB
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//设置ctx超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//关闭gin server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:" + err.Error())
	}
	logger.Info("Gin server 退出成功")

}
