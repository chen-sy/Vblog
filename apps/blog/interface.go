package blog

import (
	"context"

	"gitee.com/chensyi/vblog/common"
	"gitee.com/chensyi/vblog/exception"
)

const (
	AppName = "blog"
)

// 定义服务接口
type Service interface {
	// 创建博客
	CreateBlog(ctx context.Context, req *CreateBlogRequest) (*Blog, error)
	// 删除博客
	DeleteBlog(ctx context.Context, req *DeleteBlogRequest) error
	// 更新博客
	UpdateBlog(ctx context.Context, req *UpdateBlogRequest) (*Blog, error)
	// 获取博客详细内容
	GetBlogDetails(ctx context.Context, req *GetBlogDetailsRequest) (*Blog, error)
	// 获取创建者的博客列表
	GetBlogList(ctx context.Context, req *GetBlogListRequest) (*BlogList, error)
	// 搜索博客
	SearchBlogs(ctx context.Context, req *SearchBlogsRequest) (*BlogList, error)
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags:         map[string]string{},
		BlogType:     TYPE_ORIGINAL,
		VisibleRange: Range_ALL,
		States:       STATES_DRAFT,
	}
}

type CreateBlogRequest struct {
	// 标题
	Title string `json:"title"`
	// 内容
	Content string `json:"content"`
	// 标签
	Tags map[string]string `json:"tags" gorm:"serializer:json"`
	// 摘要，会在推荐、列表等场景外露
	Abstract string `json:"abstract"`
	// 类型，默认原创
	BlogType BlogType `json:"blogType"`
	// 类型为转载时，填写原文链接
	OriginalLink string `json:"originalLink"`
	// 可见范围，默认全部
	VisibleRange VisibleRange `json:"VisibleRange"`
	// 创建人即作者, 通过当前登录人的Token获取
	CreateBy string `json:"createBy"`
	// 状态，由用户控制是否发布，默认草稿
	States States `json:"states"`
}

// 检查参数，发布时检查
func (req *CreateBlogRequest) Validate() error {
	switch req.States {
	case STATES_DRAFT: // 草稿只需要检查标题
		if req.Title == "" {
			return exception.ValidateError("文章标题不能为空")
		}
	case STATES_PUBLISHED:
		if req.Title == "" || req.Content == "" || len(req.Tags) == 0 || req.Abstract == "" {
			return exception.ValidateError("必填项未填写")
		}
		if req.BlogType.isValid() || req.VisibleRange.isValid() {
			return exception.ValidateError("数据异常")
		}
	default:
		return exception.ValidateError("文章状态异常")
	}

	return nil
}

type DeleteBlogRequest struct {
	Id int `json:"id"`
}

type UpdateBlogRequest struct {
	Id         int        `json:"id"`
	UpdateMode UpdateMode `json:"updateMode"`
	*CreateBlogRequest
}

type GetBlogDetailsRequest struct {
	Id int `json:"id"`
}

func NewGetBlogListRequest() *GetBlogListRequest {
	return &GetBlogListRequest{
		Pagination: common.NewPagination(),
	}
}

type GetBlogListRequest struct {
	// 分页数据
	*common.Pagination
	// 基于文章标题的关键字搜索
	Keywords string `json:"keywords"`
	// 类型
	BlogType *BlogType `json:"blogType"`
	// 可见范围
	VisibleRange *VisibleRange `json:"VisibleRange"`
	// 状态
	States *States `json:"states"`
}

func NewBlogList() *BlogList {
	return &BlogList{
		Items: []*Blog{},
	}
}

type BlogList struct {
	// 文章总数
	Total int64 `json:"total"`
	// 返回的数据
	Items []*Blog `json:"items"`
}

func NewSearchBlogsRequest() *SearchBlogsRequest {
	return &SearchBlogsRequest{
		Pagination: common.NewPagination(),
		Param:      QUERY_BY_TITLE,
	}
}

type SearchBlogsRequest struct {
	// 分页数据
	*common.Pagination
	// 关键字模糊搜索，支持搜索标题或作者
	Param    QueryBy `json:"param"`
	Keywords string  `json:"keywords"`
}
