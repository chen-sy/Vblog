package api

import (
	"gitee.com/chensyi/vblog/apps/blog"
	"gitee.com/chensyi/vblog/apps/user"
	"gitee.com/chensyi/vblog/middleware"
	"gitee.com/chensyi/vblog/response"
	"github.com/gin-gonic/gin"
)

// 注册路由
func (h *apiHandler) Registry(r gin.IRouter) {
	// 创建一个路由组
	v1 := r.Group("v1").Group("blogs")
	// 设置请求url对应的函数体
	v1.GET("/search/", h.SearchBlogs)
	v1.GET("/details/:id", h.GetBlogDetails)

	// 开启鉴权
	v1.Use(middleware.NewTokenAuther().Auth)
	v1.POST("/", middleware.Required(user.ROLE_MEMBER), h.CreateBlog)
	v1.DELETE("/:id", middleware.Required(user.ROLE_MEMBER), h.DeleteBlog)
	v1.PUT("/:id", middleware.Required(user.ROLE_MEMBER), h.UpdateBlog)
	v1.PATCH("/:id", middleware.Required(user.ROLE_MEMBER), h.PatchBlog)
	v1.GET("/list/", middleware.Required(user.ROLE_MEMBER), h.GetBlogList)
}

// 创建博客
func (h *apiHandler) CreateBlog(c *gin.Context) {
	in := blog.NewCreateBlogRequest()
	// 将请求的json，转换成需要的对象
	err := c.BindJSON(in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 将gin的Context，转换成go Context
	ctx, err := middleware.NewMiddleware().NewContext(c)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.CreateBlog(ctx, in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// 删除博客
func (h *apiHandler) DeleteBlog(c *gin.Context) {
	ctx, err := middleware.NewMiddleware().NewContext(c)
	if err != nil {
		response.Failed(c, err)
		return
	}
	err = h.svc.DeleteBlog(ctx, &blog.DeleteBlogRequest{Id: c.Param("id")})
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "delete success")
}

// 更新博客(全量)
func (h *apiHandler) UpdateBlog(c *gin.Context) {
	in := blog.NewPutUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 将gin的Context，转换成go Context
	ctx, err := middleware.NewMiddleware().NewContext(c)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.UpdateBlog(ctx, in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// 更新博客(增量)
func (h *apiHandler) PatchBlog(c *gin.Context) {
	in := blog.NewPatchUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 将gin的Context，转换成go Context
	ctx, err := middleware.NewMiddleware().NewContext(c)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.UpdateBlog(ctx, in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// 获取博客详细内容，admin和创建者可直接查看，游客只能查看已发布且公开的
func (h *apiHandler) GetBlogDetails(c *gin.Context) {
	ctx, err := middleware.NewMiddleware().NewContextGeneral(c)
	if err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.GetBlogDetails(ctx, &blog.GetBlogDetailsRequest{Id: c.Param("id")})
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// 获取创建者的博客列表
func (h *apiHandler) GetBlogList(c *gin.Context) {
	in := blog.NewGetBlogListRequest()
	in.ParsePageIndex(c.Query("page_index"))
	in.ParsePageSize(c.Query("page_size"))
	in.Keywords = c.Query("keywords")
	switch c.Query("blog_type") {
	case "original":
		in.BlogType.Set(blog.TYPE_ORIGINAL)
	case "reprinted":
		in.BlogType.Set(blog.TYPE_REPRINTED)
	}
	switch c.Query("visible_range") {
	case "all":
		in.VisibleRange.Set(blog.Range_ALL)
	case "own":
		in.VisibleRange.Set(blog.Range_OWN)
	}
	switch c.Query("status") {
	case "draft":
		in.Status.Set(blog.STATUS_DRAFT)
	case "published":
		in.Status.Set(blog.STATUS_PUBLISHED)
	}
	// 将gin的Context，转换成go Context
	ctx, err := middleware.NewMiddleware().NewContext(c)
	if err != nil {
		response.Failed(c, err)
		return
	}
	list, err := h.svc.GetBlogList(ctx, in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, list)
}

// 搜索博客
func (h *apiHandler) SearchBlogs(c *gin.Context) {
	in := blog.NewSearchBlogsRequest()
	in.ParsePageIndex(c.Query("page_index"))
	in.ParsePageSize(c.Query("page_size"))
	switch c.Query("param") {
	case "title":
		in.Param = blog.QUERY_BY_TITLE
	case "author":
		in.Param = blog.QUERY_BY_AUTHOR
	}
	in.Keywords = c.Query("keywords")
	ctx, err := middleware.NewMiddleware().NewContextGeneral(c)
	if err != nil {
		response.Failed(c, err)
		return
	}
	list, err := h.svc.SearchBlogs(ctx, in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, list)
}
