package impl

import (
	"context"
	"fmt"
	"strconv"

	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/conf"
	"gorm.io/gorm"
)

// 在我们不知道结构体是否实现某接口时，可以用显示声明接口实现的语句，明确约束接口的实现，下面两种都可以检查是否实现接口
// var _ user.Service = (*UserServiceImpl)(nil)
var _ user.Service = &UserServiceImpl{}

// 构造函数
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: conf.C().MySQL.GetConn(),
	}
}

// 用户接口的实现
type UserServiceImpl struct {
	db *gorm.DB
}

// 创建用户
func (i *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	// 1. 校验用户参数
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// 2. 使用构造函数创建一个User对象
	u := user.NewUser(req)
	// 3. 保存到数据库
	if err := i.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}
	// 4. 返回结果
	return u, nil
}

// 删除用户
func (i *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) error {
	// 查询用户是否存在
	u, err := i.GetSingleUser(ctx, &user.GetSingleUserRequest{
		Param: user.QUERY_BY_ID,
		Value: strconv.Itoa(req.Id),
	})
	if err != nil {
		return err
	}
	return i.db.WithContext(ctx).Delete(u).Error
}

// 查询单个用户
func (i *UserServiceImpl) GetSingleUser(ctx context.Context, req *user.GetSingleUserRequest) (*user.User, error) {
	query := i.db.WithContext(ctx)
	// 构造查询条件
	switch req.Param {
	case user.QUERY_BY_ID:
		query = query.Where("id=?", req.Value)
	case user.QUERY_BY_USERNAME:
		query = query.Where("user_name=?", req.Value)
	default:
		return nil, fmt.Errorf("参数有误")
	}
	//db.Where("name = ?", "jinzhu").First(&user)
	//u := user.NewUser(user.NewCreateUserRequest())
	u := &user.User{}
	if err := query.First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateUser implements user.Service.
func (i *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (int64, error) {
	// 查询用户是否存在
	u, err := i.GetSingleUser(ctx, &user.GetSingleUserRequest{
		Param: user.QUERY_BY_ID,
		Value: strconv.Itoa(req.Id),
	})
	if err != nil {
		return 0, err
	}
	result := i.db.Model(u).Updates(map[string]interface{}{"user_name": req.UserName, "pass_word": req.PassWord})
	return result.RowsAffected, result.Error
}

// grom自动创建表
func MySqlAutoMigrate() error {
	db := conf.C().MySQL.GetConn()
	return db.AutoMigrate(user.User{})
}
