package middleware

import (
	"context"
	"fmt"
	"net/http"

	"gitee.com/chensyi/vblog/apps/token"
	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/ioc"
	"gitee.com/chensyi/vblog/response"
	"github.com/gin-gonic/gin"
)

var (
	ctx = context.Background()
)

func NewTokenAuther() *TokenAuther {
	return &TokenAuther{
		tokenSvc: ioc.Controller().Get(token.AppName).(token.Service),
		userSvc:  ioc.Controller().Get(user.AppName).(user.Service),
	}
}

// 鉴权中间件，基于token来鉴权
type TokenAuther struct {
	tokenSvc token.Service
	userSvc  user.Service
	role     user.Role
}

// 通过token判断是否为本系统认证的用户
func (a *TokenAuther) Auth(c *gin.Context) {
	// 获取token
	at, err := c.Cookie(token.TOKEN_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			response.Failed(c, fmt.Errorf("token Cookie Not Found"))
			return
		}
	}

	// 验证token
	in := token.NewValidateToken(at)
	tk, err := a.tokenSvc.ValidateToken(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 将鉴权后得到的token对象放到gin请求上下文中
	if c.Keys == nil {
		c.Keys = map[string]any{}
	}
	c.Keys[token.TOKEN_GIN_KEY_NAME] = tk

	// 通过token中的userid获取用户对象
	u, err := a.userSvc.GetUserByID(c.Request.Context(), tk.UserID)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 将user对象放到go请求上下文中
	ctx = context.WithValue(ctx, user.USER_KEY, u)
}

// 权限控制
func (a *TokenAuther) Perm(c *gin.Context) {
	uObj := ctx.Value(user.USER_KEY)
	if uObj == nil {
		response.Failed(c, fmt.Errorf("上下文中的user对象不存在"))
		return
	}

	u, ok := uObj.(*user.User)
	if !ok {
		response.Failed(c, fmt.Errorf("上下文中的user对象不是 *user.User"))
		return
	}

	// Admin则直接放行
	if u.Role == user.ROLE_ADMIN {
		return
	}

	// 用户角色不属于放行条件，则拦截
	if u.Role != a.role {
		response.Failed(c, fmt.Errorf("角色 %d 不属于放行队列", u.Role))
		return
	}
}

// 通过用户角色判断权限
func Required(r user.Role) gin.HandlerFunc {
	a := NewTokenAuther()
	a.role = r
	return a.Perm
}
