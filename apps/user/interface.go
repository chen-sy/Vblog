package user

import (
	"context"
	"fmt"
)

// 定义用户接口
type Service interface {
	// 创建用户
	CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(ctx context.Context, req *DeleteUserRequest) error
	// 更新用户
	UpdateUser(ctx context.Context, req *UpdateUserRequest) (int64, error)
	// 查询单个用户
	GetSingleUser(ctx context.Context, req *GetSingleUserRequest) (*User, error)
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Sex:   SEX_UNKNOWN,
		Role:  ROLE_MEMBER,
		State: STATE_NORMAL,
	}
}

// 创建用户的请求
type CreateUserRequest struct {
	// 用户名称
	UserName string `json:"username"`
	// 用户密码
	PassWord string `json:"password"`
	// 用户性别
	Sex Sex `json:"sex"`
	// 用户角色
	Role Role `json:"role"`
	// 用户状态
	State State `json:"state"`
}

// 检查参数
func (req *CreateUserRequest) Validate() error {
	if req.UserName == "" || req.PassWord == "" {
		return fmt.Errorf("用户名或密码不能为空")
	}
	return nil
}

// 删除用户的请求
type DeleteUserRequest struct {
	Id int `json:"id"`
}

// 更新用户的请求
type UpdateUserRequest struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Sex      Sex    `json:"sex"`
	State    int    `json:"state"`
}

// 查询单个用户的请求
type GetSingleUserRequest struct {
	Param QueryBy `json:"param"`
	Value string  `json:"value"`
}

func GetSingleUserByID(id int) {

}
