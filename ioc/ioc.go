package ioc

import (
	"github.com/gin-gonic/gin"
)

// 定义对象注册的容器
type IocContainter struct {
	container map[string]IocObject
}

// 初始化所有对象
func (i *IocContainter) Init() error {
	for _, o := range i.container {
		if err := o.Init(); err != nil {
			return err
		}
	}
	return nil
}

// 注册到容器
func (i *IocContainter) Registry(o IocObject) {
	i.container[o.Name()] = o
}

// 获取容器里面的对象
func (i *IocContainter) Get(name string) any {
	return i.container[name]
}

type GinApiHandler interface {
	Registry(r gin.IRouter)
}

// 把每个 ApiHandler的路由注册给Root Router
func (i *IocContainter) RouteRegistry(r gin.IRouter) {
	// 找到被托管的api handler
	for _, o := range i.container {
		// 断言是gin的handler
		if api, ok := o.(GinApiHandler); ok {
			api.Registry(r)
		}
	}
}

// 获取Ioc列表
func (i *IocContainter) List() any {
	return i.container
}
