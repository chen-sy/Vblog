package user

// 定义领域模型
type User struct {
	// 用户id
	ID int
	// 用户名称
	UserName string
	// 用户密码
	PassWord string
	// 创建时间戳
	CreateAt int
}
