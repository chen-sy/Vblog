package user

import (
	"context"

	"gitee.com/chensyi/vblog/exception"
	"golang.org/x/crypto/bcrypt"
)

const (
	AppName = "user"
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
	// 通过id查用户
	GetUserByID(ctx context.Context, id int64) (*User, error)
}

// 用户请求的构造函数，用于初始化用户请求的数据，防止空指针
func NewCreateUserRequest(req *CreateUserRequest) *CreateUserRequest {
	return &CreateUserRequest{
		UserName: req.UserName,
		PassWord: req.PassWord,
		Role:     ROLE_MEMBER,
		Status:   STATUS_NORMAL,
	}
}

// 创建用户的请求
type CreateUserRequest struct {
	// 用户名称
	UserName string `json:"username" gorm:"column:username;not null"`
	// 用户密码
	PassWord string `json:"password" gorm:"column:password;not null"`
	// 用户角色
	Role Role `json:"role"`
	// 用户状态
	Status Status `json:"status"`
}

// 检查参数
func (req *CreateUserRequest) Validate() error {
	if req.UserName == "" || req.PassWord == "" {
		return exception.ValidateError("用户名或密码不能为空")
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
	Status   int    `json:"status"`
}

// 密码加密
func (req *UpdateUserRequest) PassWordHash() {
	b, _ := bcrypt.GenerateFromPassword([]byte(req.PassWord), bcrypt.DefaultCost)
	req.PassWord = string(b)
}

// 查询单个用户的请求
type GetSingleUserRequest struct {
	Param QueryBy `json:"param"`
	Value string  `json:"value"`
}

func NewGetSingleUserByID(id string) *GetSingleUserRequest {
	return &GetSingleUserRequest{
		Param: QUERY_BY_ID,
		Value: id,
	}
}

func NewGetSingleUserByName(name string) *GetSingleUserRequest {
	return &GetSingleUserRequest{
		Param: QUERY_BY_USERNAME,
		Value: name,
	}
}
