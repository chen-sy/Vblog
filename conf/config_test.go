package conf_test

import (
	"testing"

	"gitee.com/chensyi/vblog/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}
