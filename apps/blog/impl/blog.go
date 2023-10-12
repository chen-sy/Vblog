package impl

import (
	"context"
	"fmt"
	"time"

	"dario.cat/mergo"
	"gitee.com/chensyi/vblog/apps/blog"
	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/common"
)

var _ blog.Service = &blogServiceImpl{}

// 创建博客
func (i *blogServiceImpl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	// 检查用户参数
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// 使用构造函数创建对象
	ins := blog.NewBlog(req)
	// 获取上下文中的userid
	uObj := ctx.Value(user.USER_KEY).(*user.User)
	if uObj == nil {
		return nil, fmt.Errorf("上下文中的user对象不存在")
	}
	ins.CreateBy = uObj.UserName
	// 发布博客时，添加发布时间
	if req.States == blog.STATES_PUBLISHED {
		ins.PublishedAt = time.Now().Unix()
	}
	// 保存到数据库
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// 删除博客
func (i *blogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) error {
	// 查询博客是否存在
	b, err := i.getBlog(ctx, req.Id)
	if err != nil {
		return err
	}
	//TODO删除博客需要校验是否为本人，和更新一样需要校验，可使用gorm.Scopes复用。db.Scopes(CurBlog(r)).Delete(b)
	return i.db.WithContext(ctx).Delete(b).Error
}

// 更新博客
func (i *blogServiceImpl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	// 查询博客是否存在
	ins, err := i.getBlog(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// 更新模式
	switch req.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		ins.CreateBlogRequest = req.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		err := mergo.Merge(ins.CreateBlogRequest, req.CreateBlogRequest, mergo.WithOverride)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("更新模式异常: %d", req.UpdateMode)
	}
	// 获取上下文中的userid
	uObj := ctx.Value(user.USER_KEY).(*user.User)
	if uObj == nil {
		return nil, fmt.Errorf("上下文中的user对象不存在")
	}
	err = i.db.WithContext(ctx).Model(ins).Where("id = ? and create_by = ?", ins.ID, uObj.UserName).Updates(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil

}

// 获取博客详情
func (i *blogServiceImpl) GetBlogDetails(ctx context.Context, req *blog.GetBlogDetailsRequest) (*blog.Blog, error) {
	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	ins := &blog.Blog{}
	query = query.Where("id=?", req.Id)
	if err := query.First(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

// 获取创建者的博客列表
func (i *blogServiceImpl) GetBlogList(ctx context.Context, req *blog.GetBlogListRequest) (*blog.BlogList, error) {
	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	// 初始化装载的list 防止空指针
	list := blog.NewBlogList()
	// 构造查询条件
	if req.Keywords != "" {
		query = query.Where("title LIKE ?", "%"+req.Keywords+"%")
	}
	if req.BlogType != nil {
		query = query.Where("blog_type = ?", *req.BlogType)
	}
	if req.VisibleRange != nil {
		query = query.Where("visible_range = ?", *req.VisibleRange)
	}
	if req.States != nil {
		query = query.Where("states = ?", *req.States)
	}
	// 获取上下文中的userid
	uObj := ctx.Value(user.USER_KEY).(*user.User)
	if uObj == nil {
		return nil, fmt.Errorf("上下文中的user对象不存在")
	}
	query = query.Where("create_by = ?", uObj.UserName)
	// 查询总数
	err := query.Count(&list.Total).Error
	if err != nil {
		return nil, err
	}

	// 查询list数据
	if err := query.Scopes(common.Paginate(req.PageIndex, req.PageSize)).Order("created_at desc").Find(&list.Items).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// 搜索博客
func (i *blogServiceImpl) SearchBlogs(ctx context.Context, req *blog.SearchBlogsRequest) (*blog.BlogList, error) {
	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	list := blog.NewBlogList()
	// 构造查询条件
	switch req.Param {
	case blog.QUERY_BY_TITLE:
		query = query.Where("title LIKE ?", "%"+req.Keywords+"%")
	case blog.QUERY_BY_AUTHOR:
		query = query.Where("create_by LIKE ?", "%"+req.Keywords+"%")
	default:
		return nil, fmt.Errorf("未知的关键字")
	}
	// 只能查询可见范围为全部，且状态为已发布的文章
	query = query.Where("visible_range = ? and states = ? ", blog.Range_ALL, blog.STATES_PUBLISHED)

	// 1. 查询总数量
	err := query.Count(&list.Total).Error
	if err != nil {
		return nil, err
	}

	// 2. 查询一页的数据
	err = query.Scopes(common.Paginate(req.PageIndex, req.PageSize)).Order("created_at desc").Find(&list.Items).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
