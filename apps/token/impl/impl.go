package impl

import (
	"context"
	"fmt"

	"gitee.com/chensyi/vblog/apps/token"
	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/conf"
	"gitee.com/chensyi/vblog/ioc"
	"gorm.io/gorm"
)

var _ token.Service = &TokenServiceImpl{}

func init() {
	ioc.Controller().Registry(&TokenServiceImpl{})
}

type TokenServiceImpl struct {
	// token需要存储，引入db
	db *gorm.DB
	// token依赖用户模块，引入user
	user user.Service
}

func (i *TokenServiceImpl) Name() string {
	return token.AppName
}

func (i *TokenServiceImpl) Init() error {
	i.db = conf.C().MySQL.GetConn().Debug()
	i.user = ioc.Controller().Get(user.AppName).(user.Service)
	return nil
}

// 用户登录
func (i *TokenServiceImpl) Login(ctx context.Context, req *token.LoginRequest) (*token.Token, error) {
	// 判断用户是否存在
	uReq := user.NewGetSingleUserByName(req.UserName)
	u, err := i.user.GetSingleUser(ctx, uReq)
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	// 校验密码
	err = u.CheckPassWord(req.PassWord)
	if err != nil {
		return nil, fmt.Errorf("密码错误")
	}
	// 颁发token
	t := token.NewToken()
	t.UserID = u.ID
	// 保存token
	if err = i.db.WithContext(ctx).Create(t).Error; err != nil {
		return nil, err
	}
	//颁发成功删除原来的token
	if err = i.db.WithContext(ctx).Where("user_id = ? and access_token != ?", t.UserID, t.AccessToken).Delete(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

// 用户退出
func (i *TokenServiceImpl) Logout(ctx context.Context, req *token.LogoutRequest) error {
	// 查找token
	t := token.NewToken()
	if err := i.db.WithContext(ctx).Where("access_token = ? and refresh_token=?", req.AccessToken, req.RefreshToken).First(t).Error; err != nil {
		return err
	}
	// 删除token
	return i.db.WithContext(ctx).Where("access_token = ? and refresh_token=?", req.AccessToken, req.RefreshToken).Delete(t).Error
}

// 校验token
func (i *TokenServiceImpl) ValidateToken(ctx context.Context, req *token.ValidateToken) (*token.Token, error) {
	t := token.NewToken()
	if err := i.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).First(t).Error; err != nil {
		return nil, err
	}
	if err := t.IsExpired(); err != nil {
		return nil, err
	}
	return t, nil
}

// grom自动创建表
func MySqlAutoMigrate() error {
	db := conf.C().MySQL.GetConn()
	return db.AutoMigrate(token.Token{})
}
