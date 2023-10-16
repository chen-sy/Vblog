package api

import (
	"net/http"

	"gitee.com/chensyi/vblog/apps/token"
	"github.com/gin-gonic/gin"
)

// 注册路由
func (t *apiHandler) Registry(r gin.IRouter) {
	// 创建一个路由组v1
	v1 := r.Group("v1")
	// url <=> HandleFunc
	v1.POST("/tokens/", t.Login)
	v1.DELETE("/tokens/", t.Logout)

}

// 登录
func (t *apiHandler) Login(c *gin.Context) {
	// 声明一个LoginRequest对象
	tReq := token.NewLoginRequest()
	// 将请求的json转换成LoginRequest对象
	err := c.BindJSON(tReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 将http协议的请求 ---> 控制器的请求
	ins, err := t.svc.Login(c.Request.Context(), tReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 通过gin把access_token写到浏览器
	c.SetCookie(token.TOKEN_COOKIE_NAME, ins.AccessToken, 0, "/", "localhost", false, true)
	// 返回响应
	c.JSON(http.StatusOK, ins)
}

// 退出
func (t *apiHandler) Logout(c *gin.Context) {
	// 声明一个LogoutRequest对象
	tReq := token.NewLogoutRequest()
	// 将http请求的json ---> LogoutRequest对象
	err := c.BindJSON(tReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 将http协议的请求 ---> 控制器的请求
	err = t.svc.Logout(c.Request.Context(), tReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 销毁浏览器里面存储的access_token
	c.SetCookie(token.TOKEN_COOKIE_NAME, "", -1, "/", "localhost", false, true)
	// 返回响应
	c.JSON(http.StatusOK, "已退出")
}
