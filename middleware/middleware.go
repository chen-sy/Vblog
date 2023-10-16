package middleware

import (
	"context"

	"gitee.com/chensyi/vblog/apps/token"
	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/ioc"
	"github.com/gin-gonic/gin"
)

func NewMiddleware() *Middleware {
	return &Middleware{
		tokenSvc: ioc.Controller().Get(token.AppName).(token.Service),
		userSvc:  ioc.Controller().Get(user.AppName).(user.Service),
	}
}

type Middleware struct {
	tokenSvc token.Service
	userSvc  user.Service
}

// 添加用户id和角色到上下文（兼容游客通用）
func (a *Middleware) NewContextGeneral(c *gin.Context) (ctx context.Context, err error) {
	// 获取token
	at, err := c.Cookie(token.TOKEN_COOKIE_NAME)
	if err != nil {
		// 游客访问
		return c.Request.Context(), nil
	}
	// 校验token是否有效
	in := token.NewValidateToken(at)
	tk, err := a.tokenSvc.ValidateToken(c.Request.Context(), in)
	if err != nil {
		return nil, err
	}
	// 通过token中的userid获取用户对象
	u, err := a.userSvc.GetUserByID(c.Request.Context(), tk.UserID)
	if err != nil {
		return nil, err
	}
	// 将user信息放到上下文中
	ctx = context.WithValue(c.Request.Context(), user.CTX_KEY_USERID, u.ID)
	ctx = context.WithValue(ctx, user.CTX_KEY_USERROLE, u.Role)
	return ctx, nil
}

// 添加用户到上下文（已鉴权接口使用）
func (a *Middleware) NewContext(c *gin.Context) (ctx context.Context, err error) {
	// 获取token
	tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	if tkObj == nil {
		return nil, err
	}

	tk, ok := tkObj.(*token.Token)
	if !ok {
		return nil, err
	}
	// 通过token中的userid获取用户对象
	u, err := a.userSvc.GetUserByID(c.Request.Context(), tk.UserID)
	if err != nil {
		return nil, err
	}
	// 将user信息放到上下文中
	ctx = context.WithValue(c.Request.Context(), user.CTX_KEY_USER, u)
	return ctx, nil
}
