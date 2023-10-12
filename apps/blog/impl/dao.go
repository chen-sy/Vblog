package impl

import (
	"context"

	"gitee.com/chensyi/vblog/apps/blog"
	"gitee.com/chensyi/vblog/conf"
)

func (i *blogServiceImpl) getBlog(ctx context.Context, id int) (*blog.Blog, error) {
	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	ins := &blog.Blog{}
	query = query.Where("id=?", id)
	if err := query.First(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// grom自动创建表
func MySqlAutoMigrate() error {
	db := conf.C().MySQL.GetConn()
	return db.AutoMigrate(blog.Blog{})
}
