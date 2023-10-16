package api

import (
	"gitee.com/chensyi/vblog/apps/token"
	"gitee.com/chensyi/vblog/ioc"
)

func init() {
	ioc.ApiHandler().Registry(&apiHandler{})
}

type apiHandler struct {
	svc token.Service
}

func (h *apiHandler) Name() string {
	return token.AppName
}

func (h *apiHandler) Init() error {
	h.svc = ioc.Controller().Get(token.AppName).(token.Service)
	return nil
}
