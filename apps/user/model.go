package user

import "gitee.com/chensyi/vblog/common"

// 构造user对象
func NewUser(req *CreateUserRequest) *User {
	return &User{
		Meta:              common.NewMeta(),
		CreateUserRequest: req,
	}
}

// 定义用户实体对象
type User struct {
	// 通用信息
	*common.Meta
	// 用户传递的请求
	*CreateUserRequest
}
