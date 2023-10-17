package protocol

import (
	"context"
	"fmt"
	"net/http"

	"gitee.com/chensyi/vblog/conf"
	"gitee.com/chensyi/vblog/ioc"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *HttpServer {
	// 通过ioc注册ApiHandler路由
	r := gin.Default()
	ioc.ApiHandler().RouteRegistry(r.Group("/api/vblog"))

	return &HttpServer{
		server: &http.Server{
			// 服务监听的地址
			Addr: conf.C().App.HTTPAddr(),
			// 监听关联的路由处理
			Handler: r,
		},
	}
}

// 封装http server
type HttpServer struct {
	server *http.Server
}

// 启动服务
func (s *HttpServer) Run() error {
	fmt.Printf("listen addr: %s\n", conf.C().App.HTTPAddr())
	return s.server.ListenAndServe()
}

// 优雅关闭
func (s *HttpServer) Close(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
