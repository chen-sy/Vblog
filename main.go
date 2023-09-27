package main

import (
	"fmt"
	"os"

	"gitee.com/chensyi/vblog/conf"
	"gitee.com/chensyi/vblog/ioc"
	"github.com/gin-gonic/gin"

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

	// 初始化Controller
	if err := ioc.Controller().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化ApiHandler
	if err := ioc.ApiHandler().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 通过ioc注册handler路由
	r := gin.Default()
	ioc.ApiHandler().RouteRegistry(r.Group("/api"))

	// 启动服务
	addr := conf.C().App.HTTPAddr()
	fmt.Printf("HTTP API监听地址: %s", addr)
	err = r.Run(addr)
	fmt.Println(err)
}
