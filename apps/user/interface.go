package user

import "context"

// 定义用户接口
type Service interface {
	// 创建用户
	CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(ctx context.Context, req *DeleteUserRequest) error
	// 更新用户
	UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error)
	// 查询用户
	GetUser(ctx context.Context, req *GetUserRequest) (*User, error)
}

// 创建用户的请求
type CreateUserRequest struct {
	UserName string `json:"username"` // 用户名称
	// 用户密码
	PassWord string `json:"password"`
	// 用户性别
	Sex Sex `json:"sex"`
	// 用户角色
	Role Role `json:"role"`
}

// 检查参数
func (*CreateUserRequest) Validate() error {
	return nil
}

// 删除用户的请求
type DeleteUserRequest struct {
	Id int `json:"id"`
}

// 更新用户的请求
type UpdateUserRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Sex      Sex    `json:"sex"`
	State    int    `json:"state"`
}

// 查询用户的请求
type GetUserRequest struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}
