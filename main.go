package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitee.com/chensyi/vblog/conf"
	"gitee.com/chensyi/vblog/ioc"
	"gitee.com/chensyi/vblog/protocol"

	// 注册对象
	_ "gitee.com/chensyi/vblog/apps"
)

func main() {
	// 加载配置
	err := conf.LoadConfigFromToml("etc/application.toml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 通过Ioc完成依赖装载并初始化
	if err := ioc.Controller().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := ioc.ApiHandler().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 创建一个运行在后台的http服务
	httpServer := protocol.NewHttpServer()
	go func() {
		// 开启一个goroutine启动服务
		if err := httpServer.Run(); err != nil {
			fmt.Printf("启动http服务失败, %s\n", err)
		}
	}()
	// 创建一个接收信号的通道，等待中断信号来优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	// signal.Notify把收到的信号转发给quit
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit // 阻塞，等待信号的到来，当接收到上述信号时才会往下执行
	// 创建一个30秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 30秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过30秒就超时退出
	if err := httpServer.Close(ctx); err != nil {
		fmt.Printf("关闭http服务失败, %s\n", err)
	}
	fmt.Println("服务器退出")
}
