package ioc

// 专门用于注册Controller对象
func Controller() *IocContainter {
	return controllerContainer
}

// Controller对象注册的容器
var controllerContainer = &IocContainter{
	container: map[string]IocObject{},
}
