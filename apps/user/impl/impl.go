package impl

import (
	"context"

	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/conf"
	"gorm.io/gorm"
)

// 在我们不知道结构体是否实现某接口时，可以用显示声明接口实现的语句，明确约束接口的实现，下面两种都可以检查是否实现接口
var _ user.Service = &UserServiceImpl{}

//var _ user.Service = (*UserServiceImpl)(nil)

// 用户接口的实现
type UserServiceImpl struct {
	db *gorm.DB
}

// 构造函数
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: conf.C().MySQL.GetConn(),
	}
}

// 创建用户
func (i *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	// 1. 校验用户参数
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// 2. 使用构造函数创建一个User对象
	req.UserName = "tom"
	req.PassWord = "111111"
	req.Sex = 1
	req.Role = 0
	u := user.User{
		Id:                1,
		CreatedAt:         10120021,
		UpdatedAt:         10120021,
		State:             1,
		CreateUserRequest: req,
	}
	// 3. 保存到数据库
	i.db.Create(u)

	// 4. 返回结果
	return nil, nil
}

// 删除用户
func (i *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) error {
	return nil
}

// GetUser implements user.Service.
func (*UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.User, error) {
	panic("unimplemented")
}

// UpdateUser implements user.Service.
func (*UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	panic("unimplemented")
}

// grom自动创建表
func MySqlAutoMigrate() error {
	db := conf.C().MySQL.GetConn()
	return db.AutoMigrate(user.User{})
}
