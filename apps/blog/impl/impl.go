package impl

import (
	"gitee.com/chensyi/vblog/apps/blog"
	"gitee.com/chensyi/vblog/conf"
	"gitee.com/chensyi/vblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registry(&blogServiceImpl{})
}

// blog不再暴露到外部，托管到ioc，从ioc访问
type blogServiceImpl struct {
	db *gorm.DB
}

// 托管到ioc里面的名称
func (i *blogServiceImpl) Name() string {
	return blog.AppName
}

// 实现ioc的init，让ioc来初始化对象
func (i *blogServiceImpl) Init() error {
	// 通过debug可以查看sql语句
	i.db = conf.C().MySQL.GetConn().Debug()
	return nil
}
