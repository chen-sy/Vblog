package user

import (
	"encoding/json"

	"gitee.com/chensyi/vblog/common"
	"golang.org/x/crypto/bcrypt"
)

// 构造user对象
func NewUser(req *CreateUserRequest) *User {
	return &User{
		Meta:              common.NewMeta(),
		CreateUserRequest: NewCreateUserRequest(req),
	}
}

// 定义用户实体对象
type User struct {
	// 通用信息
	*common.Meta
	// 用户传递的请求
	*CreateUserRequest
}

// 密码加密（新增更新用户密码功能，将原有CreateUserRequest上的PassWordHash，抽象到User）
func (u *User) PassWordHash(passWord string) {
	b, _ := bcrypt.GenerateFromPassword([]byte(passWord), bcrypt.DefaultCost)
	u.PassWord = string(b)
}

// 判断用户密码是否正确
func (u *User) CheckPassWord(passWord string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(passWord))
}

func (u *User) String() string {
	b, _ := json.Marshal(u)
	return string(b)
}

// gorm解析Model时会调用TableName()来获取Model对应的表名
func (u *User) TableName() string {
	return "users"
}
