package ioc

// 定义注册的对象的约束条件，必须都满足才能托管
type IocObject interface {
	// 对象初始化
	Init() error
	// 对象名称
	Name() string
}
