package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 声明默认配置
var config *Config = DefaultConfig()

// 方便外部调用配置
func C() *Config {
	return config
}

// 从toml文件加载配置
func LoadConfigFromToml(filepath string) error {
	// DecodeFile读取文件的内容并使用[Decode]进行解码
	_, err := toml.DecodeFile(filepath, config)
	return err
}

// 从环境变量中加载配置
func LoadConfigFromEnv() error {
	// Parse解析包含`env`标记的结构体，并从 环境变量
	return env.Parse(config)
}
