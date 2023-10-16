package api

import (
	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/ioc"
)

func init() {
	ioc.ApiHandler().Registry(&apiHandler{})
}

type apiHandler struct {
	svc user.Service
}

func (h *apiHandler) Name() string {
	return user.AppName
}

func (h *apiHandler) Init() error {
	h.svc = ioc.Controller().Get(user.AppName).(user.Service)
	return nil
}
