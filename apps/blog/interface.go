package blog

import "context"

// 定义服务接口
type BlogService interface {
	// 创建文章
	CreateBlog(ctx context.Context, b Blog) (*Blog, error)
	// 获取文章
	QueryBlog()
	// 修改文章
	UpdateBlog()
}
