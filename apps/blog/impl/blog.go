package impl

import (
	"context"
	"fmt"
	"time"

	"dario.cat/mergo"
	"gitee.com/chensyi/vblog/apps/blog"
	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/common"
	"gitee.com/chensyi/vblog/exception"
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
	// 获取上下文中的user对象
	uObj, ok := ctx.Value(user.CTX_KEY_USER).(*user.User)
	if !ok {
		return nil, fmt.Errorf("上下文中的user对象不存在")
	}
	ins.CreateBy = uObj.ID
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
	// 获取上下文中的user对象
	uObj, ok := ctx.Value(user.CTX_KEY_USER).(*user.User)
	if !ok {
		return fmt.Errorf("上下文中的user对象不存在")
	}
	if b.CreateBy == uObj.ID {
		return i.db.WithContext(ctx).Delete(b).Error
	} else {
		return fmt.Errorf("无权限")
	}
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
	// 获取上下文中的user对象
	uObj, ok := ctx.Value(user.CTX_KEY_USER).(*user.User)
	if !ok {
		return nil, fmt.Errorf("上下文中的user对象不存在")
	}
	err = i.db.WithContext(ctx).Model(ins).Where("id = ? and create_by = ?", ins.ID, uObj.ID).Updates(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil

}

// 获取博客详情
func (i *blogServiceImpl) GetBlogDetails(ctx context.Context, req *blog.GetBlogDetailsRequest) (*blog.Blog, error) {
	// 查询博客是否存在
	ins, err := i.getBlog(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// admin和创建者可直接查看
	id := ctx.Value(user.CTX_KEY_USERID)
	role := ctx.Value(user.CTX_KEY_USERROLE)
	if role == user.ROLE_ADMIN || ins.CreateBy == id {
		return ins, nil
	} else {
		// 非创建者只能查看已发布且公开的
		if ins.States == blog.STATES_PUBLISHED && ins.VisibleRange == blog.Range_ALL {
			return ins, nil
		}
		return nil, exception.NotExistOrNotPermission("找不到资源")
	}
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
	// 获取上下文中的user对象
	uObj, ok := ctx.Value(user.CTX_KEY_USER).(*user.User)
	if !ok {
		return nil, fmt.Errorf("上下文中的user对象不存在")
	}
	query = query.Where("create_by = ?", uObj.ID)
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
		query = query.Where("create_by in (SELECT id FROM `users` WHERE username LIKE ?)", "%"+req.Keywords+"%")
	default:
		return nil, fmt.Errorf("未知的关键字")
	}
	// 获取上下文中的角色
	id := ctx.Value(user.CTX_KEY_USERID)
	role := ctx.Value(user.CTX_KEY_USERROLE)
	if role == user.ROLE_ADMIN {
		// admin可查看全部
	} else if role == user.ROLE_MEMBER {
		// 创建者可查看自己的全部和他人已发布且公开的
		query = query.Where("(create_by=? or (visible_range = ? and states = ? ))", id, blog.Range_ALL, blog.STATES_PUBLISHED)
	} else {
		// 游客只能查看已发布且公开的
		query = query.Where("visible_range = ? and states = ? ", blog.Range_ALL, blog.STATES_PUBLISHED)
	}

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
