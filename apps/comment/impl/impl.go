package impl

import (
	"gitee.com/chensyi/vblog/apps/comment"
	"gitee.com/chensyi/vblog/conf"
	"gitee.com/chensyi/vblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registry(&serviceImpl{})
}

type serviceImpl struct {
	db *gorm.DB
}

func (i *serviceImpl) Name() string {
	return comment.AppName
}

func (i *serviceImpl) Init() error {
	i.db = conf.C().MySQL.GetConn().Debug()
	return nil
}
