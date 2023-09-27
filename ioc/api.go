package ioc

// 专门用于注册Api对象
func ApiHandler() *IocContainter {
	return apiHandlerContainer
}

// Api对象注册的容器
var apiHandlerContainer = &IocContainter{
	container: map[string]IocObject{},
}
