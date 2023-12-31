package api

import (
	"gitee.com/chensyi/vblog/apps/blog"
	"gitee.com/chensyi/vblog/ioc"
)

func init() {
	ioc.ApiHandler().Registry(&apiHandler{})
}

type apiHandler struct {
	svc blog.Service
}

func (h *apiHandler) Name() string {
	return blog.AppName
}

func (h *apiHandler) Init() error {
	h.svc = ioc.Controller().Get(blog.AppName).(blog.Service)
	return nil
}
