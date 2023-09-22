package main

import (
	"fmt"
	"os"

	tokenApi "gitee.com/chensyi/vblog/apps/token/api"
	tokenImpl "gitee.com/chensyi/vblog/apps/token/impl"
	userImpl "gitee.com/chensyi/vblog/apps/user/impl"
	"gitee.com/chensyi/vblog/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	err := conf.LoadConfigFromToml("etc/application.toml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化控制器
	usi := userImpl.NewUserServiceImpl()
	tsi := tokenImpl.NewTokenServiceImpl(usi)
	// 初始化handler
	tah := tokenApi.NewTokenApiHandler(tsi)

	// 通过gin注册handler路由
	r := gin.Default()
	tah.Registry(r.Group("/api/vblog"))

	// 启动服务
	addr := conf.C().App.HTTPAddr()
	fmt.Printf("HTTP API监听地址: %s", addr)
	err = r.Run(addr)
	fmt.Println(err)
}
