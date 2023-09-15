package test

import "gitee.com/chensyi/vblog/conf"

// 设置单元测试的配置和环境
func DevelopmentSetup() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
}
