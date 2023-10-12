package blog

import (
	"encoding/json"

	"gitee.com/chensyi/vblog/common"
)

// 构造blog对象
func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		Meta:              common.NewMeta(),
		CreateBlogRequest: req,
	}
}

// 定义文章实体对象
type Blog struct {
	// 通用信息
	*common.Meta
	// 用户传递的请求
	*CreateBlogRequest
	// 发布时间
	PublishedAt int64 `json:"published_at"`
}

func (b *Blog) String() string {
	s, _ := json.Marshal(b)
	return string(s)
}

// gorm解析Model时会调用TableName()来获取Model对应的表名
func (b *Blog) TableName() string {
	return "blogs"
}
