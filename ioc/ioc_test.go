package ioc_test

import (
	"testing"

	"gitee.com/chensyi/vblog/ioc"
	"github.com/gin-gonic/gin"

	// 注册对象
	_ "gitee.com/chensyi/vblog/apps"
)

func TestIocList(t *testing.T) {
	// 初始化Controller
	if err := ioc.Controller().Init(); err != nil {
		t.Fatal(err)
	}

	// 初始化ApiHandler
	if err := ioc.ApiHandler().Init(); err != nil {
		t.Fatal(err)
	}

	// 通过ioc注册handler路由
	r := gin.Default()
	ioc.ApiHandler().RouteRegistry(r.Group("/api"))

	l1 := ioc.ApiHandler().List()
	l2 := ioc.Controller().List()
	t.Log(l1, l2)
}
