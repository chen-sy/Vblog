package api

import (
	"net/http"

	"gitee.com/chensyi/vblog/apps/token"
	"github.com/gin-gonic/gin"
)

func NewTokenApiHandler(tokenImpl token.Service) *TokenApiHandler {
	return &TokenApiHandler{
		tokenSvc: tokenImpl,
	}
}

type TokenApiHandler struct {
	tokenSvc token.Service
}

// 注册路由
func (t *TokenApiHandler) Registry(r gin.IRouter) {
	// 创建一个路由组v1
	v1 := r.Group("v1")
	// url <=> HandleFunc
	v1.POST("/tokens/", t.Login)
	v1.DELETE("/tokens/", t.Logout)

}

// Login HandleFunc
func (t *TokenApiHandler) Login(c *gin.Context) {
	// 声明一个LoginRequest对象
	tReq := token.NewLoginRequest()
	// 将http请求的json ---> LoginRequest对象
	err := c.BindJSON(tReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 将http协议的请求 ---> 控制器的请求
	ins, err := t.tokenSvc.Login(c.Request.Context(), tReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, ins)
}

// Logout HandleFunc
func (t *TokenApiHandler) Logout(c *gin.Context) {
	// 声明一个LogoutRequest对象
	tReq := token.NewLogoutRequest()
	// 将http请求的json ---> LogoutRequest对象
	err := c.BindJSON(tReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 将http协议的请求 ---> 控制器的请求
	err = t.tokenSvc.Logout(c.Request.Context(), tReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, "已退出")
}
