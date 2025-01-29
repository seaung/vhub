package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/seaung/vhub/pkg/config"
	"github.com/seaung/vhub/pkg/models"
	"github.com/seaung/vhub/routers"
)

func main() {
	// 初始化配置
	if err := config.InitConfig(); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}

	// 初始化数据库连接
	if err := models.InitDB(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 初始化路由
	r := routers.InitRouters()

	// 创建HTTP服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.SConfig.App.Port),
		Handler: r,
	}

	// 在goroutine中启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("启动服务失败: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务...")

	// 设置关闭超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("服务关闭异常: %v", err)
	}

	// 关闭数据库连接
	if err := models.CloseDB(); err != nil {
		log.Printf("关闭数据库连接失败: %v", err)
	}

	log.Println("服务已成功关闭")
}
