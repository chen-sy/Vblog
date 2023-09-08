package user

import "context"

// 定义服务接口
type UserService interface {
	// 创建用户
	CreateUser(ctx context.Context, u User) (*User, error)
	// 获取用户
	GetUser()
}
