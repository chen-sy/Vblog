package api

import (
	"gitee.com/chensyi/vblog/apps/comment"
	"gitee.com/chensyi/vblog/ioc"
)

func init() {
	ioc.ApiHandler().Registry(&apiHandler{})
}

type apiHandler struct {
	svc comment.Service
}

func (h *apiHandler) Name() string {
	return comment.AppName
}

func (h *apiHandler) Init() error {
	h.svc = ioc.Controller().Get(comment.AppName).(comment.Service)
	return nil
}
