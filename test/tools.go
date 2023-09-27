package test

import (
	"gitee.com/chensyi/vblog/conf"
	"gitee.com/chensyi/vblog/ioc"

	// 注册对象
	_ "gitee.com/chensyi/vblog/apps"
)

// 设置单元测试的配置和环境
func DevelopmentSetup() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	if err := ioc.Controller().Init(); err != nil {
		panic(err)
	}
}
